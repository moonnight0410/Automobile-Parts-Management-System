package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"automobile-parts-backend/service"

	"github.com/gin-gonic/gin"
)

type FabricController struct {
	fabricService *service.FabricService
}

func NewFabricController(fabricService *service.FabricService) *FabricController {
	return &FabricController{
		fabricService: fabricService,
	}
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Part struct {
	PartID       string `json:"partID"`
	VIN          string `json:"vin"`
	BatchNo      string `json:"batchNo"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	CreateTime   string `json:"createTime"`
	Status       string `json:"status"`
}

type PartLifecycle struct {
	PartID       string   `json:"partID"`
	CurrentStage string   `json:"currentStage"`
	Stages       []string `json:"stages"`
	UpdateTime   string   `json:"updateTime"`
}

type BOM struct {
	BOMID        string         `json:"bomID"`
	BOMType      string         `json:"bomType"`
	ProductModel string         `json:"productModel"`
	PartBatchNo  string         `json:"partBatchNo"`
	Version      string         `json:"version"`
	Creator      string         `json:"creator"`
	CreateTime   string         `json:"createTime"`
	Status       string         `json:"status"`
	MaterialList []MaterialItem `json:"materialList"`
}

type MaterialItem struct {
	MaterialID   string `json:"materialID"`
	MaterialName string `json:"materialName"`
	Spec         string `json:"spec"`
	Quantity     int    `json:"quantity"`
	SupplierID   string `json:"supplierID"`
}

type ProductionData struct {
	ProductionID   string            `json:"productionID"`
	PartID         string            `json:"partID"`
	BatchNo        string            `json:"batchNo"`
	Params         map[string]string `json:"params"`
	ProductionLine string            `json:"productionLine"`
	Operator       string            `json:"operator"`
	FinishTime     string            `json:"finishTime"`
}

type QualityInspection struct {
	InspectionID string            `json:"inspectionID"`
	PartID       string            `json:"partID"`
	BatchNo      string            `json:"batchNo"`
	Indicators   map[string]string `json:"indicators"`
	Result       string            `json:"result"`
	Handler      string            `json:"handler"`
	HandleTime   string            `json:"handleTime"`
	Disposal     string            `json:"disposal"`
}

type SupplyOrder struct {
	OrderID    string `json:"orderID"`
	Buyer      string `json:"buyer"`
	Seller     string `json:"seller"`
	PartID     string `json:"partID"`
	Quantity   int    `json:"quantity"`
	BatchNo    string `json:"batchNo"`
	AgreedTime string `json:"agreedTime"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

type FaultReport struct {
	FaultID     string  `json:"faultID"`
	PartID      string  `json:"partID"`
	VIN         string  `json:"vin"`
	Reporter    string  `json:"reporter"`
	Description string  `json:"description"`
	FaultType   string  `json:"faultType"`
	RiskProb    float64 `json:"riskProb"`
	ReportTime  string  `json:"reportTime"`
	Status      string  `json:"status"`
}

type LogisticsData struct {
	LogisticsID string `json:"logisticsID"`
	OrderID     string `json:"orderID"`
	Carrier     string `json:"carrier"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	GPSHash     string `json:"gpsHash"`
	Receiver    string `json:"receiver"`
}

type Reconciliation struct {
	ReconID    string `json:"reconID"`
	OrderID    string `json:"orderID"`
	Amount     string `json:"amount"`
	Status     string `json:"status"`
	SettleTime string `json:"settleTime"`
}

type SupplyChainStage struct {
	ChainID          string `json:"chainID"`
	PartID           string `json:"partID"`
	BatchNo          string `json:"batchNo"`
	OrderID          string `json:"orderID"`
	LogisticsID      string `json:"logisticsID"`
	StageType        string `json:"stageType"`
	StageStatus      string `json:"stageStatus"`
	Participator     string `json:"participator"`
	ParticipatorRole string `json:"participatorRole"`
	Quantity         int    `json:"quantity"`
	OperateTime      string `json:"operateTime"`
	Operator         string `json:"operator"`
	Remark           string `json:"remark"`
}

type RecallRecord struct {
	RecallID      string   `json:"recallID"`
	BatchNos      []string `json:"batchNos"`
	Reason        string   `json:"reason"`
	AffectedParts []string `json:"affectedParts"`
	Status        string   `json:"status"`
	CreateTime    string   `json:"createTime"`
	FinishTime    string   `json:"finishTime"`
}

type AftersaleRecord struct {
	AftersaleID     string `json:"aftersaleID"`
	PartID          string `json:"partID"`
	BatchNo         string `json:"batchNo"`
	VIN             string `json:"vin"`
	FaultID         string `json:"faultID"`
	RecallID        string `json:"recallID"`
	AftersaleType   string `json:"aftersaleType"`
	AftersaleStatus string `json:"aftersaleStatus"`
	HandlerOrgID    string `json:"handlerOrgID"`
	HandlerID       string `json:"handlerID"`
	Description     string `json:"description"`
	ProcessResult   string `json:"processResult"`
	ProcessTime     string `json:"processTime"`
	Cost            string `json:"cost"`
	Remark          string `json:"remark"`
}

type UserIdentity struct {
	UserID       string   `json:"userID"`
	Org          string   `json:"org"`
	Role         string   `json:"role"`
	CertHash     string   `json:"certHash"`
	Permissions  []string `json:"permissions"`
	RegisterTime string   `json:"registerTime"`
}

type QAInteraction struct {
	QAID           string `json:"qaid"`
	UserID         string `json:"userID"`
	Question       string `json:"question"`
	Intent         string `json:"intent"`
	PartID         string `json:"partID"`
	Answer         string `json:"answer"`
	QueryTime      string `json:"queryTime"`
	ContractMethod string `json:"contractMethod"`
}

type BOMCompareRequest struct {
	ProductionBOMID string `json:"productionBOMID"`
	RDBOMID         string `json:"rdBOMID"`
}

type BOMChangeRequest struct {
	BOMID        string `json:"bomID"`
	ChangeID     string `json:"changeID"`
	ChangeReason string `json:"changeReason"`
	OldVersion   string `json:"oldVersion"`
	NewVersion   string `json:"newVersion"`
	ChangeTime   string `json:"changeTime"`
}

func (fc *FabricController) HealthCheck(c *gin.Context) {
	log.Println("[FabricController] 收到健康检查请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "unhealthy",
			"message": "Fabric服务未初始化",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	_, err := fc.fabricService.Query(ctx, "QueryPart", "TEST")
	if err != nil {
		log.Printf("[FabricController] 健康检查失败: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "unhealthy",
			"message": "Fabric服务不可用: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Fabric服务运行正常",
	})
}

func (fc *FabricController) InitLedger(c *gin.Context) {
	log.Println("[FabricController] 收到初始化账本的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "InitLedger")
	if err != nil {
		log.Printf("[FabricController] 初始化账本失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "初始化账本失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 初始化账本成功")
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "初始化账本成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreatePart(c *gin.Context) {
	log.Println("[FabricController] 收到创建零部件的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req Part
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.CreateTime == "" {
		req.CreateTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	partJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreatePart", string(partJSON))
	if err != nil {
		log.Printf("[FabricController] 创建零部件失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建零部件失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建零部件成功，PartID: %s", req.PartID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建零部件成功",
		Data:    string(result),
	})
}

func (fc *FabricController) QueryPart(c *gin.Context) {
	partID := c.Param("id")
	log.Printf("[FabricController] 收到查询零部件的请求，PartID: %s", partID)

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryPart", partID)
	if err != nil {
		log.Printf("[FabricController] 查询零部件失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "查询零部件失败: " + err.Error(),
		})
		return
	}

	var part Part
	if err := json.Unmarshal(result, &part); err != nil {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "查询零部件成功",
			Data:    string(result),
		})
		return
	}

	log.Printf("[FabricController] 查询零部件成功，PartID: %s", partID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "查询零部件成功",
		Data:    part,
	})
}

func (fc *FabricController) QueryPartLifecycle(c *gin.Context) {
	partID := c.Param("id")
	log.Printf("[FabricController] 收到查询零部件生命周期的请求，PartID: %s", partID)

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryPartLifecycle", partID)
	if err != nil {
		log.Printf("[FabricController] 查询零部件生命周期失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "查询零部件生命周期失败: " + err.Error(),
		})
		return
	}

	var lifecycle PartLifecycle
	if err := json.Unmarshal(result, &lifecycle); err != nil {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "查询零部件生命周期成功",
			Data:    string(result),
		})
		return
	}

	log.Printf("[FabricController] 查询零部件生命周期成功，PartID: %s", partID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "查询零部件生命周期成功",
		Data:    lifecycle,
	})
}

func (fc *FabricController) QueryPartByBatchNo(c *gin.Context) {
	batchNo := c.Param("batchNo")
	log.Printf("[FabricController] 收到按批次号查询零部件的请求，BatchNo: %s", batchNo)

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryPartByBatchNo", batchNo)
	if err != nil {
		log.Printf("[FabricController] 按批次号查询零部件失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "按批次号查询零部件失败: " + err.Error(),
		})
		return
	}

	var parts []Part
	if err := json.Unmarshal(result, &parts); err != nil {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "按批次号查询零部件成功",
			Data:    string(result),
		})
		return
	}

	log.Printf("[FabricController] 按批次号查询零部件成功，BatchNo: %s，数量: %d", batchNo, len(parts))
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "按批次号查询零部件成功",
		Data:    parts,
	})
}

func (fc *FabricController) QueryPartByVIN(c *gin.Context) {
	vin := c.Param("vin")
	log.Printf("[FabricController] 收到按VIN查询零部件的请求，VIN: %s", vin)

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryPartByVIN", vin)
	if err != nil {
		log.Printf("[FabricController] 按VIN查询零部件失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "按VIN查询零部件失败: " + err.Error(),
		})
		return
	}

	var parts []Part
	if err := json.Unmarshal(result, &parts); err != nil {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "按VIN查询零部件成功",
			Data:    string(result),
		})
		return
	}

	log.Printf("[FabricController] 按VIN查询零部件成功，VIN: %s，数量: %d", vin, len(parts))
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "按VIN查询零部件成功",
		Data:    parts,
	})
}

func (fc *FabricController) UpdatePartStatus(c *gin.Context) {
	partID := c.Param("id")
	log.Printf("[FabricController] 收到更新零部件状态的请求，PartID: %s", partID)

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "UpdatePartStatus", partID, req.Status)
	if err != nil {
		log.Printf("[FabricController] 更新零部件状态失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "更新零部件状态失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 更新零部件状态成功，PartID: %s，Status: %s", partID, req.Status)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "更新零部件状态成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateBOM(c *gin.Context) {
	log.Println("[FabricController] 收到创建BOM的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req BOM
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.CreateTime == "" {
		req.CreateTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	bomJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateBOM", string(bomJSON))
	if err != nil {
		log.Printf("[FabricController] 创建BOM失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建BOM失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建BOM成功，BOMID: %s", req.BOMID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建BOM成功",
		Data:    string(result),
	})
}

func (fc *FabricController) QueryBOM(c *gin.Context) {
	bomID := c.Param("id")
	log.Printf("[FabricController] 收到查询BOM的请求，BOMID: %s", bomID)

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryBOM", bomID)
	if err != nil {
		log.Printf("[FabricController] 查询BOM失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "查询BOM失败: " + err.Error(),
		})
		return
	}

	var bom BOM
	if err := json.Unmarshal(result, &bom); err != nil {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "查询BOM成功",
			Data:    string(result),
		})
		return
	}

	log.Printf("[FabricController] 查询BOM成功，BOMID: %s", bomID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "查询BOM成功",
		Data:    bom,
	})
}

func (fc *FabricController) UpdateBOM(c *gin.Context) {
	bomID := c.Param("id")
	log.Printf("[FabricController] 收到更新BOM的请求，BOMID: %s", bomID)

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req BOM
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	req.BOMID = bomID

	bomJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "UpdateBOM", string(bomJSON))
	if err != nil {
		log.Printf("[FabricController] 更新BOM失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "更新BOM失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 更新BOM成功，BOMID: %s", bomID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "更新BOM成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateProductionData(c *gin.Context) {
	log.Println("[FabricController] 收到创建生产数据的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req ProductionData
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.FinishTime == "" {
		req.FinishTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	productionJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateProductionData", string(productionJSON))
	if err != nil {
		log.Printf("[FabricController] 创建生产数据失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建生产数据失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建生产数据成功，ProductionID: %s", req.ProductionID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建生产数据成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateQualityInspection(c *gin.Context) {
	log.Println("[FabricController] 收到创建质检数据的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req QualityInspection
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.HandleTime == "" {
		req.HandleTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	qualityJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateQualityInspection", string(qualityJSON))
	if err != nil {
		log.Printf("[FabricController] 创建质检数据失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建质检数据失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建质检数据成功，InspectionID: %s", req.InspectionID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建质检数据成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateFaultReport(c *gin.Context) {
	log.Println("[FabricController] 收到创建故障报告的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req FaultReport
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.ReportTime == "" {
		req.ReportTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	faultJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateFaultReport", string(faultJSON))
	if err != nil {
		log.Printf("[FabricController] 创建故障报告失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建故障报告失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建故障报告成功，FaultID: %s", req.FaultID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建故障报告成功",
		Data:    string(result),
	})
}

func (fc *FabricController) QueryFaultProgress(c *gin.Context) {
	faultID := c.Param("id")
	log.Printf("[FabricController] 收到查询故障进度的请求，FaultID: %s", faultID)

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryFaultProgress", faultID)
	if err != nil {
		log.Printf("[FabricController] 查询故障进度失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "查询故障进度失败: " + err.Error(),
		})
		return
	}

	var fault FaultReport
	if err := json.Unmarshal(result, &fault); err != nil {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "查询故障进度成功",
			Data:    string(result),
		})
		return
	}

	log.Printf("[FabricController] 查询故障进度成功，FaultID: %s", faultID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "查询故障进度成功",
		Data:    fault,
	})
}

func (fc *FabricController) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/fabric")
	{
		api.GET("/health", fc.HealthCheck)

		api.POST("/init", fc.InitLedger)

		api.POST("/parts", fc.CreatePart)
		api.GET("/parts/:id", fc.QueryPart)
		api.GET("/parts/:id/lifecycle", fc.QueryPartLifecycle)
		api.GET("/parts/batch/:batchNo", fc.QueryPartByBatchNo)
		api.GET("/parts/vin/:vin", fc.QueryPartByVIN)
		api.PUT("/parts/:id/status", fc.UpdatePartStatus)

		api.POST("/bom", fc.CreateBOM)
		api.GET("/bom/:id", fc.QueryBOM)
		api.PUT("/bom/:id", fc.UpdateBOM)
		api.POST("/bom/compare", fc.CompareBOM)
		api.POST("/bom/change", fc.SubmitBOMChange)

		api.POST("/production", fc.CreateProductionData)

		api.POST("/quality", fc.CreateQualityInspection)

		api.POST("/supply/orders", fc.CreateSupplyOrder)
		api.POST("/supply/logistics", fc.CreateLogisticsData)
		api.PUT("/supply/stage", fc.UpdateSupplyChainStage)
		api.POST("/supply/reconciliation", fc.CreateReconciliation)

		api.POST("/faults", fc.CreateFaultReport)
		api.GET("/faults/:id", fc.QueryFaultProgress)

		api.POST("/recalls", fc.CreateRecallRecord)
		api.GET("/recalls/affected/:batchNo", fc.QueryAffectedParts)

		api.POST("/aftersales", fc.CreateAftersaleRecord)
		api.GET("/aftersales/:id", fc.QueryAftersaleRecord)

		api.POST("/users", fc.RegisterUser)
		api.GET("/users/:id/permissions", fc.QueryUserPermissions)

		api.POST("/qa", fc.RecordQAInteraction)
		api.GET("/qa/:id", fc.QueryQAInteraction)
	}
	log.Println("[FabricController] Fabric路由已注册")
}

func (fc *FabricController) CompareBOM(c *gin.Context) {
	log.Println("[FabricController] 收到BOM比较的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req BOMCompareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "CompareBOM", req.ProductionBOMID, req.RDBOMID)
	if err != nil {
		log.Printf("[FabricController] BOM比较失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "BOM比较失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] BOM比较成功，ProductionBOMID: %s, RDBOMID: %s", req.ProductionBOMID, req.RDBOMID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "BOM比较成功",
		Data:    string(result),
	})
}

func (fc *FabricController) SubmitBOMChange(c *gin.Context) {
	log.Println("[FabricController] 收到提交BOM变更的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req BOMChangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.ChangeTime == "" {
		req.ChangeTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	changeJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "SubmitBOMChange", req.BOMID, string(changeJSON))
	if err != nil {
		log.Printf("[FabricController] 提交BOM变更失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "提交BOM变更失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 提交BOM变更成功，BOMID: %s", req.BOMID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "提交BOM变更成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateSupplyOrder(c *gin.Context) {
	log.Println("[FabricController] 收到创建供应链订单的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req SupplyOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.CreateTime == "" {
		req.CreateTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	orderJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateSupplyOrder", string(orderJSON))
	if err != nil {
		log.Printf("[FabricController] 创建供应链订单失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建供应链订单失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建供应链订单成功，OrderID: %s", req.OrderID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建供应链订单成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateLogisticsData(c *gin.Context) {
	log.Println("[FabricController] 收到创建物流数据的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req LogisticsData
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	logisticsJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateLogisticsData", string(logisticsJSON))
	if err != nil {
		log.Printf("[FabricController] 创建物流数据失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建物流数据失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建物流数据成功，LogisticsID: %s", req.LogisticsID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建物流数据成功",
		Data:    string(result),
	})
}

func (fc *FabricController) UpdateSupplyChainStage(c *gin.Context) {
	log.Println("[FabricController] 收到更新供应链阶段的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req SupplyChainStage
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.OperateTime == "" {
		req.OperateTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	stageJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "UpdateSupplyChainStage", string(stageJSON))
	if err != nil {
		log.Printf("[FabricController] 更新供应链阶段失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "更新供应链阶段失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 更新供应链阶段成功，ChainID: %s", req.ChainID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "更新供应链阶段成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateReconciliation(c *gin.Context) {
	log.Println("[FabricController] 收到创建对账的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req Reconciliation
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	reconJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateReconciliation", string(reconJSON))
	if err != nil {
		log.Printf("[FabricController] 创建对账失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建对账失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建对账成功，ReconID: %s", req.ReconID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建对账成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateRecallRecord(c *gin.Context) {
	log.Println("[FabricController] 收到创建召回记录的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req RecallRecord
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.CreateTime == "" {
		req.CreateTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	recallJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateRecallRecord", string(recallJSON))
	if err != nil {
		log.Printf("[FabricController] 创建召回记录失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建召回记录失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建召回记录成功，RecallID: %s", req.RecallID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建召回记录成功",
		Data:    string(result),
	})
}

func (fc *FabricController) CreateAftersaleRecord(c *gin.Context) {
	log.Println("[FabricController] 收到创建售后记录的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req AftersaleRecord
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.ProcessTime == "" {
		req.ProcessTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	aftersaleJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "CreateAftersaleRecord", string(aftersaleJSON))
	if err != nil {
		log.Printf("[FabricController] 创建售后记录失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "创建售后记录失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 创建售后记录成功，AftersaleID: %s", req.AftersaleID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "创建售后记录成功",
		Data:    string(result),
	})
}

func (fc *FabricController) QueryAffectedParts(c *gin.Context) {
	log.Println("[FabricController] 收到查询受影响零部件的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	batchNo := c.Param("batchNo")
	if batchNo == "" {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "批次号不能为空",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryAffectedParts", batchNo)
	if err != nil {
		log.Printf("[FabricController] 查询受影响零部件失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "查询受影响零部件失败: " + err.Error(),
		})
		return
	}

	var parts []Part
	if err := json.Unmarshal(result, &parts); err != nil {
		log.Printf("[FabricController] 解析零部件列表失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "解析零部件列表失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 查询受影响零部件成功，BatchNo: %s，数量: %d", batchNo, len(parts))
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "查询受影响零部件成功",
		Data:    parts,
	})
}

func (fc *FabricController) QueryAftersaleRecord(c *gin.Context) {
	log.Println("[FabricController] 收到查询售后记录的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	aftersaleID := c.Param("id")
	if aftersaleID == "" {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "售后记录ID不能为空",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryAftersaleRecord", aftersaleID)
	if err != nil {
		log.Printf("[FabricController] 查询售后记录失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "查询售后记录失败: " + err.Error(),
		})
		return
	}

	var aftersaleRecord AftersaleRecord
	if err := json.Unmarshal(result, &aftersaleRecord); err != nil {
		log.Printf("[FabricController] 解析售后记录失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "解析售后记录失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 查询售后记录成功，AftersaleID: %s", aftersaleID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "查询售后记录成功",
		Data:    aftersaleRecord,
	})
}

func (fc *FabricController) RegisterUser(c *gin.Context) {
	log.Println("[FabricController] 收到用户注册的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req UserIdentity
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.RegisterTime == "" {
		req.RegisterTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	userJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "RegisterUser", string(userJSON))
	if err != nil {
		log.Printf("[FabricController] 用户注册失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "用户注册失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 用户注册成功，UserID: %s", req.UserID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "用户注册成功",
		Data:    string(result),
	})
}

func (fc *FabricController) QueryUserPermissions(c *gin.Context) {
	log.Println("[FabricController] 收到查询用户权限的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "用户ID不能为空",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryUserPermissions", userID)
	if err != nil {
		log.Printf("[FabricController] 查询用户权限失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "查询用户权限失败: " + err.Error(),
		})
		return
	}

	var userIdentity UserIdentity
	if err := json.Unmarshal(result, &userIdentity); err != nil {
		log.Printf("[FabricController] 解析用户身份失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "解析用户身份失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 查询用户权限成功，UserID: %s", userID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "查询用户权限成功",
		Data:    userIdentity,
	})
}

func (fc *FabricController) RecordQAInteraction(c *gin.Context) {
	log.Println("[FabricController] 收到记录QA交互的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	var req QAInteraction
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[FabricController] 请求参数解析失败: %v", err)
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if req.QueryTime == "" {
		req.QueryTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	qaJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("[FabricController] JSON序列化失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "JSON序列化失败: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Submit(ctx, "RecordQAInteraction", string(qaJSON))
	if err != nil {
		log.Printf("[FabricController] 记录QA交互失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "记录QA交互失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 记录QA交互成功，QAID: %s", req.QAID)
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "记录QA交互成功",
		Data:    string(result),
	})
}

func (fc *FabricController) QueryQAInteraction(c *gin.Context) {
	log.Println("[FabricController] 收到查询QA交互的请求")

	if fc.fabricService == nil {
		c.JSON(http.StatusServiceUnavailable, Response{
			Success: false,
			Message: "Fabric服务未初始化",
		})
		return
	}

	qaID := c.Param("id")
	if qaID == "" {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "QA交互ID不能为空",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	result, err := fc.fabricService.Query(ctx, "QueryQAInteraction", qaID)
	if err != nil {
		log.Printf("[FabricController] 查询QA交互失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "查询QA交互失败: " + err.Error(),
		})
		return
	}

	var qaInteraction QAInteraction
	if err := json.Unmarshal(result, &qaInteraction); err != nil {
		log.Printf("[FabricController] 解析QA交互失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "解析QA交互失败: " + err.Error(),
		})
		return
	}

	log.Printf("[FabricController] 查询QA交互成功，QAID: %s", qaID)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "查询QA交互成功",
		Data:    qaInteraction,
	})
}
