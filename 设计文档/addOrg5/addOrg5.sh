#!/usr/bin/env bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# 此脚本通过添加第五个组织来扩展 Hyperledger Fabric 测试网络

# 将 $PWD/../bin 添加到 PATH 以确保使用正确的二进制文件
# 如果需要，可以注释掉以解析已安装的工具版本
export PATH=${PWD}/../../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}/../../config
export VERBOSE=false

. ../scripts/utils.sh

: ${CONTAINER_CLI:="docker"}
if command -v ${CONTAINER_CLI}-compose > /dev/null 2>&1; then
    : ${CONTAINER_CLI_COMPOSE:="${CONTAINER_CLI}-compose"}
else
    : ${CONTAINER_CLI_COMPOSE:="${CONTAINER_CLI} compose"}
fi
infoln "使用 ${CONTAINER_CLI} 和 ${CONTAINER_CLI_COMPOSE}"


# 打印使用说明信息
function printHelp () {
  echo "用法: "
  echo "  addOrg5.sh up|down|generate [-c <channel name>] [-t <timeout>] [-d <delay>] [-f <docker-compose-file>] [-s <dbtype>]"
  echo "  addOrg5.sh -h|--help (打印此消息)"
  echo "    <mode> - 以下之一: 'up', 'down', 或 'generate'"
  echo "      - 'up' - 将 org5 添加到示例网络。您需要先启动测试网络并创建通道。"
  echo "      - 'down' - 关闭测试网络和 org5 节点"
  echo "      - 'generate' - 生成所需的证书和组织定义"
  echo "    -c <channel name> - 测试网络通道名称 (默认为 \"mychannel\")"
  echo "    -ca <use CA> -  使用 CA 生成加密材料"
  echo "    -t <timeout> - CLI 超时持续时间（秒）(默认为 10)"
  echo "    -d <delay> - 延迟持续时间（秒）(默认为 3)"
  echo "    -s <dbtype> - 要使用的数据库后端: goleveldb (默认) 或 couchdb"
  echo "    -verbose - 详细模式"
  echo
  echo "通常，首先生成所需的证书和创世区块，然后启动网络。例如："
  echo
  echo "	addOrg5.sh generate"
  echo "	addOrg5.sh up"
  echo "	addOrg5.sh up -c mychannel -s couchdb"
  echo "	addOrg5.sh down"
  echo
  echo "使用所有默认值："
  echo "	addOrg5.sh up"
  echo "	addOrg5.sh down"
}

# 我们使用 cryptogen 工具为新组织生成加密材料
# (x509 证书)。运行工具后，证书将
# 放在与 org1 和 org2 一起的 organizations 文件夹中

# 使用 cryptogen 或 CAs 创建组织加密材料
function generateOrg5() {
  # 使用 cryptogen 创建加密材料
  if [ "$CRYPTO" == "cryptogen" ]; then
    which cryptogen
    if [ "$?" -ne 0 ]; then
      fatalln "未找到 cryptogen 工具。正在退出"
    fi
    infoln "使用 cryptogen 工具生成证书"

    infoln "创建 Org5 身份"

    set -x
    cryptogen generate --config=org5-crypto.yaml --output="../organizations"
    res=$?
    { set +x; } 2>/dev/null
    if [ $res -ne 0 ]; then
      fatalln "生成证书失败..."
    fi

  fi

  # 使用 Fabric CA 创建加密材料
  if [ "$CRYPTO" == "Certificate Authorities" ]; then
    fabric-ca-client version > /dev/null 2>&1
    if [[ $? -ne 0 ]]; then
      echo "错误！未找到 fabric-ca-client 二进制文件.."
      echo
      echo "按照 Fabric 文档中的说明安装 Fabric 二进制文件："
      echo "https://hyperledger-fabric.readthedocs.io/en/latest/install.html"
      exit 1
    fi

    infoln "使用 Fabric CA 生成证书"
    ${CONTAINER_CLI_COMPOSE} -f ${COMPOSE_FILE_CA_BASE} -f $COMPOSE_FILE_CA_ORG5 up -d 2>&1

    . fabric-ca/registerEnroll.sh

    sleep 10

    infoln "创建 Org5 身份"
    createOrg5

  fi

  infoln "为 Org5 生成 CCP 文件"
  ./ccp-generate.sh
}

# 生成通道配置交易
function generateOrg5Definition() {
  which configtxgen
  if [ "$?" -ne 0 ]; then
    fatalln "未找到 configtxgen 工具。正在退出"
  fi
  infoln "生成 Org5 组织定义"
  export FABRIC_CFG_PATH=$PWD
  set -x
  configtxgen -printOrg Org5MSP > ../organizations/peerOrganizations/org5.example.com/org5.json
  res=$?
  { set +x; } 2>/dev/null
  if [ $res -ne 0 ]; then
    fatalln "生成 Org5 组织定义失败..."
  fi
}

function Org5Up () {
  # 启动 org5 节点

  if [ "$CONTAINER_CLI" == "podman" ]; then
    cp ../podman/core.yaml ../../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/
  fi

  if [ "${DATABASE}" == "couchdb" ]; then
    DOCKER_SOCK=${DOCKER_SOCK} ${CONTAINER_CLI_COMPOSE} -f ${COMPOSE_FILE_BASE} -f $COMPOSE_FILE_ORG5 -f ${COMPOSE_FILE_COUCH_BASE} -f $COMPOSE_FILE_COUCH_ORG5 up -d 2>&1
  else
    DOCKER_SOCK=${DOCKER_SOCK} ${CONTAINER_CLI_COMPOSE} -f ${COMPOSE_FILE_BASE} -f $COMPOSE_FILE_ORG5 up -d 2>&1
  fi
  if [ $? -ne 0 ]; then
    fatalln "错误 ！！！无法启动 Org5 网络"
  fi
}

# 生成所需的证书、创世区块并启动网络
function addOrg5 () {
  # 如果测试网络未启动，则中止
  if [ ! -d ../organizations/ordererOrganizations ]; then
    fatalln "错误：请先运行 ./network.sh up createChannel。"
  fi

  # 如果工件不存在则生成
  if [ ! -d "../organizations/peerOrganizations/org5.example.com" ]; then
    generateOrg5
    generateOrg5Definition
  fi

  infoln "启动 Org5 peer"
  Org5Up

  # 创建配置交易以将 Org5 添加到网络
  infoln "生成并提交配置交易以添加 Org5"
  export FABRIC_CFG_PATH=${PWD}/../../config
  . ../scripts/org5-scripts/updateChannelConfig.sh $CHANNEL_NAME $CLI_DELAY $CLI_TIMEOUT $VERBOSE
  if [ $? -ne 0 ]; then
    fatalln "错误 ！！！无法创建配置交易"
  fi

  infoln "将 Org5 peers 加入网络"
  . ../scripts/org5-scripts/joinChannel.sh $CHANNEL_NAME $CLI_DELAY $CLI_TIMEOUT $VERBOSE
  if [ $? -ne 0 ]; then
    fatalln "错误 ！！！无法将 Org5 peers 加入网络"
  fi
}

# 关闭运行中的网络
function networkDown () {
    cd ..
    ./network.sh down
}

# 使用 crypto vs CA。默认为 cryptogen
CRYPTO="cryptogen"
# 超时持续时间 - CLI 在放弃之前等待另一个容器响应的持续时间
CLI_TIMEOUT=10
# 默认延迟
CLI_DELAY=3
# 通道名称默认为 "mychannel"
CHANNEL_NAME="mychannel"
# 将此用作 docker compose couch 文件
COMPOSE_FILE_COUCH_BASE=compose/compose-couch-org5.yaml
COMPOSE_FILE_COUCH_ORG5=compose/${CONTAINER_CLI}/docker-compose-couch-org5.yaml
# 将此用作默认的 docker-compose yaml 定义
COMPOSE_FILE_BASE=compose/compose-org5.yaml
COMPOSE_FILE_ORG5=compose/${CONTAINER_CLI}/docker-compose-org5.yaml
# 证书机构 compose 文件
COMPOSE_FILE_CA_BASE=compose/compose-ca-org5.yaml
COMPOSE_FILE_CA_ORG5=compose/${CONTAINER_CLI}/docker-compose-ca-org5.yaml
# 数据库
DATABASE="leveldb"

# 从环境变量获取 docker sock 路径
SOCK="${DOCKER_HOST:-/var/run/docker.sock}"
DOCKER_SOCK="${SOCK##unix://}"

# 解析命令行参数

## 解析模式
if [[ $# -lt 1 ]] ; then
  printHelp
  exit 0
else
  MODE=$1
  shift
fi

# 解析标志

while [[ $# -ge 1 ]] ; do
  key="$1"
  case $key in
  -h )
    printHelp
    exit 0
    ;;
  -c )
    CHANNEL_NAME="$2"
    shift
    ;;
  -ca )
    CRYPTO="Certificate Authorities"
    ;;
  -t )
    CLI_TIMEOUT="$2"
    shift
    ;;
  -d )
    CLI_DELAY="$2"
    shift
    ;;
  -s )
    DATABASE="$2"
    shift
    ;;
  -verbose )
    VERBOSE=true
    ;;
  * )
    errorln "未知标志: $key"
    printHelp
    exit 1
    ;;
  esac
  shift
done


# 确定是启动、停止、重新启动还是生成公告
if [ "$MODE" == "up" ]; then
  infoln "将 org5 添加到通道 '${CHANNEL_NAME}'，超时时间为 '${CLI_TIMEOUT}' 秒，CLI 延迟为 '${CLI_DELAY}' 秒，使用数据库 '${DATABASE}'"
  echo
elif [ "$MODE" == "down" ]; then
  EXPMODE="停止网络"
elif [ "$MODE" == "generate" ]; then
  EXPMODE="为 Org5 生成证书和组织定义"
else
  printHelp
  exit 1
fi

# 使用 docker compose 创建网络
if [ "${MODE}" == "up" ]; then
  addOrg5
elif [ "${MODE}" == "down" ]; then ## 清除网络
  networkDown
elif [ "${MODE}" == "generate" ]; then ## 生成工件
  generateOrg5
  generateOrg5Definition
else
  printHelp
  exit 1
fi
