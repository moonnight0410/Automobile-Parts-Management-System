package model

type ProductionData struct {
	ProductionID   string            `json:"productionID"`
	PartID         string            `json:"partID"`
	BatchNo        string            `json:"batchNo"`
	Params         map[string]string `json:"params"`
	ProductionLine string            `json:"productionLine"`
	Operator       string            `json:"operator"`
	FinishTime     string            `json:"finishTime"`
}

type ProductionDataDTO struct {
	ProductionID   string            `json:"productionID" binding:"required"`
	PartID         string            `json:"partID" binding:"required"`
	BatchNo        string            `json:"batchNo" binding:"required"`
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

type QualityInspectionDTO struct {
	InspectionID string            `json:"inspectionID" binding:"required"`
	PartID       string            `json:"partID" binding:"required"`
	BatchNo      string            `json:"batchNo" binding:"required"`
	Indicators   map[string]string `json:"indicators"`
	Result       string            `json:"result" binding:"required"`
	Handler      string            `json:"handler"`
	HandleTime   string            `json:"handleTime"`
	Disposal     string            `json:"disposal"`
}
