package duoplus

import "context"

const (
	pathCloudPhoneGroupList   = "/api/v1/cloudPhone/groupList"
	pathCloudPhoneAddToGroup  = "/api/v1/cloudPhone/addToGroup"
	pathCloudPhoneMoveToGroup = "/api/v1/cloudPhone/moveToGroup"
	pathCloudPhoneCreateGroup = "/api/v1/cloudPhone/createGroup"
	pathCloudPhoneUpdateGroup = "/api/v1/cloudPhone/updateGroup"
	pathCloudPhoneDeleteGroup = "/api/v1/cloudPhone/deleteGroup"
)

type GroupListRequest struct {
	Page int `json:"page,omitempty"`
}

type CloudPhoneGroup struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	Remark string `json:"remark"`
}

type GroupListResponse struct {
	List []CloudPhoneGroup `json:"list"`
	Pagination
}

type GroupImagesRequest struct {
	ID       string   `json:"id"`
	ImageIDs []string `json:"image_ids"`
}

type GroupMutationItem struct {
	Name   string `json:"name"`
	Sort   int    `json:"sort,omitempty"`
	Remark string `json:"remark,omitempty"`
}

type CreateGroupItem struct {
	Name   string `json:"name"`
	Sort   int    `json:"sort,omitempty"`
	Remark string `json:"remark,omitempty"`
}

type UpdateGroupItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Sort   int    `json:"sort,omitempty"`
	Remark string `json:"remark,omitempty"`
}

type CreateGroupsRequest struct {
	List []CreateGroupItem `json:"list"`
}

type UpdateGroupsRequest struct {
	List []UpdateGroupItem `json:"list"`
}

type GroupMutationSuccess struct {
	Index  int    `json:"index"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	Remark string `json:"remark"`
}

type GroupMutationFailure struct {
	Index   int    `json:"index"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GroupMutationResponse struct {
	Success []GroupMutationSuccess `json:"success"`
	Fail    []GroupMutationFailure `json:"fail"`
}

func (c *Client) ListCloudPhoneGroups(ctx context.Context, req GroupListRequest) (*GroupListResponse, error) {
	var resp GroupListResponse
	if err := c.do(ctx, pathCloudPhoneGroupList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) AddCloudPhonesToGroup(ctx context.Context, groupID string, imageIDs []string) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathCloudPhoneAddToGroup, GroupImagesRequest{ID: groupID, ImageIDs: imageIDs}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) MoveCloudPhonesToGroup(ctx context.Context, groupID string, imageIDs []string) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathCloudPhoneMoveToGroup, GroupImagesRequest{ID: groupID, ImageIDs: imageIDs}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateCloudPhoneGroups(ctx context.Context, req CreateGroupsRequest) (*GroupMutationResponse, error) {
	var resp GroupMutationResponse
	if err := c.do(ctx, pathCloudPhoneCreateGroup, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateCloudPhoneGroups(ctx context.Context, req UpdateGroupsRequest) (*GroupMutationResponse, error) {
	var resp GroupMutationResponse
	if err := c.do(ctx, pathCloudPhoneUpdateGroup, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteCloudPhoneGroups(ctx context.Context, ids []string) (*ProxyOperationResponse, error) {
	var resp ProxyOperationResponse
	if err := c.do(ctx, pathCloudPhoneDeleteGroup, ProxyIDsRequest{IDs: ids}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
