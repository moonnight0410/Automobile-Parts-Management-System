package model

// AIQuestionRequest AI问题请求
type AIQuestionRequest struct {
	Question string `json:"question" binding:"required"`
	UserID   string `json:"userID"`
}

// AIQuestionResponse AI问题响应
type AIQuestionResponse struct {
	QAID           string                 `json:"qaid"`
	Question       string                 `json:"question"`
	Intent         string                 `json:"intent"`
	Confidence     float64                `json:"confidence"`
	Answer         string                 `json:"answer"`
	RelatedActions []RelatedAction         `json:"relatedActions"`
	RelatedFAQs    []string               `json:"relatedFaqs"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

// RelatedAction 相关操作
type RelatedAction struct {
	Label  string                 `json:"label"`
	Action string                 `json:"action"`
	Params map[string]interface{} `json:"params,omitempty"`
}

// ConversationMessage 对话消息
type ConversationMessage struct {
	Role      string                 `json:"role"`
	Content   string                 `json:"content"`
	Timestamp int64                  `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// ConversationHistory 对话历史
type ConversationHistory struct {
	Messages []ConversationMessage `json:"messages"`
}

// AIStats AI统计信息
type AIStats struct {
	ConversationCount int `json:"conversationCount"`
	UserCount        int `json:"userCount"`
}

// AIHealth AI健康检查
type AIHealth struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}
