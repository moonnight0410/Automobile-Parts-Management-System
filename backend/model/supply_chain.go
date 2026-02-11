package model

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

type SupplyOrderDTO struct {
	OrderID    string `json:"orderID" binding:"required"`
	Buyer      string `json:"buyer" binding:"required"`
	Seller     string `json:"seller" binding:"required"`
	PartID     string `json:"partID" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required"`
	BatchNo    string `json:"batchNo" binding:"required"`
	AgreedTime string `json:"agreedTime"`
	Status     string `json:"status"`
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

type LogisticsDataDTO struct {
	LogisticsID string `json:"logisticsID" binding:"required"`
	OrderID     string `json:"orderID" binding:"required"`
	Carrier     string `json:"carrier" binding:"required"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	GPSHash     string `json:"gpsHash"`
	Receiver    string `json:"receiver"`
}

type SupplyChainData struct {
	ChainID          string      `json:"chainID"`
	PartID           string      `json:"partID"`
	BatchNo          string      `json:"batchNo"`
	OrderID          string      `json:"orderID"`
	LogisticsID      string      `json:"logisticsID"`
	StageType        string      `json:"stageType"`
	StageStatus      string      `json:"stageStatus"`
	Participator     string      `json:"participator"`
	ParticipatorRole string      `json:"participatorRole"`
	Quantity         int         `json:"quantity"`
	OperateTime      string      `json:"operateTime"`
	Operator         string      `json:"operator"`
	AttachmentInfo   *Attachment `json:"attachmentInfo,omitempty"`
	Remark           string      `json:"remark,omitempty"`
}

type Attachment struct {
	FileStorageID  string `json:"fileStorageID"`
	FilePath       string `json:"filePath"`
	FileHash       string `json:"fileHash"`
	FileType       string `json:"fileType"`
	DatabaseSource string `json:"databaseSource"`
}
