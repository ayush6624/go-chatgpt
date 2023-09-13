package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type FineTuningJobStatus string

const (
	FINETUNINGJOBSTATUS_CREATED   FineTuningJobStatus = "created"
	FINETUNINGJOBSTATUS_PENDING   FineTuningJobStatus = "pending"
	FINETUNINGJOBSTATUS_RUNNING   FineTuningJobStatus = "running"
	FINETUNINGJOBSTATUS_SUCCEEDED FineTuningJobStatus = "succeeded"
	FINETUNINGJOBSTATUS_FAILED    FineTuningJobStatus = "failed"
	FINETUNINGJOBSTATUS_CANCELLED FineTuningJobStatus = "cancelled"
)

type FineTuningJob struct {
	Object          string              `json:"object"`
	Id              string              `json:"id"`
	Model           string              `json:"model,omitempty"`
	CreatedAt       int                 `json:"created_at"`
	FinishedAt      int                 `json:"finished_at,omitempty"`
	FineTunedModel  string              `json:"fine_tuned_model,omitempty"`
	OrganizationId  string              `json:"organization_id,omitempty"`
	ResultFiles     []string            `json:"result_files,omitempty"`
	Status          FineTuningJobStatus `json:"status,omitempty"`
	ValidationFile  interface{}         `json:"validation_file,omitempty"`
	TrainingFile    string              `json:"training_file,omitempty"`
	Hyperparameters struct {
		NEpochs int `json:"n_epochs"`
	} `json:"hyperparameters,omitempty"`
	TrainedTokens int    `json:"trained_tokens,omitempty"`
	Level         string `json:"level,omitempty"`
	Message       string `json:"message,omitempty"`
	Data          string `json:"data,omitempty"`
	Type          string `json:"type,omitempty"`
}

type FineTuningRequest struct {
	TrainingFile   string       `json:"training_file"`
	ValidationFile *string      `json:"validation_file"`
	Model          ChatGPTModel `json:"model"`
}

type FineTuningResponse struct {
	Object         string        `json:"object"`
	Id             string        `json:"id"`
	Model          string        `json:"model"`
	CreatedAt      int           `json:"created_at"`
	FineTunedModel interface{}   `json:"fine_tuned_model"`
	OrganizationId string        `json:"organization_id"`
	ResultFiles    []interface{} `json:"result_files"`
	Status         string        `json:"status"`
	ValidationFile interface{}   `json:"validation_file"`
	TrainingFile   string        `json:"training_file"`
}

type FineTuningList struct {
	Object  string          `json:"object"`
	Data    []FineTuningJob `json:"data"`
	HasMore bool            `json:"has_more"`
}

func (c *Client) CreateFineTuningRequest(ctx context.Context, req FineTuningRequest) (*FineTuningResponse, error) {
	endpoint := "/fine_tuning/jobs"

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", c.config.BaseURL+endpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	res, err := c.sendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response FineTuningResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ListFineTuningJobs(ctx context.Context, after *string, limit *int) (*FineTuningList, error) {
	endpoint := "/fine_tuning/jobs"

	queryParams := map[string]string{}

	if nil != after {
		queryParams["after"] = *after
	}

	if nil != limit {
		queryParams["limit"] = strconv.Itoa(*limit)
	}

	if len(queryParams) > 0 {
		values := url.Values{}
		for key, value := range queryParams {
			values.Add(key, value)
		}

		endpoint = endpoint + "?" + values.Encode()
	}

	httpReq, err := http.NewRequest("GET", c.config.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	res, err := c.sendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var fineTuningList FineTuningList
	if err := json.NewDecoder(res.Body).Decode(&fineTuningList); err != nil {
		return nil, err
	}

	return &fineTuningList, nil
}

func (c *Client) RetrieveFineTuningJob(ctx context.Context, fineTuningJobId string) (*FineTuningJob, error) {
	endpoint := "/fine_tuning/jobs/" + fineTuningJobId

	httpReq, err := http.NewRequest("GET", c.config.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	res, err := c.sendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var job FineTuningJob
	if err := json.NewDecoder(res.Body).Decode(&job); err != nil {
		return nil, err
	}

	return &job, nil
}

func (c *Client) CancelFineTuningJob(ctx context.Context, fineTuningJobId string) (*FineTuningJob, error) {
	endpoint := "/fine_tuning/jobs/" + fineTuningJobId + "/cancel"

	httpReq, err := http.NewRequest("POST", c.config.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	res, err := c.sendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var job FineTuningJob
	if err := json.NewDecoder(res.Body).Decode(&job); err != nil {
		return nil, err
	}

	return &job, nil
}

func (c *Client) ListFineTuningEvents(ctx context.Context, fineTuningJobId string, after *string, limit *int) (*FineTuningList, error) {
	endpoint := "/fine_tuning/jobs/" + fineTuningJobId + "/events"

	queryParams := map[string]string{}

	if nil != after {
		queryParams["after"] = *after
	}

	if nil != limit {
		queryParams["limit"] = strconv.Itoa(*limit)
	}

	if len(queryParams) > 0 {
		values := url.Values{}
		for key, value := range queryParams {
			values.Add(key, value)
		}

		endpoint = endpoint + "?" + values.Encode()
	}

	httpReq, err := http.NewRequest("GET", c.config.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	res, err := c.sendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var fineTuningList FineTuningList
	if err := json.NewDecoder(res.Body).Decode(&fineTuningList); err != nil {
		return nil, err
	}

	return &fineTuningList, nil
}
