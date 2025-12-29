package llmclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Chat 定义对话服务接口
type Chat interface {
	i()
	// Create 创建对话补全
	Create(ctx context.Context, req *ChatCompletionRequest) (*ChatCompletionResponse, error)
	// CreateStream 创建流式对话补全
	CreateStream(ctx context.Context, req *ChatCompletionRequest) (*http.Response, error)
}

// chat 实现 Chat 接口
type chat struct {
	client *Client
}

// NewChatService 创建对话服务实例
func NewChatService(client *Client) Chat {
	return &chat{client: client}
}

// i 私有防外部实现
func (s *chat) i() {}

// Create 创建对话补全
func (s *chat) Create(ctx context.Context, req *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	body, err := s.client.post(ctx, "/chat/completions", req)
	if err != nil {
		return nil, err
	}

	var result ChatCompletionResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal chat completions failed: %w", err)
	}

	return &result, nil
}

// CreateStream 创建流式对话补全
func (s *chat) CreateStream(ctx context.Context, req *ChatCompletionRequest) (*http.Response, error) {
	req.Stream = true
	return s.client.postStream(ctx, "/chat/completions", req)
}
