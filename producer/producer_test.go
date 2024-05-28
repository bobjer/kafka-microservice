package producer

import (
	"testing"
)

func TestProduceMessage(t *testing.T) {
	msg := []byte("test message")
	ProduceMessage(msg)
}
