#!/usr/bin/env bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# 创建 Org5 的函数
function createOrg5 {
	infoln "注册 CA 管理员"
	mkdir -p ../organizations/peerOrganizations/org5.example.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/../organizations/peerOrganizations/org5.example.com/

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:13054 --caname ca-org5 --tls.certfiles "${PWD}/fabric-ca/org5/tls-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-13054-ca-org5.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-13054-ca-org5.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-13054-ca-org5.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-13054-ca-org5.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/../organizations/peerOrganizations/org5.example.com/msp/config.yaml"

	infoln "注册 peer0"
  set -x
	fabric-ca-client register --caname ca-org5 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/fabric-ca/org5/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "注册用户"
  set -x
  fabric-ca-client register --caname ca-org5 --id.name user1 --id.secret user1pw --id.type client --tls.certfiles "${PWD}/fabric-ca/org5/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "注册组织管理员"
  set -x
  fabric-ca-client register --caname ca-org5 --id.name org5admin --id.secret org5adminpw --id.type admin --tls.certfiles "${PWD}/fabric-ca/org5/tls-cert.pem"
  { set +x; } 2>/dev/null

  infoln "生成 peer0 msp"
  set -x
	fabric-ca-client enroll -u https://peer0:peer0pw@localhost:13054 --caname ca-org5 -M "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/msp" --tls.certfiles "${PWD}/fabric-ca/org5/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/msp/config.yaml" "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/msp/config.yaml"

  infoln "生成 peer0-tls 证书，使用 --csr.hosts 指定主题备用名称"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:13054 --caname ca-org5 -M "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls" --enrollment.profile tls --csr.hosts peer0.org5.example.com --csr.hosts localhost --tls.certfiles "${PWD}/fabric-ca/org5/tls-cert.pem"
  { set +x; } 2>/dev/null


  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/tlscacerts/"* "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/ca.crt"
  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/signcerts/"* "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/server.crt"
  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/keystore/"* "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/server.key"

  mkdir "${PWD}/../organizations/peerOrganizations/org5.example.com/msp/tlscacerts"
  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/tlscacerts/"* "${PWD}/../organizations/peerOrganizations/org5.example.com/msp/tlscacerts/ca.crt"

  mkdir "${PWD}/../organizations/peerOrganizations/org5.example.com/tlsca"
  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/tls/tlscacerts/"* "${PWD}/../organizations/peerOrganizations/org5.example.com/tlsca/tlsca.org5.example.com-cert.pem"

  mkdir "${PWD}/../organizations/peerOrganizations/org5.example.com/ca"
  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/peers/peer0.org5.example.com/msp/cacerts/"* "${PWD}/../organizations/peerOrganizations/org5.example.com/ca/ca.org5.example.com-cert.pem"

  infoln "生成用户 msp"
  set -x
	fabric-ca-client enroll -u https://user1:user1pw@localhost:13054 --caname ca-org5 -M "${PWD}/../organizations/peerOrganizations/org5.example.com/users/User1@org5.example.com/msp" --tls.certfiles "${PWD}/fabric-ca/org5/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/msp/config.yaml" "${PWD}/../organizations/peerOrganizations/org5.example.com/users/User1@org5.example.com/msp/config.yaml"

  infoln "生成组织管理员 msp"
  set -x
	fabric-ca-client enroll -u https://org5admin:org5adminpw@localhost:13054 --caname ca-org5 -M "${PWD}/../organizations/peerOrganizations/org5.example.com/users/Admin@org5.example.com/msp" --tls.certfiles "${PWD}/fabric-ca/org5/tls-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/../organizations/peerOrganizations/org5.example.com/msp/config.yaml" "${PWD}/../organizations/peerOrganizations/org5.example.com/users/Admin@org5.example.com/msp/config.yaml"
}
