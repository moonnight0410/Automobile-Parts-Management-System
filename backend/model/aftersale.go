package model

import "encoding/json"

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

type FaultReportDTO struct {
	FaultID     string  `json:"faultID" binding:"required"`
	PartID      string  `json:"partID" binding:"required"`
	VIN         string  `json:"vin"`
	Reporter    string  `json:"reporter" binding:"required"`
	Description string  `json:"description" binding:"required"`
	FaultType   string  `json:"faultType"`
	RiskProb    float64 `json:"riskProb"`
	Status      string  `json:"status"`
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

// UnmarshalJSON 自定义反序列化，处理 affectedParts 为 null 的情况
func (r *RecallRecord) UnmarshalJSON(data []byte) error {
	type Alias RecallRecord
	aux := &struct {
		AffectedParts []string `json:"affectedParts"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.AffectedParts == nil {
		r.AffectedParts = []string{}
	} else {
		r.AffectedParts = aux.AffectedParts
	}
	return nil
}

type RecallRecordDTO struct {
	RecallID      string   `json:"recallID" binding:"required"`
	BatchNos      []string `json:"batchNos" binding:"required"`
	Reason        string   `json:"reason" binding:"required"`
	AffectedParts []string `json:"affectedParts"`
	Status        string   `json:"status"`
}

type AftersaleRecord struct {
	AftersaleID     string      `json:"aftersaleID"`
	PartID          string      `json:"partID"`
	BatchNo         string      `json:"batchNo"`
	VIN             string      `json:"vin"`
	FaultID         string      `json:"faultID"`
	RecallID        string      `json:"recallID"`
	AftersaleType   string      `json:"aftersaleType"`
	AftersaleStatus string      `json:"aftersaleStatus"`
	HandlerOrgID    string      `json:"handlerOrgID"`
	HandlerID       string      `json:"handlerID"`
	Description     string      `json:"description"`
	ProcessResult   string      `json:"processResult"`
	ProcessTime     string      `json:"processTime"`
	Cost            string      `json:"cost,omitempty"`
	AttachmentInfo  *Attachment `json:"attachmentInfo,omitempty"`
	Remark          string      `json:"remark,omitempty"`
}

type AftersaleRecordDTO struct {
	AftersaleID     string      `json:"aftersaleID" binding:"required"`
	PartID          string      `json:"partID" binding:"required"`
	BatchNo         string      `json:"batchNo"`
	VIN             string      `json:"vin"`
	FaultID         string      `json:"faultID"`
	RecallID        string      `json:"recallID"`
	AftersaleType   string      `json:"aftersaleType" binding:"required"`
	AftersaleStatus string      `json:"aftersaleStatus"`
	HandlerOrgID    string      `json:"handlerOrgID"`
	HandlerID       string      `json:"handlerID"`
	Description     string      `json:"description"`
	ProcessResult   string      `json:"processResult"`
	ProcessTime     string      `json:"processTime"`
	Cost            string      `json:"cost,omitempty"`
	AttachmentInfo  *Attachment `json:"attachmentInfo,omitempty"`
	Remark          string      `json:"remark,omitempty"`
}
