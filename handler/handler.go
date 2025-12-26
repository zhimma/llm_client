package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	llmclient "github.com/zhimma/llm_client"
	"github.com/zhimma/llm_client/service"
	"github.com/zhimma/llm_client/types"
)

// LLMHandler LLM HTTP 处理器
type LLMHandler struct {
	service service.LLMService
}

// NewHandler 创建 Handler
func NewHandler(svc service.LLMService) *LLMHandler {
	return &LLMHandler{
		service: svc,
	}
}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.RouterGroup, svc service.LLMService) {
	h := NewHandler(svc)

	// OpenAI 兼容路由
	r.POST("/chat/completions", h.ChatCompletions)
	r.POST("/chat", h.ChatCompletions) // 兼容旧路径

	// Models
	r.GET("/models", h.ListModels)
	r.GET("/models/:model", h.GetModel)

	// 扩展功能
	r.GET("/providers", h.ListProviders)
}

// ChatCompletions 处理聊天请求
func (h *LLMHandler) ChatCompletions(c *gin.Context) {
	var req types.ChatCompletionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// 调用服务
	resp, stream, err := h.service.Chat(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 非流式响应
	if !req.Stream {
		if resp == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Empty response"})
			return
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// 流式响应
	if stream == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Stream is nil"})
		return
	}
	defer stream.Close()

	h.handleStreamResponse(c, stream)
}

// ListModels 获取模型列表
func (h *LLMHandler) ListModels(c *gin.Context) {
	resp, err := h.service.ListModels(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetModel 获取单个模型
func (h *LLMHandler) GetModel(c *gin.Context) {
	modelID := c.Param("model")
	resp, err := h.service.GetModel(c.Request.Context(), modelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListProviders 获取提供商列表
func (h *LLMHandler) ListProviders(c *gin.Context) {
	resp, err := h.service.ListProviders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// handleStreamResponse 处理流式响应 (SSE)
func (h *LLMHandler) handleStreamResponse(c *gin.Context, stream *llmclient.ChatCompletionStream) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	c.Stream(func(w io.Writer) bool {
		chunk, err := stream.Recv()
		if err == io.EOF {
			c.SSEvent("message", "[DONE]")
			return false
		}
		if err != nil {
			c.SSEvent("error", err.Error())
			return false
		}

		// 发送流式数据块
		c.SSEvent("message", chunk)
		return true
	})
}

// QuickStart 快速启动 LLM 服务
// 一行代码完成 LLM 功能集成
//
// 示例:
//
//	r := gin.Default()
//	handler.QuickStart(r, &llmclient.Config{
//	    BaseURL: "http://localhost:8888/v1",
//	    APIKey:  "sk-your-key",
//	})
func QuickStart(r *gin.Engine, cfg *llmclient.Config) error {
	return QuickStartWithGroup(r.Group("/v1"), cfg)
}

// QuickStartWithGroup 在指定的路由组中快速启动
func QuickStartWithGroup(group *gin.RouterGroup, cfg *llmclient.Config) error {
	// 验证配置
	if err := cfg.Validate(); err != nil {
		return err
	}

	// 创建 Client
	client := llmclient.NewClient(cfg)

	// 创建 Service
	svc := service.NewService(client)

	// 注册路由
	RegisterRoutes(group, svc)

	return nil
}
