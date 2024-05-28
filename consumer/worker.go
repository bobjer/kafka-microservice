package consumer

import (
	"kafka-microservice/handler"
	"kafka-microservice/logger"
	"kafka-microservice/message"
	"kafka-microservice/producer"
)

func ProcessMessageInWorker(msg []byte) {
	logger.Log().Infof("Processing message: %s", string(msg))
	processedMessage, err := handler.HandleMessage(message.ProcessMessage, msg)
	if err != nil {
		logger.Log().Errorf("Error processing message: %v", err)
		return
	}

	producer.ProduceMessage(processedMessage)
}
