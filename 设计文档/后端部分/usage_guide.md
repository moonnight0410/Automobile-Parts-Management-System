# Fabric使用说明文档

## 📋 目录
- [安装依赖](#安装依赖)
- [配置步骤](#配置步骤)
- [测试步骤](#测试步骤)
- [集成步骤](#集成步骤)
- [API使用示例](#api使用示例)

---

## 📦 安装依赖

### 前置条件

- Go 1.19 或更高版本
- Hyperledger Fabric网络运行在虚拟机上 (IP: 192.168.220.129)
- SSH访问虚拟机的权限

### 安装Go依赖

在项目根目录执行:

```powershell
cd backend
go mod download
```

### 验证依赖

```powershell
go mod verify
```

### 检查依赖列表

```powershell
go list -m all
```

主要依赖:
- `github.com/hyperledger/fabric-gateway v1.7.0` - Fabric Gateway SDK
- `github.com/gin-gonic/gin` - Web框架
- `go.mongodb.org/mongo-driver` - MongoDB驱动
- `github.com/go-redis/redis` - Redis客户端

---

## ⚙️ 配置步骤

### 步骤1: 复制证书文件

#### 方法A: 使用PowerShell脚本（推荐）

1. 编辑证书复制脚本配置:

   ```powershell
   # 打开 copy_certs_from_vm.ps1
   notepad .\copy_certs_from_vm.ps1
   ```

2. 修改虚拟机连接配置:

   ```powershell
   $vm_ip = "192.168.220.129"
   $vm_user = "your_username"  # 修改为您的虚拟机用户名
   $vm_fabric_path = "~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com"
   ```

3. 执行证书复制脚本:

   ```powershell
   # 以管理员身份运行PowerShell
   Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
   
   # 执行脚本
   .\copy_certs_from_vm.ps1
   ```

4. 验证证书文件:

   ```powershell
   # 检查证书目录
   ls .\fabric-certs\org1\
   
   # 检查证书文件
   Get-Content .\fabric-certs\org1\signcerts\cert.pem
   Get-Content .\fabric-certs\org1\keystore\key.pem
   Get-Content .\fabric-certs\org1\tls\ca.crt
   ```

#### 方法B: 手动复制证书

1. 复制Admin证书:

   ```powershell
   scp your_username@192.168.220.129:~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem ./fabric-certs/org1/signcerts/cert.pem
   ```

2. 复制Admin私钥:

   ```powershell
   scp your_username@192.168.220.129:~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/*_sk ./fabric-certs/org1/keystore/key.pem
   ```

3. 复制TLS证书:

   ```powershell
   scp your_username@192.168.220.129:~/fabric/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt ./fabric-certs/org1/tls/ca.crt
   ```

### 步骤2: 配置环境变量

#### 方法A: 使用.env文件（推荐）

1. 创建.env文件:

   ```powershell
   # 在backend目录下创建.env文件
   notepad .env
   ```

2. 添加配置内容:

   ```env
   # 服务器配置
   SERVER_PORT=8080
   
   # 数据库配置
   MONGO_URL=mongodb://localhost:27017
   REDIS_URL=redis://localhost:6379
   
   # JWT认证配置
   JWT_SECRET=your-secret-key-change-this
   JWT_EXPIRE_HOURS=24
   
   # Fabric区块链配置
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

#### 方法B: 使用系统环境变量

```powershell
# 设置Fabric配置
$env:FABRIC_ENABLED = "true"
$env:FABRIC_CHANNEL = "mychannel"
$env:FABRIC_CHAINCODE = "basic"
$env:FABRIC_MSPID = "Org1MSP"
$env:FABRIC_CERT_PATH = "./fabric-certs/org1/signcerts/cert.pem"
$env:FABRIC_KEY_PATH = "./fabric-certs/org1/keystore/key.pem"
$env:FABRIC_PEER_HOST = "peer0.org1.example.com"
$env:FABRIC_PEER_ENDPOINT = "192.168.220.129:7051"
$env:FABRIC_TLS_CERT_PATH = "./fabric-certs/org1/tls/ca.crt"
$env:FABRIC_RETRY_COUNT = "3"
$env:FABRIC_RETRY_DELAY = "5"
$env:FABRIC_CONNECT_TIMEOUT = "30"
```

### 步骤3: 验证配置

1. 检查证书文件是否存在:

   ```powershell
   Test-Path .\fabric-certs\org1\signcerts\cert.pem
   Test-Path .\fabric-certs\org1\keystore\key.pem
   Test-Path .\fabric-certs\org1\tls\ca.crt
   ```

2. 检查证书文件内容:

   ```powershell
   Get-Content .\fabric-certs\org1\signcerts\cert.pem | Select-Object -First 5
   Get-Content .\fabric-certs\org1\keystore\key.pem | Select-Object -First 5
   Get-Content .\fabric-certs\org1\tls\ca.crt | Select-Object -First 5
   ```

3. 检查虚拟机网络连接:

   ```powershell
   ping 192.168.220.129
   Test-NetConnection -ComputerName 192.168.220.129 -Port 7051
   ```

---

## 🧪 测试步骤

### 步骤1: 运行连接测试

```powershell
# 在backend目录下执行
go run test_fabric_connection_complete.go
```

### 步骤2: 查看测试结果

测试脚本会输出以下信息:

```
========================================
  Fabric连接测试工具 v1.0
========================================

ℹ️  开始时间: 2024-01-01 10:00:00
ℹ️  加载配置...
ℹ️  Fabric配置:
  启用状态: true
  MSP ID: Org1MSP
  Peer端点: 192.168.220.129:7051
  Peer主机: peer0.org1.example.com
  通道名称: mychannel
  链码名称: basic
  证书路径: ./fabric-certs/org1/signcerts/cert.pem
  私钥路径: ./fabric-certs/org1/keystore/key.pem
  TLS证书路径: ./fabric-certs/org1/tls/ca.crt

========================================
  步骤1: 检查证书文件存在性
========================================

✅ 检查Admin证书文件: Admin证书文件存在: ./fabric-certs/org1/signcerts/cert.pem (大小: 1234 字节)
✅ 检查Admin私钥文件: Admin私钥文件存在: ./fabric-certs/org1/keystore/key.pem (大小: 567 字节)
✅ 检查TLS证书文件: TLS证书文件存在: ./fabric-certs/org1/tls/ca.crt (大小: 890 字节)

========================================
  步骤2: 验证证书文件内容
========================================

✅ 验证Admin证书内容: Admin证书文件格式正确
✅ 验证Admin私钥内容: Admin私钥文件格式正确
✅ 验证TLS证书内容: TLS证书文件格式正确

========================================
  步骤3: 初始化Fabric服务
========================================

ℹ️  正在连接到Fabric网络...
ℹ️  这可能需要几秒钟，请耐心等待...
[Fabric] 开始初始化Fabric服务...
[Fabric] Fabric配置检查...
[Fabric]   MSP ID: Org1MSP
[Fabric]   Peer端点: 192.168.220.129:7051
[Fabric]   通道名称: mychannel
[Fabric]   链码名称: basic
[Fabric] 步骤1: 加载并解析X509证书...
[Fabric]   证书加载成功
[Fabric] 步骤2: 创建X509身份...
[Fabric]   身份创建成功
[Fabric] 步骤3: 加载私钥...
[Fabric]   私钥加载成功
[Fabric] 步骤4: 创建签名器...
[Fabric]   签名器创建成功
[Fabric] 步骤5: 配置TLS证书...
[Fabric]   TLS证书配置成功
[Fabric] 步骤6: 建立gRPC连接到 192.168.220.129:7051...
[Fabric]   gRPC连接建立成功
[Fabric] 步骤7: 连接到Fabric Gateway...
[Fabric]   Gateway连接成功
[Fabric] Fabric服务初始化完成
✅ 初始化Fabric服务: Fabric服务初始化成功

========================================
  步骤4: 测试Fabric连接
========================================

ℹ️  尝试调用链码查询函数...
ℹ️  链码查询成功！
ℹ️  查询结果: [{"Key":"part1","Record":{"ID":"part1","Name":"发动机","Type":"核心部件",...}}]
✅ 测试Fabric连接: Fabric连接成功，链码调用正常，返回数据长度: 1234 字节

========================================
  测试结果摘要
========================================

总测试数: 7
成功: 7
失败: 0

✅ 所有测试通过！✨
ℹ️  您的后端已成功连接到Fabric网络！
ℹ️  下一步：
  1. 将Fabric服务集成到您的应用中
  2. 在controller中调用Fabric服务
  3. 测试完整的业务流程

ℹ️  结束时间: 2024-01-01 10:00:05
```

### 步骤3: 处理测试失败

如果测试失败，请参考以下步骤:

1. 检查错误信息
2. 查看[配置说明文档](config_guide.md)中的常见问题
3. 检查虚拟机网络连接
4. 确认Fabric网络正在运行

---

## 🔗 集成步骤

### 步骤1: 修改main.go

在您的main.go文件中集成Fabric服务:

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "automobile-parts-backend/config"
    "automobile-parts-backend/controller"
    "automobile-parts-backend/service"

    "github.com/gin-gonic/gin"
)

func main() {
    log.Println("========================================")
    log.Println("  汽车零件管理系统 - 后端服务")
    log.Println("========================================\n")

    log.Println("步骤1: 加载配置...")
    cfg := config.Load()
    log.Println("  配置加载成功")

    log.Println("\n步骤2: 初始化Fabric服务...")
    var fabricService *service.FabricService
    var err error

    if cfg.FabricEnabled {
        log.Println("  Fabric功能已启用，正在连接...")
        fabricService, err = service.NewFabricService(cfg)
        if err != nil {
            log.Printf("  ⚠️  Fabric服务初始化失败: %v", err)
            log.Println("  应用将以不包含Fabric功能的方式启动")
            fabricService = nil
        } else {
            log.Println("  ✅ Fabric服务初始化成功")
        }
    } else {
        log.Println("  Fabric功能未启用")
    }

    log.Println("\n步骤3: 初始化控制器...")
    fabricController := controller.NewFabricController(fabricService)
    log.Println("  ✅ 控制器初始化成功")

    log.Println("\n步骤4: 配置Gin路由...")
    router := gin.Default()

    fabricController.RegisterRoutes(router)
    log.Println("  ✅ 路由注册成功")

    log.Println("\n步骤5: 启动HTTP服务器...")
    addr := ":" + cfg.ServerPort
    if addr == ":" {
        addr = ":8080"
    }

    srv := &http.Server{
        Addr:    addr,
        Handler: router,
    }

    go func() {
        log.Printf("  🚀 服务器启动成功，监听地址: %s\n", addr)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("服务器启动失败: %v", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("\n步骤6: 正在关闭服务器...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("服务器关闭失败: %v", err)
    }

    if fabricService != nil {
        log.Println("正在关闭Fabric服务...")
        if err := fabricService.Close(); err != nil {
            log.Printf("关闭Fabric服务时出错: %v", err)
        }
    }

    log.Println("✅ 服务器已停止")
}
```

### 步骤2: 创建Fabric控制器

参考[controller/fabric_controller_example.go](controller/fabric_controller_example.go)创建您的Fabric控制器。

### 步骤3: 注册路由

在main.go中注册Fabric路由:

```go
fabricController.RegisterRoutes(router)
```

### 步骤4: 启动应用

```powershell
go run main.go
```

---

## 🌐 API使用示例

### 健康检查

检查Fabric服务是否正常运行:

```bash
curl http://localhost:8080/api/fabric/health
```

响应示例:

```json
{
  "status": "healthy",
  "message": "Fabric服务运行正常"
}
```

### 获取所有零件

获取区块链上的所有零件信息:

```bash
curl http://localhost:8080/api/fabric/parts
```

响应示例:

```json
{
  "success": true,
  "message": "获取零件信息成功",
  "data": "[{\"Key\":\"part1\",\"Record\":{\"ID\":\"part1\",\"Name\":\"发动机\",\"Type\":\"核心部件\",\"Supplier\":\"供应商A\",\"Quantity\":10,\"UnitPrice\":5000.00,\"Description\":\"汽车发动机\"}}]"
}
```

### 根据ID获取零件

获取指定ID的零件信息:

```bash
curl http://localhost:8080/api/fabric/parts/part1
```

响应示例:

```json
{
  "success": true,
  "message": "获取零件信息成功",
  "data": "{\"ID\":\"part1\",\"Name\":\"发动机\",\"Type\":\"核心部件\",\"Supplier\":\"供应商A\",\"Quantity\":10,\"UnitPrice\":5000.00,\"Description\":\"汽车发动机\"}"
}
```

### 创建新零件

在区块链上创建新零件:

```bash
curl -X POST http://localhost:8080/api/fabric/parts \
  -H "Content-Type: application/json" \
  -d '{
    "part_id": "part2",
    "part_name": "变速箱",
    "part_type": "核心部件",
    "supplier": "供应商B",
    "quantity": 5,
    "unit_price": 3000.00,
    "description": "汽车变速箱"
  }'
```

响应示例:

```json
{
  "success": true,
  "message": "创建零件成功",
  "data": "TXID: abc123..."
}
```

### 更新零件信息

更新区块链上的零件信息:

```bash
curl -X PUT http://localhost:8080/api/fabric/parts/part1 \
  -H "Content-Type: application/json" \
  -d '{
    "part_id": "part1",
    "part_name": "发动机",
    "part_type": "核心部件",
    "supplier": "供应商A",
    "quantity": 15,
    "unit_price": 5500.00,
    "description": "汽车发动机（更新）"
  }'
```

响应示例:

```json
{
  "success": true,
  "message": "更新零件成功",
  "data": "TXID: def456..."
}
```

### 删除零件

从区块链上删除零件:

```bash
curl -X DELETE http://localhost:8080/api/fabric/parts/part1
```

响应示例:

```json
{
  "success": true,
  "message": "删除零件成功",
  "data": "TXID: ghi789..."
}
```

---

## 📝 使用检查清单

在开始使用之前，请确认以下步骤已完成:

- [ ] Go依赖已安装
- [ ] 证书文件已从虚拟机复制到本地
- [ ] 环境变量已正确配置
- [ ] 虚拟机网络连接正常
- [ ] Fabric网络正在运行
- [ ] 连接测试已通过
- [ ] main.go已集成Fabric服务
- [ ] Fabric控制器已创建
- [ ] 路由已注册
- [ ] 应用已成功启动

---

## 🎯 下一步

1. 根据您的业务需求修改链码函数
2. 扩展API接口
3. 添加身份验证和授权
4. 实现错误处理和日志记录
5. 部署到生产环境

---

## 📞 获取帮助

如果遇到问题，请:

1. 查看[配置说明文档](config_guide.md)
2. 运行测试脚本诊断问题
3. 检查应用日志输出
4. 参考Hyperledger Fabric官方文档

---

## 📚 相关文档

- [配置说明文档](config_guide.md)
- [Hyperledger Fabric官方文档](https://hyperledger-fabric.readthedocs.io/)
- [Fabric Gateway SDK文档](https://github.com/hyperledger/fabric-gateway)
