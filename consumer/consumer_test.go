package consumer

import (
	"testing"
)

func TestProcessMessageInWorker(t *testing.T) {
	msg := []byte("test message")
	ProcessMessageInWorker(msg)
}
