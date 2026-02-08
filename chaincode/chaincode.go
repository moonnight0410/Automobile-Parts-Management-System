package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/msp"
)

// 全局变量：测试模式标志，用于单元测试时跳过真实的身份验证
var TestMode = false

// Part 零部件核心信息
type Part struct {
	PartID       string `json:"partID"`       // 零部件唯一标识（如"ENG-PISTON-202505"）
	VIN          string `json:"vin"`          // 关联车辆VIN码（如有）
	BatchNo      string `json:"batchNo"`      // 生产批次号
	Name         string `json:"name"`         // 零部件名称（如"发动机活塞"）
	Type         string `json:"type"`         // 类型（如"发动机部件"）
	Manufacturer string `json:"manufacturer"` // 生产厂商ID（关联身份管理）
	CreateTime   string `json:"createTime"`   // 创建时间戳（上链时间）
	Status       string `json:"status"`       // 状态（在产/在售/召回/报废等）
}

// PartLifecycle 零部件全生命周期轨迹
type PartLifecycle struct {
	PartID          string             `json:"partID"`          // 关联零部件
	BOMInfo         *BOMReference      `json:"bomInfo"`         // 关联BOM信息
	ProductionInfo  *ProductionData    `json:"productionInfo"`  // 关联生产数据
	QualityInfo     *QualityInspection `json:"qualityInfo"`     // 关联质检数据
	SupplyChainInfo []SupplyChainData  `json:"supplyChainInfo"` // 关联供应链数据（订单+物流）
	AftersaleInfo   []AftersaleRecord  `json:"aftersaleInfo"`   // 关联售后/故障数据
}

// BOMReference BOM信息引用
type BOMReference struct {
	BOMID   string `json:"bomID"`   // 关联BOM的ID
	Version string `json:"version"` // 关联BOM版本
	Type    string `json:"type"`    // 研发BOM/生产BOM
}

// BOM 物料清单主结构体（链上存储核心数据）
type BOM struct {
	// 基础标识字段
	BOMID        string `json:"bomID"`        // BOM唯一ID（如研发BOM：R-BOM-202505-001；生产BOM：P-BOM-202505-001）
	BOMType      string `json:"bomType"`      // BOM类型：研发BOM/生产BOM
	ProductModel string `json:"productModel"` // 对应车型（如2025款XX车型）
	PartBatchNo  string `json:"partBatchNo"`  // 零部件批次号（关联溯源模块）
	Version      string `json:"version"`      // BOM版本号（如V1.0、V1.1）
	Creator      string `json:"creator"`      // 创建人（企业数字身份标识）
	CreateTime   string `json:"createTime"`   // 创建时间（时间戳）
	Status       string `json:"status"`       // 状态：草稿/已发布/已变更/已作废

	// 物料清单核心字段
	MaterialList []MaterialItem `json:"materialList"` // 物料明细列表

	// 研发BOM专属：图纸/设计文件关联（适配普通数据库存储）
	RDBOMFileInfo *BOMFileInfo `json:"rdBOMFileInfo"` // 研发BOM图纸文件信息（生产BOM可为空对象）

	// 生产BOM专属：与研发BOM比对
	ProductionBOMInfo *ProductionBOMInfo `json:"productionBOMInfo"` // 生产BOM比对信息（研发BOM可为空对象）

	// BOM变更相关
	ChangeRecords []BOMChangeRecord `json:"changeRecords"` // 变更记录（版本迭代/变更申请）
}

// MaterialItem 物料明细项
type MaterialItem struct {
	MaterialID   string `json:"materialID"`   // 物料唯一ID
	MaterialName string `json:"materialName"` // 物料名称（如发动机活塞）
	Spec         string `json:"spec"`         // 规格型号
	Quantity     int    `json:"quantity"`     // 数量
	SupplierID   string `json:"supplierID"`   // 供应商ID（关联供应链模块）
}

// BOMFileInfo 研发BOM图纸/文件信息（适配普通数据库存储）
type BOMFileInfo struct {
	FileStorageID  string `json:"fileStorageID"`  // 普通数据库中的文件唯一ID（如MySQL的主键ID）
	FilePath       string `json:"filePath"`       // 文件存储路径（如：/data/bom/rd/202505/R-BOM-001-CAD.dwg）
	FileHash       string `json:"fileHash"`       // 文件内容哈希（SHA256，链上存哈希，验证数据库文件未被篡改）
	FileType       string `json:"fileType"`       // 文件类型（如CAD图纸、PNG图片、Excel清单）
	FileSize       int64  `json:"fileSize"`       // 文件大小（字节，便于前端展示）
	UploadTime     string `json:"uploadTime"`     // 文件上传时间
	DatabaseSource string `json:"databaseSource"` // 存储数据库标识（如：mysql-bom-file-01，多数据库部署时区分）
}

// ProductionBOMInfo 生产BOM专属信息
type ProductionBOMInfo struct {
	RDOMReferenceID string `json:"rdBOMReferenceID"` // 关联的研发BOM ID
	CompareResult   string `json:"compareResult"`    // 比对结果：一致/不一致
	CompareDetails  string `json:"compareDetails"`   // 比对详情（如：物料A规格不一致，研发为XX，生产为YY）
	Verifier        string `json:"verifier"`         // 核验人（数字身份）
	VerifyTime      string `json:"verifyTime"`       // 核验时间
}

// BOMChangeRecord BOM变更记录
type BOMChangeRecord struct {
	ChangeID     string      `json:"changeID"`     // 变更申请ID
	ChangeReason string      `json:"changeReason"` // 变更原因（如设计优化、物料替代）
	OldVersion   string      `json:"oldVersion"`   // 变更前版本
	NewVersion   string      `json:"newVersion"`   // 变更后版本
	AuditFlow    []AuditNode `json:"auditFlow"`    // 审批流程
	ChangeTime   string      `json:"changeTime"`   // 变更生效时间
}

// AuditNode 审批节点
type AuditNode struct {
	Auditor     string `json:"auditor"`          // 审批人
	AuditRole   string `json:"auditRole"`        // 审批角色（如研发负责人、生产负责人）
	AuditResult string `json:"auditResult"`      // 审批结果：通过/驳回
	AuditTime   string `json:"auditTime"`        // 审批时间
	Remark      string `json:"remark,omitempty"` // 审批备注
}

// ProductionData 生产数据
type ProductionData struct {
	ProductionID   string            `json:"productionID"`   // 生产记录ID
	PartID         string            `json:"partID"`         // 关联零部件
	BatchNo        string            `json:"batchNo"`        // 生产批次
	Params         map[string]string `json:"params"`         // 关键生产参数（如"温度":"180℃"、"时长":"2h"）
	ProductionLine string            `json:"productionLine"` // 生产线编号
	Operator       string            `json:"operator"`       // 操作员ID
	FinishTime     string            `json:"finishTime"`     // 生产完成时间
}

// QualityInspection 质检数据
type QualityInspection struct {
	InspectionID string            `json:"inspectionID"` // 质检记录ID
	PartID       string            `json:"partID"`       // 关联零部件
	BatchNo      string            `json:"batchNo"`      // 批次
	Indicators   map[string]string `json:"indicators"`   // 检测指标（如"尺寸":"10±0.1mm"、"硬度":"HRC50"）
	Result       string            `json:"result"`       // 结果（合格/不合格）
	Handler      string            `json:"handler"`      // 质检人员ID
	HandleTime   string            `json:"handleTime"`   // 质检时间
	Disposal     string            `json:"disposal"`     // 不合格处理（返工/报废），合格则为空
}

// SupplyOrder 采购订单
type SupplyOrder struct {
	OrderID    string `json:"orderID"`    // 订单ID
	Buyer      string `json:"buyer"`      // 采购方ID（车企）
	Seller     string `json:"seller"`     // 销售方ID（零部件厂商）
	PartID     string `json:"partID"`     // 零部件ID
	Quantity   int    `json:"quantity"`   // 数量
	BatchNo    string `json:"batchNo"`    // 涉及批次
	AgreedTime string `json:"agreedTime"` // 约定交付时间
	Status     string `json:"status"`     // 状态（待发货/已发货/已签收）
	CreateTime string `json:"createTime"` // 订单创建时间
}

// LogisticsData 物流数据
type LogisticsData struct {
	LogisticsID string `json:"logisticsID"` // 物流记录ID
	OrderID     string `json:"orderID"`     // 关联订单ID
	Carrier     string `json:"carrier"`     // 物流商ID
	StartTime   string `json:"startTime"`   // 出发时间
	EndTime     string `json:"endTime"`     // 送达时间
	GPSHash     string `json:"gpsHash"`     // GPS轨迹数据哈希（分布式存储引用）
	Receiver    string `json:"receiver"`    // 签收人ID
}

// Reconciliation 对账结算
type Reconciliation struct {
	ReconID    string `json:"reconID"`    // 对账ID
	OrderID    string `json:"orderID"`    // 关联订单
	Amount     string `json:"amount"`     // 金额
	Status     string `json:"status"`     // 状态（待对账/已确认/已结算）
	SettleTime string `json:"settleTime"` // 结算时间
}

// SupplyChainData 供应关联数据
type SupplyChainData struct {
	// 1. 关联Part/批次
	ChainID string `json:"chainID"` // 供应链环节唯一ID（区分不同环节）
	PartID  string `json:"partID"`  // 关联零部件
	BatchNo string `json:"batchNo"` // 关联批次

	// 2. 供应链模块关联订单/物流
	OrderID     string `json:"orderID"`     // 复用SupplyOrder.OrderID：关联采购订单
	LogisticsID string `json:"logisticsID"` // 复用LogisticsData.LogisticsID：关联物流记录（仅物流环节有值）

	// 3. 环节核心信息
	StageType        string `json:"stageType"`        // 环节类型：采购下单/物流发货/仓库入库/仓库出库/车企签收
	StageStatus      string `json:"stageStatus"`      // 环节状态：待处理/处理中/已完成/异常
	Participator     string `json:"participator"`     // 环节参与方ID（供应商/物流商/车企仓库）
	ParticipatorRole string `json:"participatorRole"` // 参与方角色

	// 4. 环节详情
	Quantity    int    `json:"quantity"`    // 环节涉及数量
	OperateTime string `json:"operateTime"` // 环节操作时间
	Operator    string `json:"operator"`    // 操作人ID（数字身份）

	// 5. 通用附件复用（与BOM/售后复用同一个Attachment结构体）
	AttachmentInfo *Attachment `json:"attachmentInfo,omitempty"` // 复用通用附件结构体
	Remark         string      `json:"remark,omitempty"`         // 环节备注
}

type Attachment struct {
	FileStorageID  string `json:"fileStorageID"`  // 普通数据库文件主键ID（复用规则）
	FilePath       string `json:"filePath"`       // 文件存储路径
	FileHash       string `json:"fileHash"`       // 文件哈希（SHA256，链上校验）
	FileType       string `json:"fileType"`       // 文件类型：PNG/JPG/PDF
	DatabaseSource string `json:"databaseSource"` // 存储数据库标识（复用多库规则）
}

// 售后模块结构体
// FaultReport 故障上报
type FaultReport struct {
	FaultID     string  `json:"faultID"`     // 故障ID
	PartID      string  `json:"partID"`      // 关联零部件
	VIN         string  `json:"vin"`         // 车辆VIN码
	Reporter    string  `json:"reporter"`    // 上报者ID（4S店/用户）
	Description string  `json:"description"` // 故障描述（如"发动机活塞磨损严重"）
	FaultType   string  `json:"faultType"`   // AI分类结果（如"磨损类"，NLP输出）
	RiskProb    float64 `json:"riskProb"`    // 风险概率（AI预测，0-1）
	ReportTime  string  `json:"reportTime"`  // 上报时间
	Status      string  `json:"status"`      // 处理状态（待审核/已审核/已召回）
}

// RecallRecord 召回记录
type RecallRecord struct {
	RecallID      string   `json:"recallID"`      // 召回ID
	BatchNos      []string `json:"batchNos"`      // 涉及批次列表
	Reason        string   `json:"reason"`        // 召回原因
	AffectedParts []string `json:"affectedParts"` // 受影响零部件ID列表
	Status        string   `json:"status"`        // 进度（待通知/通知中/已完成）
	CreateTime    string   `json:"createTime"`    // 召回发起时间
	FinishTime    string   `json:"finishTime"`    // 完成时间（如有）
}

// AftersaleRecord 售后记录
type AftersaleRecord struct {
	// 1. 关联Part/车辆/批次
	AftersaleID string `json:"aftersaleID"` // 售后记录唯一ID
	PartID      string `json:"partID"`      // 关联零部件
	BatchNo     string `json:"batchNo"`     // 关联批次
	VIN         string `json:"vin"`         // 关联车辆

	// 2. 关联故障/召回
	FaultID  string `json:"faultID"`  // 关联故障上报（仅故障/维修环节有值）
	RecallID string `json:"recallID"` // 关联召回记录（仅召回环节有值）

	// 3. 售后核心信息
	AftersaleType   string `json:"aftersaleType"`   // 售后类型：故障报修/维修处理/召回登记/报废处理
	AftersaleStatus string `json:"aftersaleStatus"` // 处理状态：待审核/处理中/已完成
	HandlerOrgID    string `json:"handlerOrgID"`    // 处理机构ID（4S店/售后中心）
	HandlerID       string `json:"handlerID"`       // 处理人ID

	// 4. 处理详情
	Description   string `json:"description"`   // 售后描述
	ProcessResult string `json:"processResult"` // 处理结果
	ProcessTime   string `json:"processTime"`   // 处理完成时间
	Cost          string `json:"cost"`          // 售后成本

	// 5. 通用附件复用
	AttachmentInfo *Attachment `json:"attachmentInfo"` // 复用通用附件
	Remark         string      `json:"remark"`         // 售后备注
}

// UserIdentity 数字身份
type UserIdentity struct {
	UserID       string   `json:"userID"`       // 用户唯一标识
	Org          string   `json:"org"`          // 所属组织（厂商/物流/车企/售后）
	Role         string   `json:"role"`         // 角色（如"厂商-生产员"、"车企-审核员"）
	CertHash     string   `json:"certHash"`     // 数字证书哈希（CA颁发）
	Permissions  []string `json:"permissions"`  // 权限列表（如"上传生产数据"、"查询召回"）
	RegisterTime string   `json:"registerTime"` // 注册时间
}

// QAInteraction 问答交互记录（上链存证）
type QAInteraction struct {
	QAID           string `json:"qaid"`           // 问答记录ID
	UserID         string `json:"userID"`         // 提问用户ID
	Question       string `json:"question"`       // 问题内容
	Intent         string `json:"intent"`         // 意图分类（如"故障进度查询"）
	PartID         string `json:"partID"`         // 关联零部件ID（如有）
	Answer         string `json:"answer"`         // 回答内容
	QueryTime      string `json:"queryTime"`      // 提问时间
	ContractMethod string `json:"contractMethod"` // 调用的智能合约方法（如"QueryFaultProgress"）
}

// 组织 MSP ID 常量（重新设计为 3 个组织）
const (
	MANUFACTURER_ORG_MSPID = "Org1MSP" // 零部件生产厂商组织 MSP ID
	AUTOMAKER_ORG_MSPID    = "Org2MSP" // 整车车企（采购方）组织 MSP ID
	AFTERSALE_ORG_MSPID    = "Org3MSP" // 4S店/售后中心组织 MSP ID

	// 零部件状态常量
	PART_STATUS_NORMAL            = "NORMAL"            // 正常
	PART_STATUS_IN_PRODUCTION     = "IN_PRODUCTION"     // 在产
	PART_STATUS_QUALITY_INSPECTED = "QUALITY_INSPECTED" // 质检完成
	PART_STATUS_IN_SUPPLY_CHAIN   = "IN_SUPPLY_CHAIN"   // 供应链中
	PART_STATUS_RECALLED          = "RECALLED"          // 已召回

	// BOM 状态常量
	BOM_STATUS_DRAFT     = "草稿"
	BOM_STATUS_PUBLISHED = "已发布"
	BOM_STATUS_CHANGED   = "已变更"
	BOM_STATUS_REVOKED   = "已作废"

	// 售后类型常量
	AFTERSALE_TYPE_FAULT_REPORT = "故障报修"
	AFTERSALE_TYPE_MAINTENANCE  = "维修处理"
	AFTERSALE_TYPE_RECALL       = "召回登记"
	AFTERSALE_TYPE_SCRAP        = "报废处理"

	// 售后状态常量
	AFTERSALE_STATUS_PENDING    = "待审核"
	AFTERSALE_STATUS_PROCESSING = "处理中"
	AFTERSALE_STATUS_COMPLETED  = "已完成"
)

// checkPermission 检查调用者是否具有指定权限
// 参数：
//   - ctx: 交易上下文
//   - allowedMSPID: 允许的 MSP ID
//
// 返回值：
//   - error: 无权限返回错误，有权限返回 nil
func (s *SmartContract) checkPermission(ctx contractapi.TransactionContextInterface, allowedMSPID string) error {
	clientMSPID, err := s.getClientIdentityMSPID(ctx)
	if err != nil {
		return err
	}

	if clientMSPID != allowedMSPID {
		return fmt.Errorf("权限错误：只有 %s 组织的成员才能执行此操作", allowedMSPID)
	}

	return nil
}

// SmartContract structure 提供零部件数据上链功能
type SmartContract struct {
	contractapi.Contract
}

// InitLedger 初始化账本数据
// 功能：在链码首次部署时创建测试数据，验证链码功能正常
// 参数：
//   - ctx: 交易上下文
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用（仅用于初始化）
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// 创建测试零部件数据
	testPart := Part{
		PartID:       "ENG-PISTON-001",
		VIN:          "LVX1234568789798",
		BatchNo:      "B20260127",
		Name:         "发动机活塞",
		Type:         "发动机部件",
		Manufacturer: "厂商A",
		CreateTime:   fmt.Sprintf("%d", time.Now().Unix()),
		Status:       "NORMAL",
	}

	// 序列化零部件数据
	partJson, err := json.Marshal(testPart)
	if err != nil {
		return fmt.Errorf("序列化测试零部件失败: %v", err)
	}

	// 将测试数据写入账本
	return ctx.GetStub().PutState(testPart.PartID, partJson)
}

// getClientIdentityMSPID 获取客户端身份信息
// 参数：
//   - ctx: 交易上下文
//
// 返回值：
//   - string: MSP ID
//   - error: 获取失败返回错误
func (s *SmartContract) getClientIdentityMSPID(ctx contractapi.TransactionContextInterface) (string, error) {
	// 如果是测试模式，直接解析 Creator 字节获取 MSPID
	if TestMode {
		creator, err := ctx.GetStub().GetCreator()
		if err != nil {
			return "", fmt.Errorf("获取调用者身份失败（测试模式）：%v", err)
		}

		if len(creator) == 0 {
			return "", fmt.Errorf("调用者身份为空")
		}

		// 使用 msp.SerializedIdentity 来解析 Creator 字节
		sId := &msp.SerializedIdentity{}
		if err := proto.Unmarshal(creator, sId); err != nil {
			return "", fmt.Errorf("解析身份信息失败：%v", err)
		}

		return sId.Mspid, nil
	}

	// 正常模式：使用 cid 包获取真实的 MSPID
	clientID, err := cid.New(ctx.GetStub())
	if err != nil {
		return "", fmt.Errorf("获取客户端身份信息失败：%v", err)
	}
	return clientID.GetMSPID()
}

// CreatePart 创建新零部件
// 权限要求：仅零部件生产厂商可执行
// 功能：将零部件核心信息上链存储
// 参数：
//   - ctx: 交易上下文
//   - partJSON: 零部件JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
//
// CreatePart 创建零部件记录
// 参数：
//   - ctx: 交易上下文
//   - partJSON: 零部件 JSON 字符串，包含 partID、vin、batchNo、name、type、manufacturer、createTime、status
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：仅 ManufacturerMSP 成员可调用
func (s *SmartContract) CreatePart(ctx contractapi.TransactionContextInterface, partJSON string) error {
	// 使用辅助函数检查权限
	if err := s.checkPermission(ctx, MANUFACTURER_ORG_MSPID); err != nil {
		return err
	}

	// 将前端传入的JSON字符串反序列化为Part结构体
	var part Part
	err := json.Unmarshal([]byte(partJSON), &part)
	if err != nil {
		return fmt.Errorf("反序列化零部件数据失败: %v", err)
	}

	// 校验核心字段
	if part.PartID == "" || part.BatchNo == "" {
		return fmt.Errorf("PartID和BatchNo不能为空")
	}

	// 序列化结构体并存入账本
	partBytes, err := json.Marshal(part)
	if err != nil {
		return fmt.Errorf("序列化零部件数据失败: %v", err)
	}

	// 记录操作日志
	if err := ctx.GetStub().SetEvent("CreatePart", []byte(fmt.Sprintf("创建零部件: %s", part.PartID))); err != nil {
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState(part.PartID, partBytes)
}

// QueryPart 查询零部件信息
// 功能：根据零部件ID查询零部件的详细信息
// 参数：
//   - ctx: 交易上下文
//   - partID: 零部件唯一标识
//
// 返回值：
//   - *Part: 零部件信息指针
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用
func (s *SmartContract) QueryPart(ctx contractapi.TransactionContextInterface, partID string) (*Part, error) {
	// 从账本获取核心part数据
	partBytes, err := ctx.GetStub().GetState(partID)
	if err != nil {
		return nil, fmt.Errorf("查询零部件失败: %v", err)
	}
	if partBytes == nil {
		return nil, fmt.Errorf("零部件ID %s 不存在", partID)
	}

	// 反序列化Part数据
	var part Part
	err = json.Unmarshal(partBytes, &part)
	if err != nil {
		return nil, fmt.Errorf("反序列化零部件失败: %v", err)
	}

	return &part, nil
}

// QueryPartLifecycle 查询零部件全生命周期轨迹
// 功能：根据零部件ID查询其完整的生命周期信息，包括BOM、生产、质检、供应链和售后数据
// 参数：
//   - ctx: 交易上下文
//   - partID: 零部件唯一标识
//
// 返回值：
//   - *PartLifecycle: 零部件生命周期信息指针
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用
func (s *SmartContract) QueryPartLifecycle(ctx contractapi.TransactionContextInterface, partID string) (*PartLifecycle, error) {
	// 从账本获取生命周期数据
	partLifecycleBytes, err := ctx.GetStub().GetState("LIFECYCLE_" + partID)
	if err != nil {
		return nil, fmt.Errorf("查询零部件生命周期失败: %v", err)
	}
	if partLifecycleBytes == nil {
		return nil, fmt.Errorf("零部件ID %s 的生命周期数据不存在", partID)
	}

	// 反序列化生命周期数据
	var lifecycle PartLifecycle
	err = json.Unmarshal(partLifecycleBytes, &lifecycle)
	if err != nil {
		return nil, fmt.Errorf("反序列化生命周期数据失败: %v", err)
	}

	// 初始化空切片，避免返回nil
	if lifecycle.AftersaleInfo == nil {
		lifecycle.AftersaleInfo = []AftersaleRecord{}
	}
	if lifecycle.SupplyChainInfo == nil {
		lifecycle.SupplyChainInfo = []SupplyChainData{}
	}

	return &lifecycle, nil
}

// getAllParts 获取所有零部件（辅助方法，用于leveldb支持）
// 功能：获取账本中所有零部件数据
// 参数：
//   - ctx: 交易上下文
//
// 返回值：
//   - []*Part: 零部件列表
//   - error: 操作成功返回 nil，失败返回具体错误信息
func (s *SmartContract) getAllParts(ctx contractapi.TransactionContextInterface) ([]*Part, error) {
	iterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("获取所有零部件失败: %v", err)
	}
	defer iterator.Close()

	var parts []*Part
	for iterator.HasNext() {
		queryResponse, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("迭代查询结果失败: %v", err)
		}

		key := string(queryResponse.Key)
		value := queryResponse.Value

		// 只处理零部件数据（键不包含前缀）
		if len(key) > 0 && key[0] >= 'A' && key[0] <= 'Z' && !strings.Contains(key, "_") {
			var part Part
			err = json.Unmarshal(value, &part)
			if err != nil {
				return nil, fmt.Errorf("反序列化零部件失败: %v", err)
			}
			parts = append(parts, &part)
		}
	}

	return parts, nil
}

// QueryPartByBatchNo 按批次号查询零部件
// 功能：根据生产批次号查询该批次的所有零部件信息
// 参数：
//   - ctx: 交易上下文
//   - batchNo: 生产批次号
//
// 返回值：
//   - []*Part: 零部件列表指针
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用
// 注意：此方法使用富查询，需要在支持富查询的Fabric网络环境中使用
func (s *SmartContract) QueryPartByBatchNo(ctx contractapi.TransactionContextInterface, batchNo string) ([]*Part, error) {
	// 尝试使用富查询（CouchDB支持）
	queryString := fmt.Sprintf(`{"selector":{"batchNo":"%s"}}`, batchNo)
	iterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err == nil {
		defer iterator.Close()

		var parts []*Part
		for iterator.HasNext() {
			queryResponse, err := iterator.Next()
			if err != nil {
				return nil, fmt.Errorf("迭代查询结果失败: %v", err)
			}

			var part Part
			err = json.Unmarshal(queryResponse.Value, &part)
			if err != nil {
				return nil, fmt.Errorf("反序列化零部件失败: %v", err)
			}
			parts = append(parts, &part)
		}

		return parts, nil
	}

	// 如果富查询失败（leveldb不支持），使用getAllParts辅助方法进行过滤
	log.Printf("富查询失败，使用getAllParts辅助方法进行过滤：%v", err)
	allParts, err := s.getAllParts(ctx)
	if err != nil {
		return nil, err
	}

	var filteredParts []*Part
	for _, part := range allParts {
		if part.BatchNo == batchNo {
			filteredParts = append(filteredParts, part)
		}
	}

	return filteredParts, nil
}

// QueryPartByVIN 按VIN码查询零部件
// 功能：根据车辆VIN码查询该车辆关联的所有零部件信息
// 参数：
//   - ctx: 交易上下文
//   - vin: 车辆VIN码
//
// 返回值：
//   - []*Part: 零部件列表指针
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用
// 注意：此方法使用富查询，需要在支持富查询的Fabric网络环境中使用
func (s *SmartContract) QueryPartByVIN(ctx contractapi.TransactionContextInterface, vin string) ([]*Part, error) {
	// 尝试使用富查询（CouchDB支持）
	queryString := fmt.Sprintf(`{"selector":{"vin":"%s"}}`, vin)
	iterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err == nil {
		defer iterator.Close()

		var parts []*Part
		for iterator.HasNext() {
			queryResponse, err := iterator.Next()
			if err != nil {
				return nil, fmt.Errorf("迭代查询结果失败: %v", err)
			}

			var part Part
			err = json.Unmarshal(queryResponse.Value, &part)
			if err != nil {
				return nil, fmt.Errorf("反序列化零部件失败: %v", err)
			}
			parts = append(parts, &part)
		}

		return parts, nil
	}

	// 如果富查询失败（leveldb不支持），使用getAllParts辅助方法进行过滤
	log.Printf("富查询失败，使用getAllParts辅助方法进行过滤：%v", err)
	allParts, err := s.getAllParts(ctx)
	if err != nil {
		return nil, err
	}

	var filteredParts []*Part
	for _, part := range allParts {
		if part.VIN == vin {
			filteredParts = append(filteredParts, part)
		}
	}

	return filteredParts, nil
}

// CreateBOM 创建BOM记录
// 参数：
//   - ctx: 交易上下文
//   - bomJSON: BOM JSON 字符串，包含 bomID、bomType、productModel、partBatchNo、version、creator、createTime、status、materialList、rdBOMFileInfo、productionBOMInfo、changeRecords
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：仅 ManufacturerMSP 成员可调用
func (s *SmartContract) CreateBOM(ctx contractapi.TransactionContextInterface, bomJSON string) error {
	// 使用辅助函数检查权限
	if err := s.checkPermission(ctx, MANUFACTURER_ORG_MSPID); err != nil {
		return err
	}

	// 反序列化BOM数据
	var bom BOM
	err := json.Unmarshal([]byte(bomJSON), &bom)
	if err != nil {
		return fmt.Errorf("反序列化BOM数据失败: %v", err)
	}

	// 数据验证
	if bom.BOMID == "" || bom.BOMType == "" {
		return fmt.Errorf("BOMID和BOMType不能为空")
	}

	// 初始化空对象，确保必填字段存在
	if bom.RDBOMFileInfo == nil {
		bom.RDBOMFileInfo = &BOMFileInfo{}
	}
	if bom.ProductionBOMInfo == nil {
		bom.ProductionBOMInfo = &ProductionBOMInfo{}
	}
	if bom.ChangeRecords == nil {
		bom.ChangeRecords = []BOMChangeRecord{}
	}

	bomBytes, err := json.Marshal(bom)
	if err != nil {
		return fmt.Errorf("序列化BOM数据失败: %v", err)
	}

	// 记录操作日志
	if err := ctx.GetStub().SetEvent("CreateBOM", []byte(fmt.Sprintf("创建BOM: %s", bom.BOMID))); err != nil {
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState("BOM_"+bom.BOMID, bomBytes)
}

// QueryBOM 查询BOM信息
// 功能：根据BOM ID查询BOM的详细信息
// 参数：
//   - ctx: 交易上下文
//   - bomID: BOM唯一标识
//
// 返回值：
//   - *BOM: BOM信息指针
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用
func (s *SmartContract) QueryBOM(ctx contractapi.TransactionContextInterface, bomID string) (*BOM, error) {
	// 从账本获取BOM数据
	bomBytes, err := ctx.GetStub().GetState("BOM_" + bomID)
	if err != nil {
		return nil, fmt.Errorf("查询BOM失败: %v", err)
	}
	if bomBytes == nil {
		return nil, fmt.Errorf("BOMID %s 不存在", bomID)
	}

	// 反序列化BOM数据
	var bom BOM
	err = json.Unmarshal(bomBytes, &bom)
	if err != nil {
		return nil, fmt.Errorf("反序列化BOM失败: %v", err)
	}

	return &bom, nil
}

// UpdateBOM 更新BOM
// 权限要求：仅零部件生产厂商可执行
// 功能：更新BOM信息（状态、物料清单等）
// 参数：
//   - ctx: 交易上下文
//   - bomJSON: BOM JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
func (s *SmartContract) UpdateBOM(ctx contractapi.TransactionContextInterface, bomJSON string) error {
	// 1. 校验调用者身份
	if err := s.checkPermission(ctx, MANUFACTURER_ORG_MSPID); err != nil {
		log.Printf("权限验证失败：%v", err)
		return err
	}

	// 2. 反序列化BOM数据
	var bom BOM
	err := json.Unmarshal([]byte(bomJSON), &bom)
	if err != nil {
		log.Printf("反序列化BOM数据失败：%v", err)
		return fmt.Errorf("反序列化BOM数据失败: %v", err)
	}

	// 3. 检查BOM是否存在
	existingBytes, err := ctx.GetStub().GetState("BOM_" + bom.BOMID)
	if err != nil {
		log.Printf("查询BOM失败，BOMID：%s，错误：%v", bom.BOMID, err)
		return fmt.Errorf("查询BOM失败: %v", err)
	}
	if existingBytes == nil {
		log.Printf("BOM不存在，BOMID：%s", bom.BOMID)
		return fmt.Errorf("BOMID %s 不存在", bom.BOMID)
	}

	// 4. 数据验证
	if bom.Status != "草稿" && bom.Status != "已发布" && bom.Status != "已变更" && bom.Status != "已作废" {
		log.Printf("BOM状态无效，BOMID：%s，状态：%s", bom.BOMID, bom.Status)
		return fmt.Errorf("BOM状态必须是：草稿、已发布、已变更或已作废")
	}

	// 5. 序列化并更新
	bomBytes, err := json.Marshal(bom)
	if err != nil {
		log.Printf("序列化BOM数据失败，BOMID：%s，错误：%v", bom.BOMID, err)
		return fmt.Errorf("序列化BOM数据失败: %v", err)
	}

	// 记录操作日志
	log.Printf("更新BOM成功，BOMID：%s，状态：%s", bom.BOMID, bom.Status)
	if err := ctx.GetStub().SetEvent("UpdateBOM", []byte(fmt.Sprintf("更新BOM: %s", bom.BOMID))); err != nil {
		log.Printf("设置事件失败：%v", err)
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState("BOM_"+bom.BOMID, bomBytes)
}

// CompareBOM 生产BOM与研发BOM比对
// 权限要求：零部件生产厂商和车企可执行
// 功能：比对生产BOM与研发BOM是否一致
// 参数：
//   - ctx: 交易上下文
//   - productionBOMID: 生产BOM ID
//   - rdBOMID: 研发BOM ID
//
// 返回：
//   - string: 比对结果
//   - error: 操作失败时返回错误
func (s *SmartContract) CompareBOM(ctx contractapi.TransactionContextInterface, productionBOMID string, rdBOMID string) (string, error) {
	// 1. 校验调用者身份
	clientMSPID, err := s.getClientIdentityMSPID(ctx)
	if err != nil {
		log.Printf("获取调用者身份失败：%v", err)
		return "", fmt.Errorf("获取调用者身份失败：%v", err)
	}
	if clientMSPID != MANUFACTURER_ORG_MSPID && clientMSPID != AUTOMAKER_ORG_MSPID {
		log.Printf("权限不足，MSPID：%s", clientMSPID)
		return "", fmt.Errorf("只有零部件生产厂商或整车车企组织的成员才能比对BOM")
	}

	// 2. 查询生产BOM
	prodBOMBytes, err := ctx.GetStub().GetState("BOM_" + productionBOMID)
	if err != nil {
		log.Printf("查询生产BOM失败，BOMID：%s，错误：%v", productionBOMID, err)
		return "", fmt.Errorf("查询生产BOM失败: %v", err)
	}
	if prodBOMBytes == nil {
		log.Printf("生产BOM不存在，BOMID：%s", productionBOMID)
		return "", fmt.Errorf("生产BOMID %s 不存在", productionBOMID)
	}

	// 3. 查询研发BOM
	rdBOMBytes, err := ctx.GetStub().GetState("BOM_" + rdBOMID)
	if err != nil {
		log.Printf("查询研发BOM失败，BOMID：%s，错误：%v", rdBOMID, err)
		return "", fmt.Errorf("查询研发BOM失败: %v", err)
	}
	if rdBOMBytes == nil {
		log.Printf("研发BOM不存在，BOMID：%s", rdBOMID)
		return "", fmt.Errorf("研发BOMID %s 不存在", rdBOMID)
	}

	// 4. 反序列化BOM数据
	var prodBOM, rdBOM BOM
	err = json.Unmarshal(prodBOMBytes, &prodBOM)
	if err != nil {
		log.Printf("反序列化生产BOM失败，BOMID：%s，错误：%v", productionBOMID, err)
		return "", fmt.Errorf("反序列化生产BOM失败: %v", err)
	}
	err = json.Unmarshal(rdBOMBytes, &rdBOM)
	if err != nil {
		log.Printf("反序列化研发BOM失败，BOMID：%s，错误：%v", rdBOMID, err)
		return "", fmt.Errorf("反序列化研发BOM失败: %v", err)
	}

	// 5. 比对物料数量
	if len(prodBOM.MaterialList) != len(rdBOM.MaterialList) {
		log.Printf("BOM比对结果：物料数量不同，生产BOM：%d，研发BOM：%d", len(prodBOM.MaterialList), len(rdBOM.MaterialList))
		return "不一致：物料数量不同", nil
	}

	// 6. 逐项比对物料
	for i := range prodBOM.MaterialList {
		if prodBOM.MaterialList[i].MaterialID != rdBOM.MaterialList[i].MaterialID ||
			prodBOM.MaterialList[i].Spec != rdBOM.MaterialList[i].Spec ||
			prodBOM.MaterialList[i].Quantity != rdBOM.MaterialList[i].Quantity {
			log.Printf("BOM比对结果：物料 %s 规格或数量不同", prodBOM.MaterialList[i].MaterialID)
			return fmt.Sprintf("不一致：物料 %s 规格或数量不同", prodBOM.MaterialList[i].MaterialID), nil
		}
	}

	// 记录比对结果
	log.Printf("BOM比对完成，生产BOMID：%s，研发BOMID：%s，结果：一致", productionBOMID, rdBOMID)
	return "一致", nil
}

// SubmitBOMChange 提交BOM变更申请
// 权限要求：仅零部件生产厂商可执行
// 功能：提交BOM变更申请，记录变更原因和审批流程
// 参数：
//   - ctx: 交易上下文
//   - bomID: BOM ID
//   - changeJSON: 变更记录JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
func (s *SmartContract) SubmitBOMChange(ctx contractapi.TransactionContextInterface, bomID string, changeJSON string) error {
	// 1. 校验调用者身份
	if err := s.checkPermission(ctx, MANUFACTURER_ORG_MSPID); err != nil {
		log.Printf("权限验证失败，BOMID：%s，错误：%v", bomID, err)
		return err
	}

	// 2. 反序列化变更记录
	var change BOMChangeRecord
	err := json.Unmarshal([]byte(changeJSON), &change)
	if err != nil {
		log.Printf("反序列化变更记录失败，BOMID：%s，错误：%v", bomID, err)
		return fmt.Errorf("反序列化变更记录失败: %v", err)
	}

	// 3. 查询BOM
	bomBytes, err := ctx.GetStub().GetState("BOM_" + bomID)
	if err != nil {
		log.Printf("查询BOM失败，BOMID：%s，错误：%v", bomID, err)
		return fmt.Errorf("查询BOM失败: %v", err)
	}
	if bomBytes == nil {
		log.Printf("BOM不存在，BOMID：%s", bomID)
		return fmt.Errorf("BOMID %s 不存在", bomID)
	}

	// 4. 反序列化BOM并添加变更记录
	var bom BOM
	err = json.Unmarshal(bomBytes, &bom)
	if err != nil {
		log.Printf("反序列化BOM失败，BOMID：%s，错误：%v", bomID, err)
		return fmt.Errorf("反序列化BOM失败: %v", err)
	}

	// 5. 设置变更时间
	if change.ChangeTime == "" {
		change.ChangeTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	bom.ChangeRecords = append(bom.ChangeRecords, change)
	bom.Status = "已变更"

	// 6. 序列化并更新
	updatedBOMBytes, err := json.Marshal(bom)
	if err != nil {
		log.Printf("序列化BOM数据失败，BOMID：%s，错误：%v", bomID, err)
		return fmt.Errorf("序列化BOM数据失败: %v", err)
	}

	// 记录操作日志
	log.Printf("提交BOM变更成功，BOMID：%s，变更ID：%s，版本：%s -> %s", bomID, change.ChangeID, change.OldVersion, change.NewVersion)
	if err := ctx.GetStub().SetEvent("SubmitBOMChange", []byte(fmt.Sprintf("提交BOM变更: %s", bomID))); err != nil {
		log.Printf("设置事件失败：%v", err)
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState("BOM_"+bomID, updatedBOMBytes)
}

// CreateProductionData 创建生产数据
// 权限要求：仅零部件生产厂商可执行
// 功能：记录零部件生产过程中的关键参数和操作信息
// 参数：
//   - ctx: 交易上下文
//   - productionJSON: 生产数据JSON字符串
//
// CreateProductionData 创建生产数据记录
// 参数：
//   - ctx: 交易上下文
//   - productionJSON: 生产数据 JSON 字符串，包含 productionID、partID、batchNo、params、productionLine、operator、finishTime
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：仅 ManufacturerMSP 成员可调用
func (s *SmartContract) CreateProductionData(ctx contractapi.TransactionContextInterface, productionJSON string) error {
	// 使用辅助函数检查权限
	if err := s.checkPermission(ctx, MANUFACTURER_ORG_MSPID); err != nil {
		return err
	}

	// 反序列化生产数据
	var production ProductionData
	err := json.Unmarshal([]byte(productionJSON), &production)
	if err != nil {
		return fmt.Errorf("反序列化生产数据失败: %v", err)
	}

	// 数据验证
	if production.ProductionID == "" || production.PartID == "" {
		return fmt.Errorf("ProductionID和PartID不能为空")
	}

	// 检查零部件是否存在
	partBytes, err := ctx.GetStub().GetState(production.PartID)
	if err != nil {
		return fmt.Errorf("查询零部件失败: %v", err)
	}
	if partBytes == nil {
		return fmt.Errorf("零部件ID %s 不存在", production.PartID)
	}

	// 设置完成时间
	if production.FinishTime == "" {
		production.FinishTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	// 序列化并存储生产数据
	productionBytes, err := json.Marshal(production)
	if err != nil {
		return fmt.Errorf("序列化生产数据失败: %v", err)
	}
	if err := ctx.GetStub().PutState("PROD_"+production.ProductionID, productionBytes); err != nil {
		return err
	}

	// 更新零部件生命周期
	lifecycleBytes, err := ctx.GetStub().GetState("LIFECYCLE_" + production.PartID)
	var lifecycle PartLifecycle
	if err == nil && lifecycleBytes != nil {
		json.Unmarshal(lifecycleBytes, &lifecycle)
	}
	lifecycle.PartID = production.PartID
	lifecycle.ProductionInfo = &production

	lifecycleBytes, err = json.Marshal(lifecycle)
	if err != nil {
		return fmt.Errorf("序列化生命周期数据失败: %v", err)
	}

	// 记录操作日志
	if err := ctx.GetStub().SetEvent("CreateProductionData", []byte(fmt.Sprintf("创建生产数据: %s", production.ProductionID))); err != nil {
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState("LIFECYCLE_"+production.PartID, lifecycleBytes)
}

// CreateQualityInspection 创建质检记录
// 参数：
//   - ctx: 交易上下文
//   - inspectionJSON: 质检记录 JSON 字符串，包含 inspectionID、partID、batchNo、indicators、result、handler、handleTime、disposal
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：仅 ManufacturerMSP 成员可调用
func (s *SmartContract) CreateQualityInspection(ctx contractapi.TransactionContextInterface, inspectionJSON string) error {
	// 使用辅助函数检查权限
	if err := s.checkPermission(ctx, MANUFACTURER_ORG_MSPID); err != nil {
		return err
	}

	// 反序列化质检数据
	var inspection QualityInspection
	err := json.Unmarshal([]byte(inspectionJSON), &inspection)
	if err != nil {
		return fmt.Errorf("反序列化质检数据失败: %v", err)
	}

	// 数据验证
	if inspection.InspectionID == "" || inspection.PartID == "" {
		return fmt.Errorf("InspectionID和PartID不能为空")
	}
	if inspection.Result != "合格" && inspection.Result != "不合格" {
		return fmt.Errorf("质检结果必须是：合格或不合格")
	}

	// 检查零部件是否存在
	partBytes, err := ctx.GetStub().GetState(inspection.PartID)
	if err != nil {
		return fmt.Errorf("查询零部件失败: %v", err)
	}
	if partBytes == nil {
		return fmt.Errorf("零部件ID %s 不存在", inspection.PartID)
	}

	// 设置质检时间
	if inspection.HandleTime == "" {
		inspection.HandleTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	// 序列化并存储质检数据
	inspectionBytes, err := json.Marshal(inspection)
	if err != nil {
		return fmt.Errorf("序列化质检数据失败: %v", err)
	}
	if err := ctx.GetStub().PutState("QUALITY_"+inspection.InspectionID, inspectionBytes); err != nil {
		return err
	}

	// 更新零部件生命周期
	lifecycleBytes, err := ctx.GetStub().GetState("LIFECYCLE_" + inspection.PartID)
	var lifecycle PartLifecycle
	if err == nil && lifecycleBytes != nil {
		json.Unmarshal(lifecycleBytes, &lifecycle)
	}
	lifecycle.PartID = inspection.PartID
	lifecycle.QualityInfo = &inspection

	lifecycleBytes, err = json.Marshal(lifecycle)
	if err != nil {
		return fmt.Errorf("序列化生命周期数据失败: %v", err)
	}

	// 记录操作日志
	if err := ctx.GetStub().SetEvent("CreateQualityInspection", []byte(fmt.Sprintf("创建质检记录: %s", inspection.InspectionID))); err != nil {
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState("LIFECYCLE_"+inspection.PartID, lifecycleBytes)
}

// UpdatePartStatus 更新零部件状态
// 权限要求：零部件生产厂商和车企可执行
// 功能：更新零部件状态（如从"在产"更新为"召回"）
// 参数：
//   - ctx: 交易上下文
//   - partID: 零部件ID
//   - status: 新状态
//
// 返回：
//   - error: 操作失败时返回错误
func (s *SmartContract) UpdatePartStatus(ctx contractapi.TransactionContextInterface, partID string, status string) error {
	// 1. 校验调用者身份
	clientMSPID, err := s.getClientIdentityMSPID(ctx)
	if err != nil {
		log.Printf("获取调用者身份失败，PartID：%s，错误：%v", partID, err)
		return fmt.Errorf("获取调用者身份失败：%v", err)
	}
	if clientMSPID != MANUFACTURER_ORG_MSPID && clientMSPID != AUTOMAKER_ORG_MSPID {
		log.Printf("权限不足，PartID：%s，MSPID：%s", partID, clientMSPID)
		return fmt.Errorf("只有零部件生产厂商或整车车企组织的成员才能更新零部件状态")
	}

	// 2. 数据验证
	if status != "在产" && status != "在售" && status != "召回" && status != "报废" && status != "NORMAL" {
		log.Printf("零部件状态无效，PartID：%s，状态：%s", partID, status)
		return fmt.Errorf("零部件状态必须是：在产、在售、召回、报废或NORMAL")
	}

	// 3. 查询零部件
	partBytes, err := ctx.GetStub().GetState(partID)
	if err != nil {
		log.Printf("查询零部件失败，PartID：%s，错误：%v", partID, err)
		return fmt.Errorf("查询零部件失败: %v", err)
	}
	if partBytes == nil {
		log.Printf("零部件不存在，PartID：%s", partID)
		return fmt.Errorf("零部件ID %s 不存在", partID)
	}

	// 4. 反序列化并更新状态
	var part Part
	err = json.Unmarshal(partBytes, &part)
	if err != nil {
		log.Printf("反序列化零部件失败，PartID：%s，错误：%v", partID, err)
		return fmt.Errorf("反序列化零部件失败: %v", err)
	}

	oldStatus := part.Status
	part.Status = status

	// 5. 序列化并更新
	partBytes, err = json.Marshal(part)
	if err != nil {
		log.Printf("序列化零部件数据失败，PartID：%s，错误：%v", partID, err)
		return fmt.Errorf("序列化零部件数据失败: %v", err)
	}

	// 记录操作日志
	log.Printf("更新零部件状态成功，PartID：%s，旧状态：%s，新状态：%s", partID, oldStatus, status)
	if err := ctx.GetStub().SetEvent("UpdatePartStatus", []byte(fmt.Sprintf("更新零部件状态: %s -> %s", oldStatus, status))); err != nil {
		log.Printf("设置事件失败：%v", err)
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState(partID, partBytes)
}

// CreateSupplyOrder 创建采购订单
// 权限要求：车企可执行
// 功能：创建采购订单，记录采购方、销售方、数量等信息
// 参数：
//   - ctx: 交易上下文
//   - orderJSON: 订单JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
//
// CreateSupplyOrder 创建采购订单
// 权限要求：仅整车车企可执行
// 功能：创建零部件采购订单并上链存储
// 参数：
//   - ctx: 交易上下文
//   - orderJSON: 订单JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
func (s *SmartContract) CreateSupplyOrder(ctx contractapi.TransactionContextInterface, orderJSON string) error {
	// 1. 校验调用者身份
	if err := s.checkPermission(ctx, AUTOMAKER_ORG_MSPID); err != nil {
		log.Printf("权限验证失败，错误：%v", err)
		return err
	}

	// 2. 反序列化订单数据
	var order SupplyOrder
	err := json.Unmarshal([]byte(orderJSON), &order)
	if err != nil {
		log.Printf("反序列化订单数据失败，错误：%v", err)
		return fmt.Errorf("反序列化订单数据失败: %v", err)
	}

	// 3. 数据验证
	if order.OrderID == "" || order.PartID == "" {
		log.Printf("订单数据不完整，OrderID：%s，PartID：%s", order.OrderID, order.PartID)
		return fmt.Errorf("OrderID和PartID不能为空")
	}
	if order.Quantity <= 0 {
		log.Printf("订单数量无效，OrderID：%s，数量：%d", order.OrderID, order.Quantity)
		return fmt.Errorf("订单数量必须大于0")
	}
	if order.Status != "待发货" && order.Status != "已发货" && order.Status != "已签收" {
		log.Printf("订单状态无效，OrderID：%s，状态：%s", order.OrderID, order.Status)
		return fmt.Errorf("订单状态必须是：待发货、已发货或已签收")
	}

	// 4. 检查零部件是否存在
	partBytes, err := ctx.GetStub().GetState(order.PartID)
	if err != nil {
		log.Printf("查询零部件失败，PartID：%s，错误：%v", order.PartID, err)
		return fmt.Errorf("查询零部件失败: %v", err)
	}
	if partBytes == nil {
		log.Printf("零部件不存在，PartID：%s", order.PartID)
		return fmt.Errorf("零部件ID %s 不存在", order.PartID)
	}

	// 5. 设置创建时间
	if order.CreateTime == "" {
		order.CreateTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	// 6. 序列化并存储订单
	orderBytes, err := json.Marshal(order)
	if err != nil {
		log.Printf("序列化订单数据失败，OrderID：%s，错误：%v", order.OrderID, err)
		return fmt.Errorf("序列化订单数据失败: %v", err)
	}

	// 记录操作日志
	log.Printf("创建采购订单成功，OrderID：%s，PartID：%s，数量：%d", order.OrderID, order.PartID, order.Quantity)
	if err := ctx.GetStub().SetEvent("CreateSupplyOrder", []byte(fmt.Sprintf("创建采购订单: %s", order.OrderID))); err != nil {
		log.Printf("设置事件失败：%v", err)
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState("ORDER_"+order.OrderID, orderBytes)
}

// CreateLogisticsData 创建物流数据记录
// 功能：记录零部件的物流运输信息，包括物流商、时间、GPS轨迹等
// 参数：
//   - ctx: 交易上下文
//   - logisticsJSON: 物流数据JSON字符串，包含 logisticsID、orderID、carrier、startTime、endTime、gpsHash、receiver
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用（建议由物流商或车企调用）
func (s *SmartContract) CreateLogisticsData(ctx contractapi.TransactionContextInterface, logisticsJSON string) error {
	// 反序列化物流数据
	var logistics LogisticsData
	err := json.Unmarshal([]byte(logisticsJSON), &logistics)
	if err != nil {
		return fmt.Errorf("反序列化物流数据失败: %v", err)
	}

	// 数据验证
	if logistics.LogisticsID == "" || logistics.OrderID == "" {
		return fmt.Errorf("LogisticsID和OrderID不能为空")
	}

	// 序列化并存储物流数据
	logisticsBytes, err := json.Marshal(logistics)
	if err != nil {
		return fmt.Errorf("序列化物流数据失败: %v", err)
	}

	return ctx.GetStub().PutState("LOGISTICS_"+logistics.LogisticsID, logisticsBytes)
}

// UpdateSupplyChainStage 更新供应链环节
// 功能：记录零部件在供应链中的流转信息（采购下单、物流发货、仓库入库、仓库出库、车企签收等）
// 参数：
//   - ctx: 交易上下文
//   - stageJSON: 供应链环节数据JSON字符串，包含 chainID、partID、orderID、logisticsID、stageType、stageStatus、participator、participatorRole、quantity、operateTime、operator、attachmentInfo、remark
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用（建议由各环节参与方调用）
func (s *SmartContract) UpdateSupplyChainStage(ctx contractapi.TransactionContextInterface, stageJSON string) error {
	// 反序列化供应链环节数据
	var stage SupplyChainData
	err := json.Unmarshal([]byte(stageJSON), &stage)
	if err != nil {
		return fmt.Errorf("反序列化供应链环节数据失败: %v", err)
	}

	// 数据验证
	if stage.ChainID == "" || stage.PartID == "" {
		return fmt.Errorf("ChainID和PartID不能为空")
	}

	// 序列化并存储供应链环节
	stageBytes, err := json.Marshal(stage)
	if err != nil {
		return fmt.Errorf("序列化供应链环节数据失败: %v", err)
	}

	err = ctx.GetStub().PutState("CHAIN_"+stage.ChainID, stageBytes)
	if err != nil {
		return err
	}

	// 更新零部件生命周期，添加供应链环节信息
	lifecycleBytes, err := ctx.GetStub().GetState("LIFECYCLE_" + stage.PartID)
	var lifecycle PartLifecycle
	if err == nil && lifecycleBytes != nil {
		json.Unmarshal(lifecycleBytes, &lifecycle)
	}
	lifecycle.PartID = stage.PartID
	if lifecycle.SupplyChainInfo == nil {
		lifecycle.SupplyChainInfo = []SupplyChainData{stage}
	} else {
		lifecycle.SupplyChainInfo = append(lifecycle.SupplyChainInfo, stage)
	}

	lifecycleBytes, err = json.Marshal(lifecycle)
	if err != nil {
		return fmt.Errorf("序列化生命周期数据失败: %v", err)
	}
	return ctx.GetStub().PutState("LIFECYCLE_"+stage.PartID, lifecycleBytes)
}

// CreateReconciliation 创建对账记录
// 功能：记录采购订单的对账结算信息，包括金额、状态、结算时间等
// 参数：
//   - ctx: 交易上下文
//   - reconJSON: 对账数据JSON字符串，包含 reconID、orderID、amount、status、settleTime
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用（建议由财务部门调用）
func (s *SmartContract) CreateReconciliation(ctx contractapi.TransactionContextInterface, reconJSON string) error {
	// 反序列化对账数据
	var recon Reconciliation
	err := json.Unmarshal([]byte(reconJSON), &recon)
	if err != nil {
		return fmt.Errorf("反序列化对账数据失败: %v", err)
	}

	// 数据验证
	if recon.ReconID == "" || recon.OrderID == "" {
		return fmt.Errorf("ReconID和OrderID不能为空")
	}

	// 序列化并存储对账数据
	reconBytes, err := json.Marshal(recon)
	if err != nil {
		return fmt.Errorf("序列化对账数据失败: %v", err)
	}

	return ctx.GetStub().PutState("RECON_"+recon.ReconID, reconBytes)
}

// CreateFaultReport 创建故障上报记录
// 功能：记录零部件故障信息，包括故障描述、类型、风险概率等，并自动创建对应的售后记录
// 参数：
//   - ctx: 交易上下文
//   - faultJSON: 故障数据JSON字符串，包含 faultID、partID、vin、reporter、description、faultType、riskProb、reportTime、status
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用（建议由4S店或用户调用）
func (s *SmartContract) CreateFaultReport(ctx contractapi.TransactionContextInterface, faultJSON string) error {
	// 反序列化故障数据
	var fault FaultReport
	err := json.Unmarshal([]byte(faultJSON), &fault)
	if err != nil {
		return fmt.Errorf("反序列化故障数据失败: %v", err)
	}

	// 数据验证
	if fault.FaultID == "" || fault.PartID == "" {
		return fmt.Errorf("FaultID和PartID不能为空")
	}

	// 序列化并存储故障数据
	faultBytes, err := json.Marshal(fault)
	if err != nil {
		return fmt.Errorf("序列化故障数据失败: %v", err)
	}

	err = ctx.GetStub().PutState("FAULT_"+fault.FaultID, faultBytes)
	if err != nil {
		return err
	}

	// 自动创建对应的售后记录并更新零部件生命周期
	lifecycleBytes, err := ctx.GetStub().GetState("LIFECYCLE_" + fault.PartID)
	var lifecycle PartLifecycle
	if err == nil && lifecycleBytes != nil {
		json.Unmarshal(lifecycleBytes, &lifecycle)
	}
	lifecycle.PartID = fault.PartID

	// 创建售后记录
	var aftersale AftersaleRecord
	aftersale.AftersaleID = "AFT_" + fault.FaultID
	aftersale.PartID = fault.PartID
	aftersale.VIN = fault.VIN
	aftersale.FaultID = fault.FaultID
	aftersale.AftersaleType = "故障报修"
	aftersale.AftersaleStatus = "待审核"
	aftersale.Description = fault.Description
	aftersale.ProcessTime = fault.ReportTime

	// 更新生命周期中的售后信息
	if lifecycle.AftersaleInfo == nil {
		lifecycle.AftersaleInfo = []AftersaleRecord{aftersale}
	} else {
		lifecycle.AftersaleInfo = append(lifecycle.AftersaleInfo, aftersale)
	}

	lifecycleBytes, err = json.Marshal(lifecycle)
	if err != nil {
		return fmt.Errorf("序列化生命周期数据失败: %v", err)
	}

	return ctx.GetStub().PutState("LIFECYCLE_"+fault.PartID, lifecycleBytes)
}

// CreateRecallRecord 创建召回记录
// 功能：记录零部件召回信息，包括涉及批次、原因、受影响零部件等，并自动更新相关零部件状态为"召回"
// 参数：
//   - ctx: 交易上下文
//   - recallJSON: 召回数据JSON字符串，包含 recallID、batchNos、reason、affectedParts、status、createTime、finishTime
//
// 返回值：
//   - error: 操作成功返回 nil，失败返回具体错误信息
//
// 权限要求：所有组织成员可调用（建议由车企或厂商调用）
func (s *SmartContract) CreateRecallRecord(ctx contractapi.TransactionContextInterface, recallJSON string) error {
	// 反序列化召回数据
	var recall RecallRecord
	err := json.Unmarshal([]byte(recallJSON), &recall)
	if err != nil {
		return fmt.Errorf("反序列化召回数据失败: %v", err)
	}

	// 数据验证
	if recall.RecallID == "" || len(recall.BatchNos) == 0 {
		return fmt.Errorf("RecallID和BatchNos不能为空")
	}

	// 序列化并存储召回数据
	recallBytes, err := json.Marshal(recall)
	if err != nil {
		return fmt.Errorf("序列化召回数据失败: %v", err)
	}

	err = ctx.GetStub().PutState("RECALL_"+recall.RecallID, recallBytes)
	if err != nil {
		return err
	}

	// 遍历所有涉及批次，更新相关零部件状态为"召回"
	for _, batchNo := range recall.BatchNos {
		parts, err := s.QueryPartByBatchNo(ctx, batchNo)
		if err != nil {
			continue
		}

		for _, part := range parts {
			s.UpdatePartStatus(ctx, part.PartID, "召回")
		}
	}

	return nil
}

// CreateAftersaleRecord 创建售后记录
// 权限要求：仅4S店/售后中心可执行
// 功能：创建售后记录（维修处理、召回登记、报废处理等）
// 参数：
//   - ctx: 交易上下文
//   - aftersaleJSON: 售后记录JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
//
// CreateAftersaleRecord 创建售后记录
// 权限要求：仅4S店/售后中心可执行
// 功能：创建售后记录并更新零部件生命周期
// 参数：
//   - ctx: 交易上下文
//   - aftersaleJSON: 售后记录JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
func (s *SmartContract) CreateAftersaleRecord(ctx contractapi.TransactionContextInterface, aftersaleJSON string) error {
	// 1. 校验调用者身份
	if err := s.checkPermission(ctx, AFTERSALE_ORG_MSPID); err != nil {
		log.Printf("权限验证失败，错误：%v", err)
		return err
	}

	// 2. 反序列化售后数据
	var aftersale AftersaleRecord
	err := json.Unmarshal([]byte(aftersaleJSON), &aftersale)
	if err != nil {
		log.Printf("反序列化售后数据失败，错误：%v", err)
		return fmt.Errorf("反序列化售后数据失败: %v", err)
	}

	// 3. 数据验证
	if aftersale.AftersaleID == "" || aftersale.PartID == "" {
		log.Printf("售后数据不完整，AftersaleID：%s，PartID：%s", aftersale.AftersaleID, aftersale.PartID)
		return fmt.Errorf("AftersaleID和PartID不能为空")
	}
	if aftersale.AftersaleType != "故障报修" && aftersale.AftersaleType != "维修处理" && aftersale.AftersaleType != "召回登记" && aftersale.AftersaleType != "报废处理" {
		log.Printf("售后类型无效，AftersaleID：%s，类型：%s", aftersale.AftersaleID, aftersale.AftersaleType)
		return fmt.Errorf("售后类型必须是：故障报修、维修处理、召回登记或报废处理")
	}
	if aftersale.AftersaleStatus != "待审核" && aftersale.AftersaleStatus != "处理中" && aftersale.AftersaleStatus != "已完成" {
		log.Printf("售后状态无效，AftersaleID：%s，状态：%s", aftersale.AftersaleID, aftersale.AftersaleStatus)
		return fmt.Errorf("售后状态必须是：待审核、处理中或已完成")
	}

	// 4. 检查零部件是否存在
	partBytes, err := ctx.GetStub().GetState(aftersale.PartID)
	if err != nil {
		log.Printf("查询零部件失败，PartID：%s，错误：%v", aftersale.PartID, err)
		return fmt.Errorf("查询零部件失败: %v", err)
	}
	if partBytes == nil {
		log.Printf("零部件不存在，PartID：%s", aftersale.PartID)
		return fmt.Errorf("零部件ID %s 不存在", aftersale.PartID)
	}

	// 5. 设置处理时间
	if aftersale.ProcessTime == "" {
		aftersale.ProcessTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	// 6. 序列化并存储售后记录
	aftersaleBytes, err := json.Marshal(aftersale)
	if err != nil {
		log.Printf("序列化售后数据失败，AftersaleID：%s，错误：%v", aftersale.AftersaleID, err)
		return fmt.Errorf("序列化售后数据失败: %v", err)
	}
	err = ctx.GetStub().PutState("AFTERSALE_"+aftersale.AftersaleID, aftersaleBytes)
	if err != nil {
		log.Printf("存储售后记录失败，AftersaleID：%s，错误：%v", aftersale.AftersaleID, err)
		return err
	}

	// 7. 更新零部件生命周期
	lifecycleBytes, err := ctx.GetStub().GetState("LIFECYCLE_" + aftersale.PartID)
	var lifecycle PartLifecycle
	if err == nil && lifecycleBytes != nil {
		json.Unmarshal(lifecycleBytes, &lifecycle)
	}
	lifecycle.PartID = aftersale.PartID
	if lifecycle.AftersaleInfo == nil {
		lifecycle.AftersaleInfo = []AftersaleRecord{aftersale}
	} else {
		lifecycle.AftersaleInfo = append(lifecycle.AftersaleInfo, aftersale)
	}

	lifecycleBytes, err = json.Marshal(lifecycle)
	if err != nil {
		log.Printf("序列化生命周期数据失败，PartID：%s，错误：%v", aftersale.PartID, err)
		return fmt.Errorf("序列化生命周期数据失败: %v", err)
	}

	// 记录操作日志
	log.Printf("创建售后记录成功，AftersaleID：%s，PartID：%s，类型：%s", aftersale.AftersaleID, aftersale.PartID, aftersale.AftersaleType)
	if err := ctx.GetStub().SetEvent("CreateAftersaleRecord", []byte(fmt.Sprintf("创建售后记录: %s", aftersale.AftersaleID))); err != nil {
		log.Printf("设置事件失败：%v", err)
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState("LIFECYCLE_"+aftersale.PartID, lifecycleBytes)
}

// QueryAffectedParts 查询受影响零部件
// 权限要求：所有组织可查询
// 功能：查询指定批次的所有零部件（用于召回场景）
// 参数：
//   - ctx: 交易上下文
//   - batchNo: 批次号
//
// 返回：
//   - []*Part: 零部件列表
//   - error: 查询失败时返回错误
func (s *SmartContract) QueryAffectedParts(ctx contractapi.TransactionContextInterface, batchNo string) ([]*Part, error) {
	return s.QueryPartByBatchNo(ctx, batchNo)
}

// QueryAftersaleRecord 查询售后记录
// 权限要求：所有组织可查询
// 功能：根据售后记录ID查询售后详情
// 参数：
//   - ctx: 交易上下文
//   - aftersaleID: 售后记录ID
//
// 返回：
//   - *AftersaleRecord: 售后记录
//   - error: 查询失败时返回错误
func (s *SmartContract) QueryAftersaleRecord(ctx contractapi.TransactionContextInterface, aftersaleID string) (*AftersaleRecord, error) {
	// 1. 从账本获取售后记录数据
	aftersaleBytes, err := ctx.GetStub().GetState("AFTERSALE_" + aftersaleID)
	if err != nil {
		return nil, fmt.Errorf("查询售后记录失败: %v", err)
	}
	if aftersaleBytes == nil {
		return nil, fmt.Errorf("售后记录ID %s 不存在", aftersaleID)
	}

	// 2. 反序列化售后记录数据
	var aftersale AftersaleRecord
	err = json.Unmarshal(aftersaleBytes, &aftersale)
	if err != nil {
		return nil, fmt.Errorf("反序列化售后记录失败: %v", err)
	}

	return &aftersale, nil
}

// RegisterUser 注册用户身份
// 权限要求：所有组织成员可执行
// 功能：注册用户数字身份，分配组织和权限
// 参数：
//   - ctx: 交易上下文
//   - userJSON: 用户身份JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userJSON string) error {
	// 1. 校验调用者身份
	clientMSPID, err := s.getClientIdentityMSPID(ctx)
	if err != nil {
		log.Printf("获取调用者身份失败，错误：%v", err)
		return fmt.Errorf("获取调用者身份失败：%v", err)
	}
	if clientMSPID != MANUFACTURER_ORG_MSPID && clientMSPID != AUTOMAKER_ORG_MSPID && clientMSPID != AFTERSALE_ORG_MSPID {
		log.Printf("权限不足，MSPID：%s", clientMSPID)
		return fmt.Errorf("只有零部件生产厂商、整车车企或4S店/售后中心组织的成员才能注册用户")
	}

	// 2. 反序列化用户数据
	var user UserIdentity
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		log.Printf("反序列化用户数据失败，错误：%v", err)
		return fmt.Errorf("反序列化用户数据失败: %v", err)
	}

	// 3. 数据验证
	if user.UserID == "" || user.Org == "" {
		log.Printf("用户数据不完整，UserID：%s，Org：%s", user.UserID, user.Org)
		return fmt.Errorf("UserID和Org不能为空")
	}
	if user.Org != "零部件生产厂商" && user.Org != "整车车企" && user.Org != "4S店/售后中心" {
		log.Printf("组织无效，UserID：%s，Org：%s", user.UserID, user.Org)
		return fmt.Errorf("组织必须是：零部件生产厂商、整车车企或4S店/售后中心")
	}

	// 4. 检查用户是否已存在
	existingBytes, err := ctx.GetStub().GetState("USER_" + user.UserID)
	if err != nil {
		log.Printf("查询用户失败，UserID：%s，错误：%v", user.UserID, err)
		return fmt.Errorf("查询用户失败: %v", err)
	}
	if existingBytes != nil {
		log.Printf("用户已存在，UserID：%s", user.UserID)
		return fmt.Errorf("用户ID %s 已存在", user.UserID)
	}

	// 5. 设置注册时间
	if user.RegisterTime == "" {
		user.RegisterTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	// 6. 序列化并存储用户数据
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("序列化用户数据失败，UserID：%s，错误：%v", user.UserID, err)
		return fmt.Errorf("序列化用户数据失败: %v", err)
	}

	// 记录操作日志
	log.Printf("注册用户成功，UserID：%s，Org：%s，Role：%s", user.UserID, user.Org, user.Role)
	if err := ctx.GetStub().SetEvent("RegisterUser", []byte(fmt.Sprintf("注册用户: %s", user.UserID))); err != nil {
		log.Printf("设置事件失败：%v", err)
		return fmt.Errorf("设置事件失败: %v", err)
	}

	return ctx.GetStub().PutState("USER_"+user.UserID, userBytes)
}

// QueryUserPermissions 查询用户权限
// 权限要求：所有组织可查询
// 功能：根据用户ID查询用户权限信息
// 参数：
//   - ctx: 交易上下文
//   - userID: 用户ID
//
// 返回：
//   - *UserIdentity: 用户权限信息
//   - error: 查询失败时返回错误
func (s *SmartContract) QueryUserPermissions(ctx contractapi.TransactionContextInterface, userID string) (*UserIdentity, error) {
	userBytes, err := ctx.GetStub().GetState("USER_" + userID)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("用户ID %s 不存在", userID)
	}

	var user UserIdentity
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, fmt.Errorf("反序列化用户失败: %v", err)
	}

	return &user, nil
}

// RecordQAInteraction 记录问答交互上链
// 权限要求：所有组织可执行
// 功能：记录AI智能问答的交互历史，用于审计和追溯
// 参数：
//   - ctx: 交易上下文
//   - qaJSON: 问答交互JSON字符串
//
// 返回：
//   - error: 操作失败时返回错误
func (s *SmartContract) RecordQAInteraction(ctx contractapi.TransactionContextInterface, qaJSON string) error {
	// 1. 反序列化问答数据
	var qa QAInteraction
	err := json.Unmarshal([]byte(qaJSON), &qa)
	if err != nil {
		return fmt.Errorf("反序列化问答数据失败: %v", err)
	}

	// 2. 数据验证
	if qa.QAID == "" || qa.UserID == "" {
		return fmt.Errorf("QAID和UserID不能为空")
	}

	// 3. 设置查询时间
	if qa.QueryTime == "" {
		qa.QueryTime = fmt.Sprintf("%d", time.Now().Unix())
	}

	// 4. 序列化并存储问答记录
	qaBytes, err := json.Marshal(qa)
	if err != nil {
		return fmt.Errorf("序列化问答数据失败: %v", err)
	}
	return ctx.GetStub().PutState("QA_"+qa.QAID, qaBytes)
}

// QueryQAInteraction 查询问答交互记录
// 权限要求：所有组织可查询
// 功能：根据问答ID查询问答交互记录
// 参数：
//   - ctx: 交易上下文
//   - qaID: 问答记录ID
//
// 返回：
//   - *QAInteraction: 问答交互记录
//   - error: 查询失败时返回错误
func (s *SmartContract) QueryQAInteraction(ctx contractapi.TransactionContextInterface, qaID string) (*QAInteraction, error) {
	// 1. 从账本获取问答记录数据
	qaBytes, err := ctx.GetStub().GetState("QA_" + qaID)
	if err != nil {
		return nil, fmt.Errorf("查询问答记录失败: %v", err)
	}
	if qaBytes == nil {
		return nil, fmt.Errorf("问答记录ID %s 不存在", qaID)
	}

	// 2. 反序列化问答记录数据
	var qa QAInteraction
	err = json.Unmarshal(qaBytes, &qa)
	if err != nil {
		return nil, fmt.Errorf("反序列化问答记录失败: %v", err)
	}

	return &qa, nil
}

// QueryFaultProgress 查询故障处理进度
// 权限要求：所有组织可查询
// 功能：根据故障ID查询故障处理状态（用于AI智能问答系统）
// 参数：
//   - ctx: 交易上下文
//   - faultID: 故障ID
//
// 返回：
//   - *FaultReport: 故障处理进度信息
//   - error: 查询失败时返回错误
func (s *SmartContract) QueryFaultProgress(ctx contractapi.TransactionContextInterface, faultID string) (*FaultReport, error) {
	// 1. 从账本获取故障数据
	faultBytes, err := ctx.GetStub().GetState("FAULT_" + faultID)
	if err != nil {
		return nil, fmt.Errorf("查询故障失败: %v", err)
	}
	if faultBytes == nil {
		return nil, fmt.Errorf("故障ID %s 不存在", faultID)
	}

	// 2. 反序列化故障数据
	var fault FaultReport
	err = json.Unmarshal(faultBytes, &fault)
	if err != nil {
		return nil, fmt.Errorf("反序列化故障失败: %v", err)
	}

	return &fault, nil
}

// main 主函数
// 功能：创建并启动智能合约链码
// 注意：这是链码的入口点，Fabric网络会调用此函数来启动链码
func main() {
	// 创建智能合约实例
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("创建智能合约失败：%v", err)
	}

	// 启动智能合约
	if err := chaincode.Start(); err != nil {
		log.Panicf("启动智能合约失败：%v", err)
	}
}
