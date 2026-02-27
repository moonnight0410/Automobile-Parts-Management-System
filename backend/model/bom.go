package model

type BOM struct {
	BOMID             string             `json:"bomID"`
	BOMType           string             `json:"bomType"`
	ProductModel      string             `json:"productModel"`
	PartBatchNo       string             `json:"partBatchNo"`
	Version           string             `json:"version"`
	Creator           string             `json:"creator"`
	CreateTime        string             `json:"createTime"`
	Status            string             `json:"status"`
	MaterialList      []MaterialItem     `json:"materialList"`
	RDBOMFileInfo     *BOMFileInfo       `json:"rdBOMFileInfo,omitempty"`
	ProductionBOMInfo *ProductionBOMInfo `json:"productionBOMInfo,omitempty"`
	ChangeRecords     []BOMChangeRecord  `json:"changeRecords,omitempty"`
}

type BOMDTO struct {
	BOMID        string         `json:"bomID" binding:"required"`
	BOMType      string         `json:"bomType" binding:"required"`
	ProductModel string         `json:"productModel" binding:"required"`
	PartBatchNo  string         `json:"partBatchNo" binding:"required"`
	Version      string         `json:"version" binding:"required"`
	Creator      string         `json:"creator" binding:"required"`
	CreateTime   string         `json:"createTime"`
	Status       string         `json:"status"`
	MaterialList []MaterialItem `json:"materialList" binding:"required"`
}

type MaterialItem struct {
	MaterialID   string `json:"materialID"`
	MaterialName string `json:"materialName"`
	Spec         string `json:"spec"`
	Quantity     int    `json:"quantity"`
	SupplierID   string `json:"supplierID"`
}

type BOMFileInfo struct {
	FileStorageID  string `json:"fileStorageID"`
	FilePath       string `json:"filePath"`
	FileHash       string `json:"fileHash"`
	FileType       string `json:"fileType"`
	FileSize       int64  `json:"fileSize"`
	UploadTime     string `json:"uploadTime"`
	DatabaseSource string `json:"databaseSource"`
}

type ProductionBOMInfo struct {
	RDOMReferenceID string `json:"rdBOMReferenceID"`
	CompareResult   string `json:"compareResult"`
	CompareDetails  string `json:"compareDetails"`
	Verifier        string `json:"verifier"`
	VerifyTime      string `json:"verifyTime"`
}

type BOMChangeRecord struct {
	ChangeID     string      `json:"changeID"`
	ChangeReason string      `json:"changeReason"`
	OldVersion   string      `json:"oldVersion"`
	NewVersion   string      `json:"newVersion"`
	AuditFlow    []AuditNode `json:"auditFlow"`
	ChangeTime   string      `json:"changeTime"`
}

type AuditNode struct {
	Auditor     string `json:"auditor"`
	AuditRole   string `json:"auditRole"`
	AuditResult string `json:"auditResult"`
	AuditTime   string `json:"auditTime"`
	Remark      string `json:"remark,omitempty"`
}
