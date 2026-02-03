#!/usr/bin/env bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# 此脚本设计为由 addOrg5.sh 运行，作为
# 向通道添加组织教程的第一步。
# 它创建并提交配置交易以
# 将 org5 添加到测试网络

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


# 导入
# test network home var 指向 test-network 文件夹
# 原因是我们在这里使用 var 是考虑到 org5 特定文件夹
# 当为 org5 调用此脚本时作为 test-network/scripts/org5-scripts
# 值从默认值 $PWD (test-network)
# 更改为 ${PWD}/.. 以使导入工作
export TEST_NETWORK_HOME="${PWD}/.."
. ${TEST_NETWORK_HOME}/scripts/configUpdate.sh 

infoln "创建配置交易以将 org5 添加到网络"

# 获取通道的配置，将其写入 config.json
fetchChannelConfig 1 ${CHANNEL_NAME} ${TEST_NETWORK_HOME}/channel-artifacts/config.json

# 修改配置以追加新组织
set -x
jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"Org5MSP":.[1]}}}}}' ${TEST_NETWORK_HOME}/channel-artifacts/config.json ${TEST_NETWORK_HOME}/organizations/peerOrganizations/org5.example.com/org5.json > ${TEST_NETWORK_HOME}/channel-artifacts/modified_config.json
{ set +x; } 2>/dev/null

# 计算配置更新，基于 config.json 和 modified_config.json 之间的差异，将其作为交易写入 org5_update_in_envelope.pb
createConfigUpdate ${CHANNEL_NAME} ${TEST_NETWORK_HOME}/channel-artifacts/config.json ${TEST_NETWORK_HOME}/channel-artifacts/modified_config.json ${TEST_NETWORK_HOME}/channel-artifacts/org5_update_in_envelope.pb

infoln "签名配置交易"
signConfigtxAsPeerOrg 1 ${TEST_NETWORK_HOME}/channel-artifacts/org5_update_in_envelope.pb

infoln "从另一个也签名它的 peer (peer0.org2) 提交交易"
setGlobals 2
set -x
peer channel update -f ${TEST_NETWORK_HOME}/channel-artifacts/org5_update_in_envelope.pb -c ${CHANNEL_NAME} -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "$ORDERER_CA"
{ set +x; } 2>/dev/null

successln "将 org5 添加到网络的配置交易已提交"
