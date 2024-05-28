package handler

import (
	"kafka-microservice/logger"
)

type MessageHandlerFunc func([]byte) ([]byte, error)

func HandleMessage(handler MessageHandlerFunc, msg []byte) ([]byte, error) {
	logger.Log().Infof("Handling message: %s", string(msg))
	processedMsg, err := handler(msg)
	if err != nil {
		logger.Log().Errorf("Error handling message: %v", err)
		return nil, err
	}
	logger.Log().Infof("Handled message: %s", string(processedMsg))
	return processedMsg, nil
}
