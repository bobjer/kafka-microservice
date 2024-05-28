package message

import (
	"encoding/json"
	"errors"
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
	Status    string    `json:"status,omitempty"`
}

func ValidateMessage(msg Message) error {
	if time.Since(msg.Timestamp) > 24*time.Hour {
		return errors.New("message timestamp is older than 24 hours")
	}
	return nil
}

func AddStatus(msg *Message) {
	if len(msg.Data) > 10 {
		msg.Status = "valid"
	} else {
		msg.Status = "invalid"
	}
}

func ProcessMessage(message []byte) ([]byte, error) {
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		return nil, err
	}

	if err := ValidateMessage(msg); err != nil {
		return nil, err
	}

	AddStatus(&msg)

	return json.Marshal(msg)
}
