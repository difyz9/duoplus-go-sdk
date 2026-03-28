# DuoPlus Go SDK

基于 DuoPlus OpenAPI 的非官方 Golang SDK，当前优先封装了“云手机管理”相关接口，并补充了云号码列表与云机写入短信能力。

## 已封装接口

- 云手机列表 `POST /api/v1/cloudPhone/list`
- 批量开机 `POST /api/v1/cloudPhone/powerOn`
- 批量关机 `POST /api/v1/cloudPhone/powerOff`
- 批量重启 `POST /api/v1/cloudPhone/restart`
- 云手机状态 `POST /api/v1/cloudPhone/status`
- 获取云机详情 `POST /api/v1/cloudPhone/info`
- 批量修改云机参数 `POST /api/v1/cloudPhone/update`
- 手机型号列表 `POST /api/v1/mobile/modelList`
- 批量设置 Root 权限 `POST /api/v1/cloudPhone/batchRoot`
- 执行 ADB 命令 `POST /api/v1/cloudPhone/command`
- 批量开启 ADB `POST /api/v1/cloudPhone/openAdb`
- 批量关闭 ADB `POST /api/v1/cloudPhone/closeAdb`
- 设置 ADB 连接 IP 白名单 `POST /api/v1/cloudPhone/setAdbIpWhitelist`
- 修改分享密码 `POST /api/v1/cloudPhone/updateSharePassword`
- 云手机连接用户列表 `POST /api/v1/cloudPhone/linkUserList`
- 标签列表 `POST /api/v1/cloudPhone/tagList`
- 云手机资源列表 `POST /api/v1/cloudPhone/cloudPhone`
- 分辨率列表 `POST /api/v1/cloudPhone/resolutionList`
- 购买云手机 `POST /api/v1/cloudPhone/purchase`
- 续费云手机 `POST /api/v1/cloudPhone/renewal`
- 分享云手机 `POST /api/v1/cloudPhone/share`
- 云号码列表 `POST /api/v1/cloudNumber/numberList`
- 云号码短信查询 `POST /api/v1/cloudNumber/smsList`
- 云机写入短信 `POST /api/v1/cloudNumber/imageWriteSms`
- 代理列表 `POST /api/v1/proxy/list`
- 批量添加代理 `POST /api/v1/proxy/add`
- 批量删除代理 `POST /api/v1/proxy/delete`
- 批量刷新代理 URL `POST /api/v1/proxy/refresh`
- 修改代理 `POST /api/v1/proxy/update`
- 包月开机列表 `POST /api/v1/subscriptionStartup/list`
- 包月开机购买 `POST /api/v1/subscriptionStartup/purchase`
- 包月开机续费 `POST /api/v1/subscriptionStartup/renewal`
- 云手机分组列表 `POST /api/v1/cloudPhone/groupList`
- 批量添加到分组 `POST /api/v1/cloudPhone/addToGroup`
- 批量移动到分组 `POST /api/v1/cloudPhone/moveToGroup`
- 批量创建分组 `POST /api/v1/cloudPhone/createGroup`
- 批量编辑分组 `POST /api/v1/cloudPhone/updateGroup`
- 批量删除分组 `POST /api/v1/cloudPhone/deleteGroup`
- 平台应用列表 `POST /api/v1/app/list`
- 团队应用列表 `POST /api/v1/app/teamList`
- 批量安装应用 `POST /api/v1/app/install`
- 已安装应用列表 `POST /api/v1/app/installedList`
- 批量卸载应用 `POST /api/v1/app/uninstall`
- 批量启动应用 `POST /api/v1/app/start`
- 批量关闭应用 `POST /api/v1/app/stop`
- 文件列表 `POST /api/v1/cloudDisk/list`
- 文件推送 `POST /api/v1/cloudDisk/pushFiles`
- 自定义模板列表 `POST /api/v1/automation/userTemplateList`
- 官方模板列表 `POST /api/v1/automation/officialTemplateList`
- 创建定时任务 `POST /api/v1/automation/addTask`
- 创建循环任务 `POST /api/v1/automation/addPlan`
- 循环任务列表 `POST /api/v1/automation/planList`
- 修改循环任务 `POST /api/v1/automation/savePlan`
- 暂停/开始执行循环任务 `POST /api/v1/automation/setPlanStatus`
- 删除循环任务 `POST /api/v1/automation/deletePlan`
- 定时任务列表 `POST /api/v1/automation/taskList`
- 查看定时任务报告 `POST /api/v1/automation/taskLogList`
- 修改定时任务发布时间 `POST /api/v1/automation/updateTaskTime`
- 取消/重新执行定时任务 `POST /api/v1/automation/setTaskStatus`

## 安装

```bash
go get duoplus-go-sdk
```

如果你后续会把仓库发布到 Git 服务，请将 `go.mod` 中的模块名改成你的实际仓库地址。

## 快速开始

```go
package main

import (
	"context"
	"fmt"
	"log"

	duoplus "duoplus-go-sdk"
)

func main() {
	client, err := duoplus.NewClient(
		"your-api-key",
		duoplus.WithBaseURL(duoplus.DefaultBaseURL),
		duoplus.WithLanguage("zh"),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.CloudPhones.List(context.Background(), duoplus.CloudPhoneListRequest{
		PaginationRequest: duoplus.PaginationRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range resp.List {
		fmt.Println(item.ID, item.Name, item.Status)
	}
	}
```

## 模块化访问

为避免破坏现有导入路径，SDK 仍保持根包 `duoplus`，但 `Client` 现在提供按模块归类的 service 入口：

- `client.CloudPhones`
- `client.CloudNumbers`
- `client.Groups`
- `client.Proxies`
- `client.SubscriptionStartups`
- `client.Apps`
- `client.CloudDisk`
- `client.Automation`

旧的 `client.ListCloudPhones(...)` 这类方法仍然可用；新的 service 入口更适合在业务代码里按模块组织调用。

## 示例目录

- `examples/basic`: 基础云手机列表示例
- `examples/automation`: 自动化模板列表示例
- `examples/group-disk`: 云手机分组与云盘示例

## 使用示例

### 查询云手机状态

```go
status, err := client.CloudPhones.Status(ctx, []string{"image-id-1", "image-id-2"})
if err != nil {
	return err
}

for _, item := range status.List {
	fmt.Println(item.ID, item.Status)
}
```

### 批量开机

```go
result, err := client.CloudPhones.PowerOn(ctx, []string{"image-id-1", "image-id-2"})
if err != nil {
	return err
}

fmt.Println("success:", result.Success)
fmt.Println("fail:", result.Fail)
```

### 获取云机详情

```go
detail, err := client.CloudPhones.Info(ctx, "image-id")
if err != nil {
	return err
}

fmt.Println(detail.Device.Model, detail.Locale.Language)
```

### 执行单台 ADB 命令

```go
adbResult, err := client.CloudPhones.Command(ctx, "image-id", "pm list packages")
if err != nil {
	return err
}

fmt.Println(adbResult.Content)
```

### 批量修改云机参数

```go
updateResult, err := client.CloudPhones.Update(ctx, duoplus.UpdateCloudPhonesRequest{
	Images: []duoplus.UpdateCloudPhoneItem{
		{
			ImageID: "image-id",
			Name:    "device-01",
			Proxy: &duoplus.ProxyUpdate{
				ID:  "proxy-id",
				DNS: 1,
			},
			Locale: &duoplus.LocaleUpdate{
				Type:     2,
				Timezone: "Asia/Singapore",
				Language: "en-SG",
			},
		},
	},
})
if err != nil {
	return err
}

fmt.Println(updateResult.Success)
```

### 购买云手机

```go
order, err := client.CloudPhones.Purchase(ctx, duoplus.PurchaseCloudPhoneRequest{
	OS:            "12A",
	Quantity:      1,
	Duration:      "30",
	RenewalStatus: 1,
})
if err != nil {
	return err
}

fmt.Println(order.OrderID)
```

### 代理管理

```go
proxyResp, err := client.Proxies.Add(ctx, duoplus.AddProxyRequest{
	ProxyList: []duoplus.ProxyInput{
		{
			Protocol: "socks5",
			Host:     "127.0.0.1",
			Port:     3000,
			User:     "user",
			Password: "pass",
			Name:     "proxy-1",
		},
	},
	IPScanChannel: "ip2location",
})
if err != nil {
	return err
}

fmt.Println(proxyResp.Success)
```

### 查询包月开机资源

```go
startupResp, err := client.SubscriptionStartups.List(ctx, duoplus.SubscriptionStartupListRequest{
	PaginationRequest: duoplus.PaginationRequest{Page: 1, PageSize: 10},
	SortBy:            "created_at",
	Order:             "desc",
})
if err != nil {
	return err
}

for _, item := range startupResp.List {
	fmt.Println(item.ID, item.Name, item.ExpiredAt)
}
```

### 云手机分组管理

```go
groups, err := client.Groups.List(ctx, duoplus.GroupListRequest{Page: 1})
if err != nil {
	return err
}

for _, item := range groups.List {
	fmt.Println(item.ID, item.Name)
}

_, err = client.Groups.AddPhones(ctx, "group-id", []string{"image-id-1", "image-id-2"})
if err != nil {
	return err
}
```

### 应用安装与启动

```go
_, err = client.Apps.Install(ctx, duoplus.InstallAppRequest{
	ImageIDs:     []string{"image-id"},
	AppID:        "9Jp7o#0",
	AppVersionID: "9Jp7o#0",
})
if err != nil {
	return err
}

_, err = client.Apps.Start(ctx, []string{"image-id"}, "com.zhiliaoapp.musically")
if err != nil {
	return err
}
```

### 云盘文件与推送

```go
files, err := client.CloudDisk.List(ctx, duoplus.CloudDiskListRequest{
	Keyword:           "apk",
	PaginationRequest: duoplus.PaginationRequest{Page: 1, PageSize: 20},
})
if err != nil {
	return err
}

if len(files.List) > 0 {
	_, err = client.CloudDisk.PushFiles(ctx, duoplus.PushFilesRequest{
		IDs:      []string{files.List[0].ID},
		ImageIDs: []string{"image-id"},
		DestDir:  "/sdcard/Download",
	})
	if err != nil {
		return err
	}
}
```

### 查询云号码短信

```go
smsResp, err := client.CloudNumbers.SMSList(ctx, duoplus.CloudNumberSMSListRequest{
	NumberID:          "number-id",
	PaginationRequest: duoplus.PaginationRequest{Page: 1, PageSize: 50},
})
if err != nil {
	return err
}

for _, item := range smsResp.List {
	fmt.Println(item.Code, item.ReceivedAt)
}
```

### 创建自动化定时任务

```go
_, err = client.Automation.CreateTask(ctx, duoplus.CreateAutomationTaskRequest{
	TemplateID:   "template-id",
	TemplateType: 2,
	Name:         "task-name",
	Remark:       "remark",
	Images: []duoplus.TaskImage{
		{
			ImageID: "image-id",
			IssueAt: "2026-04-01 10:00",
			Config: map[string]duoplus.AutomationConfigValue{
				"text": {
					Key:      "text",
					Value:    "hello\nworld",
					Type:     "textarea",
					Required: true,
				},
			},
		},
	},
})
if err != nil {
	return err
}
```

### 创建自动化循环任务

```go
plan, err := client.Automation.CreatePlan(ctx, duoplus.CreateAutomationTaskRequest{
	TemplateID:   "template-id",
	TemplateType: 1,
	Name:         "plan-name",
	Images: []duoplus.TaskImage{
		{
			ImageID:     "image-id",
			StartAt:     "2026-04-01 09:00",
			EndAt:       "2026-04-30 18:00",
			ExecuteType: 2,
			ExecuteTime: "10:00",
			Mode:        1,
		},
	},
})
if err != nil {
	return err
}

fmt.Println(plan.ID)
```

## 设计说明

- 默认请求域名为中国大陆地址 `https://openapi.duoplus.cn`
- 非中国大陆环境可切换为 `duoplus.DefaultIntlBaseURL`
- 默认请求语言为 `zh`
- SDK 内置最小请求间隔 1 秒，以贴合文档中的 `QPS = 1` 限制
- 所有接口统一使用 `POST + JSON`
- API 业务错误会返回 `*duoplus.APIError`

## 暂未纳入

文档目录中的“一键新机”、“高级命令”、“上传文件”、“团队管理”、“邀请成员”、“绑定邮箱”当前页面没有提供明确的 OpenAPI 请求路径与参数，因此本 SDK 暂未封装这些页面；如果 DuoPlus 后续补齐接口文档，可以继续无缝补进现有 client。
