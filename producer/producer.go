package producer

import (
	"kafka-microservice/logger"

	"github.com/IBM/sarama"
)

type Producer struct {
	Brokers    []string
	Topic      string
	NumWorkers int
	producer   sarama.SyncProducer
}

func NewProducer(brokers []string, topic string, numWorkers int) *Producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		logger.Log().Fatalf("Error creating producer: %v", err)
	}

	return &Producer{
		Brokers:    brokers,
		Topic:      topic,
		NumWorkers: numWorkers,
		producer:   producer,
	}
}

func (p *Producer) Start() {
	logger.Log().Info("Producer started")
	for i := 0; i < p.NumWorkers; i++ {
		go func() {
			for msg := range messageQueue {
				p.produceMessage(p.Topic, msg)
			}
		}()
	}
}

var messageQueue = make(chan []byte, 100)

func ProduceMessage(message []byte) {
	messageQueue <- message
	logger.Log().Infof("Message produced: %s", string(message))
}

func (p *Producer) produceMessage(topic string, message []byte) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		logger.Log().Errorf("Failed to produce message: %v", err)
		return
	}

	logger.Log().Infof("Message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset)
}
