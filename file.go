package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

type FileStatus string

const (
	FILESTATUS_UPLOADED  FileStatus = "uploaded"
	FILESTATUS_PROCESSED FileStatus = "processed"
	FILESTATUS_PENDING   FileStatus = "pending"
	FILESTATUS_ERROR     FileStatus = "error"
	FILESTATUS_DELETING  FileStatus = "deleting"
	FILESTATUS_DELETED   FileStatus = "deleted"
)

type FilePurpose string

const (
	FILEPURPOSE_FINETUNE FilePurpose = "fine-tune"
)

type File struct {
	Id            string      `json:"id"`
	Object        string      `json:"object"`
	Bytes         int         `json:"bytes"`
	CreatedAt     int         `json:"created_at"`
	Filename      string      `json:"filename"`
	Purpose       FilePurpose `json:"purpose"`
	Status        FileStatus  `json:"status"`
	StatusDetails *string     `json:"status_details"`
}

type FileList struct {
	Data   []File `json:"data"`
	Object string `json:"object"`
}

type DeleteFileResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

func (c *Client) ListFiles(ctx context.Context) (*FileList, error) {
	endpoint := "/files"
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

	var fileList FileList
	if err := json.NewDecoder(res.Body).Decode(&fileList); err != nil {
		return nil, err
	}

	return &fileList, nil
}

func (c *Client) UploadFile(ctx context.Context, file io.Reader) (*File, error) {
	endpoint := "/files"
	httpReq, err := http.NewRequest("POST", c.config.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	err = writer.WriteField("purpose", string(FILEPURPOSE_FINETUNE))
	if err != nil {
		return nil, err
	}

	part, err := writer.CreateFormFile("file", "mydata.jsonl")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	httpReq.Body = io.NopCloser(&requestBody)

	res, err := c.sendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var responseFile File
	if err := json.NewDecoder(res.Body).Decode(&responseFile); err != nil {
		return nil, err
	}

	return &responseFile, nil
}

func (c *Client) DeleteFile(ctx context.Context, fileId string) (*DeleteFileResponse, error) {
	endpoint := "/files/" + fileId
	httpReq, err := http.NewRequest("DELETE", c.config.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)

	res, err := c.sendRequest(ctx, httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var deletedFile DeleteFileResponse
	if err := json.NewDecoder(res.Body).Decode(&deletedFile); err != nil {
		return nil, err
	}

	return &deletedFile, nil
}

func (c *Client) RetrieveFile(ctx context.Context, fileId string) (*File, error) {
	endpoint := "/files/" + fileId
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

	var file File
	if err := json.NewDecoder(res.Body).Decode(&file); err != nil {
		return nil, err
	}

	return &file, nil
}

func (c *Client) RetrieveFileContent(ctx context.Context, fileId string) (*string, error) {
	endpoint := "/files/" + fileId + "/content"

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

	var fileContent string
	if err := json.NewDecoder(res.Body).Decode(&fileContent); err != nil {
		return nil, err
	}

	return &fileContent, nil
}
