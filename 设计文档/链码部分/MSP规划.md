# MSP 规划

## 组织架构

本系统采用3个组织架构，分别对应汽车零部件全生命周期的不同参与方：

| 组织名称 | MSP ID | 组织域名 | Peer端口 | 角色描述 |
|---------|---------|----------|----------|----------|
| Org1 | Org1MSP | org1.example.com | 7051 | 零部件生产厂商 |
| Org2 | Org2MSP | org2.example.com | 9051 | 整车车企（采购方） |
| Org3 | Org3MSP | org3.example.com | 11051 | 4S店/售后中心 |

## 权限分配

### Org1MSP（零部件生产厂商）

**核心业务权限：**
- 零部件信息管理
- 生产数据记录
- 质量检验记录
- BOM（物料清单）管理
- BOM变更管理

**链码操作权限：**
- `CreatePart` - 创建零部件信息
- `CreateProductionData` - 创建生产数据
- `CreateQualityInspection` - 创建质量检验记录
- `CreateBOM` - 创建BOM
- `UpdateBOM` - 更新BOM
- `CompareBOM` - 比对BOM
- `SubmitBOMChange` - 提交BOM变更
- `UpdatePartStatus` - 更新零部件状态（在产、质检合格等）

**业务场景：**
- 零部件生产完成后，记录生产数据
- 质量检验合格后，更新零部件状态
- 创建和管理BOM信息
- 当BOM需要变更时，提交变更申请

---

### Org2MSP（整车车企/采购方）

**核心业务权限：**
- 采购订单管理
- 供应链环节管理
- 物流数据记录
- 零部件溯源查询
- 对账记录

**链码操作权限：**
- `CreateSupplyOrder` - 创建采购订单
- `CreateLogisticsData` - 创建物流数据
- `UpdateSupplyChainStage` - 更新供应链环节
- `CreateReconciliation` - 创建对账记录
- `QueryPart` - 查询零部件信息
- `QueryPartLifecycle` - 查询零部件全生命周期
- `QueryPartByBatchNo` - 按批次号查询零部件
- `QueryPartByVIN` - 按VIN码查询零部件
- `QueryBOM` - 查询BOM信息

**业务场景：**
- 向生产厂商下达采购订单
- 接收零部件入库，记录供应链环节
- 跟踪物流信息
- 定期与供应商进行对账
- 溯源查询零部件的生产、质检、物流信息

---

### Org3MSP（4S店/售后中心）

**核心业务权限：**
- 故障报修记录
- 售后服务记录
- 召回记录管理
- 零部件溯源查询
- 质量问题反馈

**链码操作权限：**
- `CreateFaultReport` - 创建故障上报记录
- `CreateRecallRecord` - 创建召回记录
- `CreateAftersaleRecord` - 创建售后记录
- `QueryAffectedParts` - 查询受影响零部件
- `QueryAftersaleRecord` - 查询售后记录
- `QueryPart` - 查询零部件信息
- `QueryPartLifecycle` - 查询零部件全生命周期
- `QueryPartByVIN` - 按VIN码查询零部件

**业务场景：**
- 客户报修时，记录故障信息
- 进行维修处理，记录售后记录
- 参与召回活动，记录召回信息
- 查询零部件的完整生命周期，辅助故障诊断
- 查询受影响零部件（如召回、质量问题）

---

## 公共权限

所有组织都具备以下公共权限：

**查询权限：**
- `QueryPart` - 查询零部件信息
- `QueryPartLifecycle` - 查询零部件全生命周期
- `QueryPartByBatchNo` - 按批次号查询零部件
- `QueryPartByVIN` - 按VIN码查询零部件
- `QueryBOM` - 查询BOM信息
- `QueryAftersaleRecord` - 查询售后记录

**用户管理权限：**
- `RegisterUser` - 注册用户
- `QueryUserPermissions` - 查询用户权限

**QA交互权限：**
- `RecordQAInteraction` - 记录QA交互
- `QueryQAInteraction` - 查询QA交互
- `QueryFaultProgress` - 查询故障进度

---

## 权限控制实现

在链码中，通过以下方式实现权限控制：

### 1. MSP ID 常量定义

```go
const (
    MANUFACTURER_ORG_MSPID = "ManufacturerMSP" // 零部件生产厂商
    AUTOMAKER_ORG_MSPID    = "AutomakerMSP"    // 整车车企
    AFTERSALE_ORG_MSPID    = "AftersaleMSP"    // 4S店/售后中心
)
```

### 2. 权限检查函数

```go
// checkPermission 检查调用者是否具有指定权限
func (s *SmartContract) checkPermission(ctx contractapi.TransactionContextInterface, allowedMSPID string) error {
    clientMSPID, err := s.getClientIdentityMSPID(ctx)
    if err != nil {
        return err
    }
    if clientMSPID != allowedMSPID {
        return fmt.Errorf("权限错误：只有 %s 组织的成员才能执行此操作", allowedMSPID)
    }
    return nil
}
```

### 3. 在每个方法中调用权限检查

```go
func (s *SmartContract) CreatePart(ctx contractapi.TransactionContextInterface, partJSON string) error {
    // 权限检查：仅零部件生产厂商可执行
    if err := s.checkPermission(ctx, MANUFACTURER_ORG_MSPID); err != nil {
        return err
    }
    // ... 业务逻辑
}
```

---

## 网络配置

### 通道配置
- 通道名称：`channel1`
- 已加入组织：Org1、Org2、Org3

### 链码配置
- 链码名称：`auto-system`
- 初始版本：`1.0`
- 初始化函数：`initLedger`

### Orderer配置
- Orderer名称：`orderer.example.com`
- Orderer端口：`7050`
- TLS主机名：`orderer.example.com`

---

## 环境变量配置

详细的组织环境变量配置请参考：[环境变量配置.md](./环境变量配置.md)

### Org1（ManufacturerMSP）
```bash
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
```

### Org2（AutomakerMSP）
```bash
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```

### Org3（AftersaleMSP）
```bash
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
export CORE_PEER_ADDRESS=localhost:11051
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
```

---

## 权限矩阵

| 链码方法 | ManufacturerMSP | AutomakerMSP | AftersaleMSP | 说明 |
|-----------|-----------------|---------------|---------------|------|
| CreatePart | ✓ | ✗ | ✗ | 创建零部件信息 |
| CreateProductionData | ✓ | ✗ | ✗ | 创建生产数据 |
| CreateQualityInspection | ✓ | ✗ | ✗ | 创建质量检验记录 |
| CreateBOM | ✓ | ✗ | ✗ | 创建BOM |
| UpdateBOM | ✓ | ✗ | ✗ | 更新BOM |
| CompareBOM | ✓ | ✗ | ✗ | 比对BOM |
| SubmitBOMChange | ✓ | ✗ | ✗ | 提交BOM变更 |
| UpdatePartStatus | ✓ | ✓ | ✗ | 更新零部件状态 |
| CreateSupplyOrder | ✗ | ✓ | ✗ | 创建采购订单 |
| CreateLogisticsData | ✗ | ✓ | ✗ | 创建物流数据 |
| UpdateSupplyChainStage | ✗ | ✓ | ✗ | 更新供应链环节 |
| CreateReconciliation | ✗ | ✓ | ✗ | 创建对账记录 |
| CreateFaultReport | ✗ | ✗ | ✓ | 创建故障上报记录 |
| CreateRecallRecord | ✗ | ✗ | ✓ | 创建召回记录 |
| CreateAftersaleRecord | ✗ | ✗ | ✓ | 创建售后记录 |
| QueryPart | ✓ | ✓ | ✓ | 查询零部件信息 |
| QueryPartLifecycle | ✓ | ✓ | ✓ | 查询零部件全生命周期 |
| QueryPartByBatchNo | ✓ | ✓ | ✓ | 按批次号查询零部件 |
| QueryPartByVIN | ✓ | ✓ | ✓ | 按VIN码查询零部件 |
| QueryBOM | ✓ | ✓ | ✓ | 查询BOM信息 |
| QueryAffectedParts | ✓ | ✓ | ✓ | 查询受影响零部件 |
| QueryAftersaleRecord | ✓ | ✓ | ✓ | 查询售后记录 |
| RegisterUser | ✓ | ✓ | ✓ | 注册用户 |
| QueryUserPermissions | ✓ | ✓ | ✓ | 查询用户权限 |
| RecordQAInteraction | ✓ | ✓ | ✓ | 记录QA交互 |
| QueryQAInteraction | ✓ | ✓ | ✓ | 查询QA交互 |
| QueryFaultProgress | ✓ | ✓ | ✓ | 查询故障进度 |

**图例：**
- ✓ = 有权限
- ✗ = 无权限

---

## 总结

1. **三组织架构**：系统采用3个组织架构，分别对应零部件生产厂商、整车车企、4S店/售后中心
2. **权限分离**：每个组织拥有独立的业务权限，确保数据安全和业务隔离
3. **MSP ID对应**：
   - Org1 → ManufacturerMSP
   - Org2 → AutomakerMSP
   - Org3 → AftersaleMSP
4. **链码权限控制**：通过`checkPermission`函数实现细粒度的权限控制
5. **公共查询权限**：所有组织都可以查询零部件、BOM、售后记录等公共信息
6. **环境变量配置**：详细的环境变量配置请参考[环境变量配置.md](./环境变量配置.md)

---

## 相关文档

- [环境变量配置.md](./环境变量配置.md) - 详细的组织环境变量配置
- [部署链码流程.md](./部署链码流程.md) - 链码部署流程
- [Fabric链码更新流程.md](./Fabric链码更新流程.md) - 链码更新流程
- [毕设选题：基于区块链的汽车零部件全生命周期可信管理系统设计与实现.md](./毕设选题：基于区块链的汽车零部件全生命周期可信管理系统设计与实现.md) - 毕设需求文档
