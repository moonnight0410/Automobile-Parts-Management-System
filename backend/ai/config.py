"""
AI模块配置文件
"""
import os
from pathlib import Path

# 基础路径
BASE_DIR = Path(__file__).resolve().parent.parent
KNOWLEDGE_DIR = BASE_DIR / "knowledge"
AI_DIR = BASE_DIR / "ai"

# 模型配置
EMBEDDING_MODEL = "sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2"
EMBEDDING_DIMENSION = 384

# 向量存储配置
VECTOR_DB_PATH = AI_DIR / "vector_db"
VECTOR_INDEX_TYPE = "flat"  # flat, ivf, hnsw

# 意图识别配置
INTENT_CONFIDENCE_THRESHOLD = 0.7
INTENT_LABELS = [
    "故障上报",
    "故障查询",
    "召回查询",
    "质保咨询",
    "零部件查询",
    "进度查询",
    "维修服务",
    "纠纷处理",
    "系统使用",
    "其他"
]

# 对话管理配置
MAX_CONVERSATION_ROUNDS = 10
CONVERSATION_TIMEOUT = 1800  # 30分钟（秒）
CONVERSATION_STORAGE = "memory"  # memory, redis

# 区块链数据查询配置
BLOCKCHAIN_QUERY_ENABLED = True
BLOCKCHAIN_API_BASE = "http://localhost:8080"

# 日志配置
LOG_LEVEL = "INFO"
LOG_FILE = AI_DIR / "ai_service.log"

# 资源限制
MAX_MEMORY_USAGE = 0.8  # 最大内存使用率80%
MAX_CONCURRENT_REQUESTS = 10
