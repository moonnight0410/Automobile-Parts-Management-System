package modules

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
