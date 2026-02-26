"""
知识库管理模块
负责加载、索引和查询知识库文档
"""
import os
import json
import logging
from pathlib import Path
from typing import List, Dict, Any, Optional
from sentence_transformers import SentenceTransformer
import faiss
import numpy as np
from config import (
    KNOWLEDGE_DIR,
    EMBEDDING_MODEL,
    EMBEDDING_DIMENSION,
    VECTOR_DB_PATH,
    LOG_LEVEL
)

# 配置日志
logging.basicConfig(
    level=getattr(logging, LOG_LEVEL),
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)


class KnowledgeBase:
    """知识库管理类"""
    
    def __init__(self):
        """初始化知识库"""
        self.embedding_model = None
        self.documents = []
        self.embeddings = None
        self.index = None
        self.faq_data = []
        
        # 创建向量数据库目录
        VECTOR_DB_PATH.mkdir(parents=True, exist_ok=True)
        
        # 加载模型
        self._load_model()
        
        # 加载知识库
        self._load_knowledge_base()
        
    def _load_model(self):
        """加载嵌入模型"""
        try:
            logger.info(f"加载嵌入模型: {EMBEDDING_MODEL}")
            self.embedding_model = SentenceTransformer(EMBEDDING_MODEL)
            logger.info("嵌入模型加载成功")
        except Exception as e:
            logger.error(f"加载嵌入模型失败: {e}")
            raise
    
    def _load_knowledge_base(self):
        """加载知识库"""
        try:
            logger.info("开始加载知识库...")
            
            # 加载Markdown文档
            self._load_markdown_documents()
            
            # 加载FAQ
            self._load_faq()
            
            # 构建向量索引
            self._build_index()
            
            logger.info(f"知识库加载完成，共加载 {len(self.documents)} 个文档片段，{len(self.faq_data)} 条FAQ")
        except Exception as e:
            logger.error(f"加载知识库失败: {e}")
            raise
    
    def _load_markdown_documents(self):
        """加载Markdown文档"""
        try:
            logger.info("加载Markdown文档...")
            
            markdown_files = list(KNOWLEDGE_DIR.glob("*.md"))
            logger.info(f"找到 {len(markdown_files)} 个Markdown文件")
            
            for md_file in markdown_files:
                try:
                    # 读取文件内容
                    with open(md_file, 'r', encoding='utf-8') as f:
                        content = f.read()
                    
                    # 按标题分割文档
                    sections = self._split_markdown(content, md_file.name)
                    
                    # 添加到文档列表
                    self.documents.extend(sections)
                    
                except Exception as e:
                    logger.error(f"加载文件 {md_file} 失败: {e}")
                    continue
            
            logger.info(f"Markdown文档加载完成，共 {len(self.documents)} 个文档片段")
        except Exception as e:
            logger.error(f"加载Markdown文档失败: {e}")
            raise
    
    def _split_markdown(self, content: str, filename: str) -> List[Dict[str, Any]]:
        """分割Markdown文档为多个片段"""
        sections = []
        lines = content.split('\n')
        
        current_section = ""
        current_title = "概述"
        
        for line in lines:
            # 检测标题
            if line.startswith('#'):
                # 保存当前片段
                if current_section.strip():
                    sections.append({
                        'content': current_section.strip(),
                        'title': current_title,
                        'source': filename,
                        'type': 'document'
                    })
                
                # 开始新片段
                current_title = line.lstrip('#').strip()
                current_section = line + '\n'
            else:
                current_section += line + '\n'
        
        # 保存最后一个片段
        if current_section.strip():
            sections.append({
                'content': current_section.strip(),
                'title': current_title,
                'source': filename,
                'type': 'document'
            })
        
        return sections
    
    def _load_faq(self):
        """加载FAQ"""
        try:
            logger.info("加载FAQ...")
            
            faq_file = KNOWLEDGE_DIR / "常见问题FAQ.json"
            
            if faq_file.exists():
                with open(faq_file, 'r', encoding='utf-8') as f:
                    data = json.load(f)
                
                self.faq_data = data.get('faq', [])
                logger.info(f"FAQ加载完成，共 {len(self.faq_data)} 条")
            else:
                logger.warning(f"FAQ文件不存在: {faq_file}")
                
        except Exception as e:
            logger.error(f"加载FAQ失败: {e}")
            raise
    
    def _build_index(self):
        """构建向量索引"""
        try:
            logger.info("构建向量索引...")
            
            # 收集所有文本
            texts = []
            
            # 添加文档文本
            for doc in self.documents:
                texts.append(doc['content'])
            
            # 添加FAQ问答对
            for faq in self.faq_data:
                question = faq.get('question', '')
                answer = faq.get('answer', '')
                texts.append(f"问题: {question}\n答案: {answer}")
            
            # 生成嵌入向量
            logger.info(f"生成 {len(texts)} 个文本的嵌入向量...")
            self.embeddings = self.embedding_model.encode(
                texts,
                show_progress_bar=True,
                convert_to_numpy=True
            )
            
            # 创建FAISS索引
            self.index = faiss.IndexFlatL2(EMBEDDING_DIMENSION)
            self.index.add(self.embeddings.astype('float32'))
            
            logger.info(f"向量索引构建完成，索引大小: {self.index.ntotal}")
            
        except Exception as e:
            logger.error(f"构建向量索引失败: {e}")
            raise
    
    def search(self, query: str, top_k: int = 5) -> List[Dict[str, Any]]:
        """搜索知识库"""
        try:
            # 生成查询向量
            query_embedding = self.embedding_model.encode([query])
            
            # 搜索
            distances, indices = self.index.search(query_embedding.astype('float32'), top_k)
            
            # 构建结果
            results = []
            for i, (distance, idx) in enumerate(zip(distances[0], indices[0])):
                if idx < len(self.documents):
                    result = self.documents[idx].copy()
                    result['score'] = float(1 / (1 + distance))  # 转换为相似度分数
                    result['distance'] = float(distance)
                    results.append(result)
            
            return results
            
        except Exception as e:
            logger.error(f"搜索知识库失败: {e}")
            return []
    
    def search_faq(self, query: str, top_k: int = 3) -> List[Dict[str, Any]]:
        """搜索FAQ"""
        try:
            # 构建FAQ文本列表
            faq_texts = []
            for faq in self.faq_data:
                question = faq.get('question', '')
                answer = faq.get('answer', '')
                faq_texts.append(f"问题: {question}\n答案: {answer}")
            
            # 生成查询向量
            query_embedding = self.embedding_model.encode([query])
            
            # 生成FAQ嵌入向量
            faq_embeddings = self.embedding_model.encode(faq_texts)
            
            # 计算相似度
            similarities = np.dot(query_embedding, faq_embeddings.T)[0]
            
            # 获取top_k结果
            top_indices = np.argsort(similarities)[-top_k:][::-1]
            
            # 构建结果
            results = []
            for idx in top_indices:
                result = self.faq_data[idx].copy()
                result['score'] = float(similarities[idx])
                results.append(result)
            
            return results
            
        except Exception as e:
            logger.error(f"搜索FAQ失败: {e}")
            return []
    
    def get_document_by_id(self, doc_id: str) -> Optional[Dict[str, Any]]:
        """根据ID获取文档"""
        for doc in self.documents:
            if doc.get('id') == doc_id:
                return doc
        return None
    
    def get_faq_by_id(self, faq_id: str) -> Optional[Dict[str, Any]]:
        """根据ID获取FAQ"""
        for faq in self.faq_data:
            if faq.get('id') == faq_id:
                return faq
        return None


# 全局知识库实例
_knowledge_base = None


def get_knowledge_base() -> KnowledgeBase:
    """获取知识库实例（单例模式）"""
    global _knowledge_base
    if _knowledge_base is None:
        _knowledge_base = KnowledgeBase()
    return _knowledge_base
