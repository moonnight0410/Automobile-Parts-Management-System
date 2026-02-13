<!--
  AI智能售后助手页面
  提供智能问答服务，支持故障上报咨询、进度查询、召回查询等
-->

<template>
  <div class="ai-assistant">
    <!-- 页面标题 -->
    <div class="page-header">
      <h2>智能售后助手</h2>
      <a-tag color="blue">AI Powered</a-tag>
    </div>
    
    <a-row :gutter="16">
      <!-- 左侧对话区 -->
      <a-col :xs="24" :lg="16">
        <a-card :bordered="false" class="chat-card">
          <!-- 对话消息列表 -->
          <div class="message-list" ref="messageListRef">
            <div
              v-for="msg in messages"
              :key="msg.id"
              :class="['message-item', msg.type]"
            >
              <div class="message-avatar">
                <a-avatar v-if="msg.type === 'user'" :style="{ backgroundColor: '#1890ff' }">
                  {{ user?.username?.charAt(0).toUpperCase() }}
                </a-avatar>
                <a-avatar v-else :style="{ backgroundColor: '#52c41a' }">
                  <template #icon><RobotOutlined /></template>
                </a-avatar>
              </div>
              <div class="message-content">
                <div class="message-text">{{ msg.content }}</div>
                <!-- 快捷操作按钮 -->
                <div v-if="msg.actions && msg.actions.length" class="message-actions">
                  <a-button
                    v-for="action in msg.actions"
                    :key="action.label"
                    size="small"
                    type="primary"
                    ghost
                    @click="handleAction(action)"
                  >
                    {{ action.label }}
                  </a-button>
                </div>
                <div class="message-time">{{ msg.time }}</div>
              </div>
            </div>
            
            <!-- 加载中提示 -->
            <div v-if="sending" class="message-item bot">
              <div class="message-avatar">
                <a-avatar :style="{ backgroundColor: '#52c41a' }">
                  <template #icon><RobotOutlined /></template>
                </a-avatar>
              </div>
              <div class="message-content">
                <a-spin size="small" />
                <span style="margin-left: 8px">正在思考中...</span>
              </div>
            </div>
          </div>
          
          <!-- 输入区 -->
          <div class="input-area">
            <a-textarea
              v-model:value="inputMessage"
              placeholder="请输入您的问题，例如：如何上报零部件故障？"
              :auto-size="{ minRows: 2, maxRows: 4 }"
              @pressEnter="handleSend"
            />
            <div class="input-actions">
              <a-space>
                <a-button @click="handleClear">
                  <template #icon><ClearOutlined /></template>
                  清空对话
                </a-button>
                <a-button type="primary" @click="handleSend" :loading="sending">
                  <template #icon><SendOutlined /></template>
                  发送
                </a-button>
              </a-space>
            </div>
          </div>
        </a-card>
      </a-col>
      
      <!-- 右侧快捷问题 -->
      <a-col :xs="24" :lg="8">
        <a-card title="常见问题" :bordered="false" class="faq-card">
          <a-list :data-source="faqList" size="small">
            <template #renderItem="{ item }">
              <a-list-item>
                <a-button type="link" @click="askQuestion(item.question)">
                  {{ item.question }}
                </a-button>
              </a-list-item>
            </template>
          </a-list>
        </a-card>
        
        <a-card title="功能说明" :bordered="false" class="info-card" style="margin-top: 16px">
          <a-descriptions :column="1" size="small">
            <a-descriptions-item label="故障上报咨询">
              了解故障上报流程和所需材料
            </a-descriptions-item>
            <a-descriptions-item label="故障进度查询">
              查询故障处理进度和状态
            </a-descriptions-item>
            <a-descriptions-item label="召回相关查询">
              查询召回范围和处理流程
            </a-descriptions-item>
            <a-descriptions-item label="售后政策咨询">
              了解质保期、维修政策等
            </a-descriptions-item>
          </a-descriptions>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, nextTick, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { useAuthStore } from '../../stores'
import { sendAIQuestion, queryFaultProgressAI, queryRecallInfoAI } from '../../services/ai.service'

// ==================== 类型定义 ====================

interface Message {
  id: string
  type: 'user' | 'bot'
  content: string
  time: string
  actions?: Array<{
    label: string
    action: string
    params?: Record<string, any>
  }>
}

// ==================== 组合式API ====================

const router = useRouter()
const authStore = useAuthStore()

// ==================== 响应式状态 ====================

const messageListRef = ref<HTMLElement>()
const inputMessage = ref('')
const sending = ref(false)

// 消息列表
const messages = ref<Message[]>([
  {
    id: '1',
    type: 'bot',
    content: '您好！我是智能售后助手，可以帮您解答以下问题：\n\n• 故障上报流程咨询\n• 故障处理进度查询\n• 召回信息查询\n• 售后政策咨询\n\n请问有什么可以帮您的？',
    time: formatTime(new Date())
  }
])

// 当前用户
const user = computed(() => authStore.user)

// 常见问题列表
const faqList = ref([
  { question: '如何上报零部件故障？' },
  { question: '故障上报需要提供哪些信息？' },
  { question: '故障处理进度怎么查询？' },
  { question: '如何查询零部件是否在召回范围内？' },
  { question: '零部件质保期是多久？' },
  { question: '故障零部件维修后有质保吗？' }
])

// ==================== 方法 ====================

/**
 * 格式化时间
 */
function formatTime(date: Date): string {
  return date.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
}

/**
 * 生成唯一ID
 */
function generateId(): string {
  return Date.now().toString(36) + Math.random().toString(36).substr(2)
}

/**
 * 滚动到底部
 */
function scrollToBottom() {
  nextTick(() => {
    if (messageListRef.value) {
      messageListRef.value.scrollTop = messageListRef.value.scrollHeight
    }
  })
}

/**
 * 发送消息
 */
async function handleSend() {
  const question = inputMessage.value.trim()
  if (!question) {
    message.warning('请输入问题')
    return
  }
  
  // 添加用户消息
  messages.value.push({
    id: generateId(),
    type: 'user',
    content: question,
    time: formatTime(new Date())
  })
  
  inputMessage.value = ''
  sending.value = true
  scrollToBottom()
  
  try {
    // 调用AI服务
    const response = await sendAIQuestion({
      question,
      userID: user.value?.id || 'anonymous'
    })
    
    if (response.success && response.data) {
      // 添加AI回复
      messages.value.push({
        id: generateId(),
        type: 'bot',
        content: response.data.answer,
        time: formatTime(new Date()),
        actions: response.data.relatedActions
      })
    } else {
      throw new Error(response.message || '获取回答失败')
    }
  } catch (error: any) {
    messages.value.push({
      id: generateId(),
      type: 'bot',
      content: '抱歉，我暂时无法回答这个问题。请稍后再试或联系人工客服。',
      time: formatTime(new Date())
    })
  } finally {
    sending.value = false
    scrollToBottom()
  }
}

/**
 * 点击快捷问题
 */
function askQuestion(question: string) {
  inputMessage.value = question
  handleSend()
}

/**
 * 处理快捷操作
 */
function handleAction(action: any) {
  switch (action.action) {
    case 'navigate':
      router.push(action.params?.path || '/')
      break
    case 'query':
      if (action.params?.type === 'fault') {
        inputMessage.value = '请输入故障ID查询进度：FAULT-001'
        message.info('请输入故障ID进行查询')
      } else if (action.params?.type === 'recall') {
        inputMessage.value = '请输入批次号查询召回信息：BATCH-2024-001'
        message.info('请输入批次号进行查询')
      }
      break
    default:
      break
  }
}

/**
 * 清空对话
 */
function handleClear() {
  messages.value = [
    {
      id: generateId(),
      type: 'bot',
      content: '对话已清空。请问有什么可以帮您的？',
      time: formatTime(new Date())
    }
  ]
}

// ==================== 生命周期 ====================

onMounted(() => {
  scrollToBottom()
})
</script>

<style scoped>
.ai-assistant {
  padding: 0;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-header h2 {
  font-size: 20px;
  font-weight: 600;
  margin: 0;
}

.chat-card {
  border-radius: 8px;
  height: calc(100vh - 200px);
  display: flex;
  flex-direction: column;
}

.message-list {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  min-height: 400px;
  max-height: calc(100vh - 400px);
}

.message-item {
  display: flex;
  margin-bottom: 16px;
}

.message-item.user {
  flex-direction: row-reverse;
}

.message-avatar {
  flex-shrink: 0;
}

.message-content {
  max-width: 70%;
  margin: 0 12px;
}

.message-item.user .message-content {
  text-align: right;
}

.message-text {
  display: inline-block;
  padding: 12px 16px;
  border-radius: 12px;
  background: var(--bg-color-tertiary);
  text-align: left;
  white-space: pre-wrap;
  word-break: break-word;
}

.message-item.user .message-text {
  background: #1890ff;
  color: #fff;
}

.message-actions {
  margin-top: 8px;
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.message-time {
  font-size: 12px;
  color: var(--text-color-tertiary);
  margin-top: 4px;
}

.input-area {
  border-top: 1px solid var(--border-color-secondary);
  padding-top: 16px;
  margin-top: 16px;
}

.input-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 12px;
}

.faq-card,
.info-card {
  border-radius: 8px;
}

/* 深色主题 */
[data-theme='dark'] .message-text {
  background: var(--bg-color-tertiary);
}

[data-theme='dark'] .input-area {
  border-top-color: var(--border-color);
}
</style>
