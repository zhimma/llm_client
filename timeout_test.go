package llmclient

import (
	"encoding/json"
	"testing"
)

// TestTimeoutJSONSerialization 测试 timeout 字段的 JSON 序列化
func TestTimeoutJSONSerialization(t *testing.T) {
	// 测试 ChatCompletionRequest
	req := &ChatCompletionRequest{
		Model:   "gpt-4",
		Timeout: 900, // 900秒
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	// 验证 JSON 中的 timeout 是秒数而不是纳秒
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	timeout, ok := result["timeout"].(float64)
	if !ok {
		t.Fatal("timeout 字段类型错误")
	}

	if timeout != 900 {
		t.Errorf("期望 timeout=900, 实际得到 %v", timeout)
	}

	t.Logf("✅ JSON 序列化正确: timeout=%v (秒)", timeout)
}

// TestEmbeddingTimeoutJSONSerialization 测试 EmbeddingRequest 的 timeout 序列化
func TestEmbeddingTimeoutJSONSerialization(t *testing.T) {
	req := &EmbeddingRequest{
		Model:   "text-embedding-ada-002",
		Input:   []string{"test"},
		Timeout: 600, // 600秒
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	timeout, ok := result["timeout"].(float64)
	if !ok {
		t.Fatal("timeout 字段类型错误")
	}

	if timeout != 600 {
		t.Errorf("期望 timeout=600, 实际得到 %v", timeout)
	}

	t.Logf("✅ JSON 序列化正确: timeout=%v (秒)", timeout)
}
