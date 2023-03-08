package chatgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayush6624/go-chatgpt/utils"
)

const (
	apiURL = "https://api.openai.com/v1"
)

type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Config
	config *Config
}

type Config struct {
	// Base URL for API requests.
	BaseURL string

	// API Key (Required)
	APIKey string

	// Organization ID (Optional)
	OrganizationID string
}

func NewClient(apikey string) (*Client, error) {
	if apikey == "" {
		return nil, chatgpt_errors.ErrAPIKeyRequired
	}

	return &Client{
		client: &http.Client{},
		config: &Config{
			BaseURL: apiURL,
			APIKey:  apikey,
		},
	}, nil
}

func NewClientWithConfig(config *Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, chatgpt_errors.ErrAPIKeyRequired
	}

	return &Client{
		client: &http.Client{},
		config: config,
	}, nil
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey))
	if c.config.OrganizationID != "" {
		req.Header.Set("OpenAI-Organization", c.config.OrganizationID)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		// Parse body
		var errMessage interface{}
		if err := json.NewDecoder(res.Body).Decode(&errMessage); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("api request failed: status Code: %d %s %s Message: %+v", res.StatusCode, res.Status, res.Request.URL, errMessage)
	}

	return res, nil
}
