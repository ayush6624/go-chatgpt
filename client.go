package chatgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		return nil, fmt.Errorf("API Key is required")
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
		return nil, fmt.Errorf("API Key is required")
	}

	return &Client{
		client: &http.Client{},
		config: config,
	}, nil
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request) (error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey))
	if c.config.OrganizationID != "" {
		req.Header.Set("OpenAI-Organization", c.config.OrganizationID)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("api request failed: status Code: %d %s %s", res.StatusCode, res.Status, res.Request.URL)
	}

	var v interface{}
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	log.Printf("%+v", v)
	return nil
}
