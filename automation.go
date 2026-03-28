package duoplus

import "context"

const (
	pathAutomationUserTemplateList     = "/api/v1/automation/userTemplateList"
	pathAutomationOfficialTemplateList = "/api/v1/automation/officialTemplateList"
	pathAutomationAddTask              = "/api/v1/automation/addTask"
	pathAutomationAddPlan              = "/api/v1/automation/addPlan"
	pathAutomationPlanList             = "/api/v1/automation/planList"
	pathAutomationSavePlan             = "/api/v1/automation/savePlan"
	pathAutomationSetPlanStatus        = "/api/v1/automation/setPlanStatus"
	pathAutomationDeletePlan           = "/api/v1/automation/deletePlan"
	pathAutomationTaskList             = "/api/v1/automation/taskList"
	pathAutomationTaskLogList          = "/api/v1/automation/taskLogList"
	pathAutomationUpdateTaskTime       = "/api/v1/automation/updateTaskTime"
	pathAutomationSetTaskStatus        = "/api/v1/automation/setTaskStatus"
)

type TemplateListRequest struct {
	Name string `json:"name,omitempty"`
	PaginationRequest
}

type AutomationTemplate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type TemplateListResponse struct {
	List []AutomationTemplate `json:"list"`
	Pagination
}

type AutomationConfigValue struct {
	Key      string `json:"key"`
	Value    any    `json:"value"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

type TaskImage struct {
	ImageID        string                           `json:"image_id"`
	Config         map[string]AutomationConfigValue `json:"config,omitempty"`
	IssueAt        string                           `json:"issue_at,omitempty"`
	StartAt        string                           `json:"start_at,omitempty"`
	EndAt          string                           `json:"end_at,omitempty"`
	ExecuteType    int                              `json:"execute_type,omitempty"`
	GapTime        int                              `json:"gap_time,omitempty"`
	ExecuteTime    string                           `json:"execute_time,omitempty"`
	ExecuteEndTime string                           `json:"execute_end_time,omitempty"`
	Mode           int                              `json:"mode,omitempty"`
	Weeks          []string                         `json:"weeks,omitempty"`
	Days           []string                         `json:"days,omitempty"`
}

type CreateAutomationTaskRequest struct {
	TemplateID   string      `json:"template_id"`
	TemplateType int         `json:"template_type"`
	Name         string      `json:"name"`
	Remark       string      `json:"remark,omitempty"`
	Images       []TaskImage `json:"images"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type PlanListRequest struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Status       []int  `json:"status,omitempty"`
	TemplateType []int  `json:"template_type,omitempty"`
	Remark       string `json:"remark,omitempty"`
	PaginationRequest
}

type AutomationPlan struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Remark       string `json:"remark"`
	TaskTypeName string `json:"task_type_name"`
	Status       int    `json:"status"`
	CreatedAt    string `json:"created_at"`
}

type PlanListResponse struct {
	List []AutomationPlan `json:"list"`
	Pagination
}

type SavePlanRequest struct {
	ID     string      `json:"id"`
	Name   string      `json:"name,omitempty"`
	Remark string      `json:"remark,omitempty"`
	Images []TaskImage `json:"images"`
}

type SetPlanStatusRequest struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

type DeletePlanRequest struct {
	ID string `json:"id"`
}

type TaskListRequest struct {
	ID               string `json:"id,omitempty"`
	IssueAtStart     string `json:"issue_at_start"`
	IssueAtEnd       string `json:"issue_at_end"`
	Status           []int  `json:"status,omitempty"`
	TemplateType     []int  `json:"template_type,omitempty"`
	Name             string `json:"name,omitempty"`
	ImageName        string `json:"image_name,omitempty"`
	ImageIP          string `json:"image_ip,omitempty"`
	ExecutionAtStart string `json:"execution_at_start,omitempty"`
	ExecutionAtEnd   string `json:"execution_at_end,omitempty"`
	SortBy           string `json:"sort_by,omitempty"`
	Order            string `json:"order,omitempty"`
	PaginationRequest
}

type AutomationTask struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	TaskTypeName  string `json:"task_type_name"`
	ImageName     string `json:"image_name"`
	IP            string `json:"ip"`
	Remark        string `json:"remark"`
	Status        int    `json:"status"`
	IssueAt       string `json:"issue_at"`
	StartAt       string `json:"start_at"`
	FinishAt      string `json:"finish_at"`
	CostTime      int    `json:"cost_time"`
	ExecutionTime string `json:"execution_time"`
	CreatedAt     string `json:"created_at"`
}

type TaskListResponse struct {
	List []AutomationTask `json:"list"`
	Pagination
}

type TaskLogListRequest struct {
	TaskID   string `json:"task_id"`
	CursorID string `json:"cursor_id,omitempty"`
}

type TaskLogResultInfo struct {
	Action                 string         `json:"action"`
	Result                 bool           `json:"result"`
	ErrorMessage           string         `json:"error_message"`
	ErrorImg               string         `json:"error_img"`
	ConditionResultMessage string         `json:"condition_result_message"`
	ExtraData              map[string]any `json:"extra_data"`
}

type TaskLog struct {
	ID         string            `json:"id"`
	ResultInfo TaskLogResultInfo `json:"result_info"`
	StartAt    string            `json:"start_at"`
	FinishAt   string            `json:"finish_at"`
	CreatedAt  string            `json:"created_at"`
}

type TaskLogListResponse struct {
	List []TaskLog `json:"list"`
}

type UpdateTaskTimeRequest struct {
	ID      string `json:"id"`
	IssueAt string `json:"issue_at"`
}

type SetTaskStatusRequest struct {
	IDs    []string `json:"ids"`
	Status int      `json:"status"`
}

func (c *Client) ListUserTemplates(ctx context.Context, req TemplateListRequest) (*TemplateListResponse, error) {
	var resp TemplateListResponse
	if err := c.do(ctx, pathAutomationUserTemplateList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListOfficialTemplates(ctx context.Context, req TemplateListRequest) (*TemplateListResponse, error) {
	var resp TemplateListResponse
	if err := c.do(ctx, pathAutomationOfficialTemplateList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateAutomationTask(ctx context.Context, req CreateAutomationTaskRequest) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathAutomationAddTask, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateAutomationPlan(ctx context.Context, req CreateAutomationTaskRequest) (*IDResponse, error) {
	var resp IDResponse
	if err := c.do(ctx, pathAutomationAddPlan, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListAutomationPlans(ctx context.Context, req PlanListRequest) (*PlanListResponse, error) {
	var resp PlanListResponse
	if err := c.do(ctx, pathAutomationPlanList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SaveAutomationPlan(ctx context.Context, req SavePlanRequest) (*IDResponse, error) {
	var resp IDResponse
	if err := c.do(ctx, pathAutomationSavePlan, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SetAutomationPlanStatus(ctx context.Context, req SetPlanStatusRequest) (*IDResponse, error) {
	var resp IDResponse
	if err := c.do(ctx, pathAutomationSetPlanStatus, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteAutomationPlan(ctx context.Context, id string) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathAutomationDeletePlan, DeletePlanRequest{ID: id}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListAutomationTasks(ctx context.Context, req TaskListRequest) (*TaskListResponse, error) {
	var resp TaskListResponse
	if err := c.do(ctx, pathAutomationTaskList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListAutomationTaskLogs(ctx context.Context, req TaskLogListRequest) (*TaskLogListResponse, error) {
	var resp TaskLogListResponse
	if err := c.do(ctx, pathAutomationTaskLogList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateAutomationTaskTime(ctx context.Context, req UpdateTaskTimeRequest) (*MessageResponse, error) {
	var resp MessageResponse
	if err := c.do(ctx, pathAutomationUpdateTaskTime, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SetAutomationTaskStatus(ctx context.Context, req SetTaskStatusRequest) (*OperationResult, error) {
	var resp OperationResult
	if err := c.do(ctx, pathAutomationSetTaskStatus, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
