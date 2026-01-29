// InitLedger初始化账本
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	testPart := modules.Part{
		PartID:       "ENG-PISTON-001",
		VIN:          "LVX1234568789798",
		BatchNo:      "B20260127",
		Name:         "发动机活塞",
		Type:         "发动机部件",
		Manufacturer: "厂商A",
		CreateTime:   fmt.Sprintf("%d", time.Now().Unix()),
		Status:       "NORMAL",
	}
	partJson, err := json.Marshal(testPart)
	if err != nil {
		return fmt.Errorf("序列化测试零部件失败: %v", err)
	}

	return ctx.GetStub().PutState(testPart.PartID, partJson)
}

// 通用方法: 获取客户端身份信息
func (s *SmartContract) getClientIdentityMSPID(ctx contractapi.TransactionContextInterface) (string, error) {
	clientID, err := cid.New(ctx.GetStub())
	if err != nil {
		return "", fmt.Errorf("获取客户端身份信息失败：%v", err)
	}
	return clientID.GetMSPID()
}

// CreatePart 创建新零部件
func (s *SmartContract) CreatePart(ctx contractapi.TransactionContextInterface, partJSON string) error {
	//校验调用者身份
	clientMSPID, err := s.getClientIdentityMSPID(ctx)
	if err != nil {
		return fmt.Errorf("获取调用者身份失败：%v", err)
	}
	// 验证是否是零部件生产厂商组织的成员
	if clientMSPID != modules.MANUFACTURER_ORG_MSPID {
		return fmt.Errorf("只有零部件生产厂商组织的成员才能创建零部件信息")
	}

	//将前端传入的JSON字符串反序列化为Part结构体
	var part modules.Part
	err = json.Unmarshal([]byte(partJSON), &part)
	if err != nil {
		return fmt.Errorf("反序列化零部件数据失败: %v", err)
	}

	//校验核心字段（可以在后端部分校验）
	if part.PartID == "" || part.BatchNo == "" {
		return fmt.Errorf("PartID和BatchNo不能为空")
	}

	//序列化结构体并存入账本
	partBytes, err := json.Marshal(part)
	if err != nil {
		return fmt.Errorf("序列化零部件数据失败: %v", err)
	}
	return ctx.GetStub().PutState(part.PartID, partBytes)
}

// QueryPart 查询零部件
func (s *SmartContract) QueryPart(ctx contractapi.TransactionContextInterface, partID string) (*modules.Part, error) {
	//从账本获取核心part数据
	partBytes, err := ctx.GetStub().GetState(partID)
	if err != nil {
		return nil, fmt.Errorf("查询零部件失败: %v", err)
	}
	if partBytes == nil {
		return nil, fmt.Errorf("零部件ID %s 不存在", partID)
	}

	//反序列化Part数据
	var part modules.Part
	err = json.Unmarshal(partBytes, &part)
	if err != nil {
		return nil, fmt.Errorf("反序列化零部件失败: %v", err)
	}

	return &part, nil
}
