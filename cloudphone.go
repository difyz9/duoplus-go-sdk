package duoplus

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	pathCloudPhoneList            = "/api/v1/cloudPhone/list"
	pathCloudPhonePowerOn         = "/api/v1/cloudPhone/powerOn"
	pathCloudPhonePowerOff        = "/api/v1/cloudPhone/powerOff"
	pathCloudPhoneRestart         = "/api/v1/cloudPhone/restart"
	pathCloudPhoneStatus          = "/api/v1/cloudPhone/status"
	pathCloudPhoneInfo            = "/api/v1/cloudPhone/info"
	pathCloudPhoneUpdate          = "/api/v1/cloudPhone/update"
	pathCloudPhoneModelList       = "/api/v1/mobile/modelList"
	pathCloudPhoneBatchRoot       = "/api/v1/cloudPhone/batchRoot"
	pathCloudPhoneCommand         = "/api/v1/cloudPhone/command"
	pathCloudPhoneOpenADB         = "/api/v1/cloudPhone/openAdb"
	pathCloudPhoneCloseADB        = "/api/v1/cloudPhone/closeAdb"
	pathCloudPhoneSetADBWhitelist = "/api/v1/cloudPhone/setAdbIpWhitelist"
	pathCloudPhoneUpdateSharePass = "/api/v1/cloudPhone/updateSharePassword"
	pathCloudPhoneLinkUserList    = "/api/v1/cloudPhone/linkUserList"
	pathCloudPhoneTagList         = "/api/v1/cloudPhone/tagList"
	pathCloudPhoneResourceList    = "/api/v1/cloudPhone/cloudPhone"
	pathCloudPhoneResolutionList  = "/api/v1/cloudPhone/resolutionList"
	pathCloudPhonePurchase        = "/api/v1/cloudPhone/purchase"
	pathCloudPhoneRenewal         = "/api/v1/cloudPhone/renewal"
	pathCloudPhoneShare           = "/api/v1/cloudPhone/share"
	pathCloudNumberWriteSMS       = "/api/v1/cloudNumber/imageWriteSms"
	pathCloudNumberList           = "/api/v1/cloudNumber/numberList"
)

type PaginationRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pagesize,omitempty"`
}

type Pagination struct {
	Page      int `json:"page,omitempty"`
	PageSize  int `json:"pagesize,omitempty"`
	Total     int `json:"total,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
}

type OperationResult struct {
	Success    []string          `json:"success,omitempty"`
	Fail       []string          `json:"fail,omitempty"`
	FailReason map[string]string `json:"fail_reason,omitempty"`
	Message    string            `json:"message,omitempty"`
}

type CloudPhoneListRequest struct {
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
	PaginationRequest
}

type CloudPhone struct {
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

type CloudPhoneListResponse struct {
	List []CloudPhone `json:"list"`
	Pagination
}

type CloudPhoneStatusRequest struct {
	ImageIDs []string `json:"image_ids"`
}

type CloudPhoneStatusItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type CloudPhoneStatusResponse struct {
	List []CloudPhoneStatusItem `json:"list"`
}

type CloudPhoneInfoRequest struct {
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
	Manufacturer string `json:"manufacturer,omitempty"`
	Brand        string `json:"brand,omitempty"`
	Model        string `json:"model,omitempty"`
	IMEI         Text   `json:"imei,omitempty"`
	SerialNo     string `json:"serialno,omitempty"`
	AndroidID    string `json:"android_id,omitempty"`
	Name         string `json:"name,omitempty"`
	GSFID        string `json:"gsf_id,omitempty"`
	GAID         string `json:"gaid,omitempty"`
}

type CloudPhoneDetail struct {
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

type UpdateCloudPhonesRequest struct {
	Images []UpdateCloudPhoneItem `json:"images"`
}

type UpdateCloudPhoneItem struct {
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

type ADBCommandRequest struct {
	ImageIDs []string `json:"image_ids,omitempty"`
	ImageID  string   `json:"image_id,omitempty"`
	Command  string   `json:"command"`
}

type ADBCommandResult struct {
	Success bool   `json:"success"`
	Content string `json:"content"`
	Message string `json:"message"`
}

type SetADBIPWhitelistRequest struct {
	IPs []string `json:"ips"`
}

type MessageResponse struct {
	Message string `json:"message"`
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

type LinkUserListResponse struct {
	List []LinkUser `json:"list"`
}

type TagListRequest struct {
	Name string `json:"name,omitempty"`
	PaginationRequest
}

type Tag struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	ImageCount int    `json:"image_count"`
}

type TagListResponse struct {
	List []Tag `json:"list"`
	Pagination
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
	Pagination
}

type ResolutionListResponse struct {
	List []string `json:"list"`
}

type PurchaseCloudPhoneRequest struct {
	OS            string `json:"os"`
	Quantity      int    `json:"quantity"`
	Duration      string `json:"duration,omitempty"`
	CouponCode    string `json:"coupon_code,omitempty"`
	RenewalStatus int    `json:"renewal_status,omitempty"`
}

type RenewCloudPhoneRequest struct {
	ImageIDs   []string `json:"image_ids"`
	Duration   string   `json:"duration,omitempty"`
	CouponCode string   `json:"coupon_code,omitempty"`
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
}

type WriteSMSRequest struct {
	ImageID string    `json:"image_id"`
	SMS     []SMSItem `json:"sms"`
}

type SMSItem struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

type ShareCloudPhonesRequest struct {
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

type ShareCloudPhonesResponse map[string]string

type CloudNumberListRequest struct {
	PhoneNumber   string   `json:"phone_number,omitempty"`
	Status        []int    `json:"status,omitempty"`
	TypeIDs       []int    `json:"type_ids,omitempty"`
	RegionIDs     []string `json:"region_ids,omitempty"`
	RenewalStatus []int    `json:"renewal_status,omitempty"`
	Remark        string   `json:"remark,omitempty"`
	SortBy        string   `json:"sort_by,omitempty"`
	Order         string   `json:"order,omitempty"`
	PaginationRequest
}

type CloudNumber struct {
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

type CloudNumberListResponse struct {
	List []CloudNumber `json:"list"`
	Pagination
}

type CloudNumberSMSListRequest struct {
	NumberID string `json:"number_id"`
	PaginationRequest
}

type CloudNumberSMS struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	ReceivedAt string `json:"received_at"`
}

type CloudNumberSMSListResponse struct {
	List []CloudNumberSMS `json:"list"`
	Pagination
}

type Text string

func (t *Text) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*t = ""
		return nil
	}

	var value any
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	switch v := value.(type) {
	case string:
		*t = Text(v)
	case float64:
		*t = Text(strconv.FormatFloat(v, 'f', -1, 64))
	case bool:
		*t = Text(strconv.FormatBool(v))
	default:
		return fmt.Errorf("unsupported text value type %T", value)
	}

	return nil
}

func (c *Client) ListCloudPhones(ctx context.Context, req CloudPhoneListRequest) (*CloudPhoneListResponse, error) {
	var resp CloudPhoneListResponse
	if err := c.do(ctx, pathCloudPhoneList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PowerOnCloudPhones(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return c.cloudPhoneOperation(ctx, pathCloudPhonePowerOn, imageIDs)
}

func (c *Client) PowerOffCloudPhones(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return c.cloudPhoneOperation(ctx, pathCloudPhonePowerOff, imageIDs)
}

func (c *Client) RestartCloudPhones(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return c.cloudPhoneOperation(ctx, pathCloudPhoneRestart, imageIDs)
}

func (c *Client) GetCloudPhoneStatus(ctx context.Context, imageIDs []string) (*CloudPhoneStatusResponse, error) {
	var resp CloudPhoneStatusResponse
	if err := c.do(ctx, pathCloudPhoneStatus, CloudPhoneStatusRequest{ImageIDs: imageIDs}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetCloudPhoneInfo(ctx context.Context, imageID string) (*CloudPhoneDetail, error) {
	var resp CloudPhoneDetail
	if err := c.do(ctx, pathCloudPhoneInfo, CloudPhoneInfoRequest{ImageID: imageID}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateCloudPhones(ctx context.Context, req UpdateCloudPhonesRequest) (*OperationResult, error) {
	var resp OperationResult
	if err := c.do(ctx, pathCloudPhoneUpdate, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListPhoneModels(ctx context.Context, req ModelListRequest) (ModelListResponse, error) {
	resp := ModelListResponse{}
	if err := c.do(ctx, pathCloudPhoneModelList, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) BatchSetRoot(ctx context.Context, req BatchRootRequest) (*OperationResult, error) {
	var resp OperationResult
	if err := c.do(ctx, pathCloudPhoneBatchRoot, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ExecuteADBCommand(ctx context.Context, imageID, command string) (*ADBCommandResult, error) {
	var resp ADBCommandResult
	if err := c.do(ctx, pathCloudPhoneCommand, ADBCommandRequest{ImageID: imageID, Command: command}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ExecuteADBCommandBatch(ctx context.Context, imageIDs []string, command string) (map[string]ADBCommandResult, error) {
	resp := map[string]ADBCommandResult{}
	if err := c.do(ctx, pathCloudPhoneCommand, ADBCommandRequest{ImageIDs: imageIDs, Command: command}, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) OpenADB(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return c.cloudPhoneOperation(ctx, pathCloudPhoneOpenADB, imageIDs)
}

func (c *Client) CloseADB(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return c.cloudPhoneOperation(ctx, pathCloudPhoneCloseADB, imageIDs)
}

func (c *Client) SetADBIPWhitelist(ctx context.Context, ips []string) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathCloudPhoneSetADBWhitelist, SetADBIPWhitelistRequest{IPs: ips}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateSharePassword(ctx context.Context, req UpdateSharePasswordRequest) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathCloudPhoneUpdateSharePass, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListLinkUsers(ctx context.Context) (*LinkUserListResponse, error) {
	var resp LinkUserListResponse
	if err := c.do(ctx, pathCloudPhoneLinkUserList, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListTags(ctx context.Context, req TagListRequest) (*TagListResponse, error) {
	var resp TagListResponse
	if err := c.do(ctx, pathCloudPhoneTagList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListCloudPhoneResources(ctx context.Context) (*ResourceListResponse, error) {
	var resp ResourceListResponse
	if err := c.do(ctx, pathCloudPhoneResourceList, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListResolutions(ctx context.Context) (*ResolutionListResponse, error) {
	var resp ResolutionListResponse
	if err := c.do(ctx, pathCloudPhoneResolutionList, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PurchaseCloudPhones(ctx context.Context, req PurchaseCloudPhoneRequest) (*OrderResponse, error) {
	var resp OrderResponse
	if err := c.do(ctx, pathCloudPhonePurchase, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RenewCloudPhones(ctx context.Context, req RenewCloudPhoneRequest) (*OrderResponse, error) {
	var resp OrderResponse
	if err := c.do(ctx, pathCloudPhoneRenewal, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) WriteSMS(ctx context.Context, req WriteSMSRequest) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathCloudNumberWriteSMS, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ShareCloudPhones(ctx context.Context, req ShareCloudPhonesRequest) (ShareCloudPhonesResponse, error) {
	resp := ShareCloudPhonesResponse{}
	if err := c.do(ctx, pathCloudPhoneShare, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ListCloudNumbers(ctx context.Context, req CloudNumberListRequest) (*CloudNumberListResponse, error) {
	var resp CloudNumberListResponse
	if err := c.do(ctx, pathCloudNumberList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListCloudNumberSMS(ctx context.Context, req CloudNumberSMSListRequest) (*CloudNumberSMSListResponse, error) {
	var resp CloudNumberSMSListResponse
	if err := c.do(ctx, "/api/v1/cloudNumber/smsList", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) cloudPhoneOperation(ctx context.Context, path string, imageIDs []string) (*OperationResult, error) {
	var resp OperationResult
	if err := c.do(ctx, path, CloudPhoneStatusRequest{ImageIDs: imageIDs}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
