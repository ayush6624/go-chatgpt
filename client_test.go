package chatgpt_test

import (
	"testing"
	"github.com/ayush6624/go-chatgpt"
)

func TestNewClient(t *testing.T) {
	_, err := chatgpt.NewClient("test-apikey")
	if err != nil {
		t.Errorf("NewClient() error = %v", err)
	}
}