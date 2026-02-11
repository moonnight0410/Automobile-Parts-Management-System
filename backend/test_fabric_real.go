package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type TestResult struct {
	Name    string
	Success bool
	Message string
	Error   error
	Data    interface{}
}

type FabricConfig struct {
	MSPID        string
	PeerEndpoint string
	PeerHost     string
	Channel      string
	Chaincode    string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
}

type Part struct {
	PartID       string `json:"partID"`
	VIN          string `json:"vin"`
	BatchNo      string `json:"batchNo"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	CreateTime   string `json:"createTime"`
	Status       string `json:"status"`
}

type BOM struct {
	BOMID        string         `json:"bomID"`
	BOMType      string         `json:"bomType"`
	ProductModel string         `json:"productModel"`
	PartBatchNo  string         `json:"partBatchNo"`
	Version      string         `json:"version"`
	Creator      string         `json:"creator"`
	CreateTime   string         `json:"createTime"`
	Status       string         `json:"status"`
	MaterialList []MaterialItem `json:"materialList"`
}

type MaterialItem struct {
	MaterialID   string `json:"materialID"`
	MaterialName string `json:"materialName"`
	Spec         string `json:"spec"`
	Quantity     int    `json:"quantity"`
	SupplierID   string `json:"supplierID"`
}

type ProductionData struct {
	ProductionID   string            `json:"productionID"`
	PartID         string            `json:"partID"`
	BatchNo        string            `json:"batchNo"`
	Params         map[string]string `json:"params"`
	ProductionLine string            `json:"productionLine"`
	Operator       string            `json:"operator"`
	FinishTime     string            `json:"finishTime"`
}

type QualityInspection struct {
	InspectionID string            `json:"inspectionID"`
	PartID       string            `json:"partID"`
	BatchNo      string            `json:"batchNo"`
	Indicators   map[string]string `json:"indicators"`
	Result       string            `json:"result"`
	Handler      string            `json:"handler"`
	HandleTime   string            `json:"handleTime"`
	Disposal     string            `json:"disposal"`
}

type SupplyOrder struct {
	OrderID    string `json:"orderID"`
	Buyer      string `json:"buyer"`
	Seller     string `json:"seller"`
	PartID     string `json:"partID"`
	Quantity   int    `json:"quantity"`
	BatchNo    string `json:"batchNo"`
	AgreedTime string `json:"agreedTime"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

type FaultReport struct {
	FaultID     string  `json:"faultID"`
	PartID      string  `json:"partID"`
	VIN         string  `json:"vin"`
	Reporter    string  `json:"reporter"`
	Description string  `json:"description"`
	FaultType   string  `json:"faultType"`
	RiskProb    float64 `json:"riskProb"`
	ReportTime  string  `json:"reportTime"`
	Status      string  `json:"status"`
}

func printHeader(title string) {
	fmt.Println("")
	fmt.Println("========================================")
	fmt.Printf("  %s\n", title)
	fmt.Println("========================================")
	fmt.Println("")
}

func printSuccess(message string) {
	fmt.Printf("✅ %s\n", message)
}

func printError(message string) {
	fmt.Printf("❌ %s\n", message)
}

func printInfo(message string) {
	fmt.Printf("ℹ️  %s\n", message)
}

func printInfoWithLabel(label, message string) {
	fmt.Printf("ℹ️  %s: %s\n", label, message)
}

func printData(label string, data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		printInfoWithLabel(label, fmt.Sprintf("%v", data))
	} else {
		fmt.Printf("ℹ️  %s:\n%s\n", label, string(jsonData))
	}
}

func loadConfig() FabricConfig {
	certPath := os.Getenv("FABRIC_CERT_PATH")
	if certPath == "" {
		certPath = "fabric-certs/org1/signcerts/cert.pem"
	}

	keyPath := os.Getenv("FABRIC_KEY_PATH")
	if keyPath == "" {
		keyPath = "fabric-certs/org1/keystore/key.pem"
	}

	tlsCertPath := os.Getenv("FABRIC_TLS_CERT_PATH")
	if tlsCertPath == "" {
		tlsCertPath = "fabric-certs/org1/tls/ca.crt"
	}

	peerEndpoint := os.Getenv("FABRIC_PEER_ENDPOINT")
	if peerEndpoint == "" {
		peerEndpoint = "192.168.220.129:7051"
	}

	peerHost := os.Getenv("FABRIC_PEER_HOST")
	if peerHost == "" {
		peerHost = "peer0.org1.example.com"
	}

	mspID := os.Getenv("FABRIC_MSP_ID")
	if mspID == "" {
		mspID = "Org1MSP"
	}

	channel := os.Getenv("FABRIC_CHANNEL")
	if channel == "" {
		channel = "channel1"
	}

	chaincode := os.Getenv("FABRIC_CHAINCODE")
	if chaincode == "" {
		chaincode = "auto-system"
	}

	return FabricConfig{
		MSPID:        mspID,
		PeerEndpoint: peerEndpoint,
		PeerHost:     peerHost,
		Channel:      channel,
		Chaincode:    chaincode,
		CertPath:     certPath,
		KeyPath:      keyPath,
		TLSCertPath:  tlsCertPath,
	}
}

func checkFileExists(path string, description string) TestResult {
	_, err := os.Stat(path)
	if err != nil {
		return TestResult{
			Name:    fmt.Sprintf("检查%s文件", description),
			Success: false,
			Message: fmt.Sprintf("%s文件不存在: %s", description, path),
			Error:   err,
		}
	}

	fileInfo, _ := os.Stat(path)
	return TestResult{
		Name:    fmt.Sprintf("检查%s文件", description),
		Success: true,
		Message: fmt.Sprintf("%s文件存在: %s (大小: %d 字节)", description, path, fileInfo.Size()),
	}
}

func readX509Certificate(path string) (*x509.Certificate, error) {
	certPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, errors.New("invalid certificate pem")
	}
	return x509.ParseCertificate(block.Bytes)
}

func connectToFabricGateway(cfg FabricConfig) (*client.Gateway, *grpc.ClientConn, error) {
	printInfo("步骤1: 加载并解析X509证书...")

	cert, err := readX509Certificate(cfg.CertPath)
	if err != nil {
		return nil, nil, fmt.Errorf("读取证书失败: %w", err)
	}

	id, err := identity.NewX509Identity(cfg.MSPID, cert)
	if err != nil {
		return nil, nil, fmt.Errorf("创建身份失败: %w", err)
	}
	printSuccess("证书和身份加载成功")

	printInfo("步骤2: 加载私钥...")

	keyPEM, err := os.ReadFile(cfg.KeyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("读取私钥失败: %w", err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(keyPEM)
	if err != nil {
		return nil, nil, fmt.Errorf("解析私钥失败: %w", err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("创建签名器失败: %w", err)
	}
	printSuccess("私钥和签名器加载成功")

	printInfo("步骤3: 配置TLS证书...")

	tlsCertPEM, err := os.ReadFile(cfg.TLSCertPath)
	if err != nil {
		return nil, nil, fmt.Errorf("读取TLS证书失败: %w", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(tlsCertPEM) {
		return nil, nil, fmt.Errorf("解析TLS证书失败")
	}
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cfg.PeerHost)
	printSuccess("TLS证书配置成功")

	printInfo(fmt.Sprintf("步骤4: 建立gRPC连接到 %s...", cfg.PeerEndpoint))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	connection, err := grpc.DialContext(ctx, cfg.PeerEndpoint,
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithBlock())
	if err != nil {
		return nil, nil, fmt.Errorf("建立gRPC连接失败: %w", err)
	}
	printSuccess("gRPC连接建立成功")

	printInfo("步骤5: 连接到Fabric Gateway...")

	gateway, err := client.Connect(id, client.WithSign(sign), client.WithClientConnection(connection))
	if err != nil {
		connection.Close()
		return nil, nil, fmt.Errorf("连接Gateway失败: %w", err)
	}
	printSuccess("Gateway连接成功")

	return gateway, connection, nil
}

func testFabricConnection(cfg FabricConfig) TestResult {
	printInfo("读取TLS证书...")

	tlsCertContent, err := os.ReadFile(cfg.TLSCertPath)
	if err != nil {
		return TestResult{
			Name:    "测试Fabric网络连接",
			Success: false,
			Message: fmt.Sprintf("读取TLS证书失败: %v", err),
			Error:   err,
		}
	}

	printInfo(fmt.Sprintf("TLS证书文件大小: %d 字节", len(tlsCertContent)))

	certPool := x509.NewCertPool()

	if !certPool.AppendCertsFromPEM(tlsCertContent) {
		printError("无法使用AppendCertsFromPEM解析TLS证书")
		printInfo("尝试手动解析TLS证书...")

		cert, err := readX509Certificate(cfg.TLSCertPath)
		if err != nil {
			return TestResult{
				Name:    "测试Fabric网络连接",
				Success: false,
				Message: fmt.Sprintf("手动解析TLS证书失败: %v", err),
				Error:   err,
			}
		}

		certPool.AddCert(cert)
		printSuccess("手动解析TLS证书成功")
	} else {
		printSuccess("使用AppendCertsFromPEM解析TLS证书成功")
	}

	config := &tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	dialer := &tls.Dialer{
		Config: config,
	}

	printInfo(fmt.Sprintf("尝试连接到 %s...", cfg.PeerEndpoint))

	conn, err := dialer.Dial("tcp", cfg.PeerEndpoint)
	if err != nil {
		return TestResult{
			Name:    "测试Fabric网络连接",
			Success: false,
			Message: fmt.Sprintf("无法连接到Fabric Peer (%s): %v", cfg.PeerEndpoint, err),
			Error:   err,
		}
	}
	defer conn.Close()

	return TestResult{
		Name:    "测试Fabric网络连接",
		Success: true,
		Message: fmt.Sprintf("成功连接到Fabric Peer: %s", cfg.PeerEndpoint),
	}
}

func callChaincodeQuery(contract *client.Contract, functionName string, args ...string) TestResult {
	printInfo(fmt.Sprintf("调用链码查询方法: %s", functionName))
	printInfo(fmt.Sprintf("参数: %v", args))

	result, err := contract.EvaluateTransaction(functionName, args...)
	if err != nil {
		return TestResult{
			Name:    fmt.Sprintf("调用链码方法: %s", functionName),
			Success: false,
			Message: fmt.Sprintf("调用失败: %v", err),
			Error:   err,
		}
	}

	var data interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		data = string(result)
	}

	return TestResult{
		Name:    fmt.Sprintf("调用链码方法: %s", functionName),
		Success: true,
		Message: "调用成功",
		Data:    data,
	}
}

func callChaincodeSubmit(contract *client.Contract, functionName string, args ...string) TestResult {
	printInfo(fmt.Sprintf("调用链码提交方法: %s", functionName))
	printInfo(fmt.Sprintf("参数: %v", args))

	result, err := contract.SubmitTransaction(functionName, args...)
	if err != nil {
		return TestResult{
			Name:    fmt.Sprintf("调用链码方法: %s", functionName),
			Success: false,
			Message: fmt.Sprintf("调用失败: %v", err),
			Error:   err,
		}
	}

	var data interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		data = string(result)
	}

	return TestResult{
		Name:    fmt.Sprintf("调用链码方法: %s", functionName),
		Success: true,
		Message: "调用成功",
		Data:    data,
	}
}

func main() {
	printHeader("🚀 Fabric网络连接与链码调用测试 v2.0")

	printInfoWithLabel("开始时间", time.Now().Format("2006-01-02 15:04:05"))

	printHeader("步骤1: 加载配置")

	cfg := loadConfig()

	printInfo("Fabric配置:")
	printInfoWithLabel("MSP ID", cfg.MSPID)
	printInfoWithLabel("Peer端点", cfg.PeerEndpoint)
	printInfoWithLabel("Peer主机", cfg.PeerHost)
	printInfoWithLabel("通道名称", cfg.Channel)
	printInfoWithLabel("链码名称", cfg.Chaincode)
	printInfoWithLabel("证书路径", cfg.CertPath)
	printInfoWithLabel("私钥路径", cfg.KeyPath)
	printInfoWithLabel("TLS证书路径", cfg.TLSCertPath)

	results := []TestResult{}

	printHeader("步骤2: 检查证书文件")

	result := checkFileExists(cfg.CertPath, "Admin证书")
	results = append(results, result)
	if result.Success {
		printSuccess(result.Message)
	} else {
		printError(result.Message)
	}

	result = checkFileExists(cfg.KeyPath, "Admin私钥")
	results = append(results, result)
	if result.Success {
		printSuccess(result.Message)
	} else {
		printError(result.Message)
	}

	result = checkFileExists(cfg.TLSCertPath, "TLS证书")
	results = append(results, result)
	if result.Success {
		printSuccess(result.Message)
	} else {
		printError(result.Message)
	}

	allFilesValid := true
	for _, r := range results {
		if !r.Success {
			allFilesValid = false
			break
		}
	}

	if !allFilesValid {
		printHeader("测试结果")
		printError("证书文件检查失败，无法继续测试")
		printTestSummary(results)
		printInfoWithLabel("结束时间", time.Now().Format("2006-01-02 15:04:05"))
		os.Exit(1)
	}

	printHeader("步骤3: 测试Fabric网络连接")

	result = testFabricConnection(cfg)
	results = append(results, result)
	if result.Success {
		printSuccess(result.Message)
	} else {
		printError(result.Message)
	}

	if !result.Success {
		printHeader("测试结果")
		printError("Fabric网络连接失败，无法继续测试")
		printTestSummary(results)
		printInfoWithLabel("结束时间", time.Now().Format("2006-01-02 15:04:05"))
		os.Exit(1)
	}

	printHeader("步骤4: 连接到Fabric Gateway")

	var gateway *client.Gateway
	var connection *grpc.ClientConn
	var err error

	gateway, connection, err = connectToFabricGateway(cfg)
	if err != nil {
		printError(fmt.Sprintf("连接Fabric Gateway失败: %v", err))
		results = append(results, TestResult{
			Name:    "连接Fabric Gateway",
			Success: false,
			Message: fmt.Sprintf("连接失败: %v", err),
			Error:   err,
		})
		printTestSummary(results)
		printInfoWithLabel("结束时间", time.Now().Format("2006-01-02 15:04:05"))
		os.Exit(1)
	}
	defer func() {
		if gateway != nil {
			gateway.Close()
		}
		if connection != nil {
			connection.Close()
		}
	}()

	results = append(results, TestResult{
		Name:    "连接Fabric Gateway",
		Success: true,
		Message: "Fabric Gateway连接成功",
	})

	printHeader("步骤5: 获取通道和链码")

	network := gateway.GetNetwork(cfg.Channel)
	printSuccess(fmt.Sprintf("成功获取通道: %s", cfg.Channel))

	contract := network.GetContract(cfg.Chaincode)
	printSuccess(fmt.Sprintf("成功获取链码: %s", cfg.Chaincode))

	printHeader("步骤6: 测试链码查询方法")

	testPartID := "TEST-PART-001"

	printInfo("尝试调用链码初始化方法...")
	initResult := callChaincodeSubmit(contract, "InitLedger")
	results = append(results, initResult)
	if initResult.Success {
		printSuccess("InitLedger 调用成功")
	} else {
		printInfo(fmt.Sprintf("InitLedger 调用失败: %v", initResult.Error))
	}

	printInfo("尝试查询零部件...")
	queryResult := callChaincodeQuery(contract, "QueryPart", testPartID)
	results = append(results, queryResult)
	if queryResult.Success {
		printSuccess(fmt.Sprintf("QueryPart 调用成功"))
		if queryResult.Data != nil {
			printData(fmt.Sprintf("返回数据 (%s)", testPartID), queryResult.Data)
		}
	} else {
		printInfo(fmt.Sprintf("QueryPart 调用失败: %v", queryResult.Error))
	}

	printHeader("步骤7: 测试链码提交方法")

	printInfo("创建测试零部件...")
	testPart := Part{
		PartID:       testPartID,
		VIN:          "TEST-VIN-001",
		BatchNo:      "TEST-BATCH-001",
		Name:         "测试零部件",
		Type:         "测试类型",
		Manufacturer: "Org1MSP",
		CreateTime:   fmt.Sprintf("%d", time.Now().Unix()),
		Status:       "NORMAL",
	}

	partJSON, _ := json.Marshal(testPart)
	createPartResult := callChaincodeSubmit(contract, "CreatePart", string(partJSON))
	results = append(results, createPartResult)
	if createPartResult.Success {
		printSuccess("CreatePart 调用成功")
	} else {
		printInfo(fmt.Sprintf("CreatePart 调用失败: %v", createPartResult.Error))
	}

	printInfo("再次查询零部件...")
	queryResult = callChaincodeQuery(contract, "QueryPart", testPartID)
	results = append(results, queryResult)
	if queryResult.Success {
		printSuccess("QueryPart 调用成功")
		if queryResult.Data != nil {
			printData(fmt.Sprintf("返回数据 (%s)", testPartID), queryResult.Data)
		}
	} else {
		printInfo(fmt.Sprintf("QueryPart 调用失败: %v", queryResult.Error))
	}

	printInfo("创建测试BOM...")
	testBOM := BOM{
		BOMID:        "TEST-BOM-001",
		BOMType:      "生产BOM",
		ProductModel: "测试车型",
		PartBatchNo:  "TEST-BATCH-001",
		Version:      "V1.0",
		Creator:      "Org1MSP",
		CreateTime:   fmt.Sprintf("%d", time.Now().Unix()),
		Status:       "草稿",
		MaterialList: []MaterialItem{
			{
				MaterialID:   "MAT-001",
				MaterialName: "测试物料",
				Spec:         "测试规格",
				Quantity:     100,
				SupplierID:   "SUP-001",
			},
		},
	}

	bomJSON, _ := json.Marshal(testBOM)
	createBOMResult := callChaincodeSubmit(contract, "CreateBOM", string(bomJSON))
	results = append(results, createBOMResult)
	if createBOMResult.Success {
		printSuccess("CreateBOM 调用成功")
	} else {
		printInfo(fmt.Sprintf("CreateBOM 调用失败: %v", createBOMResult.Error))
	}

	printInfo("查询BOM...")
	queryBOMResult := callChaincodeQuery(contract, "QueryBOM", "TEST-BOM-001")
	results = append(results, queryBOMResult)
	if queryBOMResult.Success {
		printSuccess("QueryBOM 调用成功")
		if queryBOMResult.Data != nil {
			printData("返回数据 (TEST-BOM-001)", queryBOMResult.Data)
		}
	} else {
		printInfo(fmt.Sprintf("QueryBOM 调用失败: %v", queryBOMResult.Error))
	}

	printInfo("创建生产数据...")
	testProduction := ProductionData{
		ProductionID:   "PROD-001",
		PartID:         testPartID,
		BatchNo:        "TEST-BATCH-001",
		Params:         map[string]string{"温度": "180℃", "时长": "2h"},
		ProductionLine: "LINE-001",
		Operator:       "Org1MSP",
		FinishTime:     fmt.Sprintf("%d", time.Now().Unix()),
	}

	productionJSON, _ := json.Marshal(testProduction)
	createProdResult := callChaincodeSubmit(contract, "CreateProductionData", string(productionJSON))
	results = append(results, createProdResult)
	if createProdResult.Success {
		printSuccess("CreateProductionData 调用成功")
	} else {
		printInfo(fmt.Sprintf("CreateProductionData 调用失败: %v", createProdResult.Error))
	}

	printInfo("创建质检数据...")
	testQuality := QualityInspection{
		InspectionID: "QUAL-001",
		PartID:       testPartID,
		BatchNo:      "TEST-BATCH-001",
		Indicators:   map[string]string{"尺寸": "10±0.1mm", "硬度": "HRC50"},
		Result:       "合格",
		Handler:      "Org1MSP",
		HandleTime:   fmt.Sprintf("%d", time.Now().Unix()),
		Disposal:     "",
	}

	qualityJSON, _ := json.Marshal(testQuality)
	createQualityResult := callChaincodeSubmit(contract, "CreateQualityInspection", string(qualityJSON))
	results = append(results, createQualityResult)
	if createQualityResult.Success {
		printSuccess("CreateQualityInspection 调用成功")
	} else {
		printInfo(fmt.Sprintf("CreateQualityInspection 调用失败: %v", createQualityResult.Error))
	}

	printInfo("创建供应链订单...")
	testOrder := SupplyOrder{
		OrderID:    "ORDER-001",
		Buyer:      "Org2MSP",
		Seller:     "Org1MSP",
		PartID:     testPartID,
		Quantity:   10,
		BatchNo:    "TEST-BATCH-001",
		AgreedTime: fmt.Sprintf("%d", time.Now().Unix()),
		Status:     "待发货",
		CreateTime: fmt.Sprintf("%d", time.Now().Unix()),
	}

	orderJSON, _ := json.Marshal(testOrder)
	createOrderResult := callChaincodeSubmit(contract, "CreateSupplyOrder", string(orderJSON))
	results = append(results, createOrderResult)
	if createOrderResult.Success {
		printSuccess("CreateSupplyOrder 调用成功")
	} else {
		printInfo(fmt.Sprintf("CreateSupplyOrder 调用失败: %v", createOrderResult.Error))
	}

	printInfo("创建故障报告...")
	testFault := FaultReport{
		FaultID:     "FAULT-001",
		PartID:      testPartID,
		VIN:         "TEST-VIN-001",
		Reporter:    "Org2MSP",
		Description: "测试故障描述",
		FaultType:   "磨损类",
		RiskProb:    0.8,
		ReportTime:  fmt.Sprintf("%d", time.Now().Unix()),
		Status:      "待审核",
	}

	faultJSON, _ := json.Marshal(testFault)
	createFaultResult := callChaincodeSubmit(contract, "CreateFaultReport", string(faultJSON))
	results = append(results, createFaultResult)
	if createFaultResult.Success {
		printSuccess("CreateFaultReport 调用成功")
	} else {
		printInfo(fmt.Sprintf("CreateFaultReport 调用失败: %v", createFaultResult.Error))
	}

	printInfo("查询故障进度...")
	queryFaultResult := callChaincodeQuery(contract, "QueryFaultProgress", "FAULT-001")
	results = append(results, queryFaultResult)
	if queryFaultResult.Success {
		printSuccess("QueryFaultProgress 调用成功")
		if queryFaultResult.Data != nil {
			printData("返回数据 (FAULT-001)", queryFaultResult.Data)
		}
	} else {
		printInfo(fmt.Sprintf("QueryFaultProgress 调用失败: %v", queryFaultResult.Error))
	}

	printHeader("测试结果")
	printTestSummary(results)

	printInfoWithLabel("结束时间", time.Now().Format("2006-01-02 15:04:05"))

	failCount := 0
	for _, r := range results {
		if !r.Success {
			failCount++
		}
	}

	if failCount > 0 {
		os.Exit(1)
	}
}

func printTestSummary(results []TestResult) {
	successCount := 0
	failCount := 0

	for _, result := range results {
		if result.Success {
			successCount++
		} else {
			failCount++
		}
	}

	printHeader("测试结果摘要")
	printInfoWithLabel("总测试数", fmt.Sprintf("%d", len(results)))
	printInfoWithLabel("成功", fmt.Sprintf("%d", successCount))
	printInfoWithLabel("失败", fmt.Sprintf("%d", failCount))

	if failCount == 0 {
		printSuccess("所有测试通过！")
	} else {
		printError("部分测试失败")
		printInfo("失败的测试:")
		for _, result := range results {
			if !result.Success {
				printInfo(fmt.Sprintf("  - %s: %s", result.Name, result.Message))
			}
		}
	}

	printInfo("")
	printInfo("建议操作:")
	if failCount > 0 {
		printInfo("  1. 检查Fabric网络是否正在运行")
		printInfo("  2. 确认链码是否已部署")
		printInfo("  3. 验证证书文件是否正确")
		printInfo("  4. 检查通道名称和链码名称是否正确")
	} else {
		printInfo("  1. Fabric网络连接正常")
		printInfo("  2. 链码调用成功")
		printInfo("  3. 可以开始集成到现有服务")
	}
}
