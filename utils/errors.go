package chatgpt_errors

import "errors"

var (
	// ErrAPIKeyRequired is returned when the API Key is not provided
	ErrAPIKeyRequired = errors.New("API Key is required")

	// ErrInvalidModel is returned when the model is invalid
	ErrInvalidModel = errors.New("invalid model")

	// ErrNoMessages is returned when no messages are provided
	ErrNoMessages = errors.New("no messages provided")

	// ErrInvalidRole is returned when the role is invalid
	ErrInvalidRole = errors.New("invalid role. Only `user`, `system` and `assistant` are supported")

	// ErrInvalidTemperature is returned when the temperature is invalid
	ErrInvalidTemperature = errors.New("invalid temperature. 0<= temp <= 2")

	// ErrInvalidPresencePenalty
	ErrInvalidPresencePenalty = errors.New("invalid presence penalty. -2<= presence penalty <= 2")

	// ErrInvalidFrequencyPenalty
	ErrInvalidFrequencyPenalty = errors.New("invalid frequency penalty. -2<= frequency penalty <= 2")
)