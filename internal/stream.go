package internal

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zhimma/llm_client/types"
)

// StreamReader SSE 流式读取器
type StreamReader struct {
	resp    *http.Response
	scanner *bufio.Scanner
	ctx     context.Context
}

// NewStreamReader 创建流式读取器
func NewStreamReader(ctx context.Context, resp *http.Response) *StreamReader {
	return &StreamReader{
		resp:    resp,
		scanner: bufio.NewScanner(resp.Body),
		ctx:     ctx,
	}
}

// Recv 接收下一个流式响应块
func (s *StreamReader) Recv() (*types.ChatCompletionStreamResponse, error) {
	for s.scanner.Scan() {
		line := s.scanner.Text()

		// 跳过空行和注释
		if line == "" || strings.HasPrefix(line, ":") {
			continue
		}

		// 解析 SSE 数据
		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")

			// 检查结束标记
			if data == "[DONE]" {
				return nil, io.EOF
			}

			// 解析 JSON
			var chunk types.ChatCompletionStreamResponse
			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				return nil, fmt.Errorf("failed to parse stream chunk: %w", err)
			}

			return &chunk, nil
		}
	}

	if err := s.scanner.Err(); err != nil {
		return nil, fmt.Errorf("stream scanner error: %w", err)
	}

	return nil, io.EOF
}

// Close 关闭流
func (s *StreamReader) Close() error {
	if s.resp != nil && s.resp.Body != nil {
		return s.resp.Body.Close()
	}
	return nil
}
