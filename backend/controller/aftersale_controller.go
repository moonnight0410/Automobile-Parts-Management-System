package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"automobile-parts-backend/model"
	"automobile-parts-backend/service"
	"automobile-parts-backend/utils"
)

type AftersaleController struct {
	service *service.AftersaleService
}

func NewAftersaleController(service *service.AftersaleService) *AftersaleController {
	return &AftersaleController{service: service}
}

func (a *AftersaleController) CreateFaultReport(c *gin.Context) {
	var dto model.FaultReportDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := a.service.CreateFaultReport(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}

func (a *AftersaleController) CreateRecallRecord(c *gin.Context) {
	var dto model.RecallRecordDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := a.service.CreateRecallRecord(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}

func (a *AftersaleController) CreateAftersaleRecord(c *gin.Context) {
	var dto model.AftersaleRecordDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := a.service.CreateAftersaleRecord(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}

func (a *AftersaleController) ListFaultReports(c *gin.Context) {
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	var reporter string
	if role == "aftersale" {
		reporter = userID.(string)
	}

	faults, err := a.service.ListAllFaultReports(c.Request.Context(), reporter)
	if err != nil {
		log.Printf("[ERROR] ListFaultReports failed: %v", err)
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(faults, "ok"))
}

func (a *AftersaleController) ListRecallRecords(c *gin.Context) {
	recalls, err := a.service.ListAllRecallRecords(c.Request.Context())
	if err != nil {
		log.Printf("[ERROR] ListRecallRecords failed: %v", err)
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(recalls, "ok"))
}

func (a *AftersaleController) ListAftersaleRecords(c *gin.Context) {
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	var handlerOrgID string
	if role == "aftersale" {
		handlerOrgID = userID.(string)
	}

	records, err := a.service.ListAllAftersaleRecords(c.Request.Context(), handlerOrgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(records, "ok"))
}

func (a *AftersaleController) UpdateFaultReportStatus(c *gin.Context) {
	var req struct {
		FaultID string `json:"faultID" binding:"required"`
		Status   string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := a.service.UpdateFaultReportStatus(c.Request.Context(), req.FaultID, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}
