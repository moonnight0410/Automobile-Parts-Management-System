# Hyperledger Fabric v2.5 链码更新流程

本文档详细介绍如何在 Fabric v2.5 网络中更新已部署的链码。

## 目录

1. [前置准备](#前置准备)
2. [链码更新流程](#链码更新流程)
3. [验证更新](#验证更新)
4. [常见问题](#常见问题)

---

## 前置准备

### 1. 环境要求

确保以下环境已正确配置：

- Fabric v2.5.x 网络正在运行
- Fabric CLI 工具已安装
- Docker 和 Docker Compose 已安装
- Go 环境（用于编译链码）

### 2. 检查当前链码状态

```bash
#准备org1的环境变量
export PATH=~/fabric/fabric-samples/bin:$PATH
export FABRIC_CFG_PATH=~/fabric/fabric-samples/config
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
# 进入 test-network 目录
cd fabric-samples/test-network

# 查看当前已安装的链码
./network.sh down
./network.sh up createChannel

# 查看已安装的链码
peer lifecycle chaincode queryinstalled
```

---

## 链码更新流程

### 步骤 1：准备新的链码文件

#### 1.1 停止网络（如果正在运行）

```bash
cd fabric-samples/test-network
./network.sh down
```

#### 1.2 替换链码文件

将更新后的链码文件复制到 Fabric 网络的链码目录：

```bash
# 假设你的链码在本地项目目录
# 复制到 test-network 的链码目录
cp -r /path/to/your/chaincode/* fabric-samples/test-network/chaincode/
```

或者，如果使用符号链接：

```bash
# 创建符号链接（推荐，便于开发）
ln -s /path/to/your/chaincode fabric-samples/test-network/chaincode/automobile-parts
```

#### 1.3 修改链码包名称（如果需要）

编辑 `fabric-samples/test-network/network.sh` 文件，找到链码名称定义：

```bash
# 找到类似这样的行
CC_NAME="automobile-parts"
CC_SRC_PATH="../chaincode/automobile-parts"
```

确保路径指向正确的链码目录。

---

### 步骤 2：启动网络

```bash
cd fabric-samples/test-network

# 启动网络并创建通道
./network.sh up createChannel
```

---

### 步骤 3：设置环境变量

为每个组织设置环境变量：

#### 3.1 Org1（零部件生产厂商）

```bash
# 设置 Org1 环境变量
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=${PWD}/../config/

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```

#### 3.2 Org2（整车车企）

```bash
# 设置 Org2 环境变量
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
```

---

### 步骤 4：打包新的链码

#### 4.1 打包链码（在 Org1 上执行）

```bash
# 设置 Org1 环境变量（如果还未设置）
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=${PWD}/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# 打包链码
peer lifecycle chaincode package automobile-parts.tar.gz --path ../chaincode/automobile-parts --language golang --label automobile-parts_1.1

# 查看打包结果
ls -lh automobile-parts.tar.gz
```

**参数说明：**
- `--path`: 链码源代码路径
- `--language`: 链码语言（golang）
- `--label`: 链码标签（建议使用版本号，如 `automobile-parts_1.1`）

---

### 步骤 5：安装新的链码

#### 5.1 在 Org1 上安装

```bash
peer lifecycle chaincode install automobile-parts.tar.gz
```

#### 5.2 在 Org2 上安装

```bash
# 切换到 Org2 环境
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

# 安装链码
peer lifecycle chaincode install automobile-parts.tar.gz
```

#### 5.3 在 Org3 上安装（如果需要）

```bash
# 切换到 Org3 环境
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
export CORE_PEER_ADDRESS=localhost:11051

# 安装链码
peer lifecycle chaincode install automobile-parts.tar.gz
```

---

### 步骤 6：查询已安装的链码

```bash
# 查询 Org1 上已安装的链码
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

peer lifecycle chaincode queryinstalled

# 输出示例：
# Installed chaincodes on peer:
# Package ID: automobile-parts_1.1:hash值, Label: automobile-parts_1.1
# Package ID: automobile-parts_1.0:hash值, Label: automobile-parts_1.0
```

**重要**：复制新链码的 Package ID，后续步骤需要使用。

---

### 步骤 7：批准链码定义

#### 7.1 Org1 批准

```bash
# 设置 Org1 环境
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# 批准链码定义（替换 PACKAGE_ID 为实际的 Package ID）
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name automobile-parts --version 1.1 --package-id <PACKAGE_ID> --sequence 2 --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

**参数说明：**
- `-o`: Orderer 地址
- `--channelID`: 通道名称
- `--name`: 链码名称
- `--version`: 链码版本（更新为 1.1）
- `--package-id`: 新链码的 Package ID（从步骤 6 获取）
- `--sequence`: 序列号（每次更新需要递增，这里是 2）
- `--init-required`: 是否需要初始化（首次部署需要，更新时通常不需要）

#### 7.2 Org2 批准

```bash
# 设置 Org2 环境
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051

# 批准链码定义
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID test --name test1 --version 1.1 --package-id <PACKAGE_ID> --sequence 2 --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

#### 7.3 Org3 批准（如果需要）

```bash
# 设置 Org3 环境
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
export CORE_PEER_ADDRESS=localhost:11051

# 批准链码定义
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID test --name test1 --version 1.1 --package-id <PACKAGE_ID> --sequence 2 --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

---

### 步骤 8：检查提交准备情况

```bash
# 检查链码定义是否准备好提交
peer lifecycle chaincode checkcommitreadiness --channelID test --name test1 --version 1.1 --sequence 2 --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --output json

# 输出示例：
# {
#   "approvals": {
#     "Org1MSP": true,
#     "Org2MSP": true,
#     "Org3MSP": true
#   }
# }
```

**注意**：所有需要背书的组织都必须显示为 `true` 才能继续。

---

### 步骤 9：提交链码定义

```bash
# 提交链码定义到通道
peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID test --name test1 --version 1.1 --sequence 2 --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```

**参数说明：**
- `--peerAddresses`: 提交组织的 peer 地址（多个组织用空格分隔）
- `--tlsRootCertFiles`: 对应 peer 的 TLS 证书（多个组织用空格分隔）

---

### 步骤 10：查询已提交的链码定义

```bash
# 查询已提交的链码定义
peer lifecycle chaincode querycommitted --channelID test --name test1 --output json

# 输出示例：
# {
#   "Sequence": 2,
#   "Name": "test1",
#   "Version": "1.1",
#   "InitRequired": false
# }
```

---

### 步骤 11：调用链码

#### 11.1 初始化链码（仅首次部署需要）

```bash
# 如果是首次部署，需要初始化
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C test -n test1 --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"InitLedger","Args":[]}'
```

#### 11.2 调用链码方法（测试更新）

```bash
# 测试创建零部件
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n automobile-parts --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["CreatePart","{\"partID\":\"TEST-001\",\"vin\":\"VIN123\",\"batchNo\":\"B001\",\"name\":\"测试零件\",\"type\":\"测试类型\",\"manufacturer\":\"测试厂商\",\"createTime\":\"1735689600\",\"status\":\"NORMAL\"}"]}'

# 查询零部件
peer chaincode query -C mychannel -n automobile-parts -c '{"Args":["QueryPart","TEST-001"]}'
```

---

## 验证更新

### 1. 检查链码版本

```bash
# 查询链码定义，确认版本已更新
peer lifecycle chaincode querycommitted --channelID mychannel --name automobile-parts
```

### 2. 测试新功能

```bash
# 测试新的链码方法
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n automobile-parts --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["<新方法名>","<参数>"]}'

# 查询结果
peer chaincode query -C mychannel -n automobile-parts -c '{"Args":["<查询方法名>","<参数>"]}'
```

### 3. 检查日志

```bash
# 查看 peer 日志
docker logs peer0.org1.example.com

# 查看 chaincode 容器日志
docker logs dev-peer0.org1.example.com-automobile-parts-1.1-<hash>
```

---

## 常见问题

### 问题 1：链码安装失败

**错误信息**：
```
Error: chaincode install failed with status: 500 - failed to invoke backing implementation of 'InstallChaincode': could not build chaincode: docker build failed
```

**解决方案**：
1. 检查链码路径是否正确
2. 确保 `go.mod` 文件存在且依赖正确
3. 清理 Docker 缓存：
   ```bash
   docker system prune -a
   ```

### 问题 2：Package ID 不匹配

**错误信息**：
```
Error: proposal failed with status: 500 - failed to invoke backing implementation of 'ApproveChaincodeDefinitionForMyOrg': chaincode definition for 'automobile-parts' at sequence 2 on channel 'mychannel' has different package ID
```

**解决方案**：
1. 使用 `peer lifecycle chaincode queryinstalled` 查看正确的 Package ID
2. 确保批准命令中的 `--package-id` 参数正确

### 问题 3：背书策略不满足

**错误信息**：
```
Error: proposal failed with status: 500 - failed to invoke backing implementation of 'CommitChaincodeDefinition': chaincode definition for 'automobile-parts' at sequence 2 on channel 'mychannel' has not been approved by enough organizations
```

**解决方案**：
1. 使用 `peer lifecycle chaincode checkcommitreadiness` 检查批准状态
2. 确保所有需要背书的组织都已批准

### 问题 4：序列号冲突

**错误信息**：
```
Error: proposal failed with status: 500 - failed to invoke backing implementation of 'CommitChaincodeDefinition': chaincode definition for 'automobile-parts' at sequence 2 on channel 'mychannel' already exists
```

**解决方案**：
1. 使用更高的序列号（如 3）
2. 或者先查询当前序列号，然后递增

### 问题 5：链码容器未启动

**错误信息**：
```
Error: endorsement failure during invoke. response: status:500 message:"make sure the chaincode automobile-parts has been successfully defined on channel mychannel and try again"
```

**解决方案**：
1. 等待链码容器启动（通常需要几秒）
2. 检查链码定义是否已提交：
   ```bash
   peer lifecycle chaincode querycommitted --channelID mychannel --name automobile-parts
   ```
3. 查看链码容器日志：
   ```bash
   docker logs dev-peer0.org1.example.com-automobile-parts-1.1-<hash>
   ```

---

## 快速更新脚本

为了简化更新流程，可以创建以下脚本：

### update-chaincode.sh

```bash
#!/bin/bash

# 配置变量
CC_NAME="automobile-parts"
CC_SRC_PATH="../chaincode/automobile-parts"
CC_VERSION="1.1"
CC_SEQUENCE="2"
CHANNEL_NAME="mychannel"

# 设置环境变量
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=${PWD}/../config/
export CORE_PEER_TLS_ENABLED=true

# Org1 环境变量
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# 1. 打包链码
echo "=== 打包链码 ==="
peer lifecycle chaincode package ${CC_NAME}.tar.gz --path ${CC_SRC_PATH} --language golang --label ${CC_NAME}_${CC_VERSION}

# 2. 在 Org1 上安装
echo "=== 在 Org1 上安装链码 ==="
peer lifecycle chaincode install ${CC_NAME}.tar.gz

# 3. 在 Org2 上安装
echo "=== 在 Org2 上安装链码 ==="
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
peer lifecycle chaincode install ${CC_NAME}.tar.gz

# 4. 查询已安装的链码
echo "=== 查询已安装的链码 ==="
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
peer lifecycle chaincode queryinstalled

# 提示用户输入 Package ID
echo "请复制新链码的 Package ID，然后按回车继续..."
read PACKAGE_ID

# 5. Org1 批准
echo "=== Org1 批准链码定义 ==="
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version ${CC_VERSION} --package-id ${PACKAGE_ID} --sequence ${CC_SEQUENCE} --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

# 6. Org2 批准
echo "=== Org2 批准链码定义 ==="
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version ${CC_VERSION} --package-id ${PACKAGE_ID} --sequence ${CC_SEQUENCE} --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

# 7. 检查提交准备情况
echo "=== 检查提交准备情况 ==="
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version ${CC_VERSION} --sequence ${CC_SEQUENCE} --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --output json

# 8. 提交链码定义
echo "=== 提交链码定义 ==="
peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version ${CC_VERSION} --sequence ${CC_SEQUENCE} --init-required --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

# 9. 查询已提交的链码定义
echo "=== 查询已提交的链码定义 ==="
peer lifecycle chaincode querycommitted --channelID ${CHANNEL_NAME} --name ${CC_NAME}

echo "=== 链码更新完成 ==="
```

**使用方法**：
```bash
# 1. 将脚本保存到 fabric-samples/test-network 目录
# 2. 赋予执行权限
chmod +x update-chaincode.sh

# 3. 执行脚本
./update-chaincode.sh
```

---

## 注意事项

1. **序列号管理**：每次更新链码时，序列号必须递增（1, 2, 3, ...）
2. **版本号管理**：建议使用语义化版本号（1.0, 1.1, 2.0, ...）
3. **备份**：更新前建议备份当前链码和账本数据
4. **测试**：在测试网络充分测试后再部署到生产环境
5. **回滚**：Fabric 不支持直接回滚，如需回滚需要重新部署旧版本

---

## 参考资料

- [Hyperledger Fabric 官方文档](https://hyperledger-fabric.readthedocs.io/)
- [Fabric v2.5 链码生命周期](https://hyperledger-fabric.readthedocs.io/en/release-2.5/chaincode_lifecycle.html)
- [Fabric 测试网络](https://hyperledger-fabric.readthedocs.io/en/release-2.5/test_network.html)
