#!/bin/bash
# 从 Git 更新链码并部署到 Fabric 网络

# 配置变量
GIT_REPO_URL="https://github.com/moonnight0410/Automobile-Parts-Management-System.git"
LOCAL_CHAINCODE_DIR="/home/jianyu-zou/code/Automobile-Parts-Management-System"
CHANNEL_NAME="test"
CHAINCODE_NAME="test1"
NEW_VERSION="1.2"  # 每次更新需要修改版本号
CHAINCODE_PATH="${LOCAL_CHAINCODE_DIR}/chaincode"

# 进入 Fabric 网络目录
cd ~/fabric/fabric-samples/test-network

echo "================================"
echo "开始更新链码流程"
echo "时间: $(date)"
echo "================================"

# 步骤1: 从 Git 拉取最新代码
echo "步骤1: 从 Git 拉取最新链码代码..."
if [ -d "$LOCAL_CHAINCODE_DIR" ]; then
    cd "$LOCAL_CHAINCODE_DIR"
    git pull origin main
    if [ $? -ne 0 ]; then
        echo "错误: Git 拉取失败!"
        exit 1
    fi
    echo "✓ Git 拉取成功"
else
    echo "错误: 链码目录不存在: $LOCAL_CHAINCODE_DIR"
    exit 1
fi

# 返回 test-network 目录
cd ~/fabric/fabric-samples/test-network

# 步骤2: 检查链码目录是否存在
echo -e "\n步骤2: 检查链码目录..."
if [ ! -d "$CHAINCODE_PATH" ]; then
    echo "错误: 链码路径不存在: $CHAINCODE_PATH"
    exit 1
fi

ls -la "$CHAINCODE_PATH"
echo "✓ 链码目录检查完成"

# 步骤3: 打包新版本链码
echo -e "\n步骤3: 打包新版本链码 (v${NEW_VERSION})..."
PACKAGE_FILE="${CHAINCODE_NAME}_v${NEW_VERSION}.tar.gz"
peer lifecycle chaincode package "$PACKAGE_FILE" \
    --lang golang \
    --path "$CHAINCODE_PATH" \
    --label "${CHAINCODE_NAME}_${NEW_VERSION}"

if [ $? -ne 0 ]; then
    echo "错误: 链码打包失败!"
    exit 1
fi

echo "✓ 链码打包完成: $PACKAGE_FILE"

# 步骤4: 获取当前序列号并递增
echo -e "\n步骤4: 获取当前链码序列号..."
CURRENT_SEQUENCE=$(peer lifecycle chaincode querycommitted --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --output json | jq -r '.sequence')
if [ -z "$CURRENT_SEQUENCE" ] || [ "$CURRENT_SEQUENCE" = "null" ]; then
    CURRENT_SEQUENCE=0
fi

NEW_SEQUENCE=$((CURRENT_SEQUENCE + 1))
echo "当前序列号: $CURRENT_SEQUENCE"
echo "新序列号: $NEW_SEQUENCE"

# 步骤5: 设置 Org1 环境变量
echo -e "\n步骤5: 设置 Org1 环境并安装链码..."
export PATH=~/fabric/fabric-samples/bin:$PATH
export FABRIC_CFG_PATH=~/fabric/fabric-samples/config
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
# 安装到 Org1
peer lifecycle chaincode install "$PACKAGE_FILE"
if [ $? -ne 0 ]; then
    echo "错误: Org1 链码安装失败!"
    exit 1
fi
echo "✓ Org1 链码安装成功"

# 步骤6: 设置 Org2 环境变量并安装
echo -e "\n步骤6: 设置 Org2 环境并安装链码..."
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
# 安装到 Org2
peer lifecycle chaincode install "$PACKAGE_FILE"
if [ $? -ne 0 ]; then
    echo "错误: Org2 链码安装失败!"
    exit 1
fi
echo "✓ Org2 链码安装成功"

# 步骤7: 获取包ID
echo -e "\n步骤7: 获取新安装的链码包ID..."
PACKAGE_ID=$(peer lifecycle chaincode queryinstalled | grep -o "${CHAINCODE_NAME}_${NEW_VERSION}:[a-zA-Z0-9]*")
if [ -z "$PACKAGE_ID" ]; then
    echo "错误: 无法获取包ID!"
    exit 1
fi
echo "包ID: $PACKAGE_ID"

# 步骤8: Org1 批准链码定义
echo -e "\n步骤8: Org1 批准新链码定义..."
export PATH=~/fabric/fabric-samples/bin:$PATH
export FABRIC_CFG_PATH=~/fabric/fabric-samples/config
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls true --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --version "$NEW_VERSION" --sequence "$NEW_SEQUENCE" --init-required --package-id "$PACKAGE_ID" --init-required

if [ $? -ne 0 ]; then
    echo "错误: Org1 批准链码定义失败!"
    exit 1
fi
echo "✓ Org1 批准成功"

# 步骤9: Org2 批准链码定义
echo -e "\n步骤9: Org2 批准新链码定义..."
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls true --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --version "$NEW_VERSION" --sequence "$NEW_SEQUENCE" --init-required --package-id "$PACKAGE_ID" --init-required

if [ $? -ne 0 ]; then
    echo "错误: Org2 批准链码定义失败!"
    exit 1
fi
echo "✓ Org2 批准成功"

# 步骤10: 检查提交准备状态
echo -e "\n步骤10: 检查链码提交准备状态..."
peer lifecycle chaincode checkcommitreadiness --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --version "$NEW_VERSION" --sequence "$NEW_SEQUENCE" --output json

# 步骤11: 提交链码定义
echo -e "\n步骤11: 提交链码定义到通道..."
export PATH=~/fabric/fabric-samples/bin:$PATH
export FABRIC_CFG_PATH=~/fabric/fabric-samples/config
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls true --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --version "$NEW_VERSION" --sequence "$NEW_SEQUENCE" --init-required --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

if [ $? -ne 0 ]; then
    echo "错误: 链码定义提交失败!"
    exit 1
fi
echo "✓ 链码定义提交成功"

步骤12: 验证更新
echo -e "\n步骤12: 验证链码更新..."
peer lifecycle chaincode querycommitted --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" --output json
#peer lifecycle chaincode querycommitted --channelID "test" --name "test1" --output json

# # 步骤13: 初始化链码（如果需要）
# echo -e "\n步骤13: 初始化链码..."
# peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls true --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --isInit -c '{"function":"initLedger","Args":[]}'

# if [ $? -ne 0 ]; then
#     echo "警告: 链码初始化失败（可能不需要初始化）"
# else
#     echo "✓ 链码初始化成功"
# fi

echo -e "\n================================"
echo "链码更新完成!"
echo "版本: $NEW_VERSION"
echo "序列号: $NEW_SEQUENCE"
echo "时间: $(date)"
echo "================================"