package llmclient

import (
	"errors"
)

var (
	ErrMissingBaseURL = errors.New("missing BaseURL")
	ErrMissingAPIKey  = errors.New("missing APIKey")
)

// Config LLM Client 配置
type Config struct {
	// BaseURL LLM 平台地址 (例如: "http://localhost:8888/v1")
	BaseURL string

	// APIKey API 密钥
	APIKey string

	// Timeout 请求超时时间(秒),默认 600 秒
	Timeout int

	// MaxRetries 最大重试次数,默认 3 次
	MaxRetries int

	// Debug 是否启用调试模式
	Debug bool
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Timeout:    600, // 600秒 = 10分钟
		MaxRetries: 3,
		Debug:      false,
	}
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.BaseURL == "" {
		return ErrMissingBaseURL
	}
	if c.APIKey == "" {
		return ErrMissingAPIKey
	}
	if c.Timeout <= 0 {
		c.Timeout = 600 // 默认600秒
	}
	if c.MaxRetries < 0 {
		c.MaxRetries = 0
	}
	return nil
}
