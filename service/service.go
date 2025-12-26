package service

import (
	"context"

	llmclient "github.com/zhimma/llm_client"
	"github.com/zhimma/llm_client/types"
)

// LLMService LLM 服务接口
type LLMService interface {
	// Chat 聊天接口(统一处理流式和非流式)
	Chat(ctx context.Context, req *types.ChatCompletionRequest) (*types.ChatCompletionResponse, *llmclient.ChatCompletionStream, error)

	// ListModels 获取模型列表
	ListModels(ctx context.Context) (*types.ModelsList, error)

	// GetModel 获取单个模型信息
	GetModel(ctx context.Context, modelID string) (*types.Model, error)

	// ListProviders 获取提供商列表(扩展功能)
	ListProviders(ctx context.Context) (interface{}, error)
}

// llmService 服务实现
type llmService struct {
	client *llmclient.Client
}

// NewService 创建 LLM 服务
func NewService(client *llmclient.Client) LLMService {
	return &llmService{
		client: client,
	}
}

// Chat 聊天接口
func (s *llmService) Chat(ctx context.Context, req *types.ChatCompletionRequest) (*types.ChatCompletionResponse, *llmclient.ChatCompletionStream, error) {
	// 参数校验和默认值
	if req.MaxTokens <= 0 {
		req.MaxTokens = 2000
	}
	if req.Temperature < 0 || req.Temperature > 2 {
		req.Temperature = 0.7
	}

	// 根据 Stream 参数选择调用方式
	if req.Stream {
		stream, err := s.client.CreateChatCompletionStream(ctx, *req)
		return nil, stream, err
	}

	resp, err := s.client.CreateChatCompletion(ctx, *req)
	return resp, nil, err
}

// ListModels 获取模型列表
func (s *llmService) ListModels(ctx context.Context) (*types.ModelsList, error) {
	return s.client.ListModels(ctx)
}

// GetModel 获取单个模型信息
func (s *llmService) GetModel(ctx context.Context, modelID string) (*types.Model, error) {
	return s.client.GetModel(ctx, modelID)
}

// ListProviders 获取提供商列表
func (s *llmService) ListProviders(ctx context.Context) (interface{}, error) {
	// 暂时返回空,等待平台实现
	return map[string]interface{}{
		"providers": []string{},
		"count":     0,
	}, nil
}
