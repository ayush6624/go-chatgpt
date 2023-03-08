package chatgpt_test

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/ayush6624/go-chatgpt"
	"github.com/ayush6624/go-chatgpt/utils"
)

func TestNewClient(t *testing.T) {
	_, err := chatgpt.NewClient("test-apikey")
	if err != nil {
		t.Errorf("NewClient() error = %v", err)
	}

	_, err = chatgpt.NewClient("")
	assert.NotNil(t, err)
	assert.Equal(t, err, chatgpt_errors.ErrAPIKeyRequired)

	_, err = chatgpt.NewClientWithConfig(&chatgpt.Config{
	})
	assert.NotNil(t, err)
	assert.Equal(t, err, chatgpt_errors.ErrAPIKeyRequired)

	_, err = chatgpt.NewClientWithConfig(&chatgpt.Config{
		APIKey: "test-apikey",
	})
	assert.Nil(t, err)
}
