package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"automobile-parts-backend/model"
	"automobile-parts-backend/service"
	"automobile-parts-backend/utils"
)

type SupplyChainController struct {
	service *service.SupplyChainService
}

func NewSupplyChainController(service *service.SupplyChainService) *SupplyChainController {
	return &SupplyChainController{service: service}
}

func (s *SupplyChainController) CreateSupplyOrder(c *gin.Context) {
	var dto model.SupplyOrderDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := s.service.CreateSupplyOrder(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}

func (s *SupplyChainController) CreateLogisticsData(c *gin.Context) {
	var dto model.LogisticsDataDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, err.Error()))
		return
	}
	if err := s.service.CreateLogisticsData(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "ok"))
}

func (s *SupplyChainController) ListSupplyOrders(c *gin.Context) {
	// 暂时不过滤 buyer/seller，返回所有订单数据
	orders, err := s.service.ListAllSupplyOrders(c.Request.Context(), "", "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(orders, "ok"))
}

func (s *SupplyChainController) ListLogisticsData(c *gin.Context) {
	logisticsList, err := s.service.ListAllLogisticsData(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(logisticsList, "ok"))
}

func (s *SupplyChainController) DeleteLogisticsData(c *gin.Context) {
	logisticsID := c.Param("id")
	if logisticsID == "" {
		c.JSON(http.StatusBadRequest, utils.Error(http.StatusBadRequest, "物流ID不能为空"))
		return
	}
	if err := s.service.DeleteLogisticsData(c.Request.Context(), logisticsID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Success(nil, "删除成功"))
}
