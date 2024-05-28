package message

import (
	"encoding/json"
	"testing"
	"time"
)

func TestValidateMessage(t *testing.T) {
	validMsg := Message{
		ID:        "test_1",
		Timestamp: time.Now(),
		Data:      "Sample data",
	}
	oldMsg := Message{
		ID:        "test_2",
		Timestamp: time.Now().Add(-25 * time.Hour),
		Data:      "Sample data",
	}

	tests := []struct {
		name    string
		message Message
		wantErr bool
	}{
		{"Valid message", validMsg, false},
		{"Old message", oldMsg, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMessage(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddStatus(t *testing.T) {
	msg := Message{
		ID:        "test_1",
		Timestamp: time.Now(),
		Data:      "Short",
	}

	AddStatus(&msg)

	if msg.Status != "invalid" {
		t.Errorf("Expected status to be 'invalid', got %s", msg.Status)
	}

	msg.Data = "This is a long data string"

	AddStatus(&msg)

	if msg.Status != "valid" {
		t.Errorf("Expected status to be 'valid', got %s", msg.Status)
	}
}

func TestProcessMessage(t *testing.T) {
	validMsg := Message{
		ID:        "test_1",
		Timestamp: time.Now(),
		Data:      "Sample data",
	}
	validMsgBytes, _ := json.Marshal(validMsg)

	oldMsg := Message{
		ID:        "test_2",
		Timestamp: time.Now().Add(-25 * time.Hour),
		Data:      "Sample data",
	}
	oldMsgBytes, _ := json.Marshal(oldMsg)

	tests := []struct {
		name    string
		message []byte
		wantErr bool
	}{
		{"Valid message", validMsgBytes, false},
		{"Old message", oldMsgBytes, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ProcessMessage(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
