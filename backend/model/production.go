package model

import "encoding/json"

type ProductionData struct {
	ProductionID   string            `json:"productionID"`
	PartID         string            `json:"partID"`
	BatchNo        string            `json:"batchNo"`
	Params         map[string]string `json:"params"`
	ProductionLine string            `json:"productionLine"`
	Operator       string            `json:"operator"`
	FinishTime     string            `json:"finishTime"`
}

// UnmarshalJSON 处理可能为null的params字段
func (p *ProductionData) UnmarshalJSON(data []byte) error {
	type Alias ProductionData
	aux := &struct {
		Params map[string]string `json:"params"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Params == nil {
		p.Params = make(map[string]string)
	} else {
		p.Params = aux.Params
	}
	return nil
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

// UnmarshalJSON 处理可能为null的indicators字段
func (q *QualityInspection) UnmarshalJSON(data []byte) error {
	type Alias QualityInspection
	aux := &struct {
		Indicators map[string]string `json:"indicators"`
		*Alias
	}{
		Alias: (*Alias)(q),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Indicators == nil {
		q.Indicators = make(map[string]string)
	} else {
		q.Indicators = aux.Indicators
	}
	return nil
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
