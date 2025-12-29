package llmclient

import (
	"context"
	"encoding/json"
	"fmt"
)

// Models 定义模型管理服务接口
type Models interface {
	i()
	// List 获取模型列表
	List(ctx context.Context) (*ModelList, error)
}

// models 实现 Models 接口
type models struct {
	client *Client
}

// NewModelService 创建模型服务实例
func NewModelService(client *Client) Models {
	return &models{client: client}
}

// i 私有防外部实现
func (s *models) i() {}

// List 获取模型列表
func (s *models) List(ctx context.Context) (*ModelList, error) {
	body, err := s.client.get(ctx, "/models")
	if err != nil {
		return nil, err
	}

	var result ModelList
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal list models failed: %w", err)
	}

	return &result, nil
}
