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
// 验证 Org1MSP 成员能够成功创建零部件信息
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
		Mspid:   "Org1MSP",
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
// 验证非 Org1MSP 成员（如 Org2MSP）无法创建零部件，应该被拒绝
func TestCreatePart_Unauthorized(t *testing.T) {
	// 0. 设置测试模式标志
	TestMode = true
	defer func() { TestMode = false }()

	// 1. 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 构造 Org2MSP 身份对象
	// Org2MSP 代表整车车企组织，不应该有权限创建零部件
	identity := &msp.SerializedIdentity{
		Mspid:   "Org2MSP",
		IdBytes: []byte("test-user-org2"),
	}
	// 序列化身份对象
	creatorBytes, err := proto.Marshal(identity)
	if err != nil {
		t.Fatalf("序列化身份对象失败：%v", err)
	}

	// 3. 创建 MockStub 并设置 Creator
	mockStub := shimtest.NewMockStub("realty-chaincode", cc)
	// 设置为 Org2MSP 成员
	mockStub.Creator = creatorBytes

	// 4. 构造测试数据
	// 创建另一个零部件，使用不同的 ID
	testPartJSON := `{"partID":"ENG-PISTON-003","vin":"LVX432109876543211","batchNo":"B20250503","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商C","createTime":"1735689600","status":"NORMAL"}`

	// 5. 调用链码方法
	args := [][]byte{[]byte("CreatePart"), []byte(testPartJSON)}
	// 使用不同的交易ID "tx2"
	response := mockStub.MockInvoke("tx2", args)

	// 6. 验证权限控制是否生效
	// 预期调用应该失败，因为 Org2MSP 没有权限创建零部件
	if response.Status == shim.OK {
		t.Fatal("预期CreatePart应该失败（非Org1MSP成员），但执行成功")
	}
	// 输出测试通过信息，打印返回的错误消息
	t.Log("测试通过，预期失败：", response.Message)
}
