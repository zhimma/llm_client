package llmclient

import (
	"encoding/json"
	"time"
)

// Message 消息结构
type Message struct {
	Role             string      `json:"role"`
	Content          interface{} `json:"content"` // 支持 string 或 []ContentPart
	Name             string      `json:"name,omitempty"`
	ReasoningContent string      `json:"reasoning_content,omitempty"`
}

// ContentPart 多模态内容部分
type ContentPart struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

// ImageURL 图片 URL 信息
type ImageURL struct {
	URL    string `json:"url"`
	Detail string `json:"detail,omitempty"` // auto, low, high
}

// ChatCompletionRequest 聊天补全请求
type ChatCompletionRequest struct {
	Model            string                 `json:"model,omitempty"`
	Messages         []Message              `json:"messages,omitempty"`
	PromptKey        string                 `json:"prompt_key,omitempty"`
	Variables        map[string]interface{} `json:"variables,omitempty"`
	Stream           bool                   `json:"stream,omitempty"`
	Temperature      float64                `json:"temperature,omitempty"`
	MaxTokens        int                    `json:"max_tokens,omitempty"`
	TopP             float64                `json:"top_p,omitempty"`
	UseMemory        bool                   `json:"use_memory,omitempty"`
	User             string                 `json:"user,omitempty"`
	UserIdentity     *UserIdentity          `json:"user_identity,omitempty"`
	ResponseFormat   *ResponseFormat        `json:"response_format,omitempty"`
	FrequencyPenalty float64                `json:"frequency_penalty,omitempty"`
	PresencePenalty  float64                `json:"presence_penalty,omitempty"`
	Timeout          time.Duration          `json:"timeout,omitempty"`
	Files            []File                 `json:"files,omitempty"` // 多模态文件输入
}

// File 多模态文件信息
type File struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type"` // image, document, etc.
	URL  string `json:"url"`
}

// UserIdentity 用户身份标识
type UserIdentity struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// ResponseFormat 响应格式
type ResponseFormat struct {
	Type string `json:"type"` // json_object or text
}

// ChatCompletionResponse 聊天补全响应
type ChatCompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice 对话选择项
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Usage Token 使用统计
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// EmbeddingRequest 向量化请求
type EmbeddingRequest struct {
	Model   string        `json:"model"`
	Input   []string      `json:"input"`
	User    string        `json:"user,omitempty"`
	Timeout time.Duration `json:"timeout,omitempty"`
}

// EmbeddingResponse 向量化响应
type EmbeddingResponse struct {
	Object string          `json:"object"`
	Data   []EmbeddingData `json:"data"`
	Model  string          `json:"model"`
	Usage  Usage           `json:"usage"`
}

// EmbeddingData 向量数据项
type EmbeddingData struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}

// ModelList 模型列表响应
type ModelList struct {
	Object string      `json:"object"`
	Data   []ModelInfo `json:"data"`
}

// ModelInfo 模型信息
type ModelInfo struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param"`
		Code    string `json:"code"`
	} `json:"error"`
}

// ToJSON 用于调试或日志
func (r *ChatCompletionRequest) ToJSON() string {
	b, _ := json.Marshal(r)
	return string(b)
}
