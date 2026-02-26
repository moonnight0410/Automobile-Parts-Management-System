"""
问答引擎模块
负责生成回答和处理用户问题
"""
import re
import logging
from typing import Dict, List, Optional, Any
from knowledge_base import get_knowledge_base
from intent_recognizer import get_intent_recognizer
from conversation_manager import get_conversation_manager
from blockchain_query import get_blockchain_query
from config import LOG_LEVEL

# 配置日志
logging.basicConfig(
    level=getattr(logging, LOG_LEVEL),
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)


class QAEngine:
    """问答引擎类"""
    
    def __init__(self):
        """初始化问答引擎"""
        self.knowledge_base = get_knowledge_base()
        self.intent_recognizer = get_intent_recognizer()
        self.conversation_manager = get_conversation_manager()
        self.blockchain_query = get_blockchain_query()
        
        logger.info("问答引擎初始化完成")
    
    def answer(self, query: str, user_id: str) -> Dict[str, Any]:
        """
        回答用户问题
        
        Args:
            query: 用户问题
            user_id: 用户ID
            
        Returns:
            包含回答和相关信息的字典
        """
        try:
            # 添加用户消息到对话
            self.conversation_manager.add_message(user_id, "user", query)
            
            # 识别意图
            intent_result = self.intent_recognizer.recognize_with_details(query)
            intent = intent_result["intent"]
            confidence = intent_result["confidence"]
            
            logger.info(f"用户 {user_id} 问题: {query}, 意图: {intent}, 置信度: {confidence}")
            
            # 根据意图生成回答
            if intent == "故障上报":
                answer = self._answer_fault_report(query)
            elif intent == "故障查询":
                answer = self._answer_fault_query(query)
            elif intent == "召回查询":
                answer = self._answer_recall_query(query)
            elif intent == "质保咨询":
                answer = self._answer_warranty(query)
            elif intent == "零部件查询":
                answer = self._answer_part_query(query)
            elif intent == "进度查询":
                answer = self._answer_progress_query(query)
            elif intent == "维修服务":
                answer = self._answer_maintenance(query)
            elif intent == "纠纷处理":
                answer = self._answer_dispute(query)
            elif intent == "系统使用":
                answer = self._answer_system_usage(query)
            else:
                answer = self._answer_general(query)
            
            # 添加助手消息到对话
            self.conversation_manager.add_message(
                user_id,
                "assistant",
                answer["answer"],
                {
                    "intent": intent,
                    "confidence": confidence,
                    "related_actions": answer.get("related_actions", [])
                }
            )
            
            # 构建返回结果
            result = {
                "qaid": f"QA-{int(__import__('time').time() * 1000)}",
                "question": query,
                "intent": intent,
                "confidence": confidence,
                "answer": answer["answer"],
                "related_actions": answer.get("related_actions", []),
                "related_faqs": answer.get("related_faqs", [])
            }
            
            return result
            
        except Exception as e:
            logger.error(f"回答问题失败: {e}")
            return {
                "qaid": f"QA-{int(__import__('time').time() * 1000)}",
                "question": query,
                "intent": "其他",
                "confidence": 0.0,
                "answer": "抱歉，我暂时无法回答这个问题。请稍后再试或联系人工客服。",
                "related_actions": [
                    {"label": "联系人工客服", "action": "contact_support"}
                ],
                "related_faqs": []
            }
    
    def _answer_fault_report(self, query: str) -> Dict[str, Any]:
        """回答故障上报相关问题"""
        # 搜索相关知识
        results = self.knowledge_base.search(query, top_k=3)
        faq_results = self.knowledge_base.search_faq(query, top_k=2)
        
        # 构建回答
        answer_parts = []
        
        if results:
            answer_parts.append(results[0]["content"])
        
        if faq_results:
            answer_parts.append("\n\n**常见问题：**")
            for faq in faq_results:
                answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
        
        answer = "\n".join(answer_parts) if answer_parts else "关于故障上报，您可以通过以下步骤操作：\n\n1. 进入\"售后管理 > 故障报告\"页面\n2. 点击\"新建故障报告\"按钮\n3. 填写零部件ID、VIN码、故障描述等信息\n4. 上传故障照片（1-5张）\n5. 点击\"提交\"按钮"
        
        return {
            "answer": answer,
            "related_actions": [
                {"label": "立即上报故障", "action": "navigate", "params": {"path": "/aftersale/fault"}}
            ],
            "related_faqs": [faq.get("id") for faq in faq_results]
        }
    
    def _answer_fault_query(self, query: str) -> Dict[str, Any]:
        """回答故障查询相关问题"""
        # 提取故障ID
        fault_id = self._extract_fault_id(query)
        
        if fault_id:
            # 查询区块链数据
            fault_data = self.blockchain_query.query_fault_progress(fault_id)
            
            if fault_data:
                # 格式化故障进度信息
                answer = self.blockchain_query.format_fault_progress(fault_data)
            else:
                # 查询失败，提供通用回答
                answer = f"无法查询到故障 {fault_id} 的信息，可能原因：\n\n1. 故障ID不存在\n2. 区块链数据查询失败\n\n建议：\n1. 检查故障ID是否正确\n2. 在系统中输入故障ID查询\n3. 或联系客服协助查询"
        else:
            # 搜索相关知识
            results = self.knowledge_base.search(query, top_k=3)
            faq_results = self.knowledge_base.search_faq(query, top_k=2)
            
            answer_parts = []
            
            if results:
                answer_parts.append(results[0]["content"])
            
            if faq_results:
                answer_parts.append("\n\n**常见问题：**")
                for faq in faq_results:
                    answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
            
            answer = "\n".join(answer_parts) if answer_parts else "您可以通过以下方式查询故障处理进度：\n\n1. 在故障报告列表中查看\n2. 输入故障ID精确查询\n3. 通过AI助手查询（提供故障ID）"
        
        return {
            "answer": answer,
            "related_actions": [
                {"label": "查询故障进度", "action": "query", "params": {"type": "fault"}}
            ],
            "related_faqs": [faq.get("id") for faq in faq_results] if 'faq_results' in locals() else []
        }
    
    def _answer_recall_query(self, query: str) -> Dict[str, Any]:
        """回答召回查询相关问题"""
        # 提取批次号、零部件ID或VIN码
        batch_no = self._extract_batch_no(query)
        part_id = self._extract_part_id(query)
        vin = self._extract_vin(query)
        
        if batch_no or part_id or vin:
            # 查询区块链数据
            recall_data_list = self.blockchain_query.query_recall_info(batch_no, part_id, vin)
            
            if recall_data_list and len(recall_data_list) > 0:
                # 格式化召回信息
                answer_parts = []
                for recall_data in recall_data_list:
                    answer_parts.append(self.blockchain_query.format_recall_info(recall_data))
                    answer_parts.append("\n---\n")
                answer = "\n".join(answer_parts)
            else:
                # 查询失败，提供通用回答
                query_info = batch_no or part_id or vin
                answer = f"未查询到 {query_info} 的召回信息，可能原因：\n\n1. 该零部件不在召回范围内\n2. 区块链数据查询失败\n\n建议：\n1. 检查批次号、零部件ID或VIN码是否正确\n2. 在系统中输入相关信息查询\n3. 或联系客服协助查询"
        else:
            # 搜索相关知识
            results = self.knowledge_base.search(query, top_k=3)
            faq_results = self.knowledge_base.search_faq(query, top_k=2)
            
            answer_parts = []
            
            if results:
                answer_parts.append(results[0]["content"])
            
            if faq_results:
                answer_parts.append("\n\n**常见问题：**")
                for faq in faq_results:
                    answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
            
            answer = "\n".join(answer_parts) if answer_parts else "您可以通过以下方式查询召回信息：\n\n1. 输入批次号查询\n2. 输入零部件ID查询\n3. 输入VIN码查询\n\n请提供批次号、零部件ID或VIN码以获取准确的召回信息。"
        
        return {
            "answer": answer,
            "related_actions": [
                {"label": "查询召回信息", "action": "query", "params": {"type": "recall"}}
            ],
            "related_faqs": [faq.get("id") for faq in faq_results] if 'faq_results' in locals() else []
        }
    
    def _answer_warranty(self, query: str) -> Dict[str, Any]:
        """回答质保相关问题"""
        # 搜索相关知识
        results = self.knowledge_base.search(query, top_k=3)
        faq_results = self.knowledge_base.search_faq(query, top_k=2)
        
        answer_parts = []
        
        if results:
            answer_parts.append(results[0]["content"])
        
        if faq_results:
            answer_parts.append("\n\n**常见问题：**")
            for faq in faq_results:
                answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
        
        answer = "\n".join(answer_parts) if answer_parts else "关于质保，标准质保期为12个月或20,000公里（以先到为准）。您可以通过零部件ID或VIN码查询质保状态。"
        
        return {
            "answer": answer,
            "related_actions": [],
            "related_faqs": [faq.get("id") for faq in faq_results]
        }
    
    def _answer_part_query(self, query: str) -> Dict[str, Any]:
        """回答零部件查询相关问题"""
        # 提取零部件ID或批次号
        part_id = self._extract_part_id(query)
        batch_no = self._extract_batch_no(query)
        
        if part_id:
            # 查询区块链数据
            part_data = self.blockchain_query.query_part_info(part_id=part_id)
            
            if part_data:
                # 格式化零部件信息
                answer = self.blockchain_query.format_part_info(part_data)
            else:
                # 查询失败，提供通用回答
                answer = f"无法查询到零部件 {part_id} 的信息，可能原因：\n\n1. 零部件ID不存在\n2. 区块链数据查询失败\n\n建议：\n1. 检查零部件ID是否正确\n2. 在系统中输入零部件ID查询\n3. 或联系客服协助查询"
        elif batch_no:
            # 查询区块链数据
            part_data = self.blockchain_query.query_part_info(batch_no=batch_no)
            
            if part_data:
                # 格式化零部件信息
                answer = self.blockchain_query.format_part_info(part_data)
            else:
                # 查询失败，提供通用回答
                answer = f"无法查询到批次 {batch_no} 的零部件信息，可能原因：\n\n1. 批次号不存在\n2. 区块链数据查询失败\n\n建议：\n1. 检查批次号是否正确\n2. 在系统中输入批次号查询\n3. 或联系客服协助查询"
        else:
            # 搜索相关知识
            results = self.knowledge_base.search(query, top_k=3)
            faq_results = self.knowledge_base.search_faq(query, top_k=2)
            
            answer_parts = []
            
            if results:
                answer_parts.append(results[0]["content"])
            
            if faq_results:
                answer_parts.append("\n\n**常见问题：**")
                for faq in faq_results:
                    answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
            
            answer = "\n".join(answer_parts) if answer_parts else "您可以通过以下方式查询零部件信息：\n\n1. 输入零部件ID查询\n2. 输入批次号查询\n\n请提供零部件ID或批次号以获取准确的零部件信息。"
        
        return {
            "answer": answer,
            "related_actions": [
                {"label": "查询零部件信息", "action": "query", "params": {"type": "part"}}
            ],
            "related_faqs": [faq.get("id") for faq in faq_results] if 'faq_results' in locals() else []
        }
    
    def _answer_progress_query(self, query: str) -> Dict[str, Any]:
        """回答进度查询相关问题"""
        # 搜索相关知识
        results = self.knowledge_base.search(query, top_k=3)
        faq_results = self.knowledge_base.search_faq(query, top_k=2)
        
        answer_parts = []
        
        if results:
            answer_parts.append(results[0]["content"])
        
        if faq_results:
            answer_parts.append("\n\n**常见问题：**")
            for faq in faq_results:
                answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
        
        answer = "\n".join(answer_parts) if answer_parts else "故障处理时间根据故障严重程度而定：\n\n- 轻微故障：7-10个工作日\n- 一般故障：10-15个工作日\n- 严重故障：5-7个工作日（加急）\n- 危急故障：3-5个工作日（紧急）"
        
        return {
            "answer": answer,
            "related_actions": [],
            "related_faqs": [faq.get("id") for faq in faq_results]
        }
    
    def _answer_maintenance(self, query: str) -> Dict[str, Any]:
        """回答维修服务相关问题"""
        # 搜索相关知识
        results = self.knowledge_base.search(query, top_k=3)
        faq_results = self.knowledge_base.search_faq(query, top_k=2)
        
        answer_parts = []
        
        if results:
            answer_parts.append(results[0]["content"])
        
        if faq_results:
            answer_parts.append("\n\n**常见问题：**")
            for faq in faq_results:
                answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
        
        answer = "\n".join(answer_parts) if answer_parts else "维修时间根据维修类型而定：\n\n- 简单维修：1-2小时\n- 更换零部件：2-4小时\n- 复杂维修：4-8小时\n\n维修完成后建议试车确认，并提交满意度评价。"
        
        return {
            "answer": answer,
            "related_actions": [],
            "related_faqs": [faq.get("id") for faq in faq_results]
        }
    
    def _answer_dispute(self, query: str) -> Dict[str, Any]:
        """回答纠纷处理相关问题"""
        # 搜索相关知识
        results = self.knowledge_base.search(query, top_k=3)
        faq_results = self.knowledge_base.search_faq(query, top_k=2)
        
        answer_parts = []
        
        if results:
            answer_parts.append(results[0]["content"])
        
        if faq_results:
            answer_parts.append("\n\n**常见问题：**")
            for faq in faq_results:
                answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
        
        answer = "\n".join(answer_parts) if answer_parts else "纠纷处理时间根据纠纷类型而定：\n\n- 简单纠纷：5-7个工作日\n- 一般纠纷：7-10个工作日\n- 复杂纠纷：10-15个工作日\n\n纠纷处理完全免费，您可以在线提交纠纷申请。"
        
        return {
            "answer": answer,
            "related_actions": [],
            "related_faqs": [faq.get("id") for faq in faq_results]
        }
    
    def _answer_system_usage(self, query: str) -> Dict[str, Any]:
        """回答系统使用相关问题"""
        # 搜索相关知识
        results = self.knowledge_base.search(query, top_k=3)
        faq_results = self.knowledge_base.search_faq(query, top_k=2)
        
        answer_parts = []
        
        if results:
            answer_parts.append(results[0]["content"])
        
        if faq_results:
            answer_parts.append("\n\n**常见问题：**")
            for faq in faq_results:
                answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
        
        answer = "\n".join(answer_parts) if answer_parts else "系统支持Chrome、Firefox、Safari、Edge浏览器，不支持IE浏览器。系统提供7×24小时在线客服，如有问题可随时咨询。"
        
        return {
            "answer": answer,
            "related_actions": [
                {"label": "查看系统帮助", "action": "navigate", "params": {"path": "/help"}}
            ],
            "related_faqs": [faq.get("id") for faq in faq_results]
        }
    
    def _answer_general(self, query: str) -> Dict[str, Any]:
        """回答一般性问题"""
        # 搜索相关知识
        results = self.knowledge_base.search(query, top_k=3)
        faq_results = self.knowledge_base.search_faq(query, top_k=3)
        
        answer_parts = []
        
        if results:
            answer_parts.append(results[0]["content"])
        
        if faq_results:
            answer_parts.append("\n\n**相关问答：**")
            for faq in faq_results:
                answer_parts.append(f"\n**Q:** {faq['question']}\n**A:** {faq['answer']}")
        
        answer = "\n".join(answer_parts) if answer_parts else "抱歉，我没有找到相关的信息。您可以：\n\n1. 尝试重新描述您的问题\n2. 查看常见问题FAQ\n3. 联系人工客服"
        
        return {
            "answer": answer,
            "related_actions": [
                {"label": "联系人工客服", "action": "contact_support"}
            ],
            "related_faqs": [faq.get("id") for faq in faq_results]
        }
    
    def _extract_fault_id(self, text: str) -> Optional[str]:
        """提取故障ID"""
        pattern = r'FAULT[-_]?\d{8}[-_]?\d{3}'
        match = re.search(pattern, text, re.IGNORECASE)
        return match.group(0) if match else None
    
    def _extract_batch_no(self, text: str) -> Optional[str]:
        """提取批次号"""
        pattern = r'BATCH[-_]?\d{4}[-_]?\d{3}'
        match = re.search(pattern, text, re.IGNORECASE)
        return match.group(0) if match else None
    
    def _extract_part_id(self, text: str) -> Optional[str]:
        """提取零部件ID"""
        pattern = r'[A-Z]{3}[-_]?[A-Z]+[-_]?\d{6}'
        match = re.search(pattern, text)
        return match.group(0) if match else None
    
    def _extract_vin(self, text: str) -> Optional[str]:
        """提取VIN码"""
        pattern = r'[A-HJ-NPR-Z0-9]{17}'
        match = re.search(pattern, text)
        return match.group(0) if match else None


# 全局问答引擎实例
_qa_engine = None


def get_qa_engine() -> QAEngine:
    """获取问答引擎实例（单例模式）"""
    global _qa_engine
    if _qa_engine is None:
        _qa_engine = QAEngine()
    return _qa_engine
