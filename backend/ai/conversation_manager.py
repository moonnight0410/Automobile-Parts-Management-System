"""
对话管理模块
负责管理多轮对话和会话状态
"""
import time
import logging
from typing import Dict, List, Optional, Any
from collections import deque
from config import (
    MAX_CONVERSATION_ROUNDS,
    CONVERSATION_TIMEOUT,
    CONVERSATION_STORAGE,
    LOG_LEVEL
)

# 配置日志
logging.basicConfig(
    level=getattr(logging, LOG_LEVEL),
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)


class Conversation:
    """对话类"""
    
    def __init__(self, user_id: str):
        """初始化对话"""
        self.user_id = user_id
        self.session_id = f"{user_id}_{int(time.time())}"
        self.created_at = time.time()
        self.last_active = time.time()
        self.messages = deque(maxlen=MAX_CONVERSATION_ROUNDS)
        self.context = {}
        
        logger.info(f"创建新对话: {self.session_id}")
    
    def add_message(self, role: str, content: str, metadata: Optional[Dict] = None):
        """添加消息到对话"""
        message = {
            "role": role,  # user, assistant, system
            "content": content,
            "timestamp": time.time(),
            "metadata": metadata or {}
        }
        self.messages.append(message)
        self.last_active = time.time()
        
        logger.debug(f"添加消息到对话 {self.session_id}: {role} - {content[:50]}...")
    
    def get_messages(self) -> List[Dict]:
        """获取所有消息"""
        return list(self.messages)
    
    def get_recent_messages(self, n: int = 5) -> List[Dict]:
        """获取最近n条消息"""
        messages = list(self.messages)
        return messages[-n:] if len(messages) > n else messages
    
    def get_context(self, key: str) -> Any:
        """获取上下文"""
        return self.context.get(key)
    
    def set_context(self, key: str, value: Any):
        """设置上下文"""
        self.context[key] = value
        self.last_active = time.time()
    
    def is_expired(self) -> bool:
        """检查对话是否过期"""
        return (time.time() - self.last_active) > CONVERSATION_TIMEOUT
    
    def get_round_count(self) -> int:
        """获取对话轮数"""
        return len(self.messages) // 2  # 每轮对话包含用户和助手各一条消息
    
    def clear(self):
        """清空对话"""
        self.messages.clear()
        self.context.clear()
        self.last_active = time.time()
        logger.info(f"清空对话: {self.session_id}")


class ConversationManager:
    """对话管理器"""
    
    def __init__(self):
        """初始化对话管理器"""
        self.conversations: Dict[str, Conversation] = {}
        self.user_sessions: Dict[str, str] = {}  # user_id -> session_id
        
        logger.info("对话管理器初始化完成")
    
    def create_conversation(self, user_id: str) -> Conversation:
        """创建新对话"""
        conversation = Conversation(user_id)
        self.conversations[conversation.session_id] = conversation
        self.user_sessions[user_id] = conversation.session_id
        
        return conversation
    
    def get_conversation(self, session_id: str) -> Optional[Conversation]:
        """获取对话"""
        return self.conversations.get(session_id)
    
    def get_user_conversation(self, user_id: str) -> Optional[Conversation]:
        """获取用户的对话"""
        session_id = self.user_sessions.get(user_id)
        if session_id:
            return self.get_conversation(session_id)
        return None
    
    def get_or_create_conversation(self, user_id: str) -> Conversation:
        """获取或创建用户的对话"""
        conversation = self.get_user_conversation(user_id)
        
        if conversation is None or conversation.is_expired():
            # 创建新对话
            conversation = self.create_conversation(user_id)
        else:
            # 更新最后活跃时间
            conversation.last_active = time.time()
        
        return conversation
    
    def add_message(self, user_id: str, role: str, content: str, metadata: Optional[Dict] = None):
        """添加消息到用户的对话"""
        conversation = self.get_or_create_conversation(user_id)
        conversation.add_message(role, content, metadata)
        return conversation
    
    def get_conversation_history(self, user_id: str, n: int = 5) -> List[Dict]:
        """获取用户的对话历史"""
        conversation = self.get_user_conversation(user_id)
        if conversation:
            return conversation.get_recent_messages(n)
        return []
    
    def clear_conversation(self, user_id: str):
        """清空用户的对话"""
        conversation = self.get_user_conversation(user_id)
        if conversation:
            conversation.clear()
    
    def delete_conversation(self, user_id: str):
        """删除用户的对话"""
        session_id = self.user_sessions.get(user_id)
        if session_id:
            del self.conversations[session_id]
            del self.user_sessions[user_id]
            logger.info(f"删除对话: {session_id}")
    
    def cleanup_expired_conversations(self):
        """清理过期对话"""
        expired_sessions = []
        
        for session_id, conversation in self.conversations.items():
            if conversation.is_expired():
                expired_sessions.append(session_id)
        
        for session_id in expired_sessions:
            conversation = self.conversations[session_id]
            user_id = conversation.user_id
            self.delete_conversation(user_id)
        
        if expired_sessions:
            logger.info(f"清理了 {len(expired_sessions)} 个过期对话")
    
    def get_conversation_count(self) -> int:
        """获取对话数量"""
        return len(self.conversations)
    
    def get_user_count(self) -> int:
        """获取用户数量"""
        return len(self.user_sessions)


# 全局对话管理器实例
_conversation_manager = None


def get_conversation_manager() -> ConversationManager:
    """获取对话管理器实例（单例模式）"""
    global _conversation_manager
    if _conversation_manager is None:
        _conversation_manager = ConversationManager()
    return _conversation_manager
