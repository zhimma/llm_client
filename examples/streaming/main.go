package main

import (
	"context"
	"fmt"
	"io"
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

	// 创建流式对话请求
	req := types.ChatCompletionRequest{
		Model: "qwen-max",
		Messages: []types.Message{
			{Role: "user", Content: "请讲一个关于人工智能的故事"},
		},
		Stream: true,
	}

	// 发送流式请求
	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer stream.Close()

	fmt.Println("AI: ")

	// 接收流式响应
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("\n\n[完成]")
			break
		}
		if err != nil {
			log.Fatalf("Stream error: %v", err)
		}

		// 打印增量内容
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta != nil {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}
}
