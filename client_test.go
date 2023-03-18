package chatgpt

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	chatgpt_errors "github.com/ayush6624/go-chatgpt/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("test-apikey")
	if err != nil {
		t.Errorf("NewClient() error = %v", err)
	}

	_, err = NewClient("")
	assert.NotNil(t, err)
	assert.Equal(t, err, chatgpt_errors.ErrAPIKeyRequired)

	_, err = NewClientWithConfig(&Config{})
	assert.NotNil(t, err)
	assert.Equal(t, err, chatgpt_errors.ErrAPIKeyRequired)

	_, err = NewClientWithConfig(&Config{
		APIKey: "test-apikey",
	})
	assert.Nil(t, err)
}

func TestClient2_sendRequest(t *testing.T) {
	// Create a new test HTTP server and client to handle requests
	testServer, client := newTestServerAndClient()

	// Create a new request
	req, err := http.NewRequest("POST", testServer.URL, nil)
	assert.NoError(t, err)

	// Send the request using the ChatGPT client
	res, err := client.sendRequest(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	expectedHeader := map[string]string{
		"Authorization":       fmt.Sprintf("Bearer %s", "mock_api_key"),
		"OpenAI-Organization": "mock_organization_id",
		"Content-Type":        "application/json",
		"Accept":              "application/json",
	}

	// Check that the request's header is set correctly, after sendRequest was called
	for key, value := range expectedHeader {
		assert.Equal(t, req.Header.Get(key), value)
	}

	testServer, client = newTestClientWithInvalidStatusCode()
	// Prepare a test request
	req, err = http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)

	resp, err := client.sendRequest(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
