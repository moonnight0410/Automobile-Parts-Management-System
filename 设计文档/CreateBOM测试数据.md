# CreateBOM 测试数据

本文档提供了用于测试 CreateBOM 链码方法的 JSON 测试数据。

## 数据结构说明

### BOM 主结构体字段

| 字段名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|--------|------|---------|
| bomID | string | 是 | BOM唯一ID | "R-BOM-202505-001" |
| bomType | string | 是 | BOM类型（研发BOM/生产BOM） | "研发BOM" |
| productModel | string | 是 | 对应车型 | "2025款XX车型" |
| partBatchNo | string | 是 | 零部件批次号 | "B20250501" |
| version | string | 是 | BOM版本号 | "V1.0" |
| creator | string | 是 | 创建人（企业数字身份标识） | "admin-org1" |
| createTime | string | 是 | 创建时间（时间戳） | "1735689600" |
| status | string | 是 | 状态（草稿/已发布/已变更/已作废） | "已发布" |
| materialList | array | 是 | 物料明细列表 | 见下方 MaterialItem |
| rdBOMFileInfo | object | 否 | 研发BOM图纸文件信息 | 见下方 BOMFileInfo |
| productionBOMInfo | object | 否 | 生产BOM比对信息 | 见下方 ProductionBOMInfo |
| changeRecords | array | 否 | 变更记录 | 见下方 BOMChangeRecord |

### MaterialItem 物料明细项字段

| 字段名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|--------|------|---------|
| materialID | string | 是 | 物料唯一ID | "MAT-001" |
| materialName | string | 是 | 物料名称 | "发动机活塞" |
| spec | string | 是 | 规格型号 | "直径100mm，高度50mm" |
| quantity | int | 是 | 数量 | 100 |
| supplierID | string | 是 | 供应商ID | "SUP-001" |

### BOMFileInfo 研发BOM图纸文件信息字段

| 字段名 | 类型 | 必填 | 说明 | 示例值 |
|--------|------|--------|------|---------|
| fileStorageID | string | 是 | 普通数据库中的文件唯一ID | "FILE-001" |
| filePath | string | 是 | 文件存储路径 | "/data/bom/rd/202505/R-BOM-001-CAD.dwg" |
| fileHash | string | 是 | 文件内容哈希（SHA256） | "a1b2c3d4e5f6..." |
| fileType | string | 是 | 文件类型 | "CAD图纸" |
| fileSize | int64 | 是 | 文件大小（字节） | 1024000 |
| uploadTime | string | 是 | 文件上传时间 | "1735689600" |

---

## 测试数据示例

### 示例 1：研发BOM（包含图纸文件）

```json
{
  "bomID": "R-BOM-202505-001",
  "bomType": "研发BOM",
  "productModel": "2025款XX车型",
  "partBatchNo": "B20250501",
  "version": "V1.0",
  "creator": "admin-org1",
  "createTime": "1735689600",
  "status": "已发布",
  "materialList": [
    {
      "materialID": "MAT-001",
      "materialName": "发动机活塞",
      "spec": "直径100mm，高度50mm",
      "quantity": 100,
      "supplierID": "SUP-001"
    },
    {
      "materialID": "MAT-002",
      "materialName": "活塞环",
      "spec": "内径100mm，厚度2mm",
      "quantity": 200,
      "supplierID": "SUP-002"
    },
    {
      "materialID": "MAT-003",
      "materialName": "连杆",
      "spec": "长度200mm，直径30mm",
      "quantity": 100,
      "supplierID": "SUP-003"
    }
  ],
  "rdBOMFileInfo": {
    "fileStorageID": "FILE-001",
    "filePath": "/data/bom/rd/202505/R-BOM-001-CAD.dwg",
    "fileHash": "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0",
    "fileType": "CAD图纸",
    "fileSize": 1024000,
    "uploadTime": "1735689600"
  }
}
```

### 示例 2：生产BOM（包含研发BOM比对）

```json
{
  "bomID": "P-BOM-202505-001",
  "bomType": "生产BOM",
  "productModel": "2025款XX车型",
  "partBatchNo": "B20250501",
  "version": "V1.0",
  "creator": "admin-org1",
  "createTime": "1735689600",
  "status": "已发布",
  "materialList": [
    {
      "materialID": "MAT-001",
      "materialName": "发动机活塞",
      "spec": "直径100mm，高度50mm",
      "quantity": 500,
      "supplierID": "SUP-001"
    },
    {
      "materialID": "MAT-004",
      "materialName": "润滑油",
      "spec": "全合成机油，5W-30",
      "quantity": 1000,
      "supplierID": "SUP-004"
    }
  ],
  "productionBOMInfo": {
    "rdBOMID": "R-BOM-202505-001",
    "rdBOMVersion": "V1.0",
    "changeReason": "根据生产实际情况调整物料数量",
    "changeTime": "1735689600"
  }
}
```

### 示例 3：简单BOM（最小必填字段）

```json
{
  "bomID": "R-BOM-202505-002",
  "bomType": "研发BOM",
  "productModel": "2025款YY车型",
  "partBatchNo": "B20250502",
  "version": "V1.0",
  "creator": "admin-org1",
  "createTime": "1735689600",
  "status": "草稿",
  "materialList": [
    {
      "materialID": "MAT-005",
      "materialName": "刹车片",
      "spec": "前轮，厚度10mm",
      "quantity": 50,
      "supplierID": "SUP-005"
    }
  ]
}
```

### 示例 4：包含变更记录的BOM

```json
{
  "bomID": "R-BOM-202505-003",
  "bomType": "研发BOM",
  "productModel": "2025款ZZ车型",
  "partBatchNo": "B20250503",
  "version": "V2.0",
  "creator": "admin-org1",
  "createTime": "1735689600",
  "status": "已变更",
  "materialList": [
    {
      "materialID": "MAT-006",
      "materialName": "变速箱齿轮",
      "spec": "模数2.5，齿数30",
      "quantity": 80,
      "supplierID": "SUP-006"
    }
  ],
  "changeRecords": [
    {
      "changeID": "CHANGE-001",
      "changeType": "版本迭代",
      "changeReason": "根据供应商反馈优化设计",
      "changeTime": "1735689600",
      "changer": "engineer-org1"
    },
    {
      "changeID": "CHANGE-002",
      "changeType": "物料替换",
      "changeReason": "原供应商无法按时交货，更换供应商",
      "changeTime": "1735689700",
      "changer": "purchaser-org2"
    }
  ]
}
```

---

## 使用方法

### 方法 1：使用 peer CLI 命令

```bash
# 设置 Org1 环境（零部件生产厂商）
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=${PWD}/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# 调用 CreateBOM 方法
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n automobile-parts --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["CreateBOM","{\"bomID\":\"R-BOM-202505-001\",\"bomType\":\"研发BOM\",\"productModel\":\"2025款XX车型\",\"partBatchNo\":\"B20250501\",\"version\":\"V1.0\",\"creator\":\"admin-org1\",\"createTime\":\"1735689600\",\"status\":\"已发布\",\"materialList\":[{\"materialID\":\"MAT-001\",\"materialName\":\"发动机活塞\",\"spec\":\"直径100mm，高度50mm\",\"quantity\":100,\"supplierID\":\"SUP-001\"},{\"materialID\":\"MAT-002\",\"materialName\":\"活塞环\",\"spec\":\"内径100mm，厚度2mm\",\"quantity\":200,\"supplierID\":\"SUP-002\"}],\"rdBOMFileInfo\":{\"fileStorageID\":\"FILE-001\",\"filePath\":\"/data/bom/rd/202505/R-BOM-001-CAD.dwg\",\"fileHash\":\"a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0\",\"fileType\":\"CAD图纸\",\"fileSize\":1024000,\"uploadTime\":\"1735689600\"}}"]}'
```

### 方法 2：使用测试文件

#### 创建测试文件

```bash
# 创建测试数据文件
cat > test-bom.json << 'EOF'
{
  "bomID": "R-BOM-202505-001",
  "bomType": "研发BOM",
  "productModel": "2025款XX车型",
  "partBatchNo": "B20250501",
  "version": "V1.0",
  "creator": "admin-org1",
  "createTime": "1735689600",
  "status": "已发布",
  "materialList": [
    {
      "materialID": "MAT-001",
      "materialName": "发动机活塞",
      "spec": "直径100mm，高度50mm",
      "quantity": 100,
      "supplierID": "SUP-001"
    }
  ],
  "rdBOMFileInfo": {
    "fileStorageID": "FILE-001",
    "filePath": "/data/bom/rd/202505/R-BOM-001-CAD.dwg",
    "fileHash": "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0",
    "fileType": "CAD图纸",
    "fileSize": 1024000,
    "uploadTime": "1735689600"
  }
}
EOF
```

#### 读取文件并调用链码

```bash
# 读取文件内容并调用链码
BOM_DATA=$(cat test-bom.json | tr -d '\n' | tr -d ' ')

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n automobile-parts --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c "{\"Args\":[\"CreateBOM\",\"${BOM_DATA}\"]}"
```

### 方法 3：使用 Go 测试代码

```go
package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func TestCreateBOM(t *testing.T) {
	// 设置测试模式
	TestMode = true
	defer func() { TestMode = false }()

	// 创建智能合约实例
	cc, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		t.Fatalf("创建链码实例失败：%v", err)
	}

	// 创建 MockStub
	mockStub := shimtest.NewMockStub("automobile-parts", cc)

	// 构造测试数据
	testBOM := map[string]interface{}{
		"bomID":        "R-BOM-202505-001",
		"bomType":      "研发BOM",
		"productModel":  "2025款XX车型",
		"partBatchNo":  "B20250501",
		"version":       "V1.0",
		"creator":       "admin-org1",
		"createTime":    "1735689600",
		"status":        "已发布",
		"materialList": []map[string]interface{}{
			{
				"materialID":   "MAT-001",
				"materialName": "发动机活塞",
				"spec":         "直径100mm，高度50mm",
				"quantity":      100,
				"supplierID":   "SUP-001",
			},
			{
				"materialID":   "MAT-002",
				"materialName": "活塞环",
				"spec":         "内径100mm，厚度2mm",
				"quantity":      200,
				"supplierID":   "SUP-002",
			},
		},
		"rdBOMFileInfo": map[string]interface{}{
			"fileStorageID": "FILE-001",
			"filePath":      "/data/bom/rd/202505/R-BOM-001-CAD.dwg",
			"fileHash":      "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0",
			"fileType":      "CAD图纸",
			"fileSize":      1024000,
			"uploadTime":    "1735689600",
		},
	}

	// 序列化为 JSON
	bomJSON, err := json.Marshal(testBOM)
	if err != nil {
		t.Fatalf("序列化BOM数据失败：%v", err)
	}

	// 调用 CreateBOM 方法
	args := [][]byte{[]byte("CreateBOM"), bomJSON}
	response := mockStub.MockInvoke("tx1", args)

	// 验证结果
	if response.Status != shim.OK {
		t.Fatalf("CreateBOM执行失败：状态码=%d，信息=%s", response.Status, response.Message)
	}

	// 验证数据写入
	result, err := mockStub.GetState("R-BOM-202505-001")
	if err != nil || result == nil || len(result) == 0 {
		t.Fatal("BOM数据未写入账本")
	}

	t.Log("测试通过，数据：", string(result))
}
```

---

## 验证测试结果

### 查询创建的BOM

```bash
# 查询BOM
peer chaincode query -C mychannel -n automobile-parts -c '{"Args":["QueryBOM","R-BOM-202505-001"]}'

# 查询所有BOM（如果链码支持）
peer chaincode query -C mychannel -n automobile-parts -c '{"Args":["QueryAllBOMs"]}'
```

### 查询BOM的物料清单

```bash
# 查询特定BOM的物料清单
peer chaincode query -C mychannel -n automobile-parts -c '{"Args":["QueryBOMMaterials","R-BOM-202505-001"]}'
```

---

## 常见错误处理

### 错误 1：权限不足

**错误信息**：
```
Error: proposal failed with status: 500 - 只有零部件生产厂商组织的成员才能创建BOM
```

**解决方案**：
确保使用 Org1MSP（零部件生产厂商）的身份调用链码。

### 错误 2：必填字段缺失

**错误信息**：
```
Error: proposal failed with status: 500 - BOMID、BOMType、ProductModel、PartBatchNo、Version、Creator、CreateTime、Status、MaterialList不能为空
```

**解决方案**：
检查 JSON 数据是否包含所有必填字段。

### 错误 3：JSON 格式错误

**错误信息**：
```
Error: proposal failed with status: 500 - 反序列化BOM数据失败: invalid character '}' after object key
```

**解决方案**：
使用 JSON 验证工具检查 JSON 格式是否正确。

---

## 注意事项

1. **BOMID 唯一性**：确保每个 BOM 的 ID 是唯一的
2. **物料列表**：至少包含一个物料项
3. **版本号管理**：建议使用语义化版本号（V1.0, V1.1, V2.0）
4. **状态值**：使用预定义的状态值（草稿、已发布、已变更、已作废）
5. **时间戳**：使用 Unix 时间戳（秒）
6. **文件哈希**：使用 SHA256 算法计算文件哈希值

---

## 快速测试命令

```bash
# 一键测试命令（将 JSON 数据替换为实际数据）
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n automobile-parts --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["CreateBOM","<在此处粘贴JSON数据>"]}'
```
