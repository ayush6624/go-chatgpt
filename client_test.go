package chatgpt

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ayush6624/go-chatgpt/utils"
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
	// Create a new test HTTP server to handle requests
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	// Create a new client with the test server's URL and a mock API key
	client := &Client{
		client: http.DefaultClient,
		config: &Config{
			BaseURL: testServer.URL,
			APIKey:  "mock_api_key",
			OrganizationID: "mock_organization_id",
		},
	}

	// Create a new request
	req, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Send the request
	res, err := client.sendRequest(context.Background(), req)
	if err != nil {
		t.Fatalf("sendRequest failed: %v", err)
	}

	expectedHeader := map[string]string{
		"Authorization":       fmt.Sprintf("Bearer %s", "mock_api_key"),
		"OpenAI-Organization": "mock_organization_id",
		"Content-Type":        "application/json",
		"Accept":              "application/json",
	}

	// Check that the request's header is set correctly
	for key, value := range expectedHeader {
		if req.Header.Get(key) != value {
			t.Errorf("expected header %s to be %s, got %s", key, value, req.Header.Get(key))
		}
	}

	// Check that the response's status code is correct
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}
