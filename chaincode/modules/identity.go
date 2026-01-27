package modules

// UserIdentity 数字身份
type UserIdentity struct {
	UserID       string   `json:"userID"`       // 用户唯一标识
	Org          string   `json:"org"`          // 所属组织（厂商/物流/车企/售后）
	Role         string   `json:"role"`         // 角色（如"厂商-生产员"、"车企-审核员"）
	CertHash     string   `json:"certHash"`     // 数字证书哈希（CA颁发）
	Permissions  []string `json:"permissions"`  // 权限列表（如"上传生产数据"、"查询召回"）
	RegisterTime string   `json:"registerTime"` // 注册时间
}
