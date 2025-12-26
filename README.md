# LLM Client - OpenAI Compatible Go SDK

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-Private-red.svg)](LICENSE)

完整的 OpenAI 兼容 Go SDK,提供开箱即用的 LLM 功能。

## 特性

- ✅ **OpenAI 兼容** - 完全兼容 OpenAI API 规范
- ✅ **开箱即用** - 一行代码完成集成
- ✅ **完整封装** - Client + Service + Handler 三层架构
- ✅ **类型安全** - 完整的 Go 类型定义
- ✅ **流式支持** - 支持 SSE 流式响应
- ✅ **扩展功能** - 支持提示词管理、结构化提取等
- ✅ **灵活可控** - 可选择使用任意层级

## 快速开始

### 方式 1: 一键集成(最简单)

```go
package main

import (
    "github.com/gin-gonic/gin"
    llmclient "github.com/zhimma/llm_client"
    "github.com/zhimma/llm_client/handler"
)

func main() {
    r := gin.Default()

    // 一行代码完成 LLM 功能集成
    handler.QuickStart(r, &llmclient.Config{
        BaseURL: "http://localhost:8888/v1",
        APIKey:  "sk-your-api-key",
    })

    r.Run(":8080")
}
```

### 方式 2: 分层使用(灵活控制)

```go
package main

import (
    "github.com/gin-gonic/gin"
    llmclient "github.com/zhimma/llm_client"
    "github.com/zhimma/llm_client/service"
    "github.com/zhimma/llm_client/handler"
)

func main() {
    r := gin.Default()

    // 1. 创建 Client
    client := llmclient.NewClient(&llmclient.Config{
        BaseURL: "http://localhost:8888/v1",
        APIKey:  "sk-your-api-key",
    })

    // 2. 创建 Service(可以在这里添加自定义逻辑)
    svc := service.NewService(client)

    // 3. 注册路由
    llmGroup := r.Group("/v1")
    // 可以添加中间件
    llmGroup.Use(yourAuthMiddleware())
    handler.RegisterRoutes(llmGroup, svc)

    r.Run(":8080")
}
```

### 方式 3: 仅使用 Client(最灵活)

```go
package main

import (
    "context"
    "fmt"

    llmclient "github.com/zhimma/llm_client"
    "github.com/zhimma/llm_client/types"
)

func main() {
    client := llmclient.NewClient(&llmclient.Config{
        BaseURL: "http://localhost:8888/v1",
        APIKey:  "sk-your-api-key",
    })

    resp, err := client.CreateChatCompletion(context.Background(), types.ChatCompletionRequest{
        Model: "qwen-max",
        Messages: []types.Message{
            {Role: "user", Content: "你好"},
        },
    })

    fmt.Println(resp.Choices[0].Message.Content)
}
```

## 架构

```
llm_client/
├── client.go           # HTTP 客户端(底层)
├── service/           # 业务逻辑层
│   └── service.go
├── handler/           # Gin 路由处理器
│   └── handler.go
├── types/             # OpenAI 兼容类型
│   ├── chat.go
│   ├── models.go
│   └── common.go
└── internal/          # 内部实现
    ├── http.go
    └── stream.go
```

## API 文档

### Client 配置

```go
type Config struct {
    BaseURL string        // LLM 平台地址
    APIKey  string        // API Key
    Timeout time.Duration // 请求超时,默认 30s
}
```

### Service 接口

```go
type LLMService interface {
    Chat(ctx context.Context, req *types.ChatCompletionRequest) (*types.ChatCompletionResponse, *ChatCompletionStream, error)
    ListModels(ctx context.Context) (*types.ModelsList, error)
    GetModel(ctx context.Context, modelID string) (*types.Model, error)
    ListProviders(ctx context.Context) (interface{}, error)
}
```

### Handler 路由

```go
// 注册路由
handler.RegisterRoutes(r, svc)

// 提供的路由:
// POST /chat/completions  - Chat Completions (OpenAI 兼容)
// POST /chat             - Chat (兼容旧路径)
// GET  /models           - 模型列表
// GET  /models/:model    - 单个模型
// GET  /providers        - 提供商列表
```

## 扩展功能

通过 `Metadata` 字段使用平台扩展功能:

```go
resp, err := client.CreateChatCompletion(ctx, types.ChatCompletionRequest{
    Model: "qwen-max",
    Metadata: &types.Metadata{
        PromptKey: "medical_report_expert",  // 提示词版本控制
        Variables: map[string]interface{}{   // 变量替换
            "text": content,
        },
    },
})
```

## OpenAI 兼容性

| 功能 | OpenAI API | 支持状态 |
|------|-----------|---------|
| Chat Completions | `POST /v1/chat/completions` | ✅ |
| Streaming | SSE | ✅ |
| Models List | `GET /v1/models` | ✅ |
| Embeddings | `POST /v1/embeddings` | 🚧 计划中 |

## 示例

查看 `examples/` 目录获取更多示例:

- `examples/chat/` - 基础对话
- `examples/streaming/` - 流式对话
- `examples/quickstart/` - 一键集成

## 许可证

Private - 仅供内部使用
