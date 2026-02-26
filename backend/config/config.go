package config

import (
	"os"
	"strconv"
)

// Config 应用配置结构体
// 包含服务器、数据库、JWT、Fabric区块链等所有配置信息
// 配置值从环境变量加载，支持默认值

type Config struct {
	// 服务器配置
	ServerPort string // 服务器监听端口

	// 数据库配置
	MongoURL string // MongoDB连接URL
	RedisURL string // Redis连接URL

	// JWT认证配置
	JWTSecret      string // JWT签名密钥
	JWTExpireHours int    // JWT令牌过期时间（小时）

	// Fabric区块链配置
	FabricEnabled        bool   // 是否启用Fabric区块链集成
	FabricConfigPath     string // Fabric配置文件路径
	FabricChannel        string // Fabric通道名称
	FabricChaincode      string // Fabric链码名称
	FabricMSPID          string // Fabric MSP ID
	FabricCertPath       string // Fabric证书路径
	FabricKeyPath        string // Fabric密钥路径
	FabricPeerHost       string // Fabric Peer节点主机名
	FabricPeerEndpoint   string // Fabric Peer节点端点
	FabricTLSCertPath    string // Fabric TLS证书路径
	FabricRetryCount     int    // Fabric连接重试次数
	FabricRetryDelay     int    // Fabric连接重试延迟（秒）
	FabricConnectTimeout int    // Fabric连接超时时间（秒）

	// AI服务配置
	AIServiceURL string // AI服务URL
}

// Load 加载应用配置
// 从环境变量读取配置值，不存在则使用默认值
// 返回：
//   - Config: 完整的配置结构体
func Load() Config {
	return Config{
		// 服务器配置
		ServerPort: getEnv("SERVER_PORT", "8080"),

		// 数据库配置
		MongoURL: getEnv("MONGO_URL", "mongodb://localhost:27017"),
		RedisURL: getEnv("REDIS_URL", "redis://localhost:6379"),

		// JWT认证配置
		JWTSecret:      getEnv("JWT_SECRET", "change-me"),
		JWTExpireHours: getEnvInt("JWT_EXPIRE_HOURS", 24),

		// Fabric区块链配置
		FabricEnabled:        getEnvBool("FABRIC_ENABLED", false),
		FabricConfigPath:     getEnv("FABRIC_CONFIG_PATH", ""),
		FabricChannel:        getEnv("FABRIC_CHANNEL", "channel1"),
		FabricChaincode:      getEnv("FABRIC_CHAINCODE", "auto-system"),
		FabricMSPID:          getEnv("FABRIC_MSP_ID", "Org1MSP"),
		FabricCertPath:       getEnv("FABRIC_CERT_PATH", "fabric-certs/org1/signcerts/cert.pem"),
		FabricKeyPath:        getEnv("FABRIC_KEY_PATH", "fabric-certs/org1/keystore/key.pem"),
		FabricPeerHost:       getEnv("FABRIC_PEER_HOST", "peer0.org1.example.com"),
		FabricPeerEndpoint:   getEnv("FABRIC_PEER_ENDPOINT", "192.168.220.129:7051"),
		FabricTLSCertPath:    getEnv("FABRIC_TLS_CERT_PATH", "fabric-certs/org1/tls/ca.crt"),
		FabricRetryCount:     getEnvInt("FABRIC_RETRY_COUNT", 3),
		FabricRetryDelay:     getEnvInt("FABRIC_RETRY_DELAY", 5),
		FabricConnectTimeout: getEnvInt("FABRIC_CONNECT_TIMEOUT", 30),

		// AI服务配置
		AIServiceURL: getEnv("AI_SERVICE_URL", "http://localhost:5000"),
	}
}

// getEnv 获取字符串类型环境变量，不存在则返回默认值
// 参数：
//   - key: 环境变量名
//   - fallback: 默认值
//
// 返回：
//   - string: 环境变量值或默认值
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// getEnvInt 获取整数类型环境变量，不存在或解析失败则返回默认值
// 参数：
//   - key: 环境变量名
//   - fallback: 默认值
//
// 返回：
//   - int: 环境变量值或默认值
func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

// getEnvBool 获取布尔类型环境变量，不存在或解析失败则返回默认值
// 参数：
//   - key: 环境变量名
//   - fallback: 默认值
//
// 返回：
//   - bool: 环境变量值或默认值
func getEnvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}
