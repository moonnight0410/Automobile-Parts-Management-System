package service

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"automobile-parts-backend/config"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// FabricService Fabric区块链服务结构体
// 封装了与Hyperledger Fabric网络交互的核心功能

type FabricService struct {
	gateway    *client.Gateway  // Fabric Gateway客户端实例
	network    *client.Network  // Fabric通道网络实例
	contract   *client.Contract // Fabric链码合约实例
	connection *grpc.ClientConn // gRPC连接实例
}

// NewFabricService 创建新的Fabric服务实例
// 参数：
//   - cfg: 应用配置，包含Fabric相关配置
//
// 返回：
//   - *FabricService: Fabric服务实例
//   - error: 初始化过程中的错误
func NewFabricService(cfg config.Config) (*FabricService, error) {
	log.Println("[Fabric] 开始初始化Fabric服务...")

	if !cfg.FabricEnabled {
		log.Println("[Fabric] Fabric功能未启用，返回空服务实例")
		return &FabricService{}, nil
	}

	log.Printf("[Fabric] Fabric配置检查...")
	log.Printf("[Fabric]   MSP ID: %s", cfg.FabricMSPID)
	log.Printf("[Fabric]   Peer端点: %s", cfg.FabricPeerEndpoint)
	log.Printf("[Fabric]   通道名称: %s", cfg.FabricChannel)
	log.Printf("[Fabric]   链码名称: %s", cfg.FabricChaincode)

	if cfg.FabricMSPID == "" || cfg.FabricCertPath == "" || cfg.FabricKeyPath == "" ||
		cfg.FabricPeerEndpoint == "" || cfg.FabricPeerHost == "" || cfg.FabricTLSCertPath == "" {
		return nil, fmt.Errorf("fabric config incomplete: required fields missing")
	}

	var gateway *client.Gateway
	var connection *grpc.ClientConn
	var err error

	maxRetries := cfg.FabricRetryCount
	if maxRetries <= 0 {
		maxRetries = 3
	}

	retryDelay := time.Duration(cfg.FabricRetryDelay) * time.Second
	if retryDelay <= 0 {
		retryDelay = 5 * time.Second
	}

	for attempt := 1; attempt <= maxRetries; attempt++ {
		log.Printf("[Fabric] 连接尝试 %d/%d...", attempt, maxRetries)

		gateway, connection, err = connectToFabricGateway(cfg)
		if err == nil {
			log.Println("[Fabric] Fabric连接成功！")
			break
		}

		log.Printf("[Fabric] 连接失败 (尝试 %d/%d): %v", attempt, maxRetries, err)

		if attempt < maxRetries {
			log.Printf("[Fabric] 等待 %v 后重试...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Fabric after %d attempts: %w", maxRetries, err)
	}

	network := gateway.GetNetwork(cfg.FabricChannel)
	contract := network.GetContract(cfg.FabricChaincode)

	log.Println("[Fabric] Fabric服务初始化完成")

	return &FabricService{
		gateway:    gateway,
		network:    network,
		contract:   contract,
		connection: connection,
	}, nil
}

// connectToFabricGateway 连接到Fabric Gateway（内部函数）
// 包含完整的证书加载、身份创建、gRPC连接和Gateway连接
func connectToFabricGateway(cfg config.Config) (*client.Gateway, *grpc.ClientConn, error) {
	log.Println("[Fabric] 步骤1: 加载并解析X509证书...")
	cert, err := readX509Certificate(cfg.FabricCertPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read certificate: %w", err)
	}
	log.Println("[Fabric]   证书加载成功")

	log.Println("[Fabric] 步骤2: 创建X509身份...")
	id, err := identity.NewX509Identity(cfg.FabricMSPID, cert)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create identity: %w", err)
	}
	log.Println("[Fabric]   身份创建成功")

	log.Println("[Fabric] 步骤3: 加载私钥...")
	keyPEM, err := os.ReadFile(cfg.FabricKeyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read private key: %w", err)
	}
	privateKey, err := identity.PrivateKeyFromPEM(keyPEM)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	log.Println("[Fabric]   私钥加载成功")

	log.Println("[Fabric] 步骤4: 创建签名器...")
	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create signer: %w", err)
	}
	log.Println("[Fabric]   签名器创建成功")

	log.Println("[Fabric] 步骤5: 配置TLS证书...")
	tlsCert, err := readX509Certificate(cfg.FabricTLSCertPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read TLS certificate: %w", err)
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(tlsCert)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cfg.FabricPeerHost)
	log.Println("[Fabric]   TLS证书配置成功")

	log.Printf("[Fabric] 步骤6: 建立gRPC连接到 %s...", cfg.FabricPeerEndpoint)

	timeout := time.Duration(cfg.FabricConnectTimeout) * time.Second
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	connection, err := grpc.DialContext(ctx, cfg.FabricPeerEndpoint,
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithBlock())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}
	log.Println("[Fabric]   gRPC连接建立成功")

	log.Println("[Fabric] 步骤7: 连接到Fabric Gateway...")
	gateway, err := client.Connect(id, client.WithSign(sign), client.WithClientConnection(connection))
	if err != nil {
		connection.Close()
		return nil, nil, fmt.Errorf("failed to connect to gateway: %w", err)
	}
	log.Println("[Fabric]   Gateway连接成功")

	return gateway, connection, nil
}

// Query 执行链码查询操作（只读）
// 参数：
//   - ctx: 上下文，用于控制请求生命周期
//   - functionName: 链码查询函数名
//   - args: 链码函数参数
//
// 返回：
//   - []byte: 查询结果
//   - error: 查询过程中的错误
func (fs *FabricService) Query(ctx context.Context, functionName string, args ...string) ([]byte, error) {
	if fs.contract == nil {
		return nil, errors.New("fabric not initialized")
	}
	// 使用EvaluateTransaction执行只读查询
	return fs.contract.EvaluateTransaction(functionName, args...)
}

// Submit 执行链码提交操作（读写）
// 参数：
//   - ctx: 上下文，用于控制请求生命周期
//   - functionName: 链码提交函数名
//   - args: 链码函数参数
//
// 返回：
//   - []byte: 提交结果
//   - error: 提交过程中的错误
func (fs *FabricService) Submit(ctx context.Context, functionName string, args ...string) ([]byte, error) {
	if fs.contract == nil {
		return nil, errors.New("fabric not initialized")
	}
	// 使用SubmitTransaction执行读写操作
	return fs.contract.SubmitTransaction(functionName, args...)
}

// Close 关闭Fabric服务，释放资源
func (fs *FabricService) Close() error {
	log.Println("[Fabric] 关闭Fabric服务...")
	var errs []error

	if fs.gateway != nil {
		log.Println("[Fabric] 关闭Gateway连接...")
		if err := fs.gateway.Close(); err != nil {
			log.Printf("[Fabric] 关闭Gateway时出错: %v", err)
			errs = append(errs, fmt.Errorf("failed to close gateway: %w", err))
		}
	}

	if fs.connection != nil {
		log.Println("[Fabric] 关闭gRPC连接...")
		if err := fs.connection.Close(); err != nil {
			log.Printf("[Fabric] 关闭gRPC连接时出错: %v", err)
			errs = append(errs, fmt.Errorf("failed to close grpc connection: %w", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors occurred while closing Fabric service: %v", errs)
	}

	log.Println("[Fabric] Fabric服务已关闭")
	return nil
}

// readX509Certificate 从文件读取X509证书
// 参数：
//   - path: 证书文件路径
//
// 返回：
//   - *x509.Certificate: 解析后的X509证书
//   - error: 读取或解析错误
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
