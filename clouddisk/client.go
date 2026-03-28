package clouddisk

import (
	"context"

	"duoplus-go-sdk/common"
	"duoplus-go-sdk/internal/clientcore"
)

const (
	pathList      = "/api/v1/cloudDisk/list"
	pathPushFiles = "/api/v1/cloudDisk/pushFiles"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client {
	return &Client{core: core}
}

type ListRequest struct {
	Keyword string `json:"keyword,omitempty"`
	common.PaginationRequest
}

type File struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	OriginalFileName string `json:"original_file_name"`
}

type ListResponse struct {
	List   []File `json:"list"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
	common.Pagination
}

type PushFilesRequest struct {
	IDs      []string `json:"ids"`
	ImageIDs []string `json:"image_ids"`
	DestDir  string   `json:"dest_dir"`
}

type PushFilesSuccess struct {
	ImageID string `json:"image_id"`
	ID      string `json:"id"`
}

type PushFilesFailure struct {
	ImageID string `json:"image_id"`
	ID      string `json:"id"`
	Err     string `json:"err"`
}

type PushFilesResponse struct {
	Message string             `json:"message"`
	Success []PushFilesSuccess `json:"success"`
	Fail    []PushFilesFailure `json:"fail"`
}

func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, pathList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PushFiles(ctx context.Context, req PushFilesRequest) (*PushFilesResponse, error) {
	var resp PushFilesResponse
	if err := c.core.Do(ctx, pathPushFiles, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
