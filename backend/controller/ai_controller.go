package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"automobile-parts-backend/model"
	"automobile-parts-backend/service"
	"automobile-parts-backend/utils"
)

type AIController struct {
	service *service.AIService
}

func NewAIController(service *service.AIService) *AIController {
	return &AIController{service: service}
}

// AskQuestion 提问
// @Summary AI智能问答
// @Description 提交问题给AI助手，获取智能回答
// @Tags AI
// @Accept json
// @Produce json
// @Param request body model.AIQuestionRequest true "问题请求"
// @Success 200 {object} utils.Response{data=model.AIQuestionResponse}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/ai/question [post]
func (a *AIController) AskQuestion(c *gin.Context) {
	var req model.AIQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "参数错误: "+err.Error()))
		return
	}

	// 验证问题
	if req.Question == "" {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "问题不能为空"))
		return
	}

	// 设置默认用户ID
	if req.UserID == "" {
		req.UserID = "anonymous"
	}

	// 调用AI服务
	resp, err := a.service.AskQuestion(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, "AI服务错误: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Success(resp, "查询成功"))
}

// GetConversation 获取对话历史
// @Summary 获取对话历史
// @Description 获取用户的对话历史记录
// @Tags AI
// @Accept json
// @Produce json
// @Param userID query string true "用户ID"
// @Param n query int false "返回最近n条消息" default(5)
// @Success 200 {object} utils.Response{data=model.ConversationHistory}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/ai/conversation [get]
func (a *AIController) GetConversation(c *gin.Context) {
	userID := c.Query("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "用户ID不能为空"))
		return
	}

	n := 5
	if nStr := c.Query("n"); nStr != "" {
		if err := utils.ParseInt(nStr, &n); err != nil {
			c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "参数n错误"))
			return
		}
	}

	// 调用AI服务
	resp, err := a.service.GetConversation(c.Request.Context(), userID, n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, "AI服务错误: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Success(resp, "查询成功"))
}

// ClearConversation 清空对话
// @Summary 清空对话
// @Description 清空用户的对话历史
// @Tags AI
// @Accept json
// @Produce json
// @Param userID query string true "用户ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/ai/conversation [delete]
func (a *AIController) ClearConversation(c *gin.Context) {
	userID := c.Query("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "用户ID不能为空"))
		return
	}

	// 调用AI服务
	err := a.service.ClearConversation(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, "AI服务错误: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Success(nil, "清空成功"))
}

// GetStats 获取统计信息
// @Summary 获取统计信息
// @Description 获取AI服务的统计信息
// @Tags AI
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response{data=model.AIStats}
// @Failure 500 {object} utils.Response
// @Router /api/ai/stats [get]
func (a *AIController) GetStats(c *gin.Context) {
	// 调用AI服务
	resp, err := a.service.GetStats(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, "AI服务错误: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Success(resp, "查询成功"))
}

// HealthCheck 健康检查
// @Summary AI服务健康检查
// @Description 检查AI服务是否正常运行
// @Tags AI
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response{data=model.AIHealth}
// @Failure 500 {object} utils.Response
// @Router /api/ai/health [get]
func (a *AIController) HealthCheck(c *gin.Context) {
	// 调用AI服务
	resp, err := a.service.HealthCheck(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, utils.Error(http.StatusServiceUnavailable, "AI服务不可用: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Success(resp, "健康"))
}
