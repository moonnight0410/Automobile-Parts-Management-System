package modules

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

//SupplyChainData 供应关联数据
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
