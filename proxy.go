package duoplus

import "context"

const (
	pathProxyList    = "/api/v1/proxy/list"
	pathProxyAdd     = "/api/v1/proxy/add"
	pathProxyDelete  = "/api/v1/proxy/delete"
	pathProxyRefresh = "/api/v1/proxy/refresh"
	pathProxyUpdate  = "/api/v1/proxy/update"
)

type ProxyListRequest struct {
	PaginationRequest
}

type Proxy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Area string `json:"area"`
}

type ProxyListResponse struct {
	List []Proxy `json:"list"`
	Pagination
}

type AddProxyRequest struct {
	ProxyList     []ProxyInput `json:"proxy_list"`
	IPScanChannel string       `json:"ip_scan_channel,omitempty"`
}

type ProxyInput struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

type AddProxySuccess struct {
	Index int    `json:"index"`
	ID    string `json:"id"`
}

type AddProxyFailure struct {
	Index   int    `json:"index"`
	Message string `json:"message"`
}

type AddProxyResponse struct {
	Success []AddProxySuccess `json:"success"`
	Fail    []AddProxyFailure `json:"fail"`
}

type ProxyIDsRequest struct {
	IDs []string `json:"ids"`
}

type ProxyOperationResponse struct {
	Success []string `json:"success"`
	Fail    []string `json:"fail"`
}

type UpdateProxyRequest struct {
	ID            string `json:"id"`
	Host          string `json:"host,omitempty"`
	Port          int    `json:"port,omitempty"`
	User          string `json:"user,omitempty"`
	Password      string `json:"password,omitempty"`
	Name          string `json:"name,omitempty"`
	IPScanChannel string `json:"ip_scan_channel,omitempty"`
	RefreshURL    string `json:"refresh_url,omitempty"`
	ProxyURL      string `json:"proxy_url,omitempty"`
}

type UpdateProxyResponse struct {
	Message string `json:"message"`
	Result  []any  `json:"result"`
}

func (c *Client) ListProxies(ctx context.Context, req ProxyListRequest) (*ProxyListResponse, error) {
	var resp ProxyListResponse
	if err := c.do(ctx, pathProxyList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) AddProxies(ctx context.Context, req AddProxyRequest) (*AddProxyResponse, error) {
	var resp AddProxyResponse
	if err := c.do(ctx, pathProxyAdd, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteProxies(ctx context.Context, ids []string) (*ProxyOperationResponse, error) {
	var resp ProxyOperationResponse
	if err := c.do(ctx, pathProxyDelete, ProxyIDsRequest{IDs: ids}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RefreshProxyURLs(ctx context.Context, ids []string) (*ProxyOperationResponse, error) {
	var resp ProxyOperationResponse
	if err := c.do(ctx, pathProxyRefresh, ProxyIDsRequest{IDs: ids}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateProxy(ctx context.Context, req UpdateProxyRequest) (*UpdateProxyResponse, error) {
	var resp UpdateProxyResponse
	if err := c.do(ctx, pathProxyUpdate, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
