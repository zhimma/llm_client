package llmclient

import "time"

// Config LLM Client 配置
type Config struct {
	// BaseURL LLM 平台地址 (例如: "http://localhost:8888/v1")
	BaseURL string

	// APIKey API 密钥
	APIKey string

	// Timeout 请求超时时间,默认 30 秒
	Timeout time.Duration

	// MaxRetries 最大重试次数,默认 3 次
	MaxRetries int

	// Debug 是否启用调试模式
	Debug bool
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Timeout:    30 * time.Second,
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
		c.Timeout = 30 * time.Second
	}
	if c.MaxRetries < 0 {
		c.MaxRetries = 0
	}
	return nil
}
