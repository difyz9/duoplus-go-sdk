package cloudphone

import (
	"context"

	"duoplus-go-sdk/common"
	"duoplus-go-sdk/internal/clientcore"
)

const (
	pathList            = "/api/v1/cloudPhone/list"
	pathPowerOn         = "/api/v1/cloudPhone/powerOn"
	pathPowerOff        = "/api/v1/cloudPhone/powerOff"
	pathRestart         = "/api/v1/cloudPhone/restart"
	pathStatus          = "/api/v1/cloudPhone/status"
	pathInfo            = "/api/v1/cloudPhone/info"
	pathUpdate          = "/api/v1/cloudPhone/update"
	pathModelList       = "/api/v1/mobile/modelList"
	pathBatchRoot       = "/api/v1/cloudPhone/batchRoot"
	pathCommand         = "/api/v1/cloudPhone/command"
	pathOpenADB         = "/api/v1/cloudPhone/openAdb"
	pathCloseADB        = "/api/v1/cloudPhone/closeAdb"
	pathSetADBWhitelist = "/api/v1/cloudPhone/setAdbIpWhitelist"
	pathUpdateSharePass = "/api/v1/cloudPhone/updateSharePassword"
	pathLinkUserList    = "/api/v1/cloudPhone/linkUserList"
	pathTagList         = "/api/v1/cloudPhone/tagList"
	pathResourceList    = "/api/v1/cloudPhone/cloudPhone"
	pathResolutionList  = "/api/v1/cloudPhone/resolutionList"
	pathPurchase        = "/api/v1/cloudPhone/purchase"
	pathRenewal         = "/api/v1/cloudPhone/renewal"
	pathShare           = "/api/v1/cloudPhone/share"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client {
	return &Client{core: core}
}

type ListRequest struct {
	ImageIDs       []string `json:"image_id,omitempty"`
	Name           string   `json:"name,omitempty"`
	GroupID        string   `json:"group_id,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	IPs            []string `json:"ips,omitempty"`
	LinkStatus     []string `json:"link_status,omitempty"`
	ProxyID        string   `json:"proxy_id,omitempty"`
	ShareStatus    []string `json:"share_status,omitempty"`
	StartPhoneType []string `json:"start_phone_type,omitempty"`
	ADBStatus      []string `json:"adb_status,omitempty"`
	RenewalStatus  []string `json:"renewal_status,omitempty"`
	SortBy         string   `json:"sort_by,omitempty"`
	Order          string   `json:"order,omitempty"`
	UserIDs        []string `json:"user_ids,omitempty"`
	TagIDs         []string `json:"tag_ids,omitempty"`
	RegionIDs      []string `json:"region_id,omitempty"`
	common.PaginationRequest
}

type Phone struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	ADB         string `json:"adb"`
	ADBPassword string `json:"adb_password"`
	Remark      string `json:"remark"`
	CreatedAt   string `json:"created_at"`
	ExpiredAt   string `json:"expired_at"`
	IP          string `json:"ip"`
	Area        string `json:"area"`
	OS          string `json:"os"`
	Size        string `json:"size"`
}

type ListResponse struct {
	List []Phone `json:"list"`
	common.Pagination
}

type StatusRequest struct {
	ImageIDs []string `json:"image_ids"`
}

type StatusItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type StatusResponse struct {
	List []StatusItem `json:"list"`
}

type InfoRequest struct {
	ImageID string `json:"image_id"`
}

type ProxySettings struct {
	ID      string `json:"id,omitempty"`
	DNS     int    `json:"dns,omitempty"`
	IP      string `json:"ip,omitempty"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
	City    string `json:"city,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
}

type GPSSettings struct {
	Type      int     `json:"type,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

type LocaleSettings struct {
	Type     int    `json:"type,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Language string `json:"language,omitempty"`
}

type SIMSettings struct {
	Status   int    `json:"status,omitempty"`
	Country  string `json:"country,omitempty"`
	MSISDN   string `json:"msisdn,omitempty"`
	Operator string `json:"operator,omitempty"`
	MCC      string `json:"mcc,omitempty"`
	MNC      string `json:"mnc,omitempty"`
	MSIN     string `json:"msin,omitempty"`
	ICCID    string `json:"iccid,omitempty"`
	IMSI     string `json:"imsi,omitempty"`
}

type BluetoothSettings struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

type WiFiSettings struct {
	Status int    `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
	MAC    string `json:"mac,omitempty"`
	BSSID  string `json:"bssid,omitempty"`
}

type DeviceSettings struct {
	Manufacturer string      `json:"manufacturer,omitempty"`
	Brand        string      `json:"brand,omitempty"`
	Model        string      `json:"model,omitempty"`
	IMEI         common.Text `json:"imei,omitempty"`
	SerialNo     string      `json:"serialno,omitempty"`
	AndroidID    string      `json:"android_id,omitempty"`
	Name         string      `json:"name,omitempty"`
	GSFID        string      `json:"gsf_id,omitempty"`
	GAID         string      `json:"gaid,omitempty"`
}

type Detail struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Remark    string            `json:"remark"`
	OS        string            `json:"os"`
	Proxy     ProxySettings     `json:"proxy"`
	GPS       GPSSettings       `json:"gps"`
	Locale    LocaleSettings    `json:"locale"`
	SIM       SIMSettings       `json:"sim"`
	Bluetooth BluetoothSettings `json:"bluetooth"`
	WiFi      WiFiSettings      `json:"wifi"`
	Device    DeviceSettings    `json:"device"`
}

type UpdateRequest struct {
	Images []UpdateItem `json:"images"`
}

type UpdateItem struct {
	ImageID   string             `json:"image_id"`
	Name      string             `json:"name,omitempty"`
	DPIName   string             `json:"dpi_name,omitempty"`
	Remark    string             `json:"remark,omitempty"`
	Proxy     *ProxyUpdate       `json:"proxy,omitempty"`
	GPS       *GPSUpdate         `json:"gps,omitempty"`
	Locale    *LocaleUpdate      `json:"locale,omitempty"`
	SIM       *SIMUpdate         `json:"sim,omitempty"`
	Bluetooth *BluetoothSettings `json:"bluetooth,omitempty"`
	WiFi      *WiFiSettings      `json:"wifi,omitempty"`
	Device    *DeviceUpdate      `json:"device,omitempty"`
}

type ProxyUpdate struct {
	ID  string `json:"id,omitempty"`
	DNS int    `json:"dns,omitempty"`
}

type GPSUpdate struct {
	Type      int     `json:"type"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

type LocaleUpdate struct {
	Type     int    `json:"type,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Language string `json:"language,omitempty"`
}

type SIMUpdate struct {
	Status   int    `json:"status,omitempty"`
	Country  string `json:"country,omitempty"`
	MSISDN   string `json:"msisdn,omitempty"`
	Operator string `json:"operator,omitempty"`
	MCC      string `json:"mcc,omitempty"`
	MNC      string `json:"mnc,omitempty"`
	MSIN     string `json:"msin,omitempty"`
	ICCID    string `json:"iccid,omitempty"`
	IMSI     string `json:"imsi,omitempty"`
}

type DeviceUpdate struct {
	IMEI      string `json:"imei,omitempty"`
	SerialNo  string `json:"serialno,omitempty"`
	AndroidID string `json:"android_id,omitempty"`
	Name      string `json:"name,omitempty"`
	GSFID     string `json:"gsf_id,omitempty"`
	GAID      string `json:"gaid,omitempty"`
}

type ModelListRequest struct {
	OS int `json:"os"`
}

type ModelInfo struct {
	Name string `json:"name"`
}

type ModelListResponse map[string]map[string]ModelInfo

type BatchRootRequest struct {
	ImageIDs []string `json:"image_ids"`
	Status   int      `json:"status"`
	Packages []string `json:"packages,omitempty"`
}

type CommandRequest struct {
	ImageIDs []string `json:"image_ids,omitempty"`
	ImageID  string   `json:"image_id,omitempty"`
	Command  string   `json:"command"`
}

type CommandResult struct {
	Success bool   `json:"success"`
	Content string `json:"content"`
	Message string `json:"message"`
}

type SetADBIPWhitelistRequest struct {
	IPs []string `json:"ips"`
}

type UpdateSharePasswordRequest struct {
	Images []SharePasswordItem `json:"images"`
}

type SharePasswordItem struct {
	ImageID  string `json:"image_id"`
	Password string `json:"password"`
}

type LinkUser struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
}

type LinkUsersResponse struct {
	List []LinkUser `json:"list"`
}

type TagListRequest struct {
	Name string `json:"name,omitempty"`
	common.PaginationRequest
}

type Tag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	ImageCount int    `json:"image_count"`
}

type TagListResponse struct {
	List []Tag `json:"list"`
	common.Pagination
}

type Resource struct {
	Name      string `json:"name"`
	RegionID  string `json:"region_id"`
	OS        string `json:"os"`
	Count     int    `json:"count"`
	UsedCount int    `json:"used_count"`
}

type ResourceListResponse struct {
	List []Resource `json:"list"`
	common.Pagination
}

type ResolutionListResponse struct {
	List []string `json:"list"`
}

type PurchaseRequest struct {
	OS            string `json:"os"`
	Quantity      int    `json:"quantity"`
	Duration      string `json:"duration,omitempty"`
	CouponCode    string `json:"coupon_code,omitempty"`
	RenewalStatus int    `json:"renewal_status,omitempty"`
}

type RenewRequest struct {
	ImageIDs   []string `json:"image_ids"`
	Duration   string   `json:"duration,omitempty"`
	CouponCode string   `json:"coupon_code,omitempty"`
}

type ShareRequest struct {
	Share []ShareItem `json:"share"`
}

type ShareItem struct {
	ImageIDs []string     `json:"image_ids"`
	Config   *ShareConfig `json:"config,omitempty"`
}

type ShareConfig struct {
	ShareStatus    int    `json:"share_status,omitempty"`
	SharePhoneType int    `json:"share_phone_type,omitempty"`
	ShareCode      string `json:"share_code,omitempty"`
	ShareAuth      []int  `json:"share_auth,omitempty"`
}

type ShareResponse map[string]string

func (c *Client) List(ctx context.Context, req ListRequest) (*ListResponse, error) {
	var resp ListResponse
	if err := c.core.Do(ctx, pathList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PowerOn(ctx context.Context, imageIDs []string) (*common.OperationResult, error) {
	return c.operation(ctx, pathPowerOn, imageIDs)
}

func (c *Client) PowerOff(ctx context.Context, imageIDs []string) (*common.OperationResult, error) {
	return c.operation(ctx, pathPowerOff, imageIDs)
}

func (c *Client) Restart(ctx context.Context, imageIDs []string) (*common.OperationResult, error) {
	return c.operation(ctx, pathRestart, imageIDs)
}

func (c *Client) Status(ctx context.Context, imageIDs []string) (*StatusResponse, error) {
	var resp StatusResponse
	if err := c.core.Do(ctx, pathStatus, StatusRequest{ImageIDs: imageIDs}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Info(ctx context.Context, imageID string) (*Detail, error) {
	var resp Detail
	if err := c.core.Do(ctx, pathInfo, InfoRequest{ImageID: imageID}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Update(ctx context.Context, req UpdateRequest) (*common.OperationResult, error) {
	var resp common.OperationResult
	if err := c.core.Do(ctx, pathUpdate, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Models(ctx context.Context, req ModelListRequest) (ModelListResponse, error) {
	resp := ModelListResponse{}
	if err := c.core.Do(ctx, pathModelList, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) BatchRoot(ctx context.Context, req BatchRootRequest) (*common.OperationResult, error) {
	var resp common.OperationResult
	if err := c.core.Do(ctx, pathBatchRoot, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Command(ctx context.Context, imageID, command string) (*CommandResult, error) {
	var resp CommandResult
	if err := c.core.Do(ctx, pathCommand, CommandRequest{ImageID: imageID, Command: command}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CommandBatch(ctx context.Context, imageIDs []string, command string) (map[string]CommandResult, error) {
	resp := map[string]CommandResult{}
	if err := c.core.Do(ctx, pathCommand, CommandRequest{ImageIDs: imageIDs, Command: command}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) OpenADB(ctx context.Context, imageIDs []string) (*common.OperationResult, error) {
	return c.operation(ctx, pathOpenADB, imageIDs)
}

func (c *Client) CloseADB(ctx context.Context, imageIDs []string) (*common.OperationResult, error) {
	return c.operation(ctx, pathCloseADB, imageIDs)
}

func (c *Client) SetADBIPWhitelist(ctx context.Context, ips []string) (*common.MessageResponse, error) {
	var resp common.MessageResponse
	if err := c.core.Do(ctx, pathSetADBWhitelist, SetADBIPWhitelistRequest{IPs: ips}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateSharePassword(ctx context.Context, req UpdateSharePasswordRequest) (*common.MessageResponse, error) {
	var resp common.MessageResponse
	if err := c.core.Do(ctx, pathUpdateSharePass, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) LinkUsers(ctx context.Context) (*LinkUsersResponse, error) {
	var resp LinkUsersResponse
	if err := c.core.Do(ctx, pathLinkUserList, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Tags(ctx context.Context, req TagListRequest) (*TagListResponse, error) {
	var resp TagListResponse
	if err := c.core.Do(ctx, pathTagList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Resources(ctx context.Context) (*ResourceListResponse, error) {
	var resp ResourceListResponse
	if err := c.core.Do(ctx, pathResourceList, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Resolutions(ctx context.Context) (*ResolutionListResponse, error) {
	var resp ResolutionListResponse
	if err := c.core.Do(ctx, pathResolutionList, nil, &resp); err != nil {
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

func (c *Client) Share(ctx context.Context, req ShareRequest) (ShareResponse, error) {
	resp := ShareResponse{}
	if err := c.core.Do(ctx, pathShare, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) operation(ctx context.Context, path string, imageIDs []string) (*common.OperationResult, error) {
	var resp common.OperationResult
	if err := c.core.Do(ctx, path, StatusRequest{ImageIDs: imageIDs}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
