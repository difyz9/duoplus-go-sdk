package duoplus

import "context"

const (
	pathCloudDiskList      = "/api/v1/cloudDisk/list"
	pathCloudDiskPushFiles = "/api/v1/cloudDisk/pushFiles"
)

type CloudDiskListRequest struct {
	Keyword string `json:"keyword,omitempty"`
	PaginationRequest
}

type CloudDiskFile struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	OriginalFileName string `json:"original_file_name"`
}

type CloudDiskListResponse struct {
	List   []CloudDiskFile `json:"list"`
	Limit  int             `json:"limit,omitempty"`
	Offset int             `json:"offset,omitempty"`
	Pagination
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

func (c *Client) ListCloudDiskFiles(ctx context.Context, req CloudDiskListRequest) (*CloudDiskListResponse, error) {
	var resp CloudDiskListResponse
	if err := c.do(ctx, pathCloudDiskList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PushCloudDiskFiles(ctx context.Context, req PushFilesRequest) (*PushFilesResponse, error) {
	var resp PushFilesResponse
	if err := c.do(ctx, pathCloudDiskPushFiles, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
