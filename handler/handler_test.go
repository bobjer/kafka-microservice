package handler

import (
	"testing"
)

func TestHandleMessage(t *testing.T) {
	handler := func(msg []byte) ([]byte, error) {
		return msg, nil
	}

	tests := []struct {
		name    string
		message []byte
		want    []byte
		wantErr bool
	}{
		{"Valid message", []byte("test message"), []byte("test message"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleMessage(handler, tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("HandleMessage() = %s, want %s", got, tt.want)
			}
		})
	}
}
