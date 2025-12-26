package main

import (
	"context"
	"fmt"
	"log"

	llmclient "github.com/zhimma/llm_client"
	"github.com/zhimma/llm_client/types"
)

func main() {
	// 创建客户端
	client := llmclient.NewClient(&llmclient.Config{
		BaseURL: "http://localhost:8888/v1",
		APIKey:  "sk-your-api-key",
	})

	// 创建对话请求
	req := types.ChatCompletionRequest{
		Model: "qwen-max",
		Messages: []types.Message{
			{Role: "system", Content: "你是一个有帮助的助手。"},
			{Role: "user", Content: "你好,请介绍一下自己"},
		},
		Temperature: 0.7,
		MaxTokens:   2000,
	}

	// 发送请求
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 打印响应
	fmt.Printf("Model: %s\n", resp.Model)
	fmt.Printf("Response: %s\n", resp.Choices[0].Message.Content)
	fmt.Printf("Tokens Used: %d\n", resp.Usage.TotalTokens)
}
