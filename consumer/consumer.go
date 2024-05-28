package consumer

import (
	"context"
	"kafka-microservice/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type Consumer struct {
	Brokers       []string
	GroupID       string
	Topics        []string
	NumWorkers    int
	consumerGroup sarama.ConsumerGroup
}

func NewConsumer(brokers []string, groupID string, topics []string, numWorkers int) *Consumer {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		logger.Log().Fatalf("Error creating consumer group: %v", err)
	}

	return &Consumer{
		Brokers:       brokers,
		GroupID:       groupID,
		Topics:        topics,
		NumWorkers:    numWorkers,
		consumerGroup: consumerGroup,
	}
}

func (c *Consumer) Start() {
	defer c.consumerGroup.Close()

	logger.Log().Info("Consumer started")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := &consumerGroupHandler{}

	for i := 0; i < c.NumWorkers; i++ {
		go func() {
			for {
				if err := c.consumerGroup.Consume(ctx, c.Topics, handler); err != nil {
					logger.Log().Fatalf("Error during consumption: %v", err)
				}
				if ctx.Err() != nil {
					return
				}
			}
		}()
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)
	<-sigterm
	cancel()
}

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	logger.Log().Info("Consumer group setup complete")
	return nil
}

func (consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	logger.Log().Info("Consumer group cleanup complete")
	return nil
}

func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		logger.Log().Infof("Message claimed: value = %s, timestamp = %v, topic = %s", string(msg.Value), msg.Timestamp, msg.Topic)
		// Send the message to worker for processing
		logger.Log().Infof("Sending message to worker: value = %s", string(msg.Value))
		ProcessMessageInWorker(msg.Value)
		logger.Log().Infof("Message processed: value = %s", string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}
