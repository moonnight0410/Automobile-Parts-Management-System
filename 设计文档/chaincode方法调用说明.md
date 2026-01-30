# Chaincode 方法调用说明

本文档详细说明了汽车零部件全生命周期可信管理系统的所有智能合约方法，包括输入参数、输出结果和方法功能说明。

## 目录

1. [系统初始化](#系统初始化)
2. [溯源查询模块](#溯源查询模块)
3. [BOM 管理模块](#bom-管理模块)
4. [生产质量模块](#生产质量模块)
5. [供应链模块](#供应链模块)
6. [售后模块](#售后模块)
7. [身份管理模块](#身份管理模块)
8. [AI 功能模块](#ai-功能模块)

---

## 系统初始化

### InitLedger

**方法功能：** 初始化账本，设置系统初始状态

**输入参数：**
- 无

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 无

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"InitLedger","Args":[]}'
```

---

## 溯源查询模块

### CreatePart

**方法功能：** 创建新零部件信息并上链

**输入参数：**
- `partJSON` (string) - 零部件信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 零部件生产厂商组织成员

**Part 结构体字段：**
```json
{
  "partID": "ENG-PISTON-202505",
  "vin": "LSVAA4182FS123456",
  "batchNo": "BATCH-202505-001",
  "name": "发动机活塞",
  "type": "发动机部件",
  "manufacturer": "MANUFACTURER-001",
  "createTime": "1717084800",
  "status": "在产"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreatePart","Args":["{\"partID\":\"ENG-PISTON-202505\",\"vin\":\"LSVAA4182FS123456\",\"batchNo\":\"BATCH-202505-001\",\"name\":\"发动机活塞\",\"type\":\"发动机部件\",\"manufacturer\":\"MANUFACTURER-001\",\"createTime\":\"1717084800\",\"status\":\"在产\"}"]}'
```

---

### QueryPart

**方法功能：** 根据 PartID 查询零部件详细信息

**输入参数：**
- `partID` (string) - 零部件唯一标识

**输出结果：**
- `*Part` - 零部件详细信息
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryPart","Args":["ENG-PISTON-202505"]}'
```

---

### QueryPartLifecycle

**方法功能：** 查询零部件全生命周期信息，包括关联的 BOM、生产、质检、供应链和售后数据

**输入参数：**
- `partID` (string) - 零部件唯一标识

**输出结果：**
- `*PartLifecycle` - 零部件全生命周期信息
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**PartLifecycle 结构体字段：**
```json
{
  "partID": "ENG-PISTON-202505",
  "bomInfo": {
    "bomID": "P-BOM-202505-001",
    "version": "V1.0",
    "type": "生产BOM"
  },
  "productionInfo": {
    "productionID": "PROD-202505-001",
    "partID": "ENG-PISTON-202505",
    "batchNo": "BATCH-202505-001",
    "productionLine": "LINE-001",
    "worker": "WORKER-001",
    "startTime": "1717084800",
    "finishTime": "1717088400"
  },
  "qualityInfo": {
    "inspectionID": "QC-202505-001",
    "partID": "ENG-PISTON-202505",
    "batchNo": "BATCH-202505-001",
    "inspector": "INSPECTOR-001",
    "result": "合格",
    "handleTime": "1717089000"
  },
  "supplyChainInfo": [],
  "aftersaleInfo": []
}
```

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryPartLifecycle","Args":["ENG-PISTON-202505"]}'
```

---

### QueryPartByBatchNo

**方法功能：** 根据批次号查询该批次的所有零部件

**输入参数：**
- `batchNo` (string) - 生产批次号

**输出结果：**
- `[]*Part` - 零部件列表
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryPartByBatchNo","Args":["BATCH-202505-001"]}'
```

---

### QueryPartByVIN

**方法功能：** 根据 VIN 码查询关联的零部件

**输入参数：**
- `vin` (string) - 车辆 VIN 码

**输出结果：**
- `[]*Part` - 零部件列表
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryPartByVIN","Args":["LSVAA4182FS123456"]}'
```

---

## BOM 管理模块

### CreateBOM

**方法功能：** 创建新的 BOM（物料清单）信息

**输入参数：**
- `bomJSON` (string) - BOM 信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 零部件生产厂商组织成员

**BOM 结构体字段：**
```json
{
  "bomID": "P-BOM-202505-001",
  "bomType": "生产BOM",
  "productModel": "2025款XX车型",
  "partBatchNo": "BATCH-202505-001",
  "version": "V1.0",
  "creator": "MANUFACTURER-001",
  "createTime": "1717084800",
  "status": "草稿",
  "materialList": [
    {
      "materialID": "MAT-001",
      "materialName": "铝合金",
      "spec": "6061-T6",
      "quantity": 100,
      "supplierID": "SUPPLIER-001"
    }
  ]
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateBOM","Args":["{\"bomID\":\"P-BOM-202505-001\",\"bomType\":\"生产BOM\",\"productModel\":\"2025款XX车型\",\"partBatchNo\":\"BATCH-202505-001\",\"version\":\"V1.0\",\"creator\":\"MANUFACTURER-001\",\"createTime\":\"1717084800\",\"status\":\"草稿\",\"materialList\":[{\"materialID\":\"MAT-001\",\"materialName\":\"铝合金\",\"spec\":\"6061-T6\",\"quantity\":100,\"supplierID\":\"SUPPLIER-001\"}]}"]}'
```

---

### QueryBOM

**方法功能：** 根据 BOMID 查询 BOM 详细信息

**输入参数：**
- `bomID` (string) - BOM 唯一标识

**输出结果：**
- `*BOM` - BOM 详细信息
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryBOM","Args":["P-BOM-202505-001"]}'
```

---

### UpdateBOM

**方法功能：** 更新 BOM 信息，包括状态和物料清单

**输入参数：**
- `bomJSON` (string) - 更新后的 BOM 信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 零部件生产厂商组织成员

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"UpdateBOM","Args":["{\"bomID\":\"P-BOM-202505-001\",\"status\":\"已发布\"}"]}'
```

---

### CompareBOM

**方法功能：** 比对生产 BOM 和研发 BOM 的差异

**输入参数：**
- `productionBOMID` (string) - 生产 BOM ID
- `rdBOMID` (string) - 研发 BOM ID

**输出结果：**
- `string` - 比对结果（JSON 格式）
- `error` - 比对失败时返回错误信息

**权限要求：** 零部件生产厂商组织成员

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"CompareBOM","Args":["P-BOM-202505-001","R-BOM-202505-001"]}'
```

---

### SubmitBOMChange

**方法功能：** 提交 BOM 变更申请

**输入参数：**
- `bomID` (string) - BOM ID
- `changeJSON` (string) - 变更信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 零部件生产厂商组织成员

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"SubmitBOMChange","Args":["P-BOM-202505-001","{\"changeReason\":\"材料替换\",\"changeType\":\"物料变更\",\"proposer\":\"MANUFACTURER-001\",\"changeTime\":\"1717084800\"}"]}'
```

---

## 生产质量模块

### CreateProductionData

**方法功能：** 创建生产数据记录

**输入参数：**
- `productionJSON` (string) - 生产数据的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 零部件生产厂商组织成员

**ProductionData 结构体字段：**
```json
{
  "productionID": "PROD-202505-001",
  "partID": "ENG-PISTON-202505",
  "batchNo": "BATCH-202505-001",
  "productionLine": "LINE-001",
  "worker": "WORKER-001",
  "startTime": "1717084800",
  "finishTime": "1717088400"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateProductionData","Args":["{\"productionID\":\"PROD-202505-001\",\"partID\":\"ENG-PISTON-202505\",\"batchNo\":\"BATCH-202505-001\",\"productionLine\":\"LINE-001\",\"worker\":\"WORKER-001\",\"startTime\":\"1717084800\",\"finishTime\":\"1717088400\"}"]}'
```

---

### CreateQualityInspection

**方法功能：** 创建质检记录

**输入参数：**
- `inspectionJSON` (string) - 质检数据的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 零部件生产厂商组织成员

**QualityInspection 结构体字段：**
```json
{
  "inspectionID": "QC-202505-001",
  "partID": "ENG-PISTON-202505",
  "batchNo": "BATCH-202505-001",
  "inspector": "INSPECTOR-001",
  "result": "合格",
  "handleTime": "1717089000"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateQualityInspection","Args":["{\"inspectionID\":\"QC-202505-001\",\"partID\":\"ENG-PISTON-202505\",\"batchNo\":\"BATCH-202505-001\",\"inspector\":\"INSPECTOR-001\",\"result\":\"合格\",\"handleTime\":\"1717089000\"}"]}'
```

---

### UpdatePartStatus

**方法功能：** 更新零部件状态（在产、在售、召回、报废等）

**输入参数：**
- `partID` (string) - 零部件 ID
- `status` (string) - 新状态（在产/在售/召回/报废/NORMAL）

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 零部件生产厂商组织成员

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"UpdatePartStatus","Args":["ENG-PISTON-202505","在售"]}'
```

---

## 供应链模块

### CreateSupplyOrder

**方法功能：** 创建供应链订单

**输入参数：**
- `orderJSON` (string) - 订单信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 整车车企组织成员

**SupplyChainData 结构体字段：**
```json
{
  "orderID": "ORDER-202505-001",
  "partID": "ENG-PISTON-202505",
  "fromOrg": "MANUFACTURER-001",
  "toOrg": "AUTOMAKER-001",
  "quantity": 100,
  "status": "待发货",
  "createTime": "1717084800"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateSupplyOrder","Args":["{\"orderID\":\"ORDER-202505-001\",\"partID\":\"ENG-PISTON-202505\",\"fromOrg\":\"MANUFACTURER-001\",\"toOrg\":\"AUTOMAKER-001\",\"quantity\":100,\"status\":\"待发货\",\"createTime\":\"1717084800\"}"]}'
```

---

### CreateLogisticsData

**方法功能：** 创建物流数据记录

**输入参数：**
- `logisticsJSON` (string) - 物流数据的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 物流服务商组织成员

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateLogisticsData","Args":["{\"logisticsID\":\"LOG-202505-001\",\"orderID\":\"ORDER-202505-001\",\"location\":\"上海\",\"status\":\"运输中\",\"updateTime\":\"1717084800\"}"]}'
```

---

### UpdateSupplyChainStage

**方法功能：** 更新供应链阶段状态

**输入参数：**
- `stageJSON` (string) - 阶段信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 物流服务商组织成员

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"UpdateSupplyChainStage","Args":["{\"orderID\":\"ORDER-202505-001\",\"stage\":\"已签收\"}"]}'
```

---

### CreateReconciliation

**方法功能：** 创建对账记录

**输入参数：**
- `reconJSON` (string) - 对账数据的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 物流服务商组织成员

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateReconciliation","Args":["{\"reconID\":\"RECON-202505-001\",\"orderID\":\"ORDER-202505-001\",\"quantity\":100,\"reconTime\":\"1717084800\"}"]}'
```

---

## 售后模块

### CreateFaultReport

**方法功能：** 创建故障报告

**输入参数：**
- `faultJSON` (string) - 故障报告的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 4S店/售后中心组织成员

**FaultReport 结构体字段：**
```json
{
  "faultID": "FAULT-202505-001",
  "partID": "ENG-PISTON-202505",
  "vin": "LSVAA4182FS123456",
  "faultType": "性能异常",
  "faultDescription": "发动机运行不稳定",
  "reporter": "AFTERSALE-001",
  "reportTime": "1717084800"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateFaultReport","Args":["{\"faultID\":\"FAULT-202505-001\",\"partID\":\"ENG-PISTON-202505\",\"vin\":\"LSVAA4182FS123456\",\"faultType\":\"性能异常\",\"faultDescription\":\"发动机运行不稳定\",\"reporter\":\"AFTERSALE-001\",\"reportTime\":\"1717084800\"}"]}'
```

---

### CreateRecallRecord

**方法功能：** 创建召回记录

**输入参数：**
- `recallJSON` (string) - 召回信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 整车车企组织成员

**RecallRecord 结构体字段：**
```json
{
  "recallID": "RECALL-202505-001",
  "partID": "ENG-PISTON-202505",
  "batchNo": "BATCH-202505-001",
  "recallReason": "质量缺陷",
  "recallScope": "全国",
  "initiator": "AUTOMAKER-001",
  "initiateTime": "1717084800"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateRecallRecord","Args":["{\"recallID\":\"RECALL-202505-001\",\"partID\":\"ENG-PISTON-202505\",\"batchNo\":\"BATCH-202505-001\",\"recallReason\":\"质量缺陷\",\"recallScope\":\"全国\",\"initiator\":\"AUTOMAKER-001\",\"initiateTime\":\"1717084800\"}"]}'
```

---

### CreateAftersaleRecord

**方法功能：** 创建售后记录

**输入参数：**
- `aftersaleJSON` (string) - 售后记录的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 4S店/售后中心组织成员

**AftersaleRecord 结构体字段：**
```json
{
  "aftersaleID": "AFTERSALE-202505-001",
  "partID": "ENG-PISTON-202505",
  "vin": "LSVAA4182FS123456",
  "aftersaleType": "故障报修",
  "aftersaleStatus": "待审核",
  "handler": "AFTERSALE-001",
  "processTime": "1717084800"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"CreateAftersaleRecord","Args":["{\"aftersaleID\":\"AFTERSALE-202505-001\",\"partID\":\"ENG-PISTON-202505\",\"vin\":\"LSVAA4182FS123456\",\"aftersaleType\":\"故障报修\",\"aftersaleStatus\":\"待审核\",\"handler\":\"AFTERSALE-001\",\"processTime\":\"1717084800\"}"]}'
```

---

### QueryAffectedParts

**方法功能：** 根据批次号查询受影响的零部件（用于召回等场景）

**输入参数：**
- `batchNo` (string) - 批次号

**输出结果：**
- `[]*Part` - 受影响的零部件列表
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryAffectedParts","Args":["BATCH-202505-001"]}'
```

---

### QueryAftersaleRecord

**方法功能：** 查询售后记录详情

**输入参数：**
- `aftersaleID` (string) - 售后记录 ID

**输出结果：**
- `*AftersaleRecord` - 售后记录详细信息
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryAftersaleRecord","Args":["AFTERSALE-202505-001"]}'
```

---

## 身份管理模块

### RegisterUser

**方法功能：** 注册用户身份，分配组织和权限

**输入参数：**
- `userJSON` (string) - 用户身份信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 各组织成员

**UserIdentity 结构体字段：**
```json
{
  "userID": "USER-001",
  "org": "零部件生产厂商",
  "role": "管理员",
  "registerTime": "1717084800"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"RegisterUser","Args":["{\"userID\":\"USER-001\",\"org\":\"零部件生产厂商\",\"role\":\"管理员\",\"registerTime\":\"1717084800\"}"]}'
```

---

### QueryUserPermissions

**方法功能：** 查询用户权限信息

**输入参数：**
- `userID` (string) - 用户 ID

**输出结果：**
- `*UserIdentity` - 用户权限信息
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryUserPermissions","Args":["USER-001"]}'
```

---

## AI 功能模块

### RecordQAInteraction

**方法功能：** 记录 AI 智能问答的交互历史，用于审计和追溯

**输入参数：**
- `qaJSON` (string) - 问答交互信息的 JSON 字符串

**输出结果：**
- `error` - 操作失败时返回错误信息

**权限要求：** 所有组织可执行

**QAInteraction 结构体字段：**
```json
{
  "qaID": "QA-202505-001",
  "userID": "USER-001",
  "query": "如何查询零部件信息？",
  "answer": "使用 QueryPart 方法根据 PartID 查询",
  "queryTime": "1717084800"
}
```

**调用示例：**
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"function":"RecordQAInteraction","Args":["{\"qaID\":\"QA-202505-001\",\"userID\":\"USER-001\",\"query\":\"如何查询零部件信息？\",\"answer\":\"使用 QueryPart 方法根据 PartID 查询\",\"queryTime\":\"1717084800\"}"]}'
```

---

### QueryQAInteraction

**方法功能：** 查询问答交互记录

**输入参数：**
- `qaID` (string) - 问答交互 ID

**输出结果：**
- `*QAInteraction` - 问答交互详细信息
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryQAInteraction","Args":["QA-202505-001"]}'
```

---

### QueryFaultProgress

**方法功能：** 查询故障处理进度

**输入参数：**
- `faultID` (string) - 故障报告 ID

**输出结果：**
- `*FaultReport` - 故障报告详细信息
- `error` - 查询失败时返回错误信息

**权限要求：** 所有组织可查询

**调用示例：**
```bash
peer chaincode query -C mychannel -n basic -c '{"function":"QueryFaultProgress","Args":["FAULT-202505-001"]}'
```

---

## 权限说明

系统采用基于 MSP（Membership Service Provider）的权限控制机制，各组织及其权限如下：

| 组织名称 | MSPID | 主要权限 |
|---------|-------|---------|
| 零部件生产厂商 | Org1MSP | 创建零部件、创建 BOM、创建生产数据、创建质检记录、更新零部件状态 |
| 整车车企 | Org2MSP | 创建供应链订单、创建召回记录 |
| 物流服务商 | Org3MSP | 创建物流数据、更新供应链阶段、创建对账记录 |
| 4S店/售后中心 | Org4MSP | 创建故障报告、创建售后记录 |
| 行业监管机构 | Org5MSP | 查询所有数据、监管审计 |

---

## 错误码说明

| 错误信息 | 说明 |
|---------|------|
| 获取调用者身份失败 | 无法获取调用者的 MSPID |
| 只有XXX组织的成员才能执行此操作 | 调用者没有执行该操作的权限 |
| 反序列化数据失败 | JSON 格式错误或字段不匹配 |
| XXX不能为空 | 必填字段缺失 |
| XXX不存在 | 查询的数据不存在 |
| XXX必须是：XXX | 字段值不符合要求 |

---

## 注意事项

1. 所有时间戳字段使用 Unix 时间戳（秒级）
2. JSON 参数需要正确转义，建议使用工具生成
3. 查询操作使用 `peer chaincode query` 命令
4. 写入操作使用 `peer chaincode invoke` 命令
5. 调用前请确保已正确设置环境变量和证书路径
6. 建议在测试网络中先进行功能验证

---

## 附录：组织 MSPID 配置

```go
const (
	MANUFACTURER_ORG_MSPID = "Org1MSP"  // 零部件生产厂商
	AUTOMAKER_ORG_MSPID     = "Org2MSP"  // 整车车企
	LOGISTICS_ORG_MSPID     = "Org3MSP"  // 物流服务商
	AFTERSALE_ORG_MSPID     = "Org4MSP"  // 4S店/售后中心
	REGULATOR_ORG_MSPID     = "Org5MSP"  // 行业监管机构
)
```

---

**文档版本：** v1.0  
**最后更新：** 2025-01-30  
**维护者：** 汽车零部件全生命周期可信管理系统项目组
