package cloudnumber

import (
	"context"

	"github.com/difyz9/duoplus-go-sdk/common"
	"github.com/difyz9/duoplus-go-sdk/internal/clientcore"
)

const (
	pathList     = "/api/v1/cloudNumber/numberList"
	pathSMSList  = "/api/v1/cloudNumber/smsList"
	pathWriteSMS = "/api/v1/cloudNumber/imageWriteSms"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client {
	return &Client{core: core}
}

type ListRequest struct {
	PhoneNumber   string   `json:"phone_number,omitempty"`
	Status        []int    `json:"status,omitempty"`
	TypeIDs       []int    `json:"type_ids,omitempty"`
	RegionIDs     []string `json:"region_ids,omitempty"`
	RenewalStatus []int    `json:"renewal_status,omitempty"`
	Remark        string   `json:"remark,omitempty"`
	SortBy        string   `json:"sort_by,omitempty"`
	Order         string   `json:"order,omitempty"`
	common.PaginationRequest
}

type Number struct {
	ID            string `json:"id"`
	PhoneNumber   string `json:"phone_number"`
	RegionName    string `json:"region_name"`
	TypeName      string `json:"type_name"`
	StatusName    string `json:"status_name"`
	RenewalStatus int    `json:"renewal_status"`
	Remark        string `json:"remark"`
	CreatedAt     string `json:"created_at"`
	ExpiredAt     string `json:"expired_at"`
}

type ListResponse struct {
	List []Number `json:"list"`
	common.Pagination
}

type SMSListRequest struct {
	NumberID string `json:"number_id"`
	common.PaginationRequest
}

type SMS struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	ReceivedAt string `json:"received_at"`
}

type SMSListResponse struct {
	List []SMS `json:"list"`
	common.Pagination
}

type WriteSMSRequest struct {
	ImageID string    `json:"image_id"`
	SMS     []SMSItem `json:"sms"`
}

type SMSItem struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, pathList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SMSList(ctx context.Context, req SMSListRequest) (*SMSListResponse, error) {
	var resp SMSListResponse
	if err := c.core.Do(ctx, pathSMSList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) WriteSMS(ctx context.Context, req WriteSMSRequest) (*common.MessageResponse, error) {
	var resp common.MessageResponse
	if err := c.core.Do(ctx, pathWriteSMS, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
