package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract structure 提供零部件数据上链功能
type SmartContract struct {
	contractapi.Contract
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("创建智能合约失败：%v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("启动智能合约失败：%v", err)
	}
}
