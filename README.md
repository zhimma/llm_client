# LLM Client - OpenAI Compatible Go SDK

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-Private-red.svg)](LICENSE)

OpenAI 兼容的 Go SDK,用于调用统一 LLM 平台。

## 特性

- ✅ **OpenAI 兼容** - 完全兼容 OpenAI API 规范
- ✅ **类型安全** - 完整的 Go 类型定义
- ✅ **流式支持** - 支持 SSE 流式响应
- ✅ **扩展功能** - 支持提示词管理、结构化提取等扩展功能
- ✅ **易于使用** - 简洁的 API 设计

## 安装

```bash
go get github.com/zhimma/llm_client@latest
```

## 快速开始

### 基础对话

```go
package main

import (
    "context"
    "fmt"
    "log"

    llmclient "github.com/zhimma/llm_client"
    "github.com/zhimma/llm_client/types"
)

func main() {
    // 创建客户端
    client := llmclient.NewClient(&llmclient.Config{
        BaseURL: "http://localhost:8888/v1",
        APIKey:  "sk-your-api-key",
    })

    // 发送对话请求
    resp, err := client.CreateChatCompletion(context.Background(), types.ChatCompletionRequest{
        Model: "qwen-max",
        Messages: []types.Message{
            {Role: "user", Content: "你好,请介绍一下自己"},
        },
        Temperature: 0.7,
        MaxTokens:   2000,
    })

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp.Choices[0].Message.Content)
}
```

### 流式对话

```go
stream, err := client.CreateChatCompletionStream(ctx, types.ChatCompletionRequest{
    Model:  "qwen-max",
    Messages: []types.Message{{Role: "user", Content: "讲个故事"}},
    Stream: true,
})

if err != nil {
    log.Fatal(err)
}
defer stream.Close()

for {
    chunk, err := stream.Recv()
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(chunk.Choices[0].Delta.Content)
}
```

### 结构化提取(扩展功能)

```go
resp, err := client.CreateChatCompletion(ctx, types.ChatCompletionRequest{
    Model: "qwen-max",
    Metadata: &types.Metadata{
        PromptKey: "medical_report_expert",
        Variables: map[string]interface{}{
            "text": markdownContent,
        },
    },
})
```

## API 文档

### Client 配置

```go
type Config struct {
    BaseURL string        // LLM 平台地址,如 "http://localhost:8888/v1"
    APIKey  string        // API Key
    Timeout time.Duration // 请求超时时间,默认 30s
}
```

### 主要方法

#### Chat Completions

```go
// 非流式对话
func (c *Client) CreateChatCompletion(ctx context.Context, req types.ChatCompletionRequest) (*types.ChatCompletionResponse, error)

// 流式对话
func (c *Client) CreateChatCompletionStream(ctx context.Context, req types.ChatCompletionRequest) (*types.ChatCompletionStream, error)
```

#### Models

```go
// 获取模型列表
func (c *Client) ListModels(ctx context.Context) (*types.ModelsList, error)

// 获取单个模型信息
func (c *Client) GetModel(ctx context.Context, modelID string) (*types.Model, error)
```

## 扩展功能

通过 `Metadata` 字段可以使用平台的扩展功能:

```go
type Metadata struct {
    PromptKey    string                 `json:"prompt_key,omitempty"`     // 提示词版本控制
    Variables    map[string]interface{} `json:"variables,omitempty"`      // 变量替换
    UseMemory    bool                   `json:"use_memory,omitempty"`     // 启用记忆
    UserIdentity map[string]string      `json:"user_identity,omitempty"`  // 用户身份
}
```

## 示例

查看 `examples/` 目录获取更多示例:

- `examples/chat/` - 基础对话示例
- `examples/streaming/` - 流式对话示例

## OpenAI 兼容性

| 功能 | OpenAI API | 支持状态 |
|------|-----------|---------|
| Chat Completions | `POST /v1/chat/completions` | ✅ |
| Streaming | SSE | ✅ |
| Models List | `GET /v1/models` | ✅ |
| Embeddings | `POST /v1/embeddings` | 🚧 计划中 |

## 许可证

Private - 仅供内部使用
