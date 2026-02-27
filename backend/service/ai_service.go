package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"automobile-parts-backend/model"
)

// AIService AI服务
type AIService struct {
	client       *http.Client
	baseURL      string
	timeout      time.Duration
	maxRetries   int
	retryDelay   time.Duration
}

// NewAIService 创建AI服务
func NewAIService(baseURL string) *AIService {
	return &AIService{
		client: &http.Client{
			Timeout: 120 * time.Second,
		},
		baseURL:    baseURL,
		timeout:    120 * time.Second,
		maxRetries: 3,
		retryDelay: 5 * time.Second,
	}
}

// AskQuestion 提问
func (a *AIService) AskQuestion(ctx context.Context, req model.AIQuestionRequest) (*model.AIQuestionResponse, error) {
	url := fmt.Sprintf("%s/api/ai/question", a.baseURL)
	
	var resp struct {
		Success bool                    `json:"success"`
		Message string                  `json:"message"`
		Data    model.AIQuestionResponse `json:"data"`
	}
	
	err := a.doRequest(ctx, "POST", url, req, &resp)
	if err != nil {
		return nil, fmt.Errorf("请求AI服务失败: %w", err)
	}
	
	if !resp.Success {
		return nil, fmt.Errorf("AI服务返回错误: %s", resp.Message)
	}
	
	return &resp.Data, nil
}

// GetConversation 获取对话历史
func (a *AIService) GetConversation(ctx context.Context, userID string, n int) (*model.ConversationHistory, error) {
	url := fmt.Sprintf("%s/api/ai/conversation?userID=%s&n=%d", a.baseURL, userID, n)
	
	var resp struct {
		Success bool                      `json:"success"`
		Message string                    `json:"message"`
		Data    model.ConversationHistory `json:"data"`
	}
	
	err := a.doRequest(ctx, "GET", url, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("请求AI服务失败: %w", err)
	}
	
	if !resp.Success {
		return nil, fmt.Errorf("AI服务返回错误: %s", resp.Message)
	}
	
	return &resp.Data, nil
}

// ClearConversation 清空对话
func (a *AIService) ClearConversation(ctx context.Context, userID string) error {
	url := fmt.Sprintf("%s/api/ai/conversation?userID=%s", a.baseURL, userID)
	
	var resp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	
	err := a.doRequest(ctx, "DELETE", url, nil, &resp)
	if err != nil {
		return fmt.Errorf("请求AI服务失败: %w", err)
	}
	
	if !resp.Success {
		return fmt.Errorf("AI服务返回错误: %s", resp.Message)
	}
	
	return nil
}

// GetStats 获取统计信息
func (a *AIService) GetStats(ctx context.Context) (*model.AIStats, error) {
	url := fmt.Sprintf("%s/api/ai/stats", a.baseURL)
	
	var resp struct {
		Success bool         `json:"success"`
		Message string       `json:"message"`
		Data    model.AIStats `json:"data"`
	}
	
	err := a.doRequest(ctx, "GET", url, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("请求AI服务失败: %w", err)
	}
	
	if !resp.Success {
		return nil, fmt.Errorf("AI服务返回错误: %s", resp.Message)
	}
	
	return &resp.Data, nil
}

// HealthCheck 健康检查
func (a *AIService) HealthCheck(ctx context.Context) (*model.AIHealth, error) {
	url := fmt.Sprintf("%s/health", a.baseURL)
	
	var resp model.AIHealth
	
	err := a.doRequest(ctx, "GET", url, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("请求AI服务失败: %w", err)
	}
	
	return &resp, nil
}

// doRequest 执行HTTP请求（带重试）
func (a *AIService) doRequest(ctx context.Context, method, url string, body interface{}, result interface{}) error {
	var bodyReader io.Reader
	
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("序列化请求体失败: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}
	
	var lastErr error
	
	for i := 0; i < a.maxRetries; i++ {
		// 创建请求
		req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
		if err != nil {
			return fmt.Errorf("创建请求失败: %w", err)
		}
		
		// 设置请求头
		req.Header.Set("Content-Type", "application/json")
		
		// 发送请求
		resp, err := a.client.Do(req)
		if err != nil {
			lastErr = err
			log.Printf("请求失败（第%d次重试）: %v", i+1, err)
			time.Sleep(a.retryDelay)
			continue
		}
		
		// 读取响应
		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			log.Printf("读取响应失败（第%d次重试）: %v", i+1, err)
			time.Sleep(a.retryDelay)
			continue
		}
		
		// 检查HTTP状态码
		if resp.StatusCode >= 400 {
			lastErr = fmt.Errorf("HTTP错误: %s", string(respBody))
			if resp.StatusCode >= 500 {
				log.Printf("服务器错误（第%d次重试）: %v", i+1, lastErr)
				time.Sleep(a.retryDelay)
				continue
			}
			return lastErr
		}
		
		// 解析响应
		if result != nil {
			if err := json.Unmarshal(respBody, result); err != nil {
				return fmt.Errorf("解析响应失败: %w", err)
			}
		}
		
		return nil
	}
	
	return fmt.Errorf("请求失败，已重试%d次: %w", a.maxRetries, lastErr)
}
