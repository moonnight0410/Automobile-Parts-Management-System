"""
AI服务Flask API
提供RESTful API接口供Go后端调用
"""
from flask import Flask, request, jsonify
from flask_cors import CORS
import logging
from typing import Dict, Any
from qa_engine import get_qa_engine
from conversation_manager import get_conversation_manager
from config import LOG_LEVEL, MAX_CONCURRENT_REQUESTS
import threading

# 配置日志
logging.basicConfig(
    level=getattr(logging, LOG_LEVEL),
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

# 创建Flask应用
app = Flask(__name__)
CORS(app)

# 请求计数器（用于限制并发请求）
request_semaphore = threading.Semaphore(MAX_CONCURRENT_REQUESTS)


@app.route('/health', methods=['GET'])
def health_check():
    """健康检查接口"""
    return jsonify({
        "status": "healthy",
        "service": "AI Service",
        "version": "1.0.0"
    })


@app.route('/api/ai/question', methods=['POST'])
def ask_question():
    """
    提问接口
    
    请求体：
    {
        "question": "用户问题",
        "userID": "用户ID"
    }
    
    响应：
    {
        "success": true,
        "message": "查询成功",
        "data": {
            "qaid": "QA-1234567890",
            "question": "用户问题",
            "intent": "故障上报",
            "confidence": 0.85,
            "answer": "回答内容",
            "related_actions": [...],
            "related_faqs": [...]
        }
    }
    """
    try:
        # 获取请求参数
        data = request.get_json()
        question = data.get('question', '').strip()
        user_id = data.get('userID', 'anonymous')
        
        # 验证参数
        if not question:
            return jsonify({
                "success": False,
                "message": "问题不能为空"
            }), 400
        
        # 限制并发请求
        if not request_semaphore.acquire(blocking=False):
            return jsonify({
                "success": False,
                "message": "服务器繁忙，请稍后再试"
            }), 503
        
        try:
            # 获取问答引擎
            qa_engine = get_qa_engine()
            
            # 回答问题
            result = qa_engine.answer(question, user_id)
            
            # 返回结果
            return jsonify({
                "success": True,
                "message": "查询成功",
                "data": result
            })
            
        finally:
            # 释放信号量
            request_semaphore.release()
            
    except Exception as e:
        logger.error(f"处理提问失败: {e}")
        return jsonify({
            "success": False,
            "message": f"处理失败: {str(e)}"
        }), 500


@app.route('/api/ai/conversation', methods=['GET'])
def get_conversation():
    """
    获取对话历史接口
    
    请求参数：
    - userID: 用户ID
    - n: 返回最近n条消息（默认5）
    
    响应：
    {
        "success": true,
        "message": "查询成功",
        "data": {
            "messages": [...]
        }
    }
    """
    try:
        # 获取请求参数
        user_id = request.args.get('userID', 'anonymous')
        n = int(request.args.get('n', 5))
        
        # 获取对话管理器
        conversation_manager = get_conversation_manager()
        
        # 获取对话历史
        messages = conversation_manager.get_conversation_history(user_id, n)
        
        # 返回结果
        return jsonify({
            "success": True,
            "message": "查询成功",
            "data": {
                "messages": messages
            }
        })
        
    except Exception as e:
        logger.error(f"获取对话历史失败: {e}")
        return jsonify({
            "success": False,
            "message": f"查询失败: {str(e)}"
        }), 500


@app.route('/api/ai/conversation', methods=['DELETE'])
def clear_conversation():
    """
    清空对话接口
    
    请求参数：
    - userID: 用户ID
    
    响应：
    {
        "success": true,
        "message": "清空成功"
    }
    """
    try:
        # 获取请求参数
        user_id = request.args.get('userID', 'anonymous')
        
        # 获取对话管理器
        conversation_manager = get_conversation_manager()
        
        # 清空对话
        conversation_manager.clear_conversation(user_id)
        
        # 返回结果
        return jsonify({
            "success": True,
            "message": "清空成功"
        })
        
    except Exception as e:
        logger.error(f"清空对话失败: {e}")
        return jsonify({
            "success": False,
            "message": f"清空失败: {str(e)}"
        }), 500


@app.route('/api/ai/stats', methods=['GET'])
def get_stats():
    """
    获取统计信息接口
    
    响应：
    {
        "success": true,
        "message": "查询成功",
        "data": {
            "conversation_count": 10,
            "user_count": 8
        }
    }
    """
    try:
        # 获取对话管理器
        conversation_manager = get_conversation_manager()
        
        # 获取统计信息
        conversation_count = conversation_manager.get_conversation_count()
        user_count = conversation_manager.get_user_count()
        
        # 返回结果
        return jsonify({
            "success": True,
            "message": "查询成功",
            "data": {
                "conversation_count": conversation_count,
                "user_count": user_count
            }
        })
        
    except Exception as e:
        logger.error(f"获取统计信息失败: {e}")
        return jsonify({
            "success": False,
            "message": f"查询失败: {str(e)}"
        }), 500


@app.errorhandler(404)
def not_found(error):
    """404错误处理"""
    return jsonify({
        "success": False,
        "message": "接口不存在"
    }), 404


@app.errorhandler(500)
def internal_error(error):
    """500错误处理"""
    return jsonify({
        "success": False,
        "message": "服务器内部错误"
    }), 500


def main():
    """启动服务"""
    logger.info("启动AI服务...")
    
    # 启动Flask应用
    app.run(
        host='0.0.0.0',
        port=5000,
        debug=False,
        threaded=True
    )


if __name__ == '__main__':
    main()
