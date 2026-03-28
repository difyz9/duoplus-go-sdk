package duoplus

import "context"

const (
	pathAppList          = "/api/v1/app/list"
	pathAppTeamList      = "/api/v1/app/teamList"
	pathAppInstall       = "/api/v1/app/install"
	pathAppInstalledList = "/api/v1/app/installedList"
	pathAppUninstall     = "/api/v1/app/uninstall"
	pathAppStart         = "/api/v1/app/start"
	pathAppStop          = "/api/v1/app/stop"
)

type AppListRequest struct {
	PaginationRequest
}

type AppVersion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type App struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Pkg         string       `json:"pkg"`
	VersionList []AppVersion `json:"version_list"`
}

type AppListResponse struct {
	List []App `json:"list"`
	Pagination
}

type InstallAppRequest struct {
	ImageIDs     []string `json:"image_ids"`
	AppID        string   `json:"app_id"`
	AppVersionID string   `json:"app_version_id,omitempty"`
}

type InstalledAppsRequest struct {
	ImageID string `json:"image_id"`
}

type InstalledAppsResponse struct {
	List []string `json:"list"`
}

type AppPackageOperationRequest struct {
	ImageIDs []string `json:"image_ids"`
	Pkg      string   `json:"pkg"`
}

func (c *Client) ListPlatformApps(ctx context.Context, req AppListRequest) (*AppListResponse, error) {
	var resp AppListResponse
	if err := c.do(ctx, pathAppList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListTeamApps(ctx context.Context, req AppListRequest) (*AppListResponse, error) {
	var resp AppListResponse
	if err := c.do(ctx, pathAppTeamList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) InstallApp(ctx context.Context, req InstallAppRequest) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathAppInstall, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListInstalledApps(ctx context.Context, imageID string) (*InstalledAppsResponse, error) {
	var resp InstalledAppsResponse
	if err := c.do(ctx, pathAppInstalledList, InstalledAppsRequest{ImageID: imageID}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UninstallApp(ctx context.Context, imageIDs []string, pkg string) (*MessageResponse, error) {
	return c.appPackageOperation(ctx, pathAppUninstall, imageIDs, pkg)
}

func (c *Client) StartApp(ctx context.Context, imageIDs []string, pkg string) (*MessageResponse, error) {
	return c.appPackageOperation(ctx, pathAppStart, imageIDs, pkg)
}

func (c *Client) StopApp(ctx context.Context, imageIDs []string, pkg string) (*MessageResponse, error) {
	return c.appPackageOperation(ctx, pathAppStop, imageIDs, pkg)
}

func (c *Client) appPackageOperation(ctx context.Context, path string, imageIDs []string, pkg string) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, path, AppPackageOperationRequest{ImageIDs: imageIDs, Pkg: pkg}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
