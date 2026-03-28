package duoplus

import "context"

const (
	pathSubscriptionStartupList     = "/api/v1/subscriptionStartup/list"
	pathSubscriptionStartupPurchase = "/api/v1/subscriptionStartup/purchase"
	pathSubscriptionStartupRenewal  = "/api/v1/subscriptionStartup/renewal"
)

type SubscriptionStartupListRequest struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Remark        string `json:"remark,omitempty"`
	RenewalStatus int    `json:"renewal_status,omitempty"`
	SortBy        string `json:"sort_by,omitempty"`
	Order         string `json:"order,omitempty"`
	PaginationRequest
}

type SubscriptionStartup struct {
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

type SubscriptionStartupListResponse struct {
	List []SubscriptionStartup `json:"list"`
	Pagination
}

type PurchaseSubscriptionStartupRequest struct {
	Duration      string `json:"duration,omitempty"`
	Quantity      int    `json:"quantity"`
	CouponCode    string `json:"coupon_code,omitempty"`
	RenewalStatus int    `json:"renewal_status,omitempty"`
}

type RenewSubscriptionStartupRequest struct {
	PhoneIDs   []string `json:"phone_ids"`
	Duration   string   `json:"duration,omitempty"`
	CouponCode string   `json:"coupon_code,omitempty"`
}

func (c *Client) ListSubscriptionStartups(ctx context.Context, req SubscriptionStartupListRequest) (*SubscriptionStartupListResponse, error) {
	var resp SubscriptionStartupListResponse
	if err := c.do(ctx, pathSubscriptionStartupList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PurchaseSubscriptionStartups(ctx context.Context, req PurchaseSubscriptionStartupRequest) (*OrderResponse, error) {
	var resp OrderResponse
	if err := c.do(ctx, pathSubscriptionStartupPurchase, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RenewSubscriptionStartups(ctx context.Context, req RenewSubscriptionStartupRequest) (*OrderResponse, error) {
	var resp OrderResponse
	if err := c.do(ctx, pathSubscriptionStartupRenewal, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
