package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"automobile-parts-backend/model"
	"automobile-parts-backend/service"
	"automobile-parts-backend/utils"
)

type BOMController struct {
	service *service.BOMService
}

func NewBOMController(service *service.BOMService) *BOMController {
	return &BOMController{service: service}
}

func (b *BOMController) CreateBOM(c *gin.Context) {
	var dto model.BOMDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := b.service.CreateBOM(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}

func (b *BOMController) GetBOM(c *gin.Context) {
	bomID := c.Param("id")
	if bomID == "" {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "missing bomID"))
		return
	}
	bom, err := b.service.QueryBOM(c.Request.Context(), bomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(bom, "ok"))
}

func (b *BOMController) ListBOMs(c *gin.Context) {
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")
	
	var creator string
	if role == "manufacturer" {
		creator = userID.(string)
	}
	
	boms, err := b.service.ListAllBOMs(c.Request.Context(), creator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(boms, "ok"))
}
