# AI智能问答模块

## 概述

AI智能问答模块是基于知识库和意图识别的智能问答系统，为用户提供汽车零部件管理相关的智能咨询服务。

## 功能特性

1. **智能问答**: 基于知识库的语义搜索，提供准确的答案
2. **意图识别**: 自动识别用户问题意图，分类处理
3. **多轮对话**: 支持上下文相关的多轮对话
4. **区块链数据查询**: 集成区块链数据，实时查询故障进度、召回信息、零部件信息
5. **FAQ支持**: 内置常见问题库，快速回答常见问题

## 技术架构

- **前端**: Vue 3 + TypeScript + Ant Design Vue
- **后端**: Go (Gin框架) + Python (Flask)
- **AI模型**: sentence-transformers (多语言文本嵌入)
- **向量存储**: FAISS (Facebook AI Similarity Search)
- **知识库**: Markdown文档 + JSON格式FAQ

## 目录结构

```
backend/ai/
├── app.py                 # Flask API服务
├── config.py              # 配置文件
├── knowledge_base.py      # 知识库管理
├── intent_recognizer.py   # 意图识别
├── conversation_manager.py # 对话管理
├── qa_engine.py          # 问答引擎
├── blockchain_query.py   # 区块链数据查询
├── requirements.txt      # Python依赖
├── start.bat            # Windows启动脚本
├── test_ai.py           # 测试脚本
└── vector_db/           # 向量数据库（自动生成）

backend/knowledge/
├── 区块链基础知识.md
├── 故障上报流程.md
├── 故障上报材料.md
├── 故障处理进度查询.md
├── 故障分类与等级.md
├── 常见故障处理指南.md
├── 召回信息查询.md
├── 召回流程说明.md
├── 零部件质保政策.md
├── 零部件信息查询.md
├── 维修服务说明.md
├── 售后纠纷处理.md
├── 系统使用指南.md
├── 数据安全与隐私保护.md
├── 联系方式与客服.md
├── 系统更新与维护.md
├── 术语表.md
└── 常见问题FAQ.json
```

## 安装步骤

### 1. 安装Python依赖

进入AI模块目录并安装依赖：

```bash
cd backend/ai
pip install -r requirements.txt
```

### 2. 配置环境变量

编辑 `backend/.env` 文件，添加AI服务配置：

```env
# AI服务配置
AI_SERVICE_URL=http://localhost:5000
```

### 3. 启动AI服务

**Windows系统**:
双击运行 `start.bat` 或在命令行中执行：

```bash
cd backend/ai
start.bat
```

**Linux/Mac系统**:
```bash
cd backend/ai
python app.py
```

### 4. 启动Go后端

确保Go后端已启动，AI服务会自动连接到后端API。

```bash
cd backend
go run cmd/server/main.go
```

## 使用方法

### 通过API接口

#### 1. 提问

```bash
POST /api/ai/question

{
  "question": "如何上报零部件故障？",
  "userID": "user123"
}
```

响应示例：

```json
{
  "success": true,
  "message": "查询成功",
  "data": {
    "qaid": "QA-1234567890",
    "question": "如何上报零部件故障？",
    "intent": "故障上报",
    "confidence": 0.85,
    "answer": "您可以通过以下步骤上报零部件故障：\n1. 进入\"售后管理 > 故障报告\"页面\n2. 点击\"新建故障报告\"按钮\n3. 填写零部件ID、VIN码、故障描述等信息\n4. 提交后系统会自动分配故障ID并开始处理流程",
    "relatedActions": [
      {"label": "立即上报故障", "action": "navigate", "params": {"path": "/aftersale/fault"}}
    ],
    "relatedFaqs": ["faq_001", "faq_002"]
  }
}
```

#### 2. 查询对话历史

```bash
GET /api/ai/conversation?userID=user123&n=5
```

#### 3. 清空对话

```bash
DELETE /api/ai/conversation?userID=user123
```

#### 4. 查询统计信息

```bash
GET /api/ai/stats
```

#### 5. 健康检查

```bash
GET /api/ai/health
```

### 通过前端界面

1. 登录系统
2. 进入"AI智能助手"页面
3. 在输入框中输入问题
4. 点击"发送"按钮或按Enter键
5. 查看AI回答和相关操作建议

## 测试

运行测试脚本：

```bash
cd backend/ai
python test_ai.py
```

测试脚本会自动测试以下功能：
- AI服务健康检查
- 后端AI接口健康检查
- 提问功能测试（多个测试问题）
- 对话历史查询
- 统计信息查询

## 意图分类

系统支持以下意图分类：

1. **故障上报**: 上报故障、提交故障报告
2. **故障查询**: 查询故障进度、故障状态
3. **召回查询**: 查询召回信息、召回状态
4. **质保咨询**: 质保期、质保政策
5. **零部件查询**: 零部件信息、零部件状态
6. **进度查询**: 处理进度、订单状态
7. **维修服务**: 维修预约、维修流程
8. **纠纷处理**: 售后纠纷、投诉处理
9. **系统使用**: 系统功能、操作指南
10. **其他**: 其他问题

## 知识库管理

### 添加新知识

1. 在 `backend/knowledge/` 目录下创建新的Markdown文档
2. 使用标准的Markdown格式编写内容
3. 重启AI服务，知识库会自动重新加载

### 添加新FAQ

编辑 `backend/knowledge/常见问题FAQ.json`，添加新的FAQ条目：

```json
{
  "faq": [
    {
      "id": "faq_051",
      "question": "新问题",
      "category": "分类",
      "answer": "问题答案",
      "relatedActions": []
    }
  ]
}
```

## 性能优化

### 向量索引优化

修改 `config.py` 中的向量索引类型：

```python
VECTOR_INDEX_TYPE = "ivf"  # flat, ivf, hnsw
```

- `flat`: 精确搜索，速度较慢，适合小规模数据
- `ivf`: 近似搜索，速度快，适合中等规模数据
- `hnsw`: 高性能近似搜索，速度快，适合大规模数据

### 并发请求限制

修改 `config.py` 中的并发请求限制：

```python
MAX_CONCURRENT_REQUESTS = 10  # 根据服务器性能调整
```

### 对话轮数限制

修改 `config.py` 中的对话轮数限制：

```python
MAX_CONVERSATION_ROUNDS = 10  # 保留最近10轮对话
```

## 故障排查

### AI服务无法启动

1. 检查Python版本（需要3.8或更高）
2. 检查依赖是否安装完整
3. 检查端口5000是否被占用

### 知识库加载失败

1. 检查 `backend/knowledge/` 目录是否存在
2. 检查Markdown文档格式是否正确
3. 检查FAQ.json格式是否正确

### 区块链数据查询失败

1. 检查Go后端是否启动
2. 检查 `BLOCKCHAIN_API_BASE` 配置是否正确
3. 检查后端API是否正常响应

### 响应速度慢

1. 检查向量索引类型，考虑使用 `ivf` 或 `hnsw`
2. 检查并发请求限制，适当增加
3. 检查服务器资源使用情况

## 资源使用

### 内存使用

- 向量数据库: ~100-500MB（取决于知识库大小）
- 模型加载: ~500MB
- 总计: ~600MB-1GB

### CPU使用

- 查询处理: ~10-30% CPU
- 知识库加载: ~50-80% CPU（一次性）

### 磁盘使用

- 知识库文档: ~1-5MB
- 向量数据库: ~50-200MB
- 模型文件: ~400MB

## 开发说明

### 添加新的意图

1. 在 `config.py` 的 `INTENT_LABELS` 中添加新意图
2. 在 `intent_recognizer.py` 的 `_build_intent_keywords` 中添加关键词
3. 在 `qa_engine.py` 中添加对应的回答方法
4. 重启AI服务

### 修改回答模板

编辑 `qa_engine.py` 中的回答方法，修改返回的答案内容。

### 集成新的数据源

1. 在 `blockchain_query.py` 中添加新的查询方法
2. 在 `qa_engine.py` 中调用新的查询方法
3. 重启AI服务

## 许可证

本项目仅供学习和研究使用。

## 联系方式

如有问题或建议，请联系开发团队。
