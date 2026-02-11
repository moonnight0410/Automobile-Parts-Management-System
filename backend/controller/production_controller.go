package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"automobile-parts-backend/model"
	"automobile-parts-backend/service"
	"automobile-parts-backend/utils"
)

type ProductionController struct {
	service *service.ProductionService
}

func NewProductionController(service *service.ProductionService) *ProductionController {
	return &ProductionController{service: service}
}

func (p *ProductionController) CreateProductionData(c *gin.Context) {
	var dto model.ProductionDataDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := p.service.CreateProductionData(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}
