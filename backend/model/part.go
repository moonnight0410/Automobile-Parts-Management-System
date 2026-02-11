package model

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

type PartDTO struct {
	PartID       string `json:"partID" binding:"required"`
	VIN          string `json:"vin"`
	BatchNo      string `json:"batchNo" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Type         string `json:"type" binding:"required"`
	Manufacturer string `json:"manufacturer" binding:"required"`
	Status       string `json:"status"`
}

type PartLifecycle struct {
	PartID          string            `json:"partID"`
	BOMInfo         BOMReference      `json:"bomInfo"`
	ProductionInfo  ProductionData    `json:"productionInfo"`
	QualityInfo     QualityInspection `json:"qualityInfo"`
	SupplyChainInfo []SupplyChainData `json:"supplyChainInfo"`
	AftersaleInfo   []AftersaleRecord `json:"aftersaleInfo"`
}

type BOMReference struct {
	BOMID   string `json:"bomID"`
	Version string `json:"version"`
	Type    string `json:"type"`
}
