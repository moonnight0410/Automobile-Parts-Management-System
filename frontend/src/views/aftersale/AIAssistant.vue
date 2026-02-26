<template>
  <div class="ai-assistant">
    <div class="page-header">
      <div class="header-content">
        <div class="header-title">
          <div class="title-icon">
            <RobotOutlined />
          </div>
          <div class="title-text">
            <h2>智能售后助手</h2>
            <p class="subtitle">AI驱动的智能问答与售后服务支持</p>
          </div>
        </div>
      </div>
    </div>

    <a-row :gutter="20" class="chat-row">
      <a-col :xs="24" :lg="16">
        <a-card :bordered="false" class="chat-card">
          <div class="chat-header">
            <div class="chat-header-left">
              <div class="chat-avatar-wrapper">
                <a-avatar :size="40" class="bot-avatar">
                  <template #icon><RobotOutlined /></template>
                </a-avatar>
                <span class="online-dot"></span>
              </div>
              <div class="chat-info">
                <span class="chat-name">智能售后助手</span>
                <span class="chat-status">在线 · 随时为您服务</span>
              </div>
            </div>
            <div class="chat-header-right">
              <a-button class="header-action-btn" @click="handleClear">
                <template #icon><ClearOutlined /></template>
                清空对话
              </a-button>
            </div>
          </div>

          <div class="message-list" ref="messageListRef">
            <div
              v-for="msg in messages"
              :key="msg.id"
              :class="['message-item', msg.type]"
            >
              <div class="message-avatar">
                <a-avatar v-if="msg.type === 'user'" :size="36" :style="{ backgroundColor: '#6366f1' }">
                  {{ user?.username?.charAt(0).toUpperCase() || 'U' }}
                </a-avatar>
                <a-avatar v-else :size="36" :style="{ backgroundColor: '#10b981' }">
                  <template #icon><RobotOutlined /></template>
                </a-avatar>
              </div>
              <div class="message-content">
                <div class="message-bubble">
                  <div class="message-text">{{ msg.content }}</div>
                </div>
                <div v-if="msg.actions && msg.actions.length" class="message-actions">
                  <a-button
                    v-for="action in msg.actions"
                    :key="action.label"
                    size="small"
                    type="primary"
                    ghost
                    @click="handleAction(action)"
                    class="action-btn"
                  >
                    {{ action.label }}
                  </a-button>
                </div>
                <div class="message-time">{{ msg.time }}</div>
              </div>
            </div>

            <div v-if="sending" class="message-item bot">
              <div class="message-avatar">
                <a-avatar :size="36" :style="{ backgroundColor: '#10b981' }">
                  <template #icon><RobotOutlined /></template>
                </a-avatar>
              </div>
              <div class="message-content">
                <div class="message-bubble typing">
                  <div class="typing-indicator">
                    <span></span>
                    <span></span>
                    <span></span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="input-area">
            <div class="input-wrapper">
              <a-textarea
                v-model:value="inputMessage"
                placeholder="请输入您的问题，例如：如何上报零部件故障？"
                :auto-size="{ minRows: 1, maxRows: 4 }"
                class="message-input"
                @pressEnter="handleSend"
              />
              <a-button type="primary" class="send-btn" @click="handleSend" :loading="sending" :disabled="!inputMessage.trim()">
                <template #icon><SendOutlined /></template>
              </a-button>
            </div>
            <div class="input-hints">
              <span class="hint-item" v-for="hint in quickHints" :key="hint" @click="inputMessage = hint">
                {{ hint }}
              </span>
            </div>
          </div>
        </a-card>
      </a-col>

      <a-col :xs="24" :lg="8">
        <a-card :bordered="false" class="faq-card">
          <template #title>
            <div class="card-title">
              <QuestionCircleOutlined class="title-icon" />
              <span>常见问题</span>
            </div>
          </template>
          <div class="faq-list">
            <div
              v-for="(item, index) in faqList"
              :key="index"
              class="faq-item"
              @click="askQuestion(item.question)"
            >
              <div class="faq-icon">
                <MessageOutlined />
              </div>
              <div class="faq-content">
                <div class="faq-question">{{ item.question }}</div>
                <div class="faq-category">{{ item.category }}</div>
              </div>
              <RightOutlined class="faq-arrow" />
            </div>
          </div>
        </a-card>

        <a-card :bordered="false" class="feature-card">
          <template #title>
            <div class="card-title">
              <ThunderboltOutlined class="title-icon" />
              <span>功能说明</span>
            </div>
          </template>
          <div class="feature-list">
            <div class="feature-item">
              <div class="feature-icon" style="background: linear-gradient(135deg, #f97316 0%, #ea580c 100%);">
                <AlertOutlined />
              </div>
              <div class="feature-info">
                <div class="feature-name">故障上报咨询</div>
                <div class="feature-desc">了解故障上报流程和所需材料</div>
              </div>
            </div>
            <div class="feature-item">
              <div class="feature-icon" style="background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);">
                <SearchOutlined />
              </div>
              <div class="feature-info">
                <div class="feature-name">故障进度查询</div>
                <div class="feature-desc">查询故障处理进度和状态</div>
              </div>
            </div>
            <div class="feature-item">
              <div class="feature-icon" style="background: linear-gradient(135deg, #ef4444 0%, #dc2626 100%);">
                <WarningOutlined />
              </div>
              <div class="feature-info">
                <div class="feature-name">召回相关查询</div>
                <div class="feature-desc">查询召回范围和处理流程</div>
              </div>
            </div>
            <div class="feature-item">
              <div class="feature-icon" style="background: linear-gradient(135deg, #10b981 0%, #059669 100%);">
                <SafetyCertificateOutlined />
              </div>
              <div class="feature-info">
                <div class="feature-name">售后政策咨询</div>
                <div class="feature-desc">了解质保期、维修政策等</div>
              </div>
            </div>
          </div>
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
import {
  sendAIQuestion,
  getConversationHistory,
  clearConversationHistory
} from '../../services/ai.service'
import {
  RobotOutlined,
  ClearOutlined,
  SendOutlined,
  QuestionCircleOutlined,
  MessageOutlined,
  RightOutlined,
  AlertOutlined,
  SearchOutlined,
  WarningOutlined,
  SafetyCertificateOutlined,
  ThunderboltOutlined
} from '@ant-design/icons-vue'

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

const router = useRouter()
const authStore = useAuthStore()

const messageListRef = ref<HTMLElement>()
const inputMessage = ref('')
const sending = ref(false)

const messages = ref<Message[]>([
  {
    id: '1',
    type: 'bot',
    content: '您好！我是智能售后助手，可以帮您解答以下问题：\n\n• 故障上报流程咨询\n• 故障处理进度查询\n• 召回信息查询\n• 售后政策咨询\n\n请问有什么可以帮您的？',
    time: formatTime(new Date())
  }
])

const user = computed(() => authStore.user)

const quickHints = ref([
  '如何上报故障？',
  '查询故障进度',
  '召回范围查询'
])

const faqList = ref([
  { question: '如何上报零部件故障？', category: '故障上报' },
  { question: '故障上报需要提供哪些信息？', category: '故障上报' },
  { question: '故障处理进度怎么查询？', category: '进度查询' },
  { question: '如何查询零部件是否在召回范围内？', category: '召回查询' },
  { question: '零部件质保期是多久？', category: '政策咨询' },
  { question: '故障零部件维修后有质保吗？', category: '政策咨询' }
])

function formatTime(date: Date): string {
  return date.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
}

function generateId(): string {
  return Date.now().toString(36) + Math.random().toString(36).substr(2)
}

function scrollToBottom() {
  nextTick(() => {
    if (messageListRef.value) {
      messageListRef.value.scrollTop = messageListRef.value.scrollHeight
    }
  })
}

async function handleSend() {
  const question = inputMessage.value.trim()
  if (!question) {
    message.warning('请输入问题')
    return
  }

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
    const response = await sendAIQuestion({
      question,
      userID: user.value?.id || 'anonymous'
    })

    if (response.code === 0 && response.data) {
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

function askQuestion(question: string) {
  inputMessage.value = question
  handleSend()
}

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

async function handleClear() {
  try {
    const userID = user.value?.id || 'anonymous'
    const response = await clearConversationHistory(userID)
    if (response.code === 0) {
      messages.value = [
        {
          id: generateId(),
          type: 'bot',
          content: '对话已清空。请问有什么可以帮您的？',
          time: formatTime(new Date())
        }
      ]
      message.success('对话已清空')
    } else {
      message.error(response.message || '清空对话失败')
    }
  } catch (error: any) {
    message.error(error.message || '清空对话失败')
  }
}

onMounted(() => {
  scrollToBottom()
})
</script>

<style scoped>
.ai-assistant {
  padding: 0;
  min-height: 100%;
}

.page-header {
  margin-bottom: 24px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border-radius: 12px;
  padding: 24px 32px;
  box-shadow: 0 4px 20px rgba(16, 185, 129, 0.25);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 16px;
}

.title-icon {
  width: 48px;
  height: 48px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: #fff;
}

.title-text h2 {
  font-size: 24px;
  font-weight: 600;
  margin: 0;
  color: #fff;
}

.title-text .subtitle {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.85);
  margin: 4px 0 0;
}

.ai-tag {
  padding: 6px 16px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 500;
}

.chat-row {
  height: calc(100vh - 280px);
}

.chat-card {
  height: 100%;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.chat-card :deep(.ant-card-body) {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 0;
  overflow: hidden;
}

.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid #f3f4f6;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
}

.chat-header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.chat-avatar-wrapper {
  position: relative;
}

.bot-avatar {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
}

.online-dot {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 10px;
  height: 10px;
  background: #10b981;
  border: 2px solid #fff;
  border-radius: 50%;
}

.chat-info {
  display: flex;
  flex-direction: column;
}

.chat-name {
  font-size: 15px;
  font-weight: 600;
  color: #1f2937;
}

.chat-status {
  font-size: 12px;
  color: #10b981;
}

.header-action-btn {
  height: 36px;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
  color: #6b7280;
}

.header-action-btn:hover {
  border-color: #10b981;
  color: #10b981;
}

.message-list {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  background: #fafafa;
}

.message-item {
  display: flex;
  margin-bottom: 20px;
  gap: 12px;
}

.message-item.user {
  flex-direction: row-reverse;
}

.message-avatar {
  flex-shrink: 0;
}

.message-content {
  max-width: 70%;
}

.message-bubble {
  display: inline-block;
  padding: 12px 16px;
  border-radius: 16px;
  background: #fff;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.message-item.user .message-bubble {
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  border-radius: 16px 16px 4px 16px;
}

.message-item.bot .message-bubble {
  border-radius: 16px 16px 16px 4px;
}

.message-text {
  font-size: 14px;
  line-height: 1.6;
  color: #374151;
  white-space: pre-wrap;
  word-break: break-word;
}

.message-item.user .message-text {
  color: #fff;
}

.message-actions {
  margin-top: 10px;
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.action-btn {
  border-radius: 16px;
  font-size: 12px;
  height: 28px;
  border-color: #10b981;
  color: #10b981;
}

.action-btn:hover {
  background: #10b981;
  color: #fff;
}

.message-time {
  font-size: 11px;
  color: #9ca3af;
  margin-top: 6px;
}

.message-item.user .message-time {
  text-align: right;
}

.typing-indicator {
  display: flex;
  gap: 4px;
  padding: 4px 0;
}

.typing-indicator span {
  width: 8px;
  height: 8px;
  background: #10b981;
  border-radius: 50%;
  animation: typing 1.4s infinite ease-in-out;
}

.typing-indicator span:nth-child(1) {
  animation-delay: 0s;
}

.typing-indicator span:nth-child(2) {
  animation-delay: 0.2s;
}

.typing-indicator span:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes typing {
  0%, 60%, 100% {
    transform: translateY(0);
    opacity: 0.4;
  }
  30% {
    transform: translateY(-4px);
    opacity: 1;
  }
}

.input-area {
  padding: 16px 20px;
  border-top: 1px solid #f3f4f6;
  background: #fff;
}

.input-wrapper {
  display: flex;
  gap: 12px;
  align-items: flex-end;
}

.message-input {
  flex: 1;
  border-radius: 12px;
  border: 2px solid #e5e7eb;
  font-size: 14px;
  resize: none;
}

.message-input:focus {
  border-color: #10b981;
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.1);
}

.send-btn {
  height: 40px;
  width: 40px;
  border-radius: 12px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border: none;
}

.send-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, #059669 0%, #047857 100%);
}

.send-btn:disabled {
  background: #e5e7eb;
}

.input-hints {
  display: flex;
  gap: 8px;
  margin-top: 12px;
  flex-wrap: wrap;
}

.hint-item {
  padding: 6px 12px;
  background: #f3f4f6;
  border-radius: 16px;
  font-size: 12px;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.2s;
}

.hint-item:hover {
  background: #10b981;
  color: #fff;
}

.faq-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  margin-bottom: 20px;
}

.card-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 15px;
  font-weight: 600;
}

.title-icon {
  color: #10b981;
}

.faq-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.faq-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: #f8fafc;
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.2s;
}

.faq-item:hover {
  background: #f1f5f9;
  transform: translateX(4px);
}

.faq-icon {
  width: 32px;
  height: 32px;
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 14px;
}

.faq-content {
  flex: 1;
}

.faq-question {
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.faq-category {
  font-size: 11px;
  color: #9ca3af;
  margin-top: 2px;
}

.faq-arrow {
  color: #d1d5db;
  font-size: 12px;
}

.feature-card {
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
}

.feature-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.feature-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: #f8fafc;
  border-radius: 10px;
}

.feature-icon {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 16px;
}

.feature-info {
  flex: 1;
}

.feature-name {
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.feature-desc {
  font-size: 11px;
  color: #9ca3af;
  margin-top: 2px;
}

[data-theme='dark'] .page-header {
  background: linear-gradient(135deg, #065f46 0%, #047857 100%);
}

[data-theme='dark'] .chat-header {
  background: linear-gradient(135deg, #1f2937 0%, #111827 100%);
  border-bottom-color: #374151;
}

[data-theme='dark'] .message-list {
  background: #111827;
}

[data-theme='dark'] .message-bubble {
  background: #1f2937;
}

[data-theme='dark'] .message-text {
  color: #e5e7eb;
}

[data-theme='dark'] .input-area {
  background: #1f2937;
  border-top-color: #374151;
}

[data-theme='dark'] .message-input {
  background: #111827;
  border-color: #374151;
  color: #e5e7eb;
}

[data-theme='dark'] .hint-item {
  background: #374151;
  color: #9ca3af;
}

[data-theme='dark'] .hint-item:hover {
  background: #10b981;
  color: #fff;
}

[data-theme='dark'] .faq-item,
[data-theme='dark'] .feature-item {
  background: #1f2937;
}

[data-theme='dark'] .faq-question,
[data-theme='dark'] .feature-name {
  color: #e5e7eb;
}

[data-theme='dark'] .chat-name {
  color: #e5e7eb;
}
</style>
