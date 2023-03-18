package chatgpt

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	chatgpt_errors "github.com/ayush6624/go-chatgpt/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name          string
		request       *ChatCompletionRequest
		expectedError error
	}{
		{
			name:          "Valid request",
			request:       validRequest(),
			expectedError: nil,
		},
		{
			name: "Invalid model",
			request: &ChatCompletionRequest{
				Model:    "invalid-model",
				Messages: validRequest().Messages,
			},
			expectedError: chatgpt_errors.ErrInvalidModel,
		},
		{
			name:          "No messages",
			request:       &ChatCompletionRequest{},
			expectedError: chatgpt_errors.ErrNoMessages,
		},
		{
			name: "Invalid role",
			request: &ChatCompletionRequest{
				Model: GPT35Turbo,
				Messages: []ChatMessage{
					{
						Role:    "invalid-role",
						Content: "Hello",
					},
				},
			},
			expectedError: chatgpt_errors.ErrInvalidRole,
		},
		{
			name: "Invalid temperature",
			request: &ChatCompletionRequest{
				Model:       GPT35Turbo,
				Messages:    validRequest().Messages,
				Temperature: -0.5,
			},
			expectedError: chatgpt_errors.ErrInvalidTemperature,
		},
		{
			name: "Invalid presence penalty",
			request: &ChatCompletionRequest{
				Model:           GPT35Turbo,
				Messages:        validRequest().Messages,
				PresencePenalty: -3,
			},
			expectedError: chatgpt_errors.ErrInvalidPresencePenalty,
		},
		{
			name: "Invalid frequency penalty",
			request: &ChatCompletionRequest{
				Model:            GPT35Turbo,
				Messages:         validRequest().Messages,
				FrequencyPenalty: -3,
			},
			expectedError: chatgpt_errors.ErrInvalidFrequencyPenalty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.request)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func validRequest() *ChatCompletionRequest {
	return &ChatCompletionRequest{
		Model: GPT35Turbo,
		Messages: []ChatMessage{
			{
				Role:    ChatGPTModelRoleUser,
				Content: "Hello",
			},
		},
	}
}

func newTestServerAndClient() (*httptest.Server, *Client) {
	// Create a new test HTTP server to handle requests
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "id": "chatcmpl-abcd", "object": "chat.completion", "created_at": 0, "choices": [ { "index": 0, "message": { "role": "assistant", "content": "\n\n Sample response" }, "finish_reason": "stop" } ], "usage": { "prompt_tokens": 19, "completion_tokens": 47, "total_tokens": 66 }}`))
	}))

	// Create a new client with the test server's URL and a mock API key
	return testServer, &Client{
		client: http.DefaultClient,
		config: &Config{
			BaseURL:        testServer.URL,
			APIKey:         "mock_api_key",
			OrganizationID: "mock_organization_id",
		},
	}
}

func newTestClientWithInvalidResponse() (*httptest.Server, *Client) {
	// Create a new test HTTP server to handle requests
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ fakejson }`))
	}))

	// Create a new client with the test server's URL and a mock API key
	return testServer, &Client{
		client: http.DefaultClient,
		config: &Config{
			BaseURL:        testServer.URL,
			APIKey:         "mock_api_key",
			OrganizationID: "mock_organization_id",
		},
	}
}

func newTestClientWithInvalidStatusCode() (*httptest.Server, *Client) {
	// Create a new test HTTP server to handle requests
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "bad request" }`))
	}))

	// Create a new client with the test server's URL and a mock API key
	return testServer, &Client{
		client: http.DefaultClient,
		config: &Config{
			BaseURL:        testServer.URL,
			APIKey:         "mock_api_key",
			OrganizationID: "mock_organization_id",
		},
	}
}

func TestSend(t *testing.T) {
	server, client := newTestServerAndClient()
	defer server.Close()

	_, err := client.Send(context.Background(), &ChatCompletionRequest{
		Model: GPT35Turbo,
		Messages: []ChatMessage{
			{
				Role:    ChatGPTModelRoleUser,
				Content: "Hello",
			},
		},
	})
	assert.NoError(t, err)

	_, err = client.Send(context.Background(), &ChatCompletionRequest{
		Model: "invalid model",
		Messages: []ChatMessage{
			{
				Role:    ChatGPTModelRoleUser,
				Content: "Hello",
			},
		},
	})
	assert.Error(t, err)

	server, client = newTestClientWithInvalidResponse()
	defer server.Close()

	_, err = client.Send(context.Background(), &ChatCompletionRequest{
		Model: GPT35Turbo,
		Messages: []ChatMessage{
			{
				Role:    ChatGPTModelRoleUser,
				Content: "Hello",
			},
		},
	})
	assert.Error(t, err)

	server, client = newTestClientWithInvalidStatusCode()
	defer server.Close()

	_, err = client.Send(context.Background(), &ChatCompletionRequest{
		Model: GPT35Turbo,
		Messages: []ChatMessage{
			{
				Role:    ChatGPTModelRoleUser,
				Content: "Hello",
			},
		},
	})
	assert.Error(t, err)

}

func TestSimpleSend(t *testing.T) {
	server, client := newTestServerAndClient()
	defer server.Close()

	_, err := client.SimpleSend(context.Background(), "Hello")
	assert.NoError(t, err)
}
