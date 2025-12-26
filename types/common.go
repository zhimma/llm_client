package types

// Message 消息结构 (OpenAI 兼容)
type Message struct {
	Role    string `json:"role"`           // "system", "user", "assistant"
	Content string `json:"content"`        // 消息内容
	Name    string `json:"name,omitempty"` // 可选的名称
}

// Usage Token 使用量统计 (OpenAI 兼容)
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Choice 选择项 (OpenAI 兼容)
type Choice struct {
	Index        int      `json:"index"`
	Message      *Message `json:"message,omitempty"`       // 非流式响应
	Delta        *Message `json:"delta,omitempty"`         // 流式响应
	FinishReason string   `json:"finish_reason,omitempty"` // "stop", "length", "content_filter"
}

// Metadata 扩展元数据(非 OpenAI 标准,用于平台扩展功能)
type Metadata struct {
	// 提示词管理
	PromptKey string                 `json:"prompt_key,omitempty"` // 提示词版本控制
	Variables map[string]interface{} `json:"variables,omitempty"`  // 变量替换

	// 记忆管理
	UseMemory    bool              `json:"use_memory,omitempty"`    // 启用记忆
	UserIdentity map[string]string `json:"user_identity,omitempty"` // 用户身份

	// 其他扩展
	Custom map[string]interface{} `json:"custom,omitempty"` // 自定义字段
}
