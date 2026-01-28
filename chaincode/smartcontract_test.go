package main

import (
	"testing"

	"github.com/golang/protobuf/proto"                // 导入protobuf序列化包
	"github.com/hyperledger/fabric-chaincode-go/shim" // 导入shim包（用于shim.OK）
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/msp" // 身份相关包
)

func TestCreatePart(t *testing.T) {
	// 1. 创建Contract链码实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 2. 构造身份对象并序列化为字节数组（核心修复）
	identity := &msp.SerializedIdentity{
		Mspid:   "Org1MSP",                // 匹配你的Fabric网络MSP名称
		IdBytes: []byte("test-user-org1"), // 模拟用户身份
	}
	// 将身份对象序列化为[]byte
	creatorBytes, err := proto.Marshal(identity)
	if err != nil {
		t.Fatalf("序列化身份对象失败：%v", err)
	}

	// 3. 创建MockStub并设置Creator（类型匹配）
	mockStub := shimtest.NewMockStub("realty-chaincode", cc)
	mockStub.Creator = creatorBytes // 赋值[]byte类型的身份数据

	// 4. 构造参数并调用链码
	testPartJSON := `{"partID":"ENG-PISTON-002","vin":"LVX432109876543210","batchNo":"B20250502","name":"发动机活塞","type":"发动机部件","manufacturer":"厂商B","createTime":"1735689600","status":"NORMAL"}`
	args := [][]byte{[]byte("CreatePart"), []byte(testPartJSON)}
	response := mockStub.MockInvoke("tx1", args)

	// 5. 校验结果
	if response.Status != shim.OK {
		t.Fatalf("CreatePart执行失败：状态码=%d，信息=%s", response.Status, response.Message)
	}

	// 6. 验证数据写入
	result, err := mockStub.GetState("ENG-PISTON-002")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("零部件数据未写入账本")
	}
	t.Log("测试通过，数据：", string(result))
}
