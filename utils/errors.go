package chatgpt_errors

import "errors"

var (
	// ErrAPIKeyRequired is returned when the API Key is not provided
	ErrAPIKeyRequired = errors.New("API Key is required")
)