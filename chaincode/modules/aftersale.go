package modules

//售后模块结构体
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
	Description   string `json:"description"`    // 售后描述
	ProcessResult string `json:"processResult"`  // 处理结果
	ProcessTime   string `json:"processTime"`    // 处理完成时间
	Cost          string `json:"cost,omitempty"` // 售后成本

	// 5. 通用附件复用
	AttachmentInfo *Attachment `json:"attachmentInfo,omitempty"` // 复用通用附件
	Remark         string      `json:"remark,omitempty"`         // 售后备注
}
