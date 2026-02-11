package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"automobile-parts-backend/model"
	"automobile-parts-backend/service"
	"automobile-parts-backend/utils"
)

type QualityController struct {
	service *service.QualityService
}

func NewQualityController(service *service.QualityService) *QualityController {
	return &QualityController{service: service}
}

func (q *QualityController) CreateQualityInspection(c *gin.Context) {
	var dto model.QualityInspectionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := q.service.CreateQualityInspection(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}
