package types

// ChatCompletionRequest Chat Completions 请求 (OpenAI 兼容)
type ChatCompletionRequest struct {
	// 必需参数
	Model    string    `json:"model"`    // 模型名称
	Messages []Message `json:"messages"` // 消息列表

	// 可选参数 (OpenAI 标准)
	Temperature      float64  `json:"temperature,omitempty"`       // 温度 (0-2)
	TopP             float64  `json:"top_p,omitempty"`             // Top-P 采样
	N                int      `json:"n,omitempty"`                 // 生成数量
	Stream           bool     `json:"stream,omitempty"`            // 是否流式输出
	Stop             []string `json:"stop,omitempty"`              // 停止词
	MaxTokens        int      `json:"max_tokens,omitempty"`        // 最大 Token 数
	PresencePenalty  float64  `json:"presence_penalty,omitempty"`  // 存在惩罚
	FrequencyPenalty float64  `json:"frequency_penalty,omitempty"` // 频率惩罚
	User             string   `json:"user,omitempty"`              // 用户标识

	// 扩展参数 (非 OpenAI 标准)
	Metadata *Metadata `json:"metadata,omitempty"` // 扩展元数据
}

// ChatCompletionResponse Chat Completions 响应 (OpenAI 兼容)
type ChatCompletionResponse struct {
	ID      string   `json:"id"`      // 唯一标识
	Object  string   `json:"object"`  // "chat.completion"
	Created int64    `json:"created"` // 创建时间戳
	Model   string   `json:"model"`   // 使用的模型
	Choices []Choice `json:"choices"` // 选择列表
	Usage   *Usage   `json:"usage"`   // Token 使用量
}

// ChatCompletionStreamResponse 流式响应块 (OpenAI 兼容)
type ChatCompletionStreamResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"` // "chat.completion.chunk"
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}
