# Fabric配置说明文档

## 📋 目录
- [环境变量配置](#环境变量配置)
- [证书文件配置](#证书文件配置)
- [Fabric网络配置](#fabric网络配置)
- [常见问题和解决方案](#常见问题和解决方案)

---

## 🔧 环境变量配置

### 基础配置

| 环境变量 | 说明 | 默认值 | 必填 |
|---------|------|--------|------|
| `SERVER_PORT` | HTTP服务器监听端口 | `8080` | 否 |
| `MONGO_URL` | MongoDB连接URL | `mongodb://localhost:27017` | 否 |
| `REDIS_URL` | Redis连接URL | `redis://localhost:6379` | 否 |
| `JWT_SECRET` | JWT签名密钥 | `your-secret-key` | 否 |
| `JWT_EXPIRE_HOURS` | JWT令牌过期时间（小时） | `24` | 否 |

### Fabric配置

| 环境变量 | 说明 | 默认值 | 必填 |
|---------|------|--------|------|
| `FABRIC_ENABLED` | 是否启用Fabric区块链集成 | `false` | 否 |
| `FABRIC_CONFIG_PATH` | Fabric配置文件路径 | - | 否 |
| `FABRIC_CHANNEL` | Fabric通道名称 | `mychannel` | 是 |
| `FABRIC_CHAINCODE` | Fabric链码名称 | `basic` | 是 |
| `FABRIC_MSPID` | Fabric MSP ID | `Org1MSP` | 是 |
| `FABRIC_CERT_PATH` | Fabric证书路径 | `./fabric-certs/org1/signcerts/cert.pem` | 是 |
| `FABRIC_KEY_PATH` | Fabric私钥路径 | `./fabric-certs/org1/keystore/key.pem` | 是 |
| `FABRIC_PEER_HOST` | Fabric Peer节点主机名 | `peer0.org1.example.com` | 是 |
| `FABRIC_PEER_ENDPOINT` | Fabric Peer节点端点 | `localhost:7051` | 是 |
| `FABRIC_TLS_CERT_PATH` | Fabric TLS证书路径 | `./fabric-certs/org1/tls/ca.crt` | 是 |
| `FABRIC_RETRY_COUNT` | Fabric连接重试次数 | `3` | 否 |
| `FABRIC_RETRY_DELAY` | Fabric连接重试延迟（秒） | `5` | 否 |
| `FABRIC_CONNECT_TIMEOUT` | Fabric连接超时时间（秒） | `30` | 否 |

---

## 📁 证书文件配置

### 证书目录结构

```
backend/
└── fabric-certs/
    └── org1/
        ├── signcerts/
        │   └── cert.pem          # Admin用户证书
        ├── keystore/
        │   └── key.pem           # Admin用户私钥
        └── tls/
            └── ca.crt            # TLS CA证书
```

### 证书文件说明

#### 1. Admin用户证书 (`cert.pem`)

**作用**: 用于标识Fabric网络中的Admin用户身份

**格式**: PEM格式的X.509证书

**示例内容**:
```
-----BEGIN CERTIFICATE-----
MIICLDCCAdKgAwIBAgIRAJd8z...
... (证书内容) ...
-----END CERTIFICATE-----
```

**获取方式**:
- 从虚拟机复制: `~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem`

#### 2. Admin用户私钥 (`key.pem`)

**作用**: 用于签名Fabric交易

**格式**: PEM格式的私钥

**示例内容**:
```
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFA...
... (私钥内容) ...
-----END PRIVATE KEY-----
```

**获取方式**:
- 从虚拟机复制: `~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/*_sk`

#### 3. TLS CA证书 (`ca.crt`)

**作用**: 用于建立与Fabric Peer节点的安全TLS连接

**格式**: PEM格式的X.509证书

**示例内容**:
```
-----BEGIN CERTIFICATE-----
MIICLDCCAdKgAwIBAgIRAJd8z...
... (证书内容) ...
-----END CERTIFICATE-----
```

**获取方式**:
- 从虚拟机复制: `~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt`

---

## 🔌 Fabric网络配置

### 虚拟机网络配置

**虚拟机IP**: `192.168.220.129`

**Fabric网络拓扑**:
```
┌─────────────────────────────────────────┐
│         虚拟机 (192.168.220.129)        │
│                                         │
│  ┌──────────────┐    ┌──────────────┐ │
│  │  Orderer     │    │    Peer      │ │
│  │  (7050)      │    │   (7051)     │ │
│  └──────────────┘    └──────────────┘ │
│         │                    │         │
│         └────────┬───────────┘         │
│                  │                      │
│          ┌───────┴───────┐             │
│          │   Channel     │             │
│          │  (mychannel)   │             │
│          └───────────────┘             │
│                                         │
└─────────────────────────────────────────┘
                  │
                  │ gRPC
                  │
┌─────────────────┴─────────────────────┐
│         本地开发环境                   │
│                                         │
│  ┌─────────────────────────────────┐  │
│  │      Go后端应用                  │  │
│  │   (使用fabric-gateway SDK)       │  │
│  └─────────────────────────────────┘  │
│                                         │
└─────────────────────────────────────────┘
```

### 连接配置示例

#### 配置文件方式 (`.env`)

```env
# Fabric配置
FABRIC_ENABLED=true
FABRIC_CHANNEL=mychannel
FABRIC_CHAINCODE=basic
FABRIC_MSPID=Org1MSP
FABRIC_CERT_PATH=./fabric-certs/org1/signcerts/cert.pem
FABRIC_KEY_PATH=./fabric-certs/org1/keystore/key.pem
FABRIC_PEER_HOST=peer0.org1.example.com
FABRIC_PEER_ENDPOINT=192.168.220.129:7051
FABRIC_TLS_CERT_PATH=./fabric-certs/org1/tls/ca.crt
FABRIC_RETRY_COUNT=3
FABRIC_RETRY_DELAY=5
FABRIC_CONNECT_TIMEOUT=30
```

#### 环境变量方式 (PowerShell)

```powershell
$env:FABRIC_ENABLED = "true"
$env:FABRIC_CHANNEL = "mychannel"
$env:FABRIC_CHAINCODE = "basic"
$env:FABRIC_MSPID = "Org1MSP"
$env:FABRIC_CERT_PATH = "./fabric-certs/org1/signcerts/cert.pem"
$env:FABRIC_KEY_PATH = "./fabric-certs/org1/keystore/key.pem"
$env:FABRIC_PEER_HOST = "peer0.org1.example.com"
$env:FABRIC_PEER_ENDPOINT = "192.168.220.129:7051"
$env:FABRIC_TLS_CERT_PATH = "./fabric-certs/org1/tls/ca.crt"
```

---

## ❓ 常见问题和解决方案

### 问题1: 证书文件不存在

**错误信息**:
```
failed to read certificate: open ./fabric-certs/org1/signcerts/cert.pem: no such file or directory
```

**原因**: 证书文件未从虚拟机复制到本地

**解决方案**:
1. 运行证书复制脚本:
   ```powershell
   .\copy_certs_from_vm.ps1
   ```

2. 检查脚本配置:
   - 确认虚拟机IP地址正确
   - 确认虚拟机用户名正确
   - 确认Fabric网络路径正确

3. 手动复制证书:
   ```powershell
   scp your_username@192.168.220.129:~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem ./fabric-certs/org1/signcerts/cert.pem
   ```

---

### 问题2: 证书格式错误

**错误信息**:
```
failed to parse certificate: invalid PEM data
```

**原因**: 证书文件内容不是有效的PEM格式

**解决方案**:
1. 检查证书文件内容:
   ```powershell
   Get-Content ./fabric-certs/org1/signcerts/cert.pem
   ```

2. 确认证书文件以以下内容开头和结尾:
   ```
   -----BEGIN CERTIFICATE-----
   ...
   -----END CERTIFICATE-----
   ```

3. 如果证书文件内容不正确，重新从虚拟机复制

---

### 问题3: 私钥格式错误

**错误信息**:
```
failed to parse private key: invalid PEM data
```

**原因**: 私钥文件内容不是有效的PEM格式

**解决方案**:
1. 检查私钥文件内容:
   ```powershell
   Get-Content ./fabric-certs/org1/keystore/key.pem
   ```

2. 确认私钥文件以以下内容开头和结尾:
   ```
   -----BEGIN PRIVATE KEY-----
   ...
   -----END PRIVATE KEY-----
   ```
   或
   ```
   -----BEGIN EC PRIVATE KEY-----
   ...
   -----END EC PRIVATE KEY-----
   ```

---

### 问题4: 连接超时

**错误信息**:
```
failed to create gRPC connection: context deadline exceeded
```

**原因**: 无法连接到Fabric Peer节点

**解决方案**:
1. 检查虚拟机网络连接:
   ```powershell
   ping 192.168.220.129
   ```

2. 检查Fabric Peer端口是否开放:
   ```powershell
   Test-NetConnection -ComputerName 192.168.220.129 -Port 7051
   ```

3. 增加连接超时时间:
   ```env
   FABRIC_CONNECT_TIMEOUT=60
   ```

4. 检查Fabric网络是否正在运行:
   ```bash
   # 在虚拟机上执行
   docker ps | grep fabric
   ```

---

### 问题5: TLS证书验证失败

**错误信息**:
```
failed to connect to gateway: x509: certificate signed by unknown authority
```

**原因**: TLS CA证书不匹配或过期

**解决方案**:
1. 确认TLS证书路径正确:
   ```env
   FABRIC_TLS_CERT_PATH=./fabric-certs/org1/tls/ca.crt
   ```

2. 确认TLS证书是从正确的Peer节点复制的:
   ```
   ~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
   ```

3. 检查TLS证书是否过期:
   ```powershell
   # 使用OpenSSL查看证书有效期
   openssl x509 -in ./fabric-certs/org1/tls/ca.crt -noout -dates
   ```

---

### 问题6: MSP ID不匹配

**错误信息**:
```
failed to connect to gateway: MSP ID mismatch
```

**原因**: 配置的MSP ID与Fabric网络中的MSP ID不匹配

**解决方案**:
1. 检查Fabric网络中的MSP ID:
   ```bash
   # 在虚拟机上执行
   ls ~/fabric/fabric-samples/test-network/organizations/peerOrganizations/
   ```

2. 确认配置的MSP ID正确:
   ```env
   FABRIC_MSPID=Org1MSP
   ```

3. 确保证书是从正确的组织复制的

---

### 问题7: 通道不存在

**错误信息**:
```
failed to get network: channel not found
```

**原因**: 配置的通道名称在Fabric网络中不存在

**解决方案**:
1. 检查Fabric网络中的通道:
   ```bash
   # 在虚拟机上执行
   peer channel list
   ```

2. 确认配置的通道名称正确:
   ```env
   FABRIC_CHANNEL=mychannel
   ```

3. 如果通道不存在，需要先创建通道

---

### 问题8: 链码不存在

**错误信息**:
```
failed to get contract: chaincode not found
```

**原因**: 配置的链码名称在Fabric网络中不存在

**解决方案**:
1. 检查Fabric网络中的链码:
   ```bash
   # 在虚拟机上执行
   peer lifecycle chaincode queryinstalled
   ```

2. 确认配置的链码名称正确:
   ```env
   FABRIC_CHAINCODE=basic
   ```

3. 如果链码不存在，需要先安装和部署链码

---

## 📝 配置检查清单

在启动应用之前，请确认以下配置项:

- [ ] Fabric功能已启用 (`FABRIC_ENABLED=true`)
- [ ] 证书文件已从虚拟机复制到本地
- [ ] Admin证书文件存在且格式正确
- [ ] Admin私钥文件存在且格式正确
- [ ] TLS证书文件存在且格式正确
- [ ] 虚拟机IP地址正确 (`192.168.220.129`)
- [ ] Peer端点配置正确 (`192.168.220.129:7051`)
- [ ] Peer主机名配置正确 (`peer0.org1.example.com`)
- [ ] 通道名称正确 (`mychannel`)
- [ ] 链码名称正确 (`basic`)
- [ ] MSP ID正确 (`Org1MSP`)
- [ ] 虚拟机网络连接正常
- [ ] Fabric网络正在运行
- [ ] Fabric Peer端口可访问 (7051)

---

## 📞 获取帮助

如果遇到以上未列出的问题，请:

1. 检查应用日志输出
2. 运行测试脚本: `go run test_fabric_connection_complete.go`
3. 查看Fabric网络日志
4. 参考Hyperledger Fabric官方文档: https://hyperledger-fabric.readthedocs.io/
