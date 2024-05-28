package main

import (
	"kafka-microservice/consumer"
	"kafka-microservice/logger"
	"kafka-microservice/metrics"
	"kafka-microservice/producer"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	logger.Init()

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	groupID := os.Getenv("GROUP_ID")
	inputTopic := os.Getenv("INPUT_TOPIC")
	outputTopic := os.Getenv("OUTPUT_TOPIC")
	numConsumerWorkers, _ := strconv.Atoi(os.Getenv("NUM_CONSUMER_WORKERS"))
	numProducerWorkers, _ := strconv.Atoi(os.Getenv("NUM_PRODUCER_WORKERS"))

	// Initialize metrics
	metrics.Init()

	go func() {
		http.Handle("/metrics", metrics.Handler())
		http.ListenAndServe(":8080", nil)
	}()

	consumer := consumer.NewConsumer([]string{kafkaBroker}, groupID, []string{inputTopic}, numConsumerWorkers)
	go consumer.Start()

	producer := producer.NewProducer([]string{kafkaBroker}, outputTopic, numProducerWorkers)
	go producer.Start()

	// Graceful shutdown
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)
	<-sigterm
}
