package types

// Model 模型信息 (OpenAI 兼容)
type Model struct {
	ID      string `json:"id"`       // 模型 ID
	Object  string `json:"object"`   // "model"
	Created int64  `json:"created"`  // 创建时间戳
	OwnedBy string `json:"owned_by"` // 所有者

	// 扩展字段 (非 OpenAI 标准)
	Provider             string  `json:"provider,omitempty"`               // 提供商
	Category             string  `json:"category,omitempty"`               // 分类
	MaxTokens            int     `json:"max_tokens,omitempty"`             // 最大 Token 数
	InputPrice           float64 `json:"input_price,omitempty"`            // 输入价格
	OutputPrice          float64 `json:"output_price,omitempty"`           // 输出价格
	SupportsVision       bool    `json:"supports_vision,omitempty"`        // 支持视觉
	SupportsFunctionCall bool    `json:"supports_function_call,omitempty"` // 支持函数调用
}

// ModelsList 模型列表响应 (OpenAI 兼容)
type ModelsList struct {
	Object string  `json:"object"` // "list"
	Data   []Model `json:"data"`   // 模型列表
}
