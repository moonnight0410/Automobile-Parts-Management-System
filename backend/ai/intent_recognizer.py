"""
意图识别模块
负责识别用户问题的意图
"""
import logging
from typing import Dict, List, Optional, Tuple
from collections import defaultdict
from config import (
    INTENT_CONFIDENCE_THRESHOLD,
    INTENT_LABELS,
    LOG_LEVEL
)

# 配置日志
logging.basicConfig(
    level=getattr(logging, LOG_LEVEL),
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)


class IntentRecognizer:
    """意图识别类"""
    
    def __init__(self):
        """初始化意图识别器"""
        self.intent_keywords = self._build_intent_keywords()
        logger.info("意图识别器初始化完成")
    
    def _build_intent_keywords(self) -> Dict[str, List[str]]:
        """构建意图关键词字典"""
        keywords = {
            "故障上报": [
                "上报", "提交", "报告", "故障", "问题", "异常",
                "怎么上报", "如何上报", "上报流程", "上报材料",
                "故障报告", "故障申报", "故障提交"
            ],
            "故障查询": [
                "查询", "进度", "状态", "处理", "结果",
                "故障进度", "处理进度", "查询进度",
                "故障状态", "处理状态", "故障结果"
            ],
            "召回查询": [
                "召回", "批次", "范围", "通知", "执行",
                "召回信息", "召回范围", "召回通知", "召回流程",
                "召回补贴", "召回更换", "召回执行"
            ],
            "质保咨询": [
                "质保", "保修", "期限", "范围", "政策",
                "质保期", "保修期", "质保范围", "质保政策",
                "延保", "质保申请", "质保查询", "质保状态"
            ],
            "零部件查询": [
                "零部件", "零件", "配件", "信息", "查询",
                "零部件信息", "零件信息", "配件信息",
                "零部件ID", "批次号", "零部件状态"
            ],
            "进度查询": [
                "进度", "状态", "查询", "处理", "完成",
                "查询进度", "处理进度", "完成时间",
                "什么时候完成", "需要多久", "处理时间"
            ],
            "维修服务": [
                "维修", "保养", "更换", "服务", "费用",
                "维修服务", "维修费用", "维修时间",
                "维修记录", "维修质保", "维修流程"
            ],
            "纠纷处理": [
                "纠纷", "投诉", "申诉", "处理", "解决",
                "售后纠纷", "投诉建议", "纠纷处理",
                "提交纠纷", "纠纷流程", "纠纷结果"
            ],
            "系统使用": [
                "系统", "使用", "操作", "登录", "密码",
                "系统使用", "操作指南", "使用教程",
                "登录问题", "密码问题", "系统功能"
            ],
            "其他": [
                "其他", "不知道", "不清楚", "帮助", "客服"
            ]
        }
        return keywords
    
    def recognize(self, query: str) -> Tuple[str, float]:
        """
        识别用户意图
        
        Args:
            query: 用户问题
            
        Returns:
            (意图标签, 置信度)
        """
        try:
            query = query.lower()
            
            # 计算每个意图的匹配分数
            intent_scores = defaultdict(int)
            
            for intent, keywords in self.intent_keywords.items():
                for keyword in keywords:
                    if keyword in query:
                        # 关键词匹配，增加分数
                        intent_scores[intent] += 1
            
            # 如果没有匹配到任何关键词，返回"其他"
            if not intent_scores:
                return "其他", 0.0
            
            # 找到分数最高的意图
            max_intent = max(intent_scores.items(), key=lambda x: x[1])
            intent_label, score = max_intent
            
            # 计算置信度（归一化到0-1）
            max_possible_score = max(1, len(query.split()))
            confidence = min(score / max_possible_score, 1.0)
            
            # 如果置信度低于阈值，返回"其他"
            if confidence < INTENT_CONFIDENCE_THRESHOLD:
                return "其他", confidence
            
            logger.info(f"识别意图: {intent_label}, 置信度: {confidence:.2f}")
            return intent_label, confidence
            
        except Exception as e:
            logger.error(f"意图识别失败: {e}")
            return "其他", 0.0
    
    def recognize_with_details(self, query: str) -> Dict[str, any]:
        """
        识别用户意图并返回详细信息
        
        Args:
            query: 用户问题
            
        Returns:
            包含意图识别详细信息的字典
        """
        intent_label, confidence = self.recognize(query)
        
        return {
            "intent": intent_label,
            "confidence": confidence,
            "query": query,
            "matched_keywords": self._get_matched_keywords(query, intent_label)
        }
    
    def _get_matched_keywords(self, query: str, intent_label: str) -> List[str]:
        """获取匹配的关键词"""
        matched = []
        keywords = self.intent_keywords.get(intent_label, [])
        
        for keyword in keywords:
            if keyword in query.lower():
                matched.append(keyword)
        
        return matched
    
    def get_all_intents(self) -> List[str]:
        """获取所有意图标签"""
        return INTENT_LABELS
    
    def get_intent_keywords(self, intent_label: str) -> List[str]:
        """获取指定意图的关键词"""
        return self.intent_keywords.get(intent_label, [])


# 全局意图识别器实例
_intent_recognizer = None


def get_intent_recognizer() -> IntentRecognizer:
    """获取意图识别器实例（单例模式）"""
    global _intent_recognizer
    if _intent_recognizer is None:
        _intent_recognizer = IntentRecognizer()
    return _intent_recognizer
