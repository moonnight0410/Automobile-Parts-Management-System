package modules

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
	PartID          string            `json:"partID"`          // 关联零部件
	BOMInfo         BOMReference      `json:"bomInfo"`         // 关联BOM信息
	ProductionInfo  ProductionData    `json:"productionInfo"`  // 关联生产数据
	QualityInfo     QualityInspection `json:"qualityInfo"`     // 关联质检数据
	SupplyChainInfo []SupplyChainData `json:"supplyChainInfo"` // 关联供应链数据（订单+物流）
	AftersaleInfo   []AftersaleRecord `json:"aftersaleInfo"`   // 关联售后/故障数据
}

// BOMReference BOM信息引用
type BOMReference struct {
	BOMID   string `json:"bomID"`   // 关联BOM的ID
	Version string `json:"version"` // 关联BOM版本
	Type    string `json:"type"`    // 研发BOM/生产BOM
}
