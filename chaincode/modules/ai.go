package modules

// QAInteraction 问答交互记录（上链存证）
type QAInteraction struct {
	QAID           string `json:"qaid"`           // 问答记录ID
	UserID         string `json:"userID"`         // 提问用户ID
	Question       string `json:"question"`       // 问题内容
	Intent         string `json:"intent"`         // 意图分类（如"故障进度查询"）
	PartID         string `json:"partID"`         // 关联零部件ID（如有）
	Answer         string `json:"answer"`         // 回答内容
	QueryTime      string `json:"queryTime"`      // 提问时间
	ContractMethod string `json:"contractMethod"` // 调用的智能合约方法（如"QueryFaultProgress"）
}
