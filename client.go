package llmclient

import (
	"context"
	"fmt"

	"github.com/zhimma/llm_client/internal"
	"github.com/zhimma/llm_client/types"
)

// Client LLM 客户端
type Client struct {
	config     *Config
	httpClient *internal.HTTPClient
}

// NewClient 创建新的 LLM 客户端
func NewClient(config *Config) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	// 验证配置
	if err := config.Validate(); err != nil {
		panic(fmt.Sprintf("invalid config: %v", err))
	}

	return &Client{
		config: config,
		httpClient: internal.NewHTTPClient(
			config.BaseURL,
			config.APIKey,
			config.Timeout,
			config.Debug,
		),
	}
}

// CreateChatCompletion 创建 Chat Completion (非流式)
func (c *Client) CreateChatCompletion(ctx context.Context, req types.ChatCompletionRequest) (*types.ChatCompletionResponse, error) {
	// 确保不是流式请求
	req.Stream = false

	resp, err := c.httpClient.Post(ctx, "/chat/completions", req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	var result types.ChatCompletionResponse
	if err := internal.DecodeResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidResponse, err)
	}

	return &result, nil
}

// CreateChatCompletionStream 创建 Chat Completion (流式)
func (c *Client) CreateChatCompletionStream(ctx context.Context, req types.ChatCompletionRequest) (*ChatCompletionStream, error) {
	// 确保是流式请求
	req.Stream = true

	resp, err := c.httpClient.Post(ctx, "/chat/completions", req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	// 检查响应状态
	if resp.StatusCode >= 400 {
		resp.Body.Close()
		return nil, fmt.Errorf("%w: status code %d", ErrRequestFailed, resp.StatusCode)
	}

	return &ChatCompletionStream{
		reader: internal.NewStreamReader(ctx, resp),
	}, nil
}

// ListModels 获取模型列表
func (c *Client) ListModels(ctx context.Context) (*types.ModelsList, error) {
	resp, err := c.httpClient.Get(ctx, "/models")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	var result types.ModelsList
	if err := internal.DecodeResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidResponse, err)
	}

	return &result, nil
}

// GetModel 获取单个模型信息
func (c *Client) GetModel(ctx context.Context, modelID string) (*types.Model, error) {
	resp, err := c.httpClient.Get(ctx, "/models/"+modelID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	var result types.Model
	if err := internal.DecodeResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidResponse, err)
	}

	return &result, nil
}

// ChatCompletionStream 流式响应封装
type ChatCompletionStream struct {
	reader *internal.StreamReader
}

// Recv 接收下一个流式响应块
func (s *ChatCompletionStream) Recv() (*types.ChatCompletionStreamResponse, error) {
	return s.reader.Recv()
}

// Close 关闭流
func (s *ChatCompletionStream) Close() error {
	return s.reader.Close()
}
