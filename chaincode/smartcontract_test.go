package main

import (
	"testing"

	"github.com/golang/protobuf/proto"                // 导入protobuf序列化包
	"github.com/hyperledger/fabric-chaincode-go/shim" // 导入shim包（用于shim.OK）
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/msp"
)

// TestCreatePart 测试创建零部件功能（正常流程）
// 验证 ManufacturerMSP 成员能够成功创建零部件信息
func TestCreatePart(t *testing.T) {
	// 0. 设置测试模式标志
	// 在测试开始前设置全局 TestMode 为 true，使链码使用测试环境的身份验证
	TestMode = true
	// 测试结束后恢复为 false，避免影响其他测试
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 构造身份对象
	// 使用 msp.SerializedIdentity 结构体模拟 Fabric 网络中的身份
	identity := &msp.SerializedIdentity{
		Mspid:   "ManufacturerMSP",
		IdBytes: []byte("test-user-org1"),
	}
	// 将身份对象序列化为字节数组，用于设置 MockStub.Creator
	creatorBytes, err := proto.Marshal(identity)
	if err != nil {
		t.Fatalf("序列化身份对象失败：%v", err)
	}

	// 3. 创建 MockStub 并设置 Creator
	// MockStub 是 Fabric 提供的测试工具，用于模拟链码执行环境
	mockStub := shimtest.NewMockStub("realty-chaincode", cc)
	// 设置 Creator 字段，模拟交易发起者的身份信息
	mockStub.Creator = creatorBytes

	// 4. 构造测试数据
	// 创建一个零部件的 JSON 字符串，包含所有必需字段
	testPartJSON := `{"partID":"ENG-PISTON-002","vin":"LVX432109876543210","batchNo":"B20250502","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商B","createTime":"1735689600","status":"NORMAL"}`

	// 5. 调用链码方法
	// 构造参数数组：第一个元素是方法名，第二个元素是方法参数
	args := [][]byte{[]byte("CreatePart"), []byte(testPartJSON)}
	// 使用 MockInvoke 模拟链码调用，传入交易ID "tx1"
	response := mockStub.MockInvoke("tx1", args)

	// 6. 校验调用结果
	// 检查返回状态是否为 OK (200)
	if response.Status != shim.OK {
		t.Fatalf("CreatePart执行失败：状态码=%d，信息=%s", response.Status, response.Message)
	}

	// 7. 验证数据是否写入账本
	// 使用 GetState 查询刚才创建的零部件数据
	result, err := mockStub.GetState("ENG-PISTON-002")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("零部件数据未写入账本")
	}
	// 输出测试通过信息，打印写入的数据
	t.Log("测试通过，数据：", string(result))
}

// TestCreatePart_Unauthorized 测试创建零部件的权限控制
// 验证非 ManufacturerMSP 成员（如 AutomakerMSP）无法创建零部件，应该被拒绝
func TestCreatePart_Unauthorized(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 构造 AutomakerMSP 身份对象
	// AutomakerMSP 代表整车车企组织，不应该有权限创建零部件
	identity := &msp.SerializedIdentity{
		Mspid:   "AutomakerMSP",
		IdBytes: []byte("test-user-org2"),
	}
	// 序列化身份对象
	creatorBytes, err := proto.Marshal(identity)
	if err != nil {
		t.Fatalf("序列化身份对象失败：%v", err)
	}

	// 3. 创建 MockStub 并设置 Creator
	mockStub := shimtest.NewMockStub("realty-chaincode", cc)
	// 设置为 AutomakerMSP 成员
	mockStub.Creator = creatorBytes

	// 4. 构造测试数据
	// 创建另一个零部件，使用不同的 ID
	testPartJSON := `{"partID":"ENG-PISTON-003","vin":"LVX432109876543211","batchNo":"B20250503","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商C","createTime":"1735689600","status":"NORMAL"}`

	// 5. 调用链码方法
	args := [][]byte{[]byte("CreatePart"), []byte(testPartJSON)}
	// 使用不同的交易ID "tx2"
	response := mockStub.MockInvoke("tx2", args)

	// 6. 验证权限控制是否生效
	// 预期调用应该失败，因为 AutomakerMSP 没有权限创建零部件
	if response.Status == shim.OK {
		t.Fatal("预期CreatePart应该失败（非ManufacturerMSP成员），但执行成功")
	}
	// 输出测试通过信息，打印返回的错误消息
	t.Log("测试通过，预期失败：", response.Message)
}

// TestCreateLogisticsData 测试创建物流数据功能（正常流程）
// 验证 AutomakerMSP 成员（整车车企）能够成功创建物流数据
func TestCreateLogisticsData(t *testing.T) {
	// 0. 设置测试模式标志
	// 在测试开始前设置全局 TestMode 为 true，使链码使用测试环境的身份验证
	TestMode = true
	// 测试结束后恢复为 false，避免影响其他测试
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AutomakerMSP 身份
	// AutomakerMSP 代表整车车企组织，有权限创建物流数据
	mockStub, err := createMockStubWithIdentity("AutomakerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	// 创建一个物流数据的 JSON 字符串，包含物流ID、订单ID、承运商、时间等信息
	logisticsJSON := `{"logisticsID":"LOGISTICS-001","orderID":"ORDER-001","carrier":"物流商A","startTime":"1735689600","endTime":"1735689700","gpsHash":"hash123","receiver":"车企A"}`

	// 4. 调用链码方法
	// 构造参数数组：第一个元素是方法名，第二个元素是方法参数
	args := [][]byte{[]byte("CreateLogisticsData"), []byte(logisticsJSON)}
	// 使用 MockInvoke 模拟链码调用
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	// 检查返回状态是否为 OK (200)
	if response.Status != shim.OK {
		t.Fatalf("CreateLogisticsData执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	// 使用 GetState 查询刚才创建的物流数据，键格式为 "LOGISTICS_" + logisticsID
	result, err := mockStub.GetState("LOGISTICS_LOGISTICS-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("物流数据未写入账本")
	}
	// 输出测试通过信息，打印写入的数据
	t.Log("测试通过，数据：", string(result))
}

// TestUpdateSupplyChainStage 测试更新供应链阶段功能（正常流程）
// 验证 AutomakerMSP 成员（整车车企）能够成功更新供应链阶段信息
func TestUpdateSupplyChainStage(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AutomakerMSP 身份
	// AutomakerMSP 代表整车车企组织，有权限更新供应链阶段
	mockStub, err := createMockStubWithIdentity("AutomakerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	// 创建一个供应链阶段的 JSON 字符串，包含阶段ID、零部件ID、阶段类型、状态等信息
	stageJSON := `{"chainID":"CHAIN-001","partID":"ENG-PISTON-001","batchNo":"B20250502","orderID":"ORDER-001","logisticsID":"","stageType":"采购下单","stageStatus":"已完成","participator":"车企A","participatorRole":"采购方","quantity":100,"operateTime":"1735689600","operator":"OP-001","attachmentInfo":null,"remark":""}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("UpdateSupplyChainStage"), []byte(stageJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("UpdateSupplyChainStage执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	// 键格式为 "CHAIN_" + chainID
	result, err := mockStub.GetState("CHAIN_CHAIN-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("供应链数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestCreateReconciliation 测试创建对账记录功能（正常流程）
// 验证 AutomakerMSP 成员（整车车企）能够成功创建对账记录
func TestCreateReconciliation(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AutomakerMSP 身份
	// AutomakerMSP 代表整车车企组织，有权限创建对账记录
	mockStub, err := createMockStubWithIdentity("AutomakerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	// 创建一个对账记录的 JSON 字符串，包含对账ID、订单ID、金额、状态等信息
	reconJSON := `{"reconID":"RECON-001","orderID":"ORDER-001","amount":"100000.00","status":"已结算","settleTime":"1735689600"}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("CreateReconciliation"), []byte(reconJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("CreateReconciliation执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	// 键格式为 "RECON_" + reconID
	result, err := mockStub.GetState("RECON_RECON-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("对账数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestCreateFaultReport 测试创建故障报告功能（正常流程）
// 验证 AftersaleMSP 成员（4S店/售后中心）能够成功创建故障报告
func TestCreateFaultReport(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AftersaleMSP 身份
	// AftersaleMSP 代表4S店/售后中心组织，有权限创建故障报告
	mockStub, err := createMockStubWithIdentity("AftersaleMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	// 创建一个故障报告的 JSON 字符串，包含故障ID、零部件ID、VIN码、描述等信息
	faultJSON := `{"faultID":"FAULT-001","partID":"ENG-PISTON-001","vin":"LVX1234568789798","reporter":"4S店A","description":"发动机活塞磨损严重","faultType":"磨损类","riskProb":0.85,"reportTime":"1735689600","status":"待审核"}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("CreateFaultReport"), []byte(faultJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("CreateFaultReport执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	// 键格式为 "FAULT_" + faultID
	result, err := mockStub.GetState("FAULT_FAULT-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("故障数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestCreateRecallRecord 测试创建召回记录功能（正常流程）
// 验证 AftersaleMSP 成员（4S店/售后中心）能够成功创建召回记录
func TestCreateRecallRecord(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AftersaleMSP 身份
	// AftersaleMSP 代表4S店/售后中心组织，有权限创建召回记录
	mockStub, err := createMockStubWithIdentity("AftersaleMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	// 创建一个召回记录的 JSON 字符串，包含召回ID、批次号、原因、受影响零部件等信息
	recallJSON := `{"recallID":"RECALL-001","batchNos":["B20250502","B20250503"],"reason":"质量问题","affectedParts":["ENG-PISTON-001","ENG-PISTON-002"],"status":"待通知","createTime":"1735689600","finishTime":""}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("CreateRecallRecord"), []byte(recallJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("CreateRecallRecord执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	// 键格式为 "RECALL_" + recallID
	result, err := mockStub.GetState("RECALL_RECALL-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("召回数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestCreateAftersaleRecord 测试创建售后记录功能（正常流程）
// 验证 AftersaleMSP 成员（4S店/售后中心）能够成功创建售后记录
func TestCreateAftersaleRecord(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AftersaleMSP 身份
	// AftersaleMSP 代表4S店/售后中心组织，有权限创建售后记录
	mockStub, err := createMockStubWithIdentity("AftersaleMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 预先写入零部件数据到账本
	// 由于 CreateAftersaleRecord 需要验证零部件是否存在，而创建零部件需要 ManufacturerMSP 身份
	// 因此使用 MockTransactionStart/End 直接写入数据，绕过身份验证
	mockStub.MockTransactionStart("tx0")
	partJSON := `{"partID":"ENG-PISTON-001","vin":"LVX1234568789798","batchNo":"B20250502","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	mockStub.PutState("ENG-PISTON-001", []byte(partJSON))
	mockStub.MockTransactionEnd("tx0")

	// 4. 构造测试数据
	// 创建一个售后记录的 JSON 字符串，包含售后ID、零部件ID、类型、处理结果等信息
	aftersaleJSON := `{"aftersaleID":"AFTER-001","partID":"ENG-PISTON-001","batchNo":"B20250502","vin":"LVX1234568789798","faultID":"FAULT-001","recallID":"","aftersaleType":"故障报修","aftersaleStatus":"已完成","handlerOrgID":"4S店A","handlerID":"TECH-001","description":"更换活塞","processResult":"维修完成","processTime":"1735689600","cost":"500.00","attachmentInfo":null,"remark":""}`

	// 5. 调用链码方法
	args := [][]byte{[]byte("CreateAftersaleRecord"), []byte(aftersaleJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 6. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("CreateAftersaleRecord执行失败：%s", response.Message)
	}

	// 7. 验证数据是否写入账本
	// 键格式为 "AFTERSALE_" + aftersaleID
	result, err := mockStub.GetState("AFTERSALE_AFTER-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("售后数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestRegisterUser 测试注册用户功能（正常流程）
// 验证各组织成员能够成功注册用户
func TestRegisterUser(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	// ManufacturerMSP 代表零部件生产厂商组织，有权限注册用户
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	// 创建一个用户信息的 JSON 字符串，包含用户ID、组织、角色、权限等信息
	// 注意：org 字段必须是"零部件生产厂商"、"整车车企"、"物流服务商"、"4S店/售后中心"或"行业监管机构"之一
	userJSON := `{"userID":"USER-001","org":"零部件生产厂商","role":"生产员","certHash":"hash123","permissions":["上传生产数据","查询零部件"],"registerTime":"1735689600"}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("RegisterUser"), []byte(userJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("RegisterUser执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	// 键格式为 "USER_" + userID
	result, err := mockStub.GetState("USER_USER-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("用户数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestQueryUserPermissions 测试查询用户权限功能（正常流程）
// 验证能够成功查询已注册用户的权限信息
func TestQueryUserPermissions(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先注册一个用户
	userJSON := `{"userID":"USER-001","org":"零部件生产厂商","role":"生产员","certHash":"hash123","permissions":["上传生产数据","查询零部件"],"registerTime":"1735689600"}`
	createArgs := [][]byte{[]byte("RegisterUser"), []byte(userJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryUserPermissions"), []byte("USER-001")}
	queryResponse := mockStub.MockInvoke("tx2", queryArgs)

	// 5. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryUserPermissions执行失败：%s", queryResponse.Message)
	}

	// 6. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}

// TestRecordQAInteraction 测试记录QA交互功能（正常流程）
// 验证能够成功记录用户与智能助手的问答交互
func TestRecordQAInteraction(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	// 创建一个QA交互记录的 JSON 字符串，包含QAID、用户ID、问题、答案、调用的方法等信息
	qaJSON := `{"qaid":"QA-001","userID":"USER-001","question":"如何查询零部件状态？","intent":"查询","partID":"","answer":"请使用QueryPart方法查询零部件状态","queryTime":"1735689600","contractMethod":"QueryPart"}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("RecordQAInteraction"), []byte(qaJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("RecordQAInteraction执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	// 键格式为 "QA_" + qaid
	result, err := mockStub.GetState("QA_QA-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("QA数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestQueryQAInteraction 测试查询QA交互功能（正常流程）
// 验证能够成功查询已记录的QA交互信息
func TestQueryQAInteraction(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先记录一个QA交互
	qaJSON := `{"qaid":"QA-001","userID":"USER-001","question":"如何查询零部件状态？","intent":"查询","partID":"","answer":"请使用QueryPart方法查询零部件状态","queryTime":"1735689600","contractMethod":"QueryPart"}`
	createArgs := [][]byte{[]byte("RecordQAInteraction"), []byte(qaJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryQAInteraction"), []byte("QA-001")}
	queryResponse := mockStub.MockInvoke("tx2", queryArgs)

	// 5. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryQAInteraction执行失败：%s", queryResponse.Message)
	}

	// 6. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}

// TestQueryFaultProgress 测试查询故障进度功能（正常流程）
// 验证能够成功查询故障报告的处理进度
func TestQueryFaultProgress(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AftersaleMSP 身份
	mockStub, err := createMockStubWithIdentity("AftersaleMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个故障报告
	faultJSON := `{"faultID":"FAULT-001","partID":"ENG-PISTON-001","vin":"LVX1234568789798","reporter":"4S店A","description":"发动机活塞磨损严重","faultType":"磨损类","riskProb":0.85,"reportTime":"1735689600","status":"待审核"}`
	createArgs := [][]byte{[]byte("CreateFaultReport"), []byte(faultJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryFaultProgress"), []byte("FAULT-001")}
	queryResponse := mockStub.MockInvoke("tx2", queryArgs)

	// 5. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryFaultProgress执行失败：%s", queryResponse.Message)
	}

	// 6. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}

// createMockStubWithIdentity 创建带有指定身份的 MockStub 辅助函数
// 参数：
//   - mspid: 组织 MSP ID（如 "ManufacturerMSP"、"AutomakerMSP" 等）
//   - cc: 智能合约实例
//
// 返回：
//   - *shimtest.MockStub: 配置好身份的 MockStub 实例
//   - error: 序列化失败时返回错误
func createMockStubWithIdentity(mspid string, cc *contractapi.ContractChaincode) (*shimtest.MockStub, error) {
	// 1. 构造身份对象
	// 使用 msp.SerializedIdentity 结构体模拟 Fabric 网络中的身份
	identity := &msp.SerializedIdentity{
		Mspid:   mspid,
		IdBytes: []byte("test-user-" + mspid),
	}
	// 2. 将身份对象序列化为字节数组
	creatorBytes, err := proto.Marshal(identity)
	if err != nil {
		return nil, err
	}

	// 3. 创建 MockStub 并设置 Creator
	// MockStub 是 Fabric 提供的测试工具，用于模拟链码执行环境
	mockStub := shimtest.NewMockStub("realty-chaincode", cc)
	// 设置 Creator 字段，模拟交易发起者的身份信息
	mockStub.Creator = creatorBytes
	return mockStub, nil
}

// TestCreateBOM 测试创建BOM功能（正常流程）
// 验证 ManufacturerMSP 成员（零部件生产厂商）能够成功创建BOM
func TestCreateBOM(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	// ManufacturerMSP 代表零部件生产厂商组织，有权限创建BOM
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	// 创建一个研发BOM的 JSON 字符串，包含BOMID、类型、物料清单等信息
	// 注意：rdBOMFileInfo、productionBOMInfo、changeRecords 字段需要满足schema要求
	// MaterialItem 结构体字段：materialID, materialName, spec, quantity, supplierID
	bomJSON := `{"bomID":"R-BOM-202505-001","bomType":"研发BOM","productModel":"2025款XX车型","partBatchNo":"B20250501","version":"V1.0","creator":"厂商A","createTime":"1735689600","status":"草稿","materialList":[{"materialID":"MAT-001","materialName":"发动机活塞","spec":"V1.0","quantity":100,"supplierID":"SUP-001"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("CreateBOM"), []byte(bomJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("CreateBOM执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	// 键格式为 "BOM_" + bomID
	result, err := mockStub.GetState("BOM_R-BOM-202505-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("BOM数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestCreateBOM_Unauthorized 测试创建BOM的权限控制
// 验证非 ManufacturerMSP 成员无法创建BOM，应该被拒绝
func TestCreateBOM_Unauthorized(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AutomakerMSP 身份
	// AutomakerMSP 代表整车车企组织，不应该有权限创建BOM
	mockStub, err := createMockStubWithIdentity("AutomakerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	bomJSON := `{"bomID":"R-BOM-202505-002","bomType":"研发BOM","productModel":"2025款XX车型","partBatchNo":"B20250502","version":"V1.0","creator":"厂商B","createTime":"1735689600","status":"草稿","materialList":[{"materialID":"MAT-002","materialName":"发动机活塞","spec":"V1.0","quantity":200,"supplierID":"SUP-002"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("CreateBOM"), []byte(bomJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 验证权限控制是否生效
	// 预期调用应该失败，因为 AutomakerMSP 没有权限创建BOM
	if response.Status == shim.OK {
		t.Fatal("预期CreateBOM应该失败（非ManufacturerMSP成员），但执行成功")
	}
	t.Log("测试通过，预期失败：", response.Message)
}

// TestQueryBOM 测试查询BOM功能（正常流程）
// 验证能够成功查询已创建的BOM信息
func TestQueryBOM(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个BOM
	// 注意：rdBOMFileInfo、productionBOMInfo 是结构体类型，需要用 {} 表示空对象
	// changeRecords 是数组类型，需要用 [] 表示空数组
	bomJSON := `{"bomID":"R-BOM-202505-003","bomType":"研发BOM","productModel":"2025款XX车型","partBatchNo":"B20250503","version":"V1.0","creator":"厂商A","createTime":"1735689600","status":"草稿","materialList":[{"materialID":"MAT-003","materialName":"发动机活塞","spec":"V1.0","quantity":150,"supplierID":"SUP-001"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`
	createArgs := [][]byte{[]byte("CreateBOM"), []byte(bomJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryBOM"), []byte("R-BOM-202505-003")}
	queryResponse := mockStub.MockInvoke("tx2", queryArgs)

	// 5. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryBOM执行失败：%s", queryResponse.Message)
	}

	// 6. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}

// TestUpdateBOM 测试更新BOM功能（正常流程）
// 验证 ManufacturerMSP 成员能够成功更新BOM信息
func TestUpdateBOM(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个BOM
	bomJSON := `{"bomID":"R-BOM-202505-004","bomType":"研发BOM","productModel":"2025款XX车型","partBatchNo":"B20250504","version":"V1.0","creator":"厂商A","createTime":"1735689600","status":"草稿","materialList":[{"materialID":"MAT-004","materialName":"发动机活塞","spec":"V1.0","quantity":100,"supplierID":"SUP-001"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`
	createArgs := [][]byte{[]byte("CreateBOM"), []byte(bomJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 构造更新数据
	// 更新状态为"已发布"，版本为"V1.1"
	updateBOMJSON := `{"bomID":"R-BOM-202505-004","bomType":"研发BOM","productModel":"2025款XX车型","partBatchNo":"B20250504","version":"V1.1","creator":"厂商A","createTime":"1735689600","status":"已发布","materialList":[{"materialID":"MAT-004","materialName":"发动机活塞","spec":"V1.0","quantity":120,"supplierID":"SUP-001"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`

	// 5. 调用更新方法
	updateArgs := [][]byte{[]byte("UpdateBOM"), []byte(updateBOMJSON)}
	updateResponse := mockStub.MockInvoke("tx2", updateArgs)

	// 6. 校验更新结果
	if updateResponse.Status != shim.OK {
		t.Fatalf("UpdateBOM执行失败：%s", updateResponse.Message)
	}

	// 7. 验证数据是否更新
	result, err := mockStub.GetState("BOM_R-BOM-202505-004")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("BOM数据未更新")
	}
	t.Log("测试通过，更新后数据：", string(result))
}

// TestUpdateBOM_Unauthorized 测试更新BOM的权限控制
// 验证非 ManufacturerMSP 成员无法更新BOM，应该被拒绝
func TestUpdateBOM_Unauthorized(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AutomakerMSP 身份
	mockStub, err := createMockStubWithIdentity("AutomakerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造更新数据
	updateBOMJSON := `{"bomID":"R-BOM-202505-005","bomType":"研发BOM","productModel":"2025款XX车型","partBatchNo":"B20250505","version":"V1.1","creator":"厂商A","createTime":"1735689600","status":"已发布","materialList":[{"materialID":"MAT-005","materialName":"发动机活塞","spec":"V1.0","quantity":120,"supplierID":"SUP-001"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`

	// 4. 调用更新方法
	updateArgs := [][]byte{[]byte("UpdateBOM"), []byte(updateBOMJSON)}
	updateResponse := mockStub.MockInvoke("tx1", updateArgs)

	// 5. 验证权限控制是否生效
	// 预期调用应该失败，因为 AutomakerMSP 没有权限更新BOM
	if updateResponse.Status == shim.OK {
		t.Fatal("预期UpdateBOM应该失败（非ManufacturerMSP成员），但执行成功")
	}
	t.Log("测试通过，预期失败：", updateResponse.Message)
}

// TestCompareBOM 测试比较BOM功能（正常流程）
// 验证能够成功比较研发BOM和生产BOM的差异
func TestCompareBOM(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建两个BOM（研发BOM和生产BOM）
	rdBOMJSON := `{"bomID":"R-BOM-202505-006","bomType":"研发BOM","productModel":"2025款XX车型","partBatchNo":"B20250506","version":"V1.0","creator":"厂商A","createTime":"1735689600","status":"已发布","materialList":[{"materialID":"MAT-006","materialName":"发动机活塞","spec":"V1.0","quantity":100,"supplierID":"SUP-001"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`
	prodBOMJSON := `{"bomID":"P-BOM-202505-006","bomType":"生产BOM","productModel":"2025款XX车型","partBatchNo":"B20250506","version":"V1.0","creator":"厂商A","createTime":"1735689600","status":"已发布","materialList":[{"materialID":"MAT-006","materialName":"发动机活塞","spec":"V1.0","quantity":100,"supplierID":"SUP-001"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`

	createArgs1 := [][]byte{[]byte("CreateBOM"), []byte(rdBOMJSON)}
	createArgs2 := [][]byte{[]byte("CreateBOM"), []byte(prodBOMJSON)}
	mockStub.MockInvoke("tx1", createArgs1)
	mockStub.MockInvoke("tx2", createArgs2)

	// 4. 调用比较方法
	compareArgs := [][]byte{[]byte("CompareBOM"), []byte("P-BOM-202505-006"), []byte("R-BOM-202505-006")}
	compareResponse := mockStub.MockInvoke("tx3", compareArgs)

	// 5. 校验比较结果
	if compareResponse.Status != shim.OK {
		t.Fatalf("CompareBOM执行失败：%s", compareResponse.Message)
	}

	// 6. 输出比较结果
	t.Log("测试通过，比较结果：", string(compareResponse.Payload))
}

// TestSubmitBOMChange 测试提交BOM变更功能（正常流程）
// 验证 ManufacturerMSP 成员能够成功提交BOM变更记录
func TestSubmitBOMChange(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个BOM
	bomJSON := `{"bomID":"R-BOM-202505-007","bomType":"研发BOM","productModel":"2025款XX车型","partBatchNo":"B20250507","version":"V1.0","creator":"厂商A","createTime":"1735689600","status":"已发布","materialList":[{"materialID":"MAT-007","materialName":"发动机活塞","spec":"V1.0","quantity":100,"supplierID":"SUP-001"}],"rdBOMFileInfo":{},"productionBOMInfo":{},"changeRecords":[]}`
	createArgs := [][]byte{[]byte("CreateBOM"), []byte(bomJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 构造变更记录数据
	// BOMChangeRecord 结构体字段：changeID, changeReason, oldVersion, newVersion, auditFlow, changeTime
	changeJSON := `{"changeID":"CHANGE-001","changeReason":"供应商变更","oldVersion":"V1.0","newVersion":"V1.1","auditFlow":[{"auditor":"工程师A","auditRole":"研发负责人","auditResult":"通过","auditTime":"1735689700","remark":"审核通过"}],"changeTime":"1735689700"}`

	// 5. 调用提交变更方法
	changeArgs := [][]byte{[]byte("SubmitBOMChange"), []byte("R-BOM-202505-007"), []byte(changeJSON)}
	changeResponse := mockStub.MockInvoke("tx2", changeArgs)

	// 6. 校验变更结果
	if changeResponse.Status != shim.OK {
		t.Fatalf("SubmitBOMChange执行失败：%s", changeResponse.Message)
	}

	// 7. 验证数据是否更新
	result, err := mockStub.GetState("BOM_R-BOM-202505-007")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("BOM变更数据未更新")
	}
	t.Log("测试通过，变更后数据：", string(result))
}

// TestQueryPartLifecycle 测试查询零部件生命周期功能（正常流程）
// 验证能够成功查询零部件的全生命周期轨迹
// 注意：生命周期数据是通过其他方法（如CreateProductionData）自动创建和更新的
func TestQueryPartLifecycle(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个零部件
	partJSON := `{"partID":"ENG-PISTON-LIFECYCLE","vin":"LVX1234568789799","batchNo":"B20250508","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	createArgs := [][]byte{[]byte("CreatePart"), []byte(partJSON)}
	createResponse := mockStub.MockInvoke("tx1", createArgs)
	if createResponse.Status != shim.OK {
		t.Fatalf("CreatePart执行失败：%s", createResponse.Message)
	}

	// 4. 创建生产数据（这会自动创建生命周期数据）
	productionJSON := `{"productionID":"PROD-LIFECYCLE","partID":"ENG-PISTON-LIFECYCLE","batchNo":"B20250508","params":{"温度":"180℃","时长":"2h"},"productionLine":"LINE-001","operator":"操作员A","finishTime":"1735689600"}`
	productionArgs := [][]byte{[]byte("CreateProductionData"), []byte(productionJSON)}
	productionResponse := mockStub.MockInvoke("tx2", productionArgs)
	if productionResponse.Status != shim.OK {
		t.Fatalf("CreateProductionData执行失败：%s", productionResponse.Message)
	}

	// 5. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryPartLifecycle"), []byte("ENG-PISTON-LIFECYCLE")}
	queryResponse := mockStub.MockInvoke("tx3", queryArgs)

	// 6. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryPartLifecycle执行失败：%s", queryResponse.Message)
	}

	// 7. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}

// TestQueryPartByBatchNo 测试按批次号查询零部件功能（正常流程）
// 验证能够成功查询指定批次的所有零部件
// 注意：此测试被跳过，因为 MockStub 不支持 GetQueryResult 方法（用于富查询）
// 需要在实际 Fabric 网络环境中测试
func TestQueryPartByBatchNo(t *testing.T) {
	t.Skip("MockStub 不支持 GetQueryResult 方法，需要在实际 Fabric 网络环境中测试")

	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建多个同批次的零部件
	part1JSON := `{"partID":"ENG-PISTON-BATCH-001","vin":"LVX1234568789800","batchNo":"BATCH-202505","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	part2JSON := `{"partID":"ENG-PISTON-BATCH-002","vin":"LVX1234568789801","batchNo":"BATCH-202505","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	mockStub.MockInvoke("tx1", [][]byte{[]byte("CreatePart"), []byte(part1JSON)})
	mockStub.MockInvoke("tx2", [][]byte{[]byte("CreatePart"), []byte(part2JSON)})

	// 4. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryPartByBatchNo"), []byte("BATCH-202505")}
	queryResponse := mockStub.MockInvoke("tx3", queryArgs)

	// 5. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryPartByBatchNo执行失败：%s", queryResponse.Message)
	}

	// 6. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}

// TestQueryPartByVIN 测试按VIN码查询零部件功能（正常流程）
// 验证能够成功查询指定VIN码的所有零部件
// 注意：此测试被跳过，因为 MockStub 不支持 GetQueryResult 方法（用于富查询）
// 需要在实际 Fabric 网络环境中测试
func TestQueryPartByVIN(t *testing.T) {
	t.Skip("MockStub 不支持 GetQueryResult 方法，需要在实际 Fabric 网络环境中测试")

	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建多个同VIN码的零部件
	part1JSON := `{"partID":"ENG-PISTON-VIN-001","vin":"LVX1234568789802","batchNo":"B20250509","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	part2JSON := `{"partID":"ENG-CYLINDER-VIN-001","vin":"LVX1234568789802","batchNo":"B20250509","name":"发动机气缸","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	mockStub.MockInvoke("tx1", [][]byte{[]byte("CreatePart"), []byte(part1JSON)})
	mockStub.MockInvoke("tx2", [][]byte{[]byte("CreatePart"), []byte(part2JSON)})

	// 4. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryPartByVIN"), []byte("LVX1234568789802")}
	queryResponse := mockStub.MockInvoke("tx3", queryArgs)

	// 5. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryPartByVIN执行失败：%s", queryResponse.Message)
	}

	// 6. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}

// TestCreateProductionData 测试创建生产数据功能（正常流程）
// 验证 ManufacturerMSP 成员（零部件生产厂商）能够成功创建生产数据
func TestCreateProductionData(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个零部件
	partJSON := `{"partID":"ENG-PISTON-PROD","vin":"LVX1234568789808","batchNo":"B20250510","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	createArgs := [][]byte{[]byte("CreatePart"), []byte(partJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 构造测试数据
	// ProductionData 结构体字段：productionID, partID, batchNo, params, productionLine, operator, finishTime
	productionJSON := `{"productionID":"PROD-001","partID":"ENG-PISTON-PROD","batchNo":"B20250510","params":{"温度":"180℃","时长":"2h","压力":"100MPa"},"productionLine":"LINE-001","operator":"操作员A","finishTime":"1735689600"}`

	// 5. 调用链码方法
	args := [][]byte{[]byte("CreateProductionData"), []byte(productionJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("CreateProductionData执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	result, err := mockStub.GetState("PROD_PROD-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("生产数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestCreateProductionData_Unauthorized 测试创建生产数据功能（未授权流程）
// 验证非 ManufacturerMSP 成员无法创建生产数据
func TestCreateProductionData_Unauthorized(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AutomakerMSP 身份（整车车企，无权限）
	mockStub, err := createMockStubWithIdentity("AutomakerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	productionJSON := `{"productionID":"PROD-002","partID":"ENG-PISTON-PROD","batchNo":"B20250511","params":{"温度":"180℃","时长":"2h"},"productionLine":"LINE-001","operator":"操作员A","finishTime":"1735689600"}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("CreateProductionData"), []byte(productionJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果（预期失败）
	if response.Status == shim.OK {
		t.Fatal("预期CreateProductionData应该失败，但执行成功了")
	}
	t.Logf("测试通过，预期失败：%s", response.Message)
}

// TestCreateQualityInspection 测试创建质检记录功能（正常流程）
// 验证 ManufacturerMSP 成员（零部件生产厂商）能够成功创建质检记录
func TestCreateQualityInspection(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个零部件
	partJSON := `{"partID":"ENG-PISTON-INSPECT","vin":"LVX1234568789809","batchNo":"B20250512","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	createArgs := [][]byte{[]byte("CreatePart"), []byte(partJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 构造测试数据
	// QualityInspection 结构体字段：inspectionID, partID, batchNo, indicators, result, handler, handleTime, disposal
	inspectionJSON := `{"inspectionID":"INSPECT-001","partID":"ENG-PISTON-INSPECT","batchNo":"B20250512","indicators":{"尺寸":"10±0.1mm","硬度":"HRC50","重量":"500g"},"result":"合格","handler":"质检员A","handleTime":"1735689600","disposal":""}`

	// 5. 调用链码方法
	args := [][]byte{[]byte("CreateQualityInspection"), []byte(inspectionJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("CreateQualityInspection执行失败：%s", response.Message)
	}

	// 6. 验证数据是否写入账本
	result, err := mockStub.GetState("QUALITY_INSPECT-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("质检记录未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestCreateQualityInspection_Unauthorized 测试创建质检记录功能（未授权流程）
// 验证非 ManufacturerMSP 成员无法创建质检记录
func TestCreateQualityInspection_Unauthorized(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AutomakerMSP 身份（整车车企，无权限）
	mockStub, err := createMockStubWithIdentity("AutomakerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	inspectionJSON := `{"inspectionID":"INSPECT-002","partID":"ENG-PISTON-INSPECT","batchNo":"B20250513","indicators":{"尺寸":"10±0.1mm","硬度":"HRC50"},"result":"合格","handler":"质检员A","handleTime":"1735689600","disposal":""}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("CreateQualityInspection"), []byte(inspectionJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果（预期失败）
	if response.Status == shim.OK {
		t.Fatal("预期CreateQualityInspection应该失败，但执行成功了")
	}
	t.Logf("测试通过，预期失败：%s", response.Message)
}

// TestUpdatePartStatus 测试更新零部件状态功能（正常流程）
// 验证 ManufacturerMSP 或 AutomakerMSP 成员能够成功更新零部件状态
func TestUpdatePartStatus(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个零部件
	partJSON := `{"partID":"ENG-PISTON-STATUS","vin":"LVX1234568789803","batchNo":"B20250514","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	createArgs := [][]byte{[]byte("CreatePart"), []byte(partJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 调用更新方法，将状态更新为"在产"
	updateArgs := [][]byte{[]byte("UpdatePartStatus"), []byte("ENG-PISTON-STATUS"), []byte("在产")}
	updateResponse := mockStub.MockInvoke("tx2", updateArgs)

	// 5. 校验更新结果
	if updateResponse.Status != shim.OK {
		t.Fatalf("UpdatePartStatus执行失败：%s", updateResponse.Message)
	}

	// 6. 验证状态是否更新
	result, err := mockStub.GetState("ENG-PISTON-STATUS")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("零部件数据未找到")
	}
	t.Log("测试通过，更新后数据：", string(result))
}

// TestUpdatePartStatus_Unauthorized 测试更新零部件状态功能（未授权流程）
// 验证非 ManufacturerMSP 或 AutomakerMSP 成员无法更新零部件状态
func TestUpdatePartStatus_Unauthorized(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 AftersaleMSP 身份（4S店/售后中心，无权限）
	mockStub, err := createMockStubWithIdentity("AftersaleMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个零部件
	partJSON := `{"partID":"ENG-PISTON-STATUS-2","vin":"LVX1234568789804","batchNo":"B20250515","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	createArgs := [][]byte{[]byte("CreatePart"), []byte(partJSON)}
	mockStub.MockInvoke("tx1", createArgs)

	// 4. 调用更新方法
	updateArgs := [][]byte{[]byte("UpdatePartStatus"), []byte("ENG-PISTON-STATUS-2"), []byte("在产")}
	updateResponse := mockStub.MockInvoke("tx2", updateArgs)

	// 5. 校验更新结果（预期失败）
	if updateResponse.Status == shim.OK {
		t.Fatal("预期UpdatePartStatus应该失败，但执行成功了")
	}
	t.Logf("测试通过，预期失败：%s", updateResponse.Message)
}

// TestCreateSupplyOrder 测试创建采购订单功能（正常流程）
// 验证 AutomakerMSP 成员（整车车企）能够成功创建采购订单
func TestCreateSupplyOrder(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份（先创建零部件）
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个零部件
	partJSON := `{"partID":"ENG-PISTON-ORDER","vin":"LVX1234568789810","batchNo":"B20250516","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	createArgs := [][]byte{[]byte("CreatePart"), []byte(partJSON)}
	createResponse := mockStub.MockInvoke("tx1", createArgs)
	if createResponse.Status != shim.OK {
		t.Fatalf("CreatePart执行失败：%s", createResponse.Message)
	}

	// 4. 切换身份为 AutomakerMSP（整车车企）来创建订单
	automakerIdentity := &msp.SerializedIdentity{
		Mspid:   "AutomakerMSP",
		IdBytes: []byte("test-user-automaker"),
	}
	automakerCreatorBytes, _ := proto.Marshal(automakerIdentity)
	mockStub.Creator = automakerCreatorBytes

	// 5. 构造测试数据
	// SupplyOrder 结构体字段：orderID, buyer, seller, partID, quantity, batchNo, agreedTime, status, createTime
	orderJSON := `{"orderID":"ORDER-001","buyer":"车企A","seller":"厂商A","partID":"ENG-PISTON-ORDER","quantity":100,"batchNo":"B20250516","agreedTime":"1735689600","status":"待发货","createTime":"1735689600"}`

	// 6. 调用链码方法
	args := [][]byte{[]byte("CreateSupplyOrder"), []byte(orderJSON)}
	response := mockStub.MockInvoke("tx2", args)

	// 7. 校验调用结果
	if response.Status != shim.OK {
		t.Fatalf("CreateSupplyOrder执行失败：%s", response.Message)
	}

	// 8. 验证数据是否写入账本
	result, err := mockStub.GetState("ORDER_ORDER-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("采购订单未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}

// TestCreateSupplyOrder_Unauthorized 测试创建采购订单功能（未授权流程）
// 验证非 AutomakerMSP 成员无法创建采购订单
func TestCreateSupplyOrder_Unauthorized(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份（零部件生产厂商，无权限）
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 构造测试数据
	orderJSON := `{"orderID":"ORDER-002","buyer":"车企A","seller":"厂商A","partID":"ENG-PISTON-ORDER","quantity":100,"batchNo":"B20250517","agreedTime":"1735689600","status":"待发货","createTime":"1735689600"}`

	// 4. 调用链码方法
	args := [][]byte{[]byte("CreateSupplyOrder"), []byte(orderJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验调用结果（预期失败）
	if response.Status == shim.OK {
		t.Fatal("预期CreateSupplyOrder应该失败，但执行成功了")
	}
	t.Logf("测试通过，预期失败：%s", response.Message)
}

// TestQueryAffectedParts 测试查询受影响零部件功能（正常流程）
// 验证能够成功查询指定批次的所有受影响零部件
// 注意：此测试被跳过，因为 MockStub 不支持 GetQueryResult 方法（用于富查询）
// 需要在实际 Fabric 网络环境中测试
func TestQueryAffectedParts(t *testing.T) {
	t.Skip("MockStub 不支持 GetQueryResult 方法，需要在实际 Fabric 网络环境中测试")

	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建多个同批次的零部件
	part1JSON := `{"partID":"ENG-PISTON-AFFECTED-001","vin":"LVX1234568789805","batchNo":"AFFECTED-BATCH","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	part2JSON := `{"partID":"ENG-PISTON-AFFECTED-002","vin":"LVX1234568789806","batchNo":"AFFECTED-BATCH","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	mockStub.MockInvoke("tx1", [][]byte{[]byte("CreatePart"), []byte(part1JSON)})
	mockStub.MockInvoke("tx2", [][]byte{[]byte("CreatePart"), []byte(part2JSON)})

	// 4. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryAffectedParts"), []byte("AFFECTED-BATCH")}
	queryResponse := mockStub.MockInvoke("tx3", queryArgs)

	// 5. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryAffectedParts执行失败：%s", queryResponse.Message)
	}

	// 6. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}

// TestQueryAftersaleRecord 测试查询售后记录功能（正常流程）
// 验证能够成功查询售后记录详情
func TestQueryAftersaleRecord(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 创建 MockStub 并设置 ManufacturerMSP 身份
	mockStub, err := createMockStubWithIdentity("ManufacturerMSP", cc)
	if err != nil {
		t.Fatalf("创建MockStub失败：%v", err)
	}

	// 3. 先创建一个零部件
	partJSON := `{"partID":"ENG-PISTON-QUERY","vin":"LVX1234568789807","batchNo":"B20250518","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商A","createTime":"1735689600","status":"NORMAL"}`
	createArgs := [][]byte{[]byte("CreatePart"), []byte(partJSON)}
	createResponse := mockStub.MockInvoke("tx1", createArgs)
	if createResponse.Status != shim.OK {
		t.Fatalf("CreatePart执行失败：%s", createResponse.Message)
	}

	// 4. 切换身份为 AftersaleMSP（4S店/售后中心）来创建售后记录
	aftersaleIdentity := &msp.SerializedIdentity{
		Mspid:   "AftersaleMSP",
		IdBytes: []byte("test-user-aftersale"),
	}
	aftersaleCreatorBytes, _ := proto.Marshal(aftersaleIdentity)
	mockStub.Creator = aftersaleCreatorBytes

	// 5. 创建一个售后记录
	aftersaleJSON := `{"aftersaleID":"AFTERSALE-QUERY-001","partID":"ENG-PISTON-QUERY","vin":"LVX1234568789807","batchNo":"B20250518","aftersaleType":"故障报修","aftersaleStatus":"处理中","handlerOrgID":"4S店A","handlerID":"技师A","description":"发动机活塞磨损严重","processResult":"更换新活塞","processTime":"1735689700","cost":"1000","attachmentInfo":{},"remark":"已更换原厂配件"}`
	aftersaleArgs := [][]byte{[]byte("CreateAftersaleRecord"), []byte(aftersaleJSON)}
	aftersaleResponse := mockStub.MockInvoke("tx2", aftersaleArgs)
	if aftersaleResponse.Status != shim.OK {
		t.Fatalf("CreateAftersaleRecord执行失败：%s", aftersaleResponse.Message)
	}

	// 6. 调用查询方法
	queryArgs := [][]byte{[]byte("QueryAftersaleRecord"), []byte("AFTERSALE-QUERY-001")}
	queryResponse := mockStub.MockInvoke("tx3", queryArgs)

	// 7. 校验查询结果
	if queryResponse.Status != shim.OK {
		t.Fatalf("QueryAftersaleRecord执行失败：%s", queryResponse.Message)
	}

	// 8. 输出查询结果
	t.Log("测试通过，查询结果：", string(queryResponse.Payload))
}
