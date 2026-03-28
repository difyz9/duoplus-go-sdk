# DuoPlus Go SDK

基于 DuoPlus OpenAPI 的非官方 Golang SDK。

## 功能概览

已覆盖 DuoPlus API 文档中公开可识别的 OpenAPI 页面：

- 云手机管理
- 云号码
- 云手机分组
- 代理管理
- 包月开机管理
- 应用管理
- 云盘管理
- 自动化

当前未纳入的页面：

- 一键新机
- 高级命令
- 上传文件
- 团队管理
- 邀请成员
- 绑定邮箱

这些页面在文档中是产品操作说明，没有公开 `POST /api/v1/...` 接口路径与请求结构，因此 SDK 不对其做猜测性封装。

## 项目结构

```text
github.com/difyz9/duoplus-go-sdk/
├── app/                    # 应用管理模块
├── automation/             # 自动化模块
├── clouddisk/              # 云盘管理模块
├── cloudnumber/            # 云号码模块
├── cloudphone/             # 云手机管理模块
├── common/                 # 通用请求/响应类型
├── group/                  # 云手机分组模块
├── internal/clientcore/    # HTTP 核心、鉴权、限流、错误处理
├── proxy/                  # 代理管理模块
├── subscriptionstartup/    # 包月开机模块
├── examples/               # 可运行示例
├── client.go               # 根客户端入口
└── readme.md
```

## 安装

```bash
go get github.com/difyz9/duoplus-go-sdk
```

## 设计说明

- 默认请求域名为中国大陆地址 `https://openapi.duoplus.cn`
- 非中国大陆环境可切换为 `duoplus.DefaultIntlBaseURL`
- 默认请求语言为 `zh`
- 默认按文档限制内置 1 秒最小请求间隔
- 所有接口统一使用 `POST + JSON`
- 业务错误返回 `*duoplus.APIError`
- 通用类型集中放在 `common` 包，模块请求/响应类型放在对应子包中

## 快速开始

```go
package main

import (
	"context"
	"fmt"
	"log"

	duoplus "github.com/difyz9/duoplus-go-sdk"
	"github.com/difyz9/duoplus-go-sdk/cloudphone"
	"github.com/difyz9/duoplus-go-sdk/common"
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

	resp, err := client.CloudPhones.List(context.Background(), cloudphone.ListRequest{
		PaginationRequest: common.PaginationRequest{Page: 1, PageSize: 10},
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range resp.List {
		fmt.Println(item.ID, item.Name, item.Status)
	}
}
```

## 如何使用

### 1. 初始化客户端

```go
client, err := duoplus.NewClient(
	apiKey,
	duoplus.WithBaseURL(duoplus.DefaultBaseURL),
	duoplus.WithLanguage("zh"),
)
```

可选项：

- `duoplus.WithBaseURL(...)`
- `duoplus.WithLanguage(...)`
- `duoplus.WithHTTPClient(...)`
- `duoplus.WithMinInterval(...)`

### 2. 从模块入口访问功能

根客户端上挂载了按模块划分的 service：

- `client.CloudPhones`
- `client.CloudNumbers`
- `client.Groups`
- `client.Proxies`
- `client.SubscriptionStartups`
- `client.Apps`
- `client.CloudDisk`
- `client.Automation`

### 3. 从对应 package 引入请求/响应类型

例如云手机列表使用：

- `cloudphone.ListRequest`
- `cloudphone.ListResponse`

分页参数和通用返回类型则来自：

- `common.PaginationRequest`
- `common.Pagination`
- `common.MessageResponse`

## 功能列表

### 云手机管理

- 列表查询
- 开机、关机、重启
- 状态查询
- 详情查询
- 参数批量修改
- 手机型号列表
- Root 权限设置
- ADB 命令执行
- ADB 开关与 IP 白名单
- 分享与分享密码管理
- 标签、连接用户、资源列表、分辨率
- 购买与续费

对应 package：`cloudphone`

### 云号码

- 云号码列表
- 云号码短信查询
- 写入短信到云机

对应 package：`cloudnumber`

### 云手机分组

- 分组列表
- 添加到分组
- 移动到分组
- 创建分组
- 编辑分组
- 删除分组

对应 package：`group`

### 代理管理

- 代理列表
- 批量添加
- 批量删除
- 批量刷新 URL
- 修改代理

对应 package：`proxy`

### 包月开机管理

- 列表
- 购买
- 续费

对应 package：`subscriptionstartup`

### 应用管理

- 平台应用列表
- 团队应用列表
- 批量安装
- 已安装应用列表
- 批量卸载
- 批量启动
- 批量关闭

对应 package：`app`

### 云盘管理

- 文件列表
- 推送文件到云机

对应 package：`clouddisk`

### 自动化

- 自定义模板列表
- 官方模板列表
- 创建定时任务
- 创建循环任务
- 循环任务列表
- 修改循环任务
- 暂停/开始执行循环任务
- 删除循环任务
- 定时任务列表
- 查看任务日志
- 修改任务发布时间
- 取消/重新执行任务

对应 package：`automation`

## 使用案例

### 案例 1：查询云手机状态

```go
status, err := client.CloudPhones.Status(ctx, []string{"image-id-1", "image-id-2"})
if err != nil {
	return err
}

for _, item := range status.List {
	fmt.Println(item.ID, item.Status)
}
```

### 案例 2：批量添加代理

```go
resp, err := client.Proxies.Add(ctx, proxy.AddRequest{
	ProxyList: []proxy.Input{
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

fmt.Println(resp.Success)
```

### 案例 3：查询云盘文件并推送到云机

```go
files, err := client.CloudDisk.List(ctx, clouddisk.ListRequest{
	Keyword:           "apk",
	PaginationRequest: common.PaginationRequest{Page: 1, PageSize: 20},
})
if err != nil {
	return err
}

if len(files.List) > 0 {
	_, err = client.CloudDisk.PushFiles(ctx, clouddisk.PushFilesRequest{
		IDs:      []string{files.List[0].ID},
		ImageIDs: []string{"image-id"},
		DestDir:  "/sdcard/Download",
	})
	if err != nil {
		return err
	}
}
```

### 案例 4：创建自动化定时任务

```go
_, err := client.Automation.CreateTask(ctx, automation.CreateTaskRequest{
	TemplateID:   "template-id",
	TemplateType: 2,
	Name:         "task-name",
	Remark:       "remark",
	Images: []automation.TaskImage{
		{
			ImageID: "image-id",
			IssueAt: "2026-04-01 10:00",
			Config: map[string]automation.ConfigValue{
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

### 案例 5：购买云手机

```go
order, err := client.CloudPhones.Purchase(ctx, cloudphone.PurchaseRequest{
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

## 示例目录

- `examples/basic`: 基础云手机列表示例
- `examples/automation`: 自动化模板列表示例
- `examples/group-disk`: 云手机分组与云盘示例
- `examples/power-on-wait`: 开机后轮询状态直到完成的真实业务脚本示例
- `examples/power-on-adb-check`: 开机后执行 ADB 命令并校验输出的真实业务脚本示例
- `examples/real-workflow`: 基于真实资源自动选择云手机和云盘文件的完整业务流程示例

运行方式：

```bash
export DUOPLUS_API_KEY=your-api-key
go run ./examples/basic
go run ./examples/automation
go run ./examples/group-disk
```

### 真实业务脚本示例

#### 示例 1：开机后轮询直到完成

默认只做 dry-run，不会真的开机：

```bash
export DUOPLUS_API_KEY=your-api-key
go run ./examples/power-on-wait
```

执行真实开机并轮询状态：

```bash
export DUOPLUS_API_KEY=your-api-key
export DUOPLUS_EXECUTE=1
go run ./examples/power-on-wait
```

可选地指定目标云手机：

```bash
export DUOPLUS_TARGET_IMAGE_ID=your-image-id
```

#### 示例 2：开机后执行 ADB 命令并校验结果

默认也是 dry-run：

```bash
export DUOPLUS_API_KEY=your-api-key
go run ./examples/power-on-adb-check
```

执行真实流程：

```bash
export DUOPLUS_API_KEY=your-api-key
export DUOPLUS_EXECUTE=1
go run ./examples/power-on-adb-check
```

可选环境变量：

```bash
export DUOPLUS_TARGET_IMAGE_ID=your-image-id
export DUOPLUS_ADB_COMMAND='getprop ro.product.model'
export DUOPLUS_EXPECT_SUBSTRING='Pixel'
```

这个脚本会：

- 自动选择一台云手机，优先选择关机中的机器
- 如有需要先开机，并轮询到开机完成
- 执行一条 ADB shell 命令
- 对返回内容做简单校验，适合接业务前置健康检查

#### 示例 3：更贴近业务的完整流程

这个脚本会：

- 自动选择一台云手机
- 自动选择云盘中的一个文件
- 如果云手机处于关机状态，则先开机并轮询到完成
- 再把文件推送到云机的目标目录

默认也是 dry-run：

```bash
export DUOPLUS_API_KEY=your-api-key
go run ./examples/real-workflow
```

执行真实流程：

```bash
export DUOPLUS_API_KEY=your-api-key
export DUOPLUS_EXECUTE=1
go run ./examples/real-workflow
```

可选地自定义推送目录：

```bash
export DUOPLUS_DEST_DIR=/sdcard/Download
```

说明：

- 真实执行可能触发云手机开机计费
- 脚本默认 dry-run，就是为了避免误操作
- 我当前测试到的真实资源状态是：有 2 台已关机云手机、1 个云盘文件、0 个云手机分组，因此这个示例会优先选取关机云机和首个云盘文件

## 错误处理

当接口返回业务错误时，会返回 `*duoplus.APIError`。

```go
resp, err := client.CloudPhones.List(ctx, cloudphone.ListRequest{})
if err != nil {
	var apiErr *duoplus.APIError
	if errors.As(err, &apiErr) {
		fmt.Println(apiErr.Code, apiErr.Message)
	}
	return err
}

_ = resp
```
