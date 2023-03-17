package chatgpt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	chatgpt_errors "github.com/ayush6624/go-chatgpt/utils"
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
				Model:            GPT35Turbo,
				Messages:         validRequest().Messages,
				PresencePenalty:  -3,
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
