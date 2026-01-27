package modules

// BOM 物料清单主结构体（链上存储核心数据）
type BOM struct {
	// 基础标识字段
	BOMID        string `json:"bomID"`        // BOM唯一ID（如研发BOM：R-BOM-202505-001；生产BOM：P-BOM-202505-001）
	BOMType      string `json:"bomType"`      // BOM类型：研发BOM/生产BOM
	ProductModel string `json:"productModel"` // 对应车型（如2025款XX车型）
	PartBatchNo  string `json:"partBatchNo"`  // 零部件批次号（关联溯源模块）
	Version      string `json:"version"`      // BOM版本号（如V1.0、V1.1）
	Creator      string `json:"creator"`      // 创建人（企业数字身份标识）
	CreateTime   string `json:"createTime"`   // 创建时间（时间戳）
	Status       string `json:"status"`       // 状态：草稿/已发布/已变更/已作废

	// 物料清单核心字段
	MaterialList []MaterialItem `json:"materialList"` // 物料明细列表

	// 研发BOM专属：图纸/设计文件关联（适配普通数据库存储）
	RDBOMFileInfo *BOMFileInfo `json:"rdBOMFileInfo,omitempty"` // 研发BOM图纸文件信息（omitempty：生产BOM可忽略）

	// 生产BOM专属：与研发BOM比对
	ProductionBOMInfo *ProductionBOMInfo `json:"productionBOMInfo,omitempty"` // 生产BOM比对信息

	// BOM变更相关
	ChangeRecords []BOMChangeRecord `json:"changeRecords,omitempty"` // 变更记录（版本迭代/变更申请）
}

// MaterialItem 物料明细项
type MaterialItem struct {
	MaterialID   string `json:"materialID"`   // 物料唯一ID
	MaterialName string `json:"materialName"` // 物料名称（如发动机活塞）
	Spec         string `json:"spec"`         // 规格型号
	Quantity     int    `json:"quantity"`     // 数量
	SupplierID   string `json:"supplierID"`   // 供应商ID（关联供应链模块）
}

// BOMFileInfo 研发BOM图纸/文件信息（适配普通数据库存储）
type BOMFileInfo struct {
	FileStorageID  string `json:"fileStorageID"`  // 普通数据库中的文件唯一ID（如MySQL的主键ID）
	FilePath       string `json:"filePath"`       // 文件存储路径（如：/data/bom/rd/202505/R-BOM-001-CAD.dwg）
	FileHash       string `json:"fileHash"`       // 文件内容哈希（SHA256，链上存哈希，验证数据库文件未被篡改）
	FileType       string `json:"fileType"`       // 文件类型（如CAD图纸、PNG图片、Excel清单）
	FileSize       int64  `json:"fileSize"`       // 文件大小（字节，便于前端展示）
	UploadTime     string `json:"uploadTime"`     // 文件上传时间
	DatabaseSource string `json:"databaseSource"` // 存储数据库标识（如：mysql-bom-file-01，多数据库部署时区分）
}

// ProductionBOMInfo 生产BOM专属信息
type ProductionBOMInfo struct {
	RDOMReferenceID string `json:"rdBOMReferenceID"` // 关联的研发BOM ID
	CompareResult   string `json:"compareResult"`    // 比对结果：一致/不一致
	CompareDetails  string `json:"compareDetails"`   // 比对详情（如：物料A规格不一致，研发为XX，生产为YY）
	Verifier        string `json:"verifier"`         // 核验人（数字身份）
	VerifyTime      string `json:"verifyTime"`       // 核验时间
}

// BOMChangeRecord BOM变更记录
type BOMChangeRecord struct {
	ChangeID     string      `json:"changeID"`     // 变更申请ID
	ChangeReason string      `json:"changeReason"` // 变更原因（如设计优化、物料替代）
	OldVersion   string      `json:"oldVersion"`   // 变更前版本
	NewVersion   string      `json:"newVersion"`   // 变更后版本
	AuditFlow    []AuditNode `json:"auditFlow"`    // 审批流程
	ChangeTime   string      `json:"changeTime"`   // 变更生效时间
}

// AuditNode 审批节点
type AuditNode struct {
	Auditor     string `json:"auditor"`          // 审批人
	AuditRole   string `json:"auditRole"`        // 审批角色（如研发负责人、生产负责人）
	AuditResult string `json:"auditResult"`      // 审批结果：通过/驳回
	AuditTime   string `json:"auditTime"`        // 审批时间
	Remark      string `json:"remark,omitempty"` // 审批备注
}
