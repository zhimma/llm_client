package main

import (
	"log"

	"github.com/gin-gonic/gin"
	llmclient "github.com/zhimma/llm_client"
	"github.com/zhimma/llm_client/handler"
)

func main() {
	r := gin.Default()

	// 一行代码完成 LLM 功能集成
	if err := handler.QuickStart(r, &llmclient.Config{
		BaseURL: "http://localhost:8888/v1",
		APIKey:  "sk-your-api-key",
	}); err != nil {
		log.Fatal(err)
	}

	log.Println("LLM Server started on :8080")
	log.Println("Try: curl http://localhost:8080/v1/models")
	r.Run(":8080")
}
