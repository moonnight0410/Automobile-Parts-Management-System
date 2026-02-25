# 汽车零部件管理系统 API 接口文档

## 目录

1. [概述](#概述)
2. [认证说明](#认证说明)
3. [通用响应格式](#通用响应格式)
4. [接口列表](#接口列表)
   - [认证接口](#认证接口)
   - [零部件接口](#零部件接口)
   - [BOM接口](#bom接口)
   - [生产数据接口](#生产数据接口)
   - [质检数据接口](#质检数据接口)
   - [采购订单接口](#采购订单接口)
   - [物流数据接口](#物流数据接口)
   - [故障报告接口](#故障报告接口)
   - [召回记录接口](#召回记录接口)
   - [售后记录接口](#售后记录接口)
5. [数据权限说明](#数据权限说明)
6. [错误码说明](#错误码说明)

---

## 概述

本文档描述汽车零部件管理系统的后端API接口规范。系统基于Hyperledger Fabric区块链技术，提供零部件全生命周期管理功能。

**基础URL**: `http://localhost:8080`

**支持的角色**:
- `manufacturer`: 制造商
- `automaker`: 整车车企
- `aftersale`: 售后中心

---

## 认证说明

### JWT Token 认证

除认证接口外，所有API请求都需要在请求头中携带JWT Token：

```
Authorization: Bearer <token>
```

Token通过登录接口获取，有效期为24小时。

---

## 通用响应格式

所有API响应均采用JSON格式，结构如下：

### 成功响应

```json
{
  "code": 200,
  "message": "操作成功",
  "data": { ... }
}
```

### 错误响应

```json
{
  "code": 400,
  "message": "错误描述",
  "data": null
}
```

---

## 接口列表

### 认证接口

#### 1. 用户登录

**接口地址**: `POST /api/auth/login`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 |

**请求示例**:
```json
{
  "username": "manufacturer1",
  "password": "password123"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "userID": "manufacturer1",
      "role": "manufacturer",
      "organization": "Org1"
    }
  }
}
```

#### 2. 用户注册

**接口地址**: `POST /api/auth/register`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 |
| role | string | 是 | 角色（manufacturer/automaker/aftersale） |

**请求示例**:
```json
{
  "username": "newuser",
  "password": "password123",
  "role": "manufacturer"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "注册成功",
  "data": null
}
```

---

### 零部件接口

#### 1. 创建零部件

**接口地址**: `POST /api/parts`

**权限要求**: 制造商

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| partID | string | 是 | 零部件ID |
| partName | string | 是 | 零部件名称 |
| partType | string | 是 | 零部件类型 |
| batchNo | string | 是 | 批次号 |
| manufacturer | string | 是 | 制造商ID |
| productionDate | string | 是 | 生产日期 |
| vin | string | 否 | VIN码 |
| status | string | 否 | 状态（默认：正常） |

**请求示例**:
```json
{
  "partID": "PART-001",
  "partName": "制动盘",
  "partType": "制动系统",
  "batchNo": "BATCH-2024-001",
  "manufacturer": "manufacturer1",
  "productionDate": "2024-01-15",
  "vin": "",
  "status": "正常"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取零部件列表

**接口地址**: `GET /api/parts`

**权限要求**: 所有已认证用户

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| batchNo | string | 否 | 批次号 |
| vin | string | 否 | VIN码 |

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "partID": "PART-001",
      "partName": "制动盘",
      "partType": "制动系统",
      "batchNo": "BATCH-2024-001",
      "manufacturer": "manufacturer1",
      "productionDate": "2024-01-15",
      "vin": "",
      "status": "正常"
    }
  ]
}
```

#### 3. 获取我的零部件列表

**接口地址**: `GET /api/parts/my`

**权限要求**: 所有已认证用户

**说明**: 
- 制造商角色：返回该制造商生产的所有零部件
- 其他角色：返回所有零部件

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "partID": "PART-001",
      "partName": "制动盘",
      "partType": "制动系统",
      "batchNo": "BATCH-2024-001",
      "manufacturer": "manufacturer1",
      "productionDate": "2024-01-15",
      "vin": "",
      "status": "正常"
    }
  ]
}
```

#### 4. 获取单个零部件

**接口地址**: `GET /api/parts/:id`

**权限要求**: 所有已认证用户

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | string | 是 | 零部件ID |

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": {
    "partID": "PART-001",
    "partName": "制动盘",
    "partType": "制动系统",
    "batchNo": "BATCH-2024-001",
    "manufacturer": "manufacturer1",
    "productionDate": "2024-01-15",
    "vin": "",
    "status": "正常"
  }
}
```

---

### BOM接口

#### 1. 创建BOM

**接口地址**: `POST /api/boms`

**权限要求**: 制造商

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| bomID | string | 是 | BOM ID |
| bomType | string | 是 | BOM类型（生产BOM/研发BOM/工程BOM） |
| productModel | string | 是 | 产品型号 |
| version | string | 是 | 版本号 |
| components | array | 是 | 组件列表 |
| status | string | 否 | 状态（默认：草稿） |

**请求示例**:
```json
{
  "bomID": "BOM-001",
  "bomType": "生产BOM",
  "productModel": "Model-A",
  "version": "v1.0",
  "components": [
    {
      "partID": "PART-001",
      "quantity": 2
    }
  ],
  "status": "草稿"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取BOM列表

**接口地址**: `GET /api/boms`

**权限要求**: 所有已认证用户

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "bomID": "BOM-001",
      "bomType": "生产BOM",
      "productModel": "Model-A",
      "version": "v1.0",
      "status": "已发布",
      "createTime": "2024-01-15"
    }
  ]
}
```

#### 3. 获取单个BOM

**接口地址**: `GET /api/boms/:id`

**权限要求**: 所有已认证用户

**路径参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | string | 是 | BOM ID |

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": {
    "bomID": "BOM-001",
    "bomType": "生产BOM",
    "productModel": "Model-A",
    "version": "v1.0",
    "components": [
      {
        "partID": "PART-001",
        "quantity": 2
      }
    ],
    "status": "已发布",
    "createTime": "2024-01-15"
  }
}
```

---

### 生产数据接口

#### 1. 创建生产数据

**接口地址**: `POST /api/production`

**权限要求**: 制造商

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| recordID | string | 是 | 记录ID |
| partID | string | 是 | 零部件ID |
| batchNo | string | 是 | 批次号 |
| productionLine | string | 是 | 生产线 |
| operator | string | 是 | 操作员 |
| operationTime | string | 是 | 操作时间 |

**请求示例**:
```json
{
  "recordID": "PROD-001",
  "partID": "PART-001",
  "batchNo": "BATCH-2024-001",
  "productionLine": "A线",
  "operator": "张三",
  "operationTime": "2024-01-15 10:30:00"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取生产数据列表

**接口地址**: `GET /api/production`

**权限要求**: 所有已认证用户

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "recordID": "PROD-001",
      "partID": "PART-001",
      "batchNo": "BATCH-2024-001",
      "productionLine": "A线",
      "operator": "张三",
      "operationTime": "2024-01-15 10:30:00"
    }
  ]
}
```

---

### 质检数据接口

#### 1. 创建质检数据

**接口地址**: `POST /api/quality`

**权限要求**: 制造商

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| inspectionID | string | 是 | 质检ID |
| partID | string | 是 | 零部件ID |
| batchNo | string | 是 | 批次号 |
| inspectionResult | string | 是 | 质检结果（合格/不合格） |
| handler | string | 是 | 处理人 |
| inspectionTime | string | 是 | 质检时间 |

**请求示例**:
```json
{
  "inspectionID": "QC-001",
  "partID": "PART-001",
  "batchNo": "BATCH-2024-001",
  "inspectionResult": "合格",
  "handler": "李四",
  "inspectionTime": "2024-01-15 14:00:00"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取质检数据列表

**接口地址**: `GET /api/quality`

**权限要求**: 所有已认证用户

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "inspectionID": "QC-001",
      "partID": "PART-001",
      "batchNo": "BATCH-2024-001",
      "inspectionResult": "合格",
      "handler": "李四",
      "inspectionTime": "2024-01-15 14:00:00"
    }
  ]
}
```

---

### 采购订单接口

#### 1. 创建采购订单

**接口地址**: `POST /api/orders`

**权限要求**: 整车车企

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| orderID | string | 是 | 订单ID |
| buyerOrg | string | 是 | 采购方组织 |
| sellerOrg | string | 是 | 销售方组织 |
| partID | string | 是 | 零部件ID |
| quantity | number | 是 | 数量 |
| deliveryDate | string | 是 | 交付日期 |
| status | string | 否 | 状态（默认：待处理） |

**请求示例**:
```json
{
  "orderID": "ORDER-001",
  "buyerOrg": "automaker1",
  "sellerOrg": "manufacturer1",
  "partID": "PART-001",
  "quantity": 100,
  "deliveryDate": "2024-02-15",
  "status": "待处理"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取采购订单列表

**接口地址**: `GET /api/orders`

**权限要求**: 所有已认证用户

**说明**:
- 整车车企：返回该车企创建的所有订单
- 制造商：返回该制造商作为销售方的所有订单

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "orderID": "ORDER-001",
      "buyerOrg": "automaker1",
      "sellerOrg": "manufacturer1",
      "partID": "PART-001",
      "quantity": 100,
      "deliveryDate": "2024-02-15",
      "status": "待处理",
      "createTime": "2024-01-15"
    }
  ]
}
```

---

### 物流数据接口

#### 1. 创建物流数据

**接口地址**: `POST /api/logistics`

**权限要求**: 整车车企

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| logisticsID | string | 是 | 物流ID |
| orderID | string | 是 | 订单ID |
| carrier | string | 是 | 物流商 |
| departureTime | string | 是 | 出发时间 |
| arrivalTime | string | 否 | 到达时间 |
| receiver | string | 否 | 签收人 |
| status | string | 否 | 状态（默认：待发货） |

**请求示例**:
```json
{
  "logisticsID": "LOG-001",
  "orderID": "ORDER-001",
  "carrier": "顺丰物流",
  "departureTime": "2024-01-15 10:00:00",
  "arrivalTime": "",
  "receiver": "",
  "status": "运输中"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取物流数据列表

**接口地址**: `GET /api/logistics`

**权限要求**: 所有已认证用户

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "logisticsID": "LOG-001",
      "orderID": "ORDER-001",
      "carrier": "顺丰物流",
      "departureTime": "2024-01-15 10:00:00",
      "arrivalTime": "2024-01-17 15:00:00",
      "receiver": "王五",
      "status": "已送达"
    }
  ]
}
```

---

### 故障报告接口

#### 1. 创建故障报告

**接口地址**: `POST /api/faults`

**权限要求**: 售后中心

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| faultID | string | 是 | 故障ID |
| partID | string | 是 | 零部件ID |
| vin | string | 是 | VIN码 |
| faultType | string | 是 | 故障类型 |
| riskProbability | number | 是 | 风险概率（0-100） |
| description | string | 否 | 故障描述 |
| status | string | 否 | 状态（默认：待审核） |

**请求示例**:
```json
{
  "faultID": "FAULT-001",
  "partID": "PART-001",
  "vin": "VIN-001",
  "faultType": "制动故障",
  "riskProbability": 85,
  "description": "制动效能下降",
  "status": "待审核"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取故障报告列表

**接口地址**: `GET /api/faults`

**权限要求**: 所有已认证用户

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "faultID": "FAULT-001",
      "partID": "PART-001",
      "vin": "VIN-001",
      "faultType": "制动故障",
      "riskProbability": 85,
      "description": "制动效能下降",
      "status": "待审核",
      "reportTime": "2024-01-15 10:30:00"
    }
  ]
}
```

---

### 召回记录接口

#### 1. 创建召回记录

**接口地址**: `POST /api/recalls`

**权限要求**: 售后中心

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| recallID | string | 是 | 召回ID |
| affectedBatchNos | array | 是 | 受影响批次号列表 |
| reason | string | 是 | 召回原因 |
| affectedPartsCount | number | 是 | 受影响零部件数量 |
| progress | number | 否 | 进度（0-100） |
| status | string | 否 | 状态（默认：进行中） |

**请求示例**:
```json
{
  "recallID": "RECALL-001",
  "affectedBatchNos": ["BATCH-2024-001", "BATCH-2024-002"],
  "reason": "制动系统潜在隐患",
  "affectedPartsCount": 150,
  "progress": 0,
  "status": "进行中"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取召回记录列表

**接口地址**: `GET /api/recalls`

**权限要求**: 所有已认证用户

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "recallID": "RECALL-001",
      "affectedBatchNos": ["BATCH-2024-001", "BATCH-2024-002"],
      "reason": "制动系统潜在隐患",
      "affectedPartsCount": 150,
      "progress": 75,
      "status": "进行中",
      "createTime": "2024-01-15 10:30:00"
    }
  ]
}
```

---

### 售后记录接口

#### 1. 创建售后记录

**接口地址**: `POST /api/aftersale-records`

**权限要求**: 售后中心

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| recordID | string | 是 | 记录ID |
| partID | string | 是 | 零部件ID |
| vin | string | 是 | VIN码 |
| serviceType | string | 是 | 服务类型 |
| description | string | 是 | 服务描述 |
| cost | number | 否 | 费用 |
| status | string | 否 | 状态 |

**请求示例**:
```json
{
  "recordID": "AFTER-001",
  "partID": "PART-001",
  "vin": "VIN-001",
  "serviceType": "维修",
  "description": "更换制动盘",
  "cost": 500,
  "status": "已完成"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": null
}
```

#### 2. 获取售后记录列表

**接口地址**: `GET /api/aftersale-records`

**权限要求**: 所有已认证用户

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": [
    {
      "recordID": "AFTER-001",
      "partID": "PART-001",
      "vin": "VIN-001",
      "serviceType": "维修",
      "description": "更换制动盘",
      "cost": 500,
      "status": "已完成",
      "createTime": "2024-01-15 10:30:00"
    }
  ]
}
```

---

## 数据权限说明

系统实现了基于角色的数据权限控制：

### 制造商（manufacturer）
- **零部件**：只能查看和管理自己生产的零部件
- **BOM**：可以创建和管理自己的BOM
- **生产数据**：可以创建生产数据
- **质检数据**：可以创建质检数据
- **采购订单**：只能查看自己作为销售方的订单（只读）

### 整车车企（automaker）
- **零部件**：可以查看所有零部件信息
- **BOM**：可以查看所有BOM信息（只读）
- **采购订单**：可以创建和管理自己的采购订单
- **物流数据**：可以创建和管理物流数据

### 售后中心（aftersale）
- **零部件**：可以查看与自己相关的零部件信息
- **故障报告**：可以创建和管理故障报告
- **召回记录**：可以创建和管理召回记录
- **售后记录**：可以创建和管理售后记录

---

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权（未登录或Token无效） |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

---

## 健康检查

**接口地址**: `GET /api/health`

**说明**: 用于监控系统运行状态，无需认证

**响应示例**:
```json
{
  "code": 200,
  "message": "ok",
  "data": {
    "status": "ok"
  }
}
```

---

*文档版本: 1.0.0*
*最后更新: 2024年*
