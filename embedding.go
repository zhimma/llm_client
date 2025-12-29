package llmclient

import (
	"context"
	"encoding/json"
	"fmt"
)

// Embeddings 定义向量化服务接口
type Embeddings interface {
	i()
	// Create 创建向量嵌入
	Create(ctx context.Context, req *EmbeddingRequest) (*EmbeddingResponse, error)
}

// embeddings 实现 Embeddings 接口
type embeddings struct {
	client *Client
}

// NewEmbeddingService 创建向量服务实例
func NewEmbeddingService(client *Client) Embeddings {
	return &embeddings{client: client}
}

// i 私有防外部实现
func (s *embeddings) i() {}

// Create 创建向量嵌入
func (s *embeddings) Create(ctx context.Context, req *EmbeddingRequest) (*EmbeddingResponse, error) {
	body, err := s.client.post(ctx, "/embeddings", req)
	if err != nil {
		return nil, err
	}

	var result EmbeddingResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal embeddings failed: %w", err)
	}

	return &result, nil
}
