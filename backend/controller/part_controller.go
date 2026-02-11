package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"automobile-parts-backend/model"
	"automobile-parts-backend/service"
	"automobile-parts-backend/utils"
)

type PartController struct {
	service *service.PartService
}

func NewPartController(service *service.PartService) *PartController {
	return &PartController{service: service}
}

func (p *PartController) CreatePart(c *gin.Context) {
	var dto model.PartDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := p.service.CreatePart(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}

func (p *PartController) GetPart(c *gin.Context) {
	partID := c.Param("id")
	if partID == "" {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "missing partID"))
		return
	}
	part, err := p.service.QueryPart(c.Request.Context(), partID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(part, "ok"))
}

func (p *PartController) ListParts(c *gin.Context) {
	batchNo := c.Query("batchNo")
	vin := c.Query("vin")
	if batchNo == "" && vin == "" {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "missing query parameter"))
		return
	}
	if batchNo != "" {
		parts, err := p.service.QueryPartByBatchNo(c.Request.Context(), batchNo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
			return
		}
		c.JSON(http.StatusOK, utils.Success(parts, "ok"))
		return
	}
	parts, err := p.service.QueryPartByVIN(c.Request.Context(), vin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(parts, "ok"))
}
