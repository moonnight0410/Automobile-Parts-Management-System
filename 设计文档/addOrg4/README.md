## 将 Org4 添加到测试网络

您可以使用 `addOrg4.sh` 脚本将第四个组织添加到 Fabric 测试网络。`addOrg4.sh` 脚本生成 Org4 的加密材料，创建 Org4 组织定义，并将 Org4 添加到测试网络上的通道中。

在使用 `addOrg4.sh` 脚本之前，您需要先在 `test-network` 目录中运行 `./network.sh up createChannel`。

```
./network.sh up createChannel
cd addOrg4
./addOrg4.sh up
```

如果您使用 `network.sh` 创建了默认 `mychannel` 以外的通道，则需要将该通道名称传递给 `addOrg4.sh` 脚本。

```
./network.sh up createChannel -c channel1
cd addOrg4
./addOrg4.sh up -c channel1
```

您还可以重新运行 `addOrg4.sh` 脚本将 Org4 添加到其他通道。

```
cd ..
./network.sh createChannel -c channel2
cd addOrg4
./addOrg4.sh up -c channel2
```

有关更多信息，请使用 `./addOrg4.sh -h` 查看 `addOrg4.sh` 帮助文本。

### 端口分配

Org4 使用以下端口：
- Peer 端口：12051
- CA 端口：12054
- CouchDB 端口：9985

### 文件结构

```
addOrg4/
├── addOrg4.sh                    # 主脚本，用于添加 Org4 到网络
├── configtx.yaml                 # Org4 的配置定义
├── org4-crypto.yaml              # Org4 的加密配置
├── ccp-generate.sh              # 生成连接配置文件的脚本
├── ccp-template.json            # JSON 格式的连接配置模板
├── ccp-template.yaml            # YAML 格式的连接配置模板
├── compose/                     # Docker Compose 配置文件
│   ├── compose-ca-org4.yaml     # CA 服务配置
│   ├── compose-couch-org4.yaml  # CouchDB 服务配置
│   ├── compose-org4.yaml        # Peer 服务配置
│   ├── docker/                  # Docker 特定配置
│   │   ├── docker-compose-ca-org4.yaml
│   │   ├── docker-compose-couch-org4.yaml
│   │   └── docker-compose-org4.yaml
│   └── podman/                  # Podman 特定配置
│       ├── podman-compose-ca-org4.yaml
│       ├── podman-compose-couch-org4.yaml
│       └── podman-compose-org4.yaml
└── fabric-ca/                    # Fabric CA 配置
    ├── org4/
    │   └── fabric-ca-server-config.yaml
    └── registerEnroll.sh        # 注册和注册脚本
```

### 使用 CA 生成证书

如果您想使用 Fabric CA 而不是 cryptogen 生成证书，可以使用 `-ca` 标志：

```
./addOrg4.sh up -ca
```

### 使用 CouchDB

如果您想使用 CouchDB 作为状态数据库，可以使用 `-s couchdb` 标志：

```
./addOrg4.sh up -s couchdb
```

### 清理网络

要关闭网络并删除 Org4 节点：

```
./addOrg4.sh down
```
