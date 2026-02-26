"""
区块链数据查询模块
负责查询区块链上的故障进度、召回信息和零部件信息
"""
import logging
import requests
from typing import Dict, List, Optional, Any
from config import (
    BLOCKCHAIN_QUERY_ENABLED,
    BLOCKCHAIN_API_BASE,
    LOG_LEVEL
)

# 配置日志
logging.basicConfig(
    level=getattr(logging, LOG_LEVEL),
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)


class BlockchainQuery:
    """区块链数据查询类"""
    
    def __init__(self):
        """初始化区块链查询器"""
        self.enabled = BLOCKCHAIN_QUERY_ENABLED
        self.api_base = BLOCKCHAIN_API_BASE
        
        if self.enabled:
            logger.info("区块链数据查询已启用")
        else:
            logger.info("区块链数据查询未启用")
    
    def query_fault_progress(self, fault_id: str) -> Optional[Dict[str, Any]]:
        """
        查询故障进度
        
        Args:
            fault_id: 故障ID
            
        Returns:
            故障进度信息
        """
        if not self.enabled:
            logger.warning("区块链数据查询未启用")
            return None
        
        try:
            url = f"{self.api_base}/api/faults/{fault_id}"
            response = requests.get(url, timeout=10)
            
            if response.status_code == 200:
                data = response.json()
                if data.get('success'):
                    return data.get('data')
                else:
                    logger.warning(f"查询故障进度失败: {data.get('message')}")
                    return None
            else:
                logger.warning(f"查询故障进度失败，状态码: {response.status_code}")
                return None
                
        except Exception as e:
            logger.error(f"查询故障进度异常: {e}")
            return None
    
    def query_recall_info(self, batch_no: str = None, part_id: str = None, vin: str = None) -> Optional[List[Dict[str, Any]]]:
        """
        查询召回信息
        
        Args:
            batch_no: 批次号
            part_id: 零部件ID
            vin: VIN码
            
        Returns:
            召回信息列表
        """
        if not self.enabled:
            logger.warning("区块链数据查询未启用")
            return None
        
        try:
            params = {}
            if batch_no:
                params['batchNo'] = batch_no
            if part_id:
                params['partID'] = part_id
            if vin:
                params['vin'] = vin
            
            if not params:
                logger.warning("查询召回信息需要提供至少一个参数")
                return None
            
            url = f"{self.api_base}/api/recalls"
            response = requests.get(url, params=params, timeout=10)
            
            if response.status_code == 200:
                data = response.json()
                if data.get('success'):
                    return data.get('data', [])
                else:
                    logger.warning(f"查询召回信息失败: {data.get('message')}")
                    return None
            else:
                logger.warning(f"查询召回信息失败，状态码: {response.status_code}")
                return None
                
        except Exception as e:
            logger.error(f"查询召回信息异常: {e}")
            return None
    
    def query_part_info(self, part_id: str = None, batch_no: str = None) -> Optional[Dict[str, Any]]:
        """
        查询零部件信息
        
        Args:
            part_id: 零部件ID
            batch_no: 批次号
            
        Returns:
            零部件信息
        """
        if not self.enabled:
            logger.warning("区块链数据查询未启用")
            return None
        
        try:
            if part_id:
                url = f"{self.api_base}/api/parts/{part_id}"
                response = requests.get(url, timeout=10)
            elif batch_no:
                url = f"{self.api_base}/api/parts"
                response = requests.get(url, params={'batchNo': batch_no}, timeout=10)
            else:
                logger.warning("查询零部件信息需要提供part_id或batch_no")
                return None
            
            if response.status_code == 200:
                data = response.json()
                if data.get('success'):
                    return data.get('data')
                else:
                    logger.warning(f"查询零部件信息失败: {data.get('message')}")
                    return None
            else:
                logger.warning(f"查询零部件信息失败，状态码: {response.status_code}")
                return None
                
        except Exception as e:
            logger.error(f"查询零部件信息异常: {e}")
            return None
    
    def format_fault_progress(self, fault_data: Dict[str, Any]) -> str:
        """
        格式化故障进度信息
        
        Args:
            fault_data: 故障数据
            
        Returns:
            格式化后的故障进度信息
        """
        try:
            fault_id = fault_data.get('faultID', '')
            status = fault_data.get('status', '')
            created_at = fault_data.get('createdAt', '')
            updated_at = fault_data.get('updatedAt', '')
            
            status_map = {
                'pending': '待审核',
                'approved': '审核通过',
                'analyzing': '技术分析中',
                'waiting': '等待处理',
                'processing': '处理中',
                'completed': '已完成',
                'rejected': '已驳回',
                'closed': '已关闭'
            }
            
            status_text = status_map.get(status, status)
            
            result = f"**故障ID**: {fault_id}\n"
            result += f"**当前状态**: {status_text}\n"
            result += f"**上报时间**: {created_at}\n"
            result += f"**最后更新**: {updated_at}\n"
            
            # 添加处理记录
            records = fault_data.get('records', [])
            if records:
                result += "\n**处理记录**:\n"
                for record in records:
                    record_time = record.get('time', '')
                    record_action = record.get('action', '')
                    record_note = record.get('note', '')
                    result += f"- {record_time}: {record_action}"
                    if record_note:
                        result += f" ({record_note})"
                    result += "\n"
            
            return result
            
        except Exception as e:
            logger.error(f"格式化故障进度信息失败: {e}")
            return "格式化故障进度信息失败"
    
    def format_recall_info(self, recall_data: Dict[str, Any]) -> str:
        """
        格式化召回信息
        
        Args:
            recall_data: 召回数据
            
        Returns:
            格式化后的召回信息
        """
        try:
            recall_id = recall_data.get('recallID', '')
            batch_no = recall_data.get('batchNo', '')
            part_id = recall_data.get('partID', '')
            status = recall_data.get('status', '')
            reason = recall_data.get('reason', '')
            measure = recall_data.get('measure', '')
            start_time = recall_data.get('startTime', '')
            end_time = recall_data.get('endTime', '')
            
            status_map = {
                'normal': '正常',
                'recalling': '召回中',
                'pending': '待召回',
                'completed': '已完成',
                'cancelled': '已取消'
            }
            
            status_text = status_map.get(status, status)
            
            result = f"**召回编号**: {recall_id}\n"
            result += f"**批次号**: {batch_no}\n"
            if part_id:
                result += f"**零部件ID**: {part_id}\n"
            result += f"**召回状态**: {status_text}\n"
            result += f"**召回原因**: {reason}\n"
            result += f"**召回措施**: {measure}\n"
            result += f"**召回时间**: {start_time} - {end_time}\n"
            
            return result
            
        except Exception as e:
            logger.error(f"格式化召回信息失败: {e}")
            return "格式化召回信息失败"
    
    def format_part_info(self, part_data: Dict[str, Any]) -> str:
        """
        格式化零部件信息
        
        Args:
            part_data: 零部件数据
            
        Returns:
            格式化后的零部件信息
        """
        try:
            part_id = part_data.get('partID', '')
            part_name = part_data.get('partName', '')
            part_type = part_data.get('partType', '')
            batch_no = part_data.get('batchNo', '')
            manufacturer = part_data.get('manufacturer', '')
            production_date = part_data.get('productionDate', '')
            status = part_data.get('status', '')
            
            status_map = {
                'in_stock': '在库',
                'in_transit': '在途',
                'installed': '已安装',
                'repaired': '已维修',
                'scrapped': '已报废'
            }
            
            status_text = status_map.get(status, status)
            
            # 质保信息
            warranty_info = part_data.get('warranty', {})
            warranty_status = warranty_info.get('status', '')
            warranty_start = warranty_info.get('startDate', '')
            warranty_end = warranty_info.get('endDate', '')
            
            warranty_status_map = {
                'in_warranty': '在质保期内',
                'out_of_warranty': '已过质保期'
            }
            
            warranty_status_text = warranty_status_map.get(warranty_status, warranty_status)
            
            result = f"**零部件ID**: {part_id}\n"
            result += f"**零部件名称**: {part_name}\n"
            result += f"**零部件类型**: {part_type}\n"
            result += f"**批次号**: {batch_no}\n"
            result += f"**制造商**: {manufacturer}\n"
            result += f"**生产日期**: {production_date}\n"
            result += f"**当前状态**: {status_text}\n"
            result += f"**质保状态**: {warranty_status_text}\n"
            if warranty_start and warranty_end:
                result += f"**质保期**: {warranty_start} - {warranty_end}\n"
            
            return result
            
        except Exception as e:
            logger.error(f"格式化零部件信息失败: {e}")
            return "格式化零部件信息失败"


# 全局区块链查询器实例
_blockchain_query = None


def get_blockchain_query() -> BlockchainQuery:
    """获取区块链查询器实例（单例模式）"""
    global _blockchain_query
    if _blockchain_query is None:
        _blockchain_query = BlockchainQuery()
    return _blockchain_query
