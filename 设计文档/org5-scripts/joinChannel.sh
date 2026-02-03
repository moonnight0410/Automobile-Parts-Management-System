#!/usr/bin/env bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# 此脚本设计为由 addOrg5.sh 运行，作为
# 向通道添加组织教程的第二步。
# 它将 org5 peers 加入到之前在
# 测试网络教程中设置的通道。

CHANNEL_NAME="$1"
DELAY="$2"
TIMEOUT="$3"
VERBOSE="$4"
: ${CHANNEL_NAME:="mychannel"}
: ${DELAY:="3"}
: ${TIMEOUT:="10"}
: ${VERBOSE:="false"}
COUNTER=1
MAX_RETRY=5

# 导入环境变量
# test network home var 指向 test-network 文件夹
# 原因是我们在这里使用 var 是考虑到 org5 特定文件夹
# 当为 org5 调用此脚本时作为 test-network/scripts/org5-scripts
# 值从默认值 $PWD (test-network)
# 更改为 ${PWD}/.. 以使导入工作
export TEST_NETWORK_HOME="${PWD}/.."
. ${TEST_NETWORK_HOME}/scripts/envVar.sh

# joinChannel ORG
joinChannel() {
  ORG=$1
  setGlobals $ORG
  local rc=1
  local COUNTER=1
  ## 有时加入需要时间，因此重试
  while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ] ; do
    sleep $DELAY
    set -x
    peer channel join -b $BLOCKFILE >&log.txt
    res=$?
    { set +x; } 2>/dev/null
    let rc=$res
    COUNTER=$(expr $COUNTER + 1)
	done
	cat log.txt
	verifyResult $res "在 $MAX_RETRY 次尝试后，peer0.org${ORG} 未能加入通道 '$CHANNEL_NAME' "
}

setAnchorPeer() {
  ORG=$1
  ${TEST_NETWORK_HOME}/scripts/setAnchorPeer.sh $ORG $CHANNEL_NAME
}

setGlobals 5
BLOCKFILE="${TEST_NETWORK_HOME}/channel-artifacts/${CHANNEL_NAME}.block"

echo "从 orderer 获取通道配置块..."
set -x
peer channel fetch 0 $BLOCKFILE -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com -c $CHANNEL_NAME --tls --cafile "$ORDERER_CA" >&log.txt
res=$?
{ set +x; } 2>/dev/null
cat log.txt
verifyResult $res "从 orderer 获取配置块失败"

infoln "将 org5 peer 加入通道..."
joinChannel 5

infoln "为 org5 设置锚节点..."
setAnchorPeer 5

successln "通道 '$CHANNEL_NAME' 已加入"
successln "Org5 peer 已成功添加到网络"
