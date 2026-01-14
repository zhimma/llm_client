package llmclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client LLM HTTP 客户端 (作为服务入口)
type Client struct {
	config     *Config
	httpClient *http.Client

	// 子服务封装 (使用接口)
	Chat       Chat
	Embeddings Embeddings
	Models     Models
}

// NewClient 创建新的 LLM 客户端并初始化子服务
func NewClient(config *Config) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	if err := config.Validate(); err != nil {
		panic(fmt.Sprintf("invalid config: %v", err))
	}

	c := &Client{
		config: config,
		httpClient: &http.Client{
			// 不要在这里设置 Timeout，使用 Context 进行更细粒度的控制
			Timeout: 0,
		},
	}

	// 初始化子服务 (注入 client 引用)
	c.Chat = NewChatService(c)
	c.Embeddings = NewEmbeddingService(c)
	c.Models = NewModelService(c)

	return c
}

// --- 通用底层请求方法 (仅限包内使用) ---

// get 发送 GET 请求, 返回原始响应体
func (c *Client) get(ctx context.Context, path string) (body []byte, err error) {
	baseURL := strings.TrimSuffix(c.config.BaseURL, "/")
	url := baseURL + path

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	req.Header.Set("X-API-Key", c.config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, c.handleError(resp.StatusCode, body)
	}

	return body, nil
}

// post 发送 POST 请求, 返回原始响应体
func (c *Client) post(ctx context.Context, path string, data interface{}) (body []byte, err error) {
	baseURL := strings.TrimSuffix(c.config.BaseURL, "/")
	url := baseURL + path

	// 🎯 动态处理超时
	// 如果数据中指定了超时,则使用该超时;否则使用默认的 600s
	timeoutSeconds := c.config.Timeout
	if timeoutSeconds <= 0 {
		timeoutSeconds = 600
	}

	// 尝试从不同的请求结构中提取超时设置
	switch v := data.(type) {
	case *ChatCompletionRequest:
		if v.Timeout > 0 {
			timeoutSeconds = v.Timeout
		}
	case *EmbeddingRequest:
		if v.Timeout > 0 {
			timeoutSeconds = v.Timeout
		}
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	req.Header.Set("X-API-Key", c.config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, c.handleError(resp.StatusCode, body)
	}

	return body, nil
}

// postStream 发送 POST 请求并返回流式响应
func (c *Client) postStream(ctx context.Context, path string, data interface{}) (resp *http.Response, err error) {
	baseURL := strings.TrimSuffix(c.config.BaseURL, "/")
	url := baseURL + path

	// 🎯 动态处理超时 (流式请求通常需要更长的生命周期,但仍受 Context 约束)
	timeoutSeconds := c.config.Timeout
	if timeoutSeconds <= 0 {
		timeoutSeconds = 600
	}

	switch v := data.(type) {
	case *ChatCompletionRequest:
		if v.Timeout > 0 {
			timeoutSeconds = v.Timeout
		}
	}

	// 注意:流式请求不能在方法层面就结束 Context,需要由调用者管理
	// 这里通过 timeout 设置的是请求建立的阶段,而非整个流的耗时
	// 为了资源安全,我们为请求建立设置一个防御性超时
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	// 在流式请求中，如果请求失败，我们需要显式调用 cancel
	// 如果请求成功，cancel 将通过某种方式透传或在适当时候关闭（通常由调用者处理或通过 Body 代理）
	// 但在此底层方法中，我们至少确保在 Do(req) 完成前或发生错误时进行保护
	defer func() {
		if err != nil && cancel != nil {
			cancel()
		}
	}()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	req.Header.Set("X-API-Key", c.config.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, c.handleError(resp.StatusCode, body)
	}

	return resp, nil
}

// handleError 处理错误响应
func (c *Client) handleError(statusCode int, body []byte) error {
	var errResp ErrorResponse
	if err := json.Unmarshal(body, &errResp); err == nil && errResp.Error.Message != "" {
		return fmt.Errorf("LLM API Error (HTTP %d): %s [%s]", statusCode, errResp.Error.Message, errResp.Error.Code)
	}
	return fmt.Errorf("LLM API Error (HTTP %d): %s", statusCode, string(body))
}
