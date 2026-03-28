package duoplus

import "context"

type CloudPhoneService struct {
	client *Client
}

type CloudNumberService struct {
	client *Client
}

type GroupService struct {
	client *Client
}

type ProxyService struct {
	client *Client
}

type SubscriptionStartupService struct {
	client *Client
}

type AppService struct {
	client *Client
}

type CloudDiskService struct {
	client *Client
}

type AutomationService struct {
	client *Client
}

func (s *CloudPhoneService) List(ctx context.Context, req CloudPhoneListRequest) (*CloudPhoneListResponse, error) {
	return s.client.ListCloudPhones(ctx, req)
}

func (s *CloudPhoneService) PowerOn(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return s.client.PowerOnCloudPhones(ctx, imageIDs)
}

func (s *CloudPhoneService) PowerOff(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return s.client.PowerOffCloudPhones(ctx, imageIDs)
}

func (s *CloudPhoneService) Restart(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return s.client.RestartCloudPhones(ctx, imageIDs)
}

func (s *CloudPhoneService) Status(ctx context.Context, imageIDs []string) (*CloudPhoneStatusResponse, error) {
	return s.client.GetCloudPhoneStatus(ctx, imageIDs)
}

func (s *CloudPhoneService) Info(ctx context.Context, imageID string) (*CloudPhoneDetail, error) {
	return s.client.GetCloudPhoneInfo(ctx, imageID)
}

func (s *CloudPhoneService) Update(ctx context.Context, req UpdateCloudPhonesRequest) (*OperationResult, error) {
	return s.client.UpdateCloudPhones(ctx, req)
}

func (s *CloudPhoneService) Models(ctx context.Context, req ModelListRequest) (ModelListResponse, error) {
	return s.client.ListPhoneModels(ctx, req)
}

func (s *CloudPhoneService) BatchRoot(ctx context.Context, req BatchRootRequest) (*OperationResult, error) {
	return s.client.BatchSetRoot(ctx, req)
}

func (s *CloudPhoneService) Command(ctx context.Context, imageID, command string) (*ADBCommandResult, error) {
	return s.client.ExecuteADBCommand(ctx, imageID, command)
}

func (s *CloudPhoneService) CommandBatch(ctx context.Context, imageIDs []string, command string) (map[string]ADBCommandResult, error) {
	return s.client.ExecuteADBCommandBatch(ctx, imageIDs, command)
}

func (s *CloudPhoneService) OpenADB(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return s.client.OpenADB(ctx, imageIDs)
}

func (s *CloudPhoneService) CloseADB(ctx context.Context, imageIDs []string) (*OperationResult, error) {
	return s.client.CloseADB(ctx, imageIDs)
}

func (s *CloudPhoneService) SetADBIPWhitelist(ctx context.Context, ips []string) (*MessageResponse, error) {
	return s.client.SetADBIPWhitelist(ctx, ips)
}

func (s *CloudPhoneService) UpdateSharePassword(ctx context.Context, req UpdateSharePasswordRequest) (*MessageResponse, error) {
	return s.client.UpdateSharePassword(ctx, req)
}

func (s *CloudPhoneService) LinkUsers(ctx context.Context) (*LinkUserListResponse, error) {
	return s.client.ListLinkUsers(ctx)
}

func (s *CloudPhoneService) Tags(ctx context.Context, req TagListRequest) (*TagListResponse, error) {
	return s.client.ListTags(ctx, req)
}

func (s *CloudPhoneService) Resources(ctx context.Context) (*ResourceListResponse, error) {
	return s.client.ListCloudPhoneResources(ctx)
}

func (s *CloudPhoneService) Resolutions(ctx context.Context) (*ResolutionListResponse, error) {
	return s.client.ListResolutions(ctx)
}

func (s *CloudPhoneService) Purchase(ctx context.Context, req PurchaseCloudPhoneRequest) (*OrderResponse, error) {
	return s.client.PurchaseCloudPhones(ctx, req)
}

func (s *CloudPhoneService) Renew(ctx context.Context, req RenewCloudPhoneRequest) (*OrderResponse, error) {
	return s.client.RenewCloudPhones(ctx, req)
}

func (s *CloudPhoneService) Share(ctx context.Context, req ShareCloudPhonesRequest) (ShareCloudPhonesResponse, error) {
	return s.client.ShareCloudPhones(ctx, req)
}

func (s *CloudNumberService) List(ctx context.Context, req CloudNumberListRequest) (*CloudNumberListResponse, error) {
	return s.client.ListCloudNumbers(ctx, req)
}

func (s *CloudNumberService) SMSList(ctx context.Context, req CloudNumberSMSListRequest) (*CloudNumberSMSListResponse, error) {
	return s.client.ListCloudNumberSMS(ctx, req)
}

func (s *CloudNumberService) WriteSMS(ctx context.Context, req WriteSMSRequest) (*MessageResponse, error) {
	return s.client.WriteSMS(ctx, req)
}

func (s *GroupService) List(ctx context.Context, req GroupListRequest) (*GroupListResponse, error) {
	return s.client.ListCloudPhoneGroups(ctx, req)
}

func (s *GroupService) AddPhones(ctx context.Context, groupID string, imageIDs []string) (*MessageResponse, error) {
	return s.client.AddCloudPhonesToGroup(ctx, groupID, imageIDs)
}

func (s *GroupService) MovePhones(ctx context.Context, groupID string, imageIDs []string) (*MessageResponse, error) {
	return s.client.MoveCloudPhonesToGroup(ctx, groupID, imageIDs)
}

func (s *GroupService) Create(ctx context.Context, req CreateGroupsRequest) (*GroupMutationResponse, error) {
	return s.client.CreateCloudPhoneGroups(ctx, req)
}

func (s *GroupService) Update(ctx context.Context, req UpdateGroupsRequest) (*GroupMutationResponse, error) {
	return s.client.UpdateCloudPhoneGroups(ctx, req)
}

func (s *GroupService) Delete(ctx context.Context, ids []string) (*ProxyOperationResponse, error) {
	return s.client.DeleteCloudPhoneGroups(ctx, ids)
}

func (s *ProxyService) List(ctx context.Context, req ProxyListRequest) (*ProxyListResponse, error) {
	return s.client.ListProxies(ctx, req)
}

func (s *ProxyService) Add(ctx context.Context, req AddProxyRequest) (*AddProxyResponse, error) {
	return s.client.AddProxies(ctx, req)
}

func (s *ProxyService) Delete(ctx context.Context, ids []string) (*ProxyOperationResponse, error) {
	return s.client.DeleteProxies(ctx, ids)
}

func (s *ProxyService) RefreshURLs(ctx context.Context, ids []string) (*ProxyOperationResponse, error) {
	return s.client.RefreshProxyURLs(ctx, ids)
}

func (s *ProxyService) Update(ctx context.Context, req UpdateProxyRequest) (*UpdateProxyResponse, error) {
	return s.client.UpdateProxy(ctx, req)
}

func (s *SubscriptionStartupService) List(ctx context.Context, req SubscriptionStartupListRequest) (*SubscriptionStartupListResponse, error) {
	return s.client.ListSubscriptionStartups(ctx, req)
}

func (s *SubscriptionStartupService) Purchase(ctx context.Context, req PurchaseSubscriptionStartupRequest) (*OrderResponse, error) {
	return s.client.PurchaseSubscriptionStartups(ctx, req)
}

func (s *SubscriptionStartupService) Renew(ctx context.Context, req RenewSubscriptionStartupRequest) (*OrderResponse, error) {
	return s.client.RenewSubscriptionStartups(ctx, req)
}

func (s *AppService) ListPlatform(ctx context.Context, req AppListRequest) (*AppListResponse, error) {
	return s.client.ListPlatformApps(ctx, req)
}

func (s *AppService) ListTeam(ctx context.Context, req AppListRequest) (*AppListResponse, error) {
	return s.client.ListTeamApps(ctx, req)
}

func (s *AppService) Install(ctx context.Context, req InstallAppRequest) (*MessageResponse, error) {
	return s.client.InstallApp(ctx, req)
}

func (s *AppService) Installed(ctx context.Context, imageID string) (*InstalledAppsResponse, error) {
	return s.client.ListInstalledApps(ctx, imageID)
}

func (s *AppService) Uninstall(ctx context.Context, imageIDs []string, pkg string) (*MessageResponse, error) {
	return s.client.UninstallApp(ctx, imageIDs, pkg)
}

func (s *AppService) Start(ctx context.Context, imageIDs []string, pkg string) (*MessageResponse, error) {
	return s.client.StartApp(ctx, imageIDs, pkg)
}

func (s *AppService) Stop(ctx context.Context, imageIDs []string, pkg string) (*MessageResponse, error) {
	return s.client.StopApp(ctx, imageIDs, pkg)
}

func (s *CloudDiskService) List(ctx context.Context, req CloudDiskListRequest) (*CloudDiskListResponse, error) {
	return s.client.ListCloudDiskFiles(ctx, req)
}

func (s *CloudDiskService) PushFiles(ctx context.Context, req PushFilesRequest) (*PushFilesResponse, error) {
	return s.client.PushCloudDiskFiles(ctx, req)
}

func (s *AutomationService) ListUserTemplates(ctx context.Context, req TemplateListRequest) (*TemplateListResponse, error) {
	return s.client.ListUserTemplates(ctx, req)
}

func (s *AutomationService) ListOfficialTemplates(ctx context.Context, req TemplateListRequest) (*TemplateListResponse, error) {
	return s.client.ListOfficialTemplates(ctx, req)
}

func (s *AutomationService) CreateTask(ctx context.Context, req CreateAutomationTaskRequest) (*MessageResponse, error) {
	return s.client.CreateAutomationTask(ctx, req)
}

func (s *AutomationService) CreatePlan(ctx context.Context, req CreateAutomationTaskRequest) (*IDResponse, error) {
	return s.client.CreateAutomationPlan(ctx, req)
}

func (s *AutomationService) ListPlans(ctx context.Context, req PlanListRequest) (*PlanListResponse, error) {
	return s.client.ListAutomationPlans(ctx, req)
}

func (s *AutomationService) SavePlan(ctx context.Context, req SavePlanRequest) (*IDResponse, error) {
	return s.client.SaveAutomationPlan(ctx, req)
}

func (s *AutomationService) SetPlanStatus(ctx context.Context, req SetPlanStatusRequest) (*IDResponse, error) {
	return s.client.SetAutomationPlanStatus(ctx, req)
}

func (s *AutomationService) DeletePlan(ctx context.Context, id string) (*MessageResponse, error) {
	return s.client.DeleteAutomationPlan(ctx, id)
}

func (s *AutomationService) ListTasks(ctx context.Context, req TaskListRequest) (*TaskListResponse, error) {
	return s.client.ListAutomationTasks(ctx, req)
}

func (s *AutomationService) ListTaskLogs(ctx context.Context, req TaskLogListRequest) (*TaskLogListResponse, error) {
	return s.client.ListAutomationTaskLogs(ctx, req)
}

func (s *AutomationService) UpdateTaskTime(ctx context.Context, req UpdateTaskTimeRequest) (*MessageResponse, error) {
	return s.client.UpdateAutomationTaskTime(ctx, req)
}

func (s *AutomationService) SetTaskStatus(ctx context.Context, req SetTaskStatusRequest) (*OperationResult, error) {
	return s.client.SetAutomationTaskStatus(ctx, req)
}