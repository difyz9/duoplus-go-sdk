package subscriptionstartup

import (
	"context"

	"github.com/difyz9/duoplus-go-sdk/common"
	"github.com/difyz9/duoplus-go-sdk/internal/clientcore"
)

const (
	pathList     = "/api/v1/subscriptionStartup/list"
	pathPurchase = "/api/v1/subscriptionStartup/purchase"
	pathRenewal  = "/api/v1/subscriptionStartup/renewal"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client {
	return &Client{core: core}
}

type ListRequest struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Remark        string `json:"remark,omitempty"`
	RenewalStatus int    `json:"renewal_status,omitempty"`
	SortBy        string `json:"sort_by,omitempty"`
	Order         string `json:"order,omitempty"`
	common.PaginationRequest
}

type Item struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CPU           string `json:"cpu"`
	RAM           string `json:"ram"`
	ROM           string `json:"rom"`
	RenewalStatus int    `json:"renewal_status"`
	Remark        string `json:"remark"`
	ExpiredAt     string `json:"expired_at"`
	CreatedAt     string `json:"created_at"`
	NeedRenewal   bool   `json:"need_renewal"`
}

type ListResponse struct {
	List []Item `json:"list"`
	common.Pagination
}

type PurchaseRequest struct {
	Duration      string `json:"duration,omitempty"`
	Quantity      int    `json:"quantity"`
	CouponCode    string `json:"coupon_code,omitempty"`
	RenewalStatus int    `json:"renewal_status,omitempty"`
}

type RenewRequest struct {
	PhoneIDs   []string `json:"phone_ids"`
	Duration   string   `json:"duration,omitempty"`
	CouponCode string   `json:"coupon_code,omitempty"`
}

func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, pathList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Purchase(ctx context.Context, req PurchaseRequest) (*common.OrderResponse, error) {
	var resp common.OrderResponse
	if err := c.core.Do(ctx, pathPurchase, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Renew(ctx context.Context, req RenewRequest) (*common.OrderResponse, error) {
	var resp common.OrderResponse
	if err := c.core.Do(ctx, pathRenewal, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
