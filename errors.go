package llmclient

import "errors"

var (
	// ErrMissingBaseURL BaseURL 未配置
	ErrMissingBaseURL = errors.New("llm_client: missing base URL")

	// ErrMissingAPIKey API Key 未配置
	ErrMissingAPIKey = errors.New("llm_client: missing API key")

	// ErrInvalidRequest 无效的请求参数
	ErrInvalidRequest = errors.New("llm_client: invalid request")

	// ErrRequestFailed 请求失败
	ErrRequestFailed = errors.New("llm_client: request failed")

	// ErrStreamClosed 流已关闭
	ErrStreamClosed = errors.New("llm_client: stream closed")

	// ErrInvalidResponse 无效的响应
	ErrInvalidResponse = errors.New("llm_client: invalid response")
)

// APIError API 错误响应
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	Code       string `json:"code"`
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	return e.Message
}
