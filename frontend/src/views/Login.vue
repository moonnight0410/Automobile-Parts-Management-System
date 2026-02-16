<!--
  登录页面
  用户登录和注册入口
-->

<template>
  <div class="login-container">
    <div class="background-decoration">
      <div class="circle circle-1"></div>
      <div class="circle circle-2"></div>
      <div class="circle circle-3"></div>
      <div class="grid-pattern"></div>
    </div>
    
    <div class="login-wrapper">
      <div class="left-panel">
        <div class="brand-section">
          <div class="brand-icon">
            <CarOutlined />
          </div>
          <h1 class="brand-title">汽车零部件区块链管理系统</h1>
          <p class="brand-subtitle">Automobile Parts Blockchain Management System</p>
        </div>
        
        <div class="features-section">
          <div class="feature-item">
            <div class="feature-icon">
              <SafetyCertificateOutlined />
            </div>
            <div class="feature-content">
              <h3>区块链溯源</h3>
              <p>全程可追溯的零部件生命周期管理</p>
            </div>
          </div>
          <div class="feature-item">
            <div class="feature-icon">
              <ClusterOutlined />
            </div>
            <div class="feature-content">
              <h3>多方协作</h3>
              <p>制造商、车企、售后中心高效协同</p>
            </div>
          </div>
          <div class="feature-item">
            <div class="feature-icon">
              <CloudSyncOutlined />
            </div>
            <div class="feature-content">
              <h3>实时同步</h3>
              <p>数据实时上链，信息透明可信</p>
            </div>
          </div>
        </div>
        
        <div class="tech-stack">
          <span class="tech-label">技术支持</span>
          <div class="tech-icons">
            <span class="tech-item">Hyperledger Fabric</span>
            <span class="tech-item">Vue 3</span>
            <span class="tech-item">TypeScript</span>
          </div>
        </div>
      </div>
      
      <div class="right-panel">
        <a-card class="login-card" :bordered="false">
          <button class="theme-toggle" @click="toggleTheme" :title="isDark ? '切换到浅色模式' : '切换到深色模式'">
            <svg v-if="isDark" class="theme-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle cx="12" cy="12" r="5" stroke="currentColor" stroke-width="2"/>
              <line x1="12" y1="1" x2="12" y2="3" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="12" y1="21" x2="12" y2="23" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="4.22" y1="4.22" x2="5.64" y2="5.64" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="18.36" y1="18.36" x2="19.78" y2="19.78" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="1" y1="12" x2="3" y2="12" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="21" y1="12" x2="23" y2="12" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="4.22" y1="19.78" x2="5.64" y2="18.36" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <line x1="18.36" y1="5.64" x2="19.78" y2="4.22" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
            <svg v-else class="theme-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </button>
          
          <div class="login-header">
            <div class="header-icon">
              <UserOutlined />
            </div>
            <h2>{{ activeTab === 'login' ? '欢迎回来' : '创建账户' }}</h2>
            <p>{{ activeTab === 'login' ? '请登录您的账户以继续' : '填写以下信息完成注册' }}</p>
          </div>
          
          <div class="mode-switch">
            <button 
              class="mode-btn" 
              :class="{ active: activeTab === 'login' }"
              @click="activeTab = 'login'"
            >
              <LoginOutlined />
              <span>登录</span>
            </button>
            <button 
              class="mode-btn" 
              :class="{ active: activeTab === 'register' }"
              @click="activeTab = 'register'"
            >
              <UserAddOutlined />
              <span>注册</span>
            </button>
          </div>
          
          <a-form
            v-show="activeTab === 'login'"
            :model="loginForm"
            :rules="loginRules"
            @finish="handleLogin"
            layout="vertical"
            class="login-form"
          >
            <a-form-item name="username">
              <div class="input-wrapper">
                <label class="input-label">
                  <UserOutlined class="label-icon" />
                  <span>用户名</span>
                </label>
                <a-input
                  v-model:value="loginForm.username"
                  size="large"
                  placeholder="请输入用户名"
                />
              </div>
            </a-form-item>
            
            <a-form-item name="password">
              <div class="input-wrapper">
                <label class="input-label">
                  <LockOutlined class="label-icon" />
                  <span>密码</span>
                </label>
                <a-input-password
                  v-model:value="loginForm.password"
                  size="large"
                  placeholder="请输入密码"
                />
              </div>
            </a-form-item>
            
            <div class="form-options">
              <a-checkbox v-model:checked="rememberMe">记住我</a-checkbox>
              <a href="#" class="forgot-link">忘记密码？</a>
            </div>
            
            <a-form-item>
              <a-button
                type="primary"
                html-type="submit"
                size="large"
                block
                :loading="loading"
                class="submit-btn"
              >
                <template #icon><LoginOutlined /></template>
                登录系统
              </a-button>
            </a-form-item>
          </a-form>
          
          <div v-show="activeTab === 'login'" class="quick-login">
            <div class="divider">
              <span>快速登录（演示）</span>
            </div>
            <div class="quick-buttons">
              <button class="quick-btn manufacturer" @click="quickLogin('manufacturer')">
                <span class="quick-icon"><ToolOutlined /></span>
                <span>制造商</span>
              </button>
              <button class="quick-btn automaker" @click="quickLogin('automaker')">
                <span class="quick-icon"><CarOutlined /></span>
                <span>车企</span>
              </button>
              <button class="quick-btn aftersale" @click="quickLogin('aftersale')">
                <span class="quick-icon"><CustomerServiceOutlined /></span>
                <span>售后中心</span>
              </button>
            </div>
          </div>
          
          <a-form
            v-show="activeTab === 'register'"
            :model="registerForm"
            :rules="registerRules"
            @finish="handleRegister"
            layout="vertical"
            class="register-form"
          >
            <a-form-item name="username">
              <div class="input-wrapper">
                <label class="input-label">
                  <UserOutlined class="label-icon" />
                  <span>用户名</span>
                </label>
                <a-input
                  v-model:value="registerForm.username"
                  size="large"
                  placeholder="请输入用户名"
                />
                <span class="field-hint">至少3个字符</span>
              </div>
            </a-form-item>
            
            <a-form-item name="password">
              <div class="input-wrapper">
                <label class="input-label">
                  <LockOutlined class="label-icon" />
                  <span>密码</span>
                </label>
                <a-input-password
                  v-model:value="registerForm.password"
                  size="large"
                  placeholder="请输入密码"
                />
                <span class="field-hint">至少6位字符</span>
              </div>
            </a-form-item>
            
            <a-form-item name="confirmPassword">
              <div class="input-wrapper">
                <label class="input-label">
                  <SafetyOutlined class="label-icon" />
                  <span>确认密码</span>
                </label>
                <a-input-password
                  v-model:value="registerForm.confirmPassword"
                  size="large"
                  placeholder="请再次输入密码"
                />
                <span class="field-hint">需与密码一致</span>
              </div>
            </a-form-item>
            
            <a-form-item name="org">
              <div class="input-wrapper">
                <label class="input-label">
                  <ApartmentOutlined class="label-icon" />
                  <span>组织</span>
                </label>
                <a-select
                  v-model:value="registerForm.org"
                  size="large"
                  placeholder="请选择组织"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                >
                  <a-select-option value="Org1MSP">制造商组织</a-select-option>
                  <a-select-option value="Org2MSP">车企组织</a-select-option>
                  <a-select-option value="Org3MSP">售后组织</a-select-option>
                </a-select>
              </div>
            </a-form-item>
            
            <a-form-item name="role">
              <div class="input-wrapper">
                <label class="input-label">
                  <IdcardOutlined class="label-icon" />
                  <span>角色</span>
                </label>
                <a-select
                  v-model:value="registerForm.role"
                  size="large"
                  placeholder="请选择角色"
                  :get-popup-container="(triggerNode: any) => triggerNode.parentNode"
                >
                  <a-select-option value="manufacturer">制造商</a-select-option>
                  <a-select-option value="automaker">车企</a-select-option>
                  <a-select-option value="aftersale">售后中心</a-select-option>
                </a-select>
              </div>
            </a-form-item>
            
            <a-form-item>
              <a-button
                type="primary"
                html-type="submit"
                size="large"
                block
                :loading="loading"
                class="submit-btn"
              >
                <template #icon><UserAddOutlined /></template>
                创建账户
              </a-button>
            </a-form-item>
          </a-form>
          
          <div class="login-footer">
            <p>© 2026 汽车零部件区块链管理系统</p>
          </div>
        </a-card>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message } from 'ant-design-vue'
import { useAuthStore } from '../stores/auth.store'
import { UserRole } from '../types'
import {
  CarOutlined,
  SafetyCertificateOutlined,
  ClusterOutlined,
  CloudSyncOutlined,
  UserOutlined,
  LoginOutlined,
  UserAddOutlined,
  LockOutlined,
  ToolOutlined,
  CustomerServiceOutlined,
  SafetyOutlined,
  ApartmentOutlined,
  IdcardOutlined
} from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const activeTab = ref('login')
const loading = ref(false)
const rememberMe = ref(false)
const isDark = ref(false)

const toggleTheme = () => {
  isDark.value = !isDark.value
  const html = document.documentElement
  const newTheme = isDark.value ? 'dark' : 'light'
  html.setAttribute('data-theme', newTheme)
  localStorage.setItem('theme', newTheme)
}

const initTheme = () => {
  const savedTheme = localStorage.getItem('theme')
  const html = document.documentElement
  if (savedTheme) {
    isDark.value = savedTheme === 'dark'
    html.setAttribute('data-theme', savedTheme)
  } else {
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
    isDark.value = prefersDark
    html.setAttribute('data-theme', prefersDark ? 'dark' : 'light')
  }
}

initTheme()

const loginForm = reactive({
  username: '',
  password: ''
})

const registerForm = reactive({
  username: '',
  password: '',
  confirmPassword: '',
  org: '',
  role: '' as UserRole | ''
})

const validateConfirmPassword = async (_rule: any, value: string) => {
  if (value !== registerForm.password) {
    return Promise.reject('两次输入的密码不一致')
  }
  return Promise.resolve()
}

const loginRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码至少6位', trigger: 'blur' }
  ]
}

const registerRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, message: '用户名至少3个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码至少6位', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    { validator: validateConfirmPassword, trigger: 'blur' }
  ],
  org: [
    { required: true, message: '请选择组织', trigger: 'change' }
  ],
  role: [
    { required: true, message: '请选择角色', trigger: 'change' }
  ]
}

async function handleLogin() {
  loading.value = true
  
  try {
    await authStore.login(loginForm.username, loginForm.password)
    message.success('登录成功')
    
    const redirect = (route.query.redirect as string) || '/'
    router.push(redirect)
  } catch (error: any) {
    message.error(error.message || '登录失败')
  } finally {
    loading.value = false
  }
}

async function handleRegister() {
  loading.value = true
  
  try {
    await authStore.register({
      userID: registerForm.username,
      password: registerForm.password,
      org: registerForm.org,
      role: registerForm.role as UserRole
    })
    
    message.success({
      content: '注册成功！请使用新账号登录',
      duration: 3
    })
    
    activeTab.value = 'login'
    loginForm.username = registerForm.username
    loginForm.password = ''
    
    registerForm.username = ''
    registerForm.password = ''
    registerForm.confirmPassword = ''
    registerForm.org = ''
    registerForm.role = ''
  } catch (error: any) {
    const errorMsg = error.response?.data?.message || error.message || '注册失败，请稍后重试'
    message.error({
      content: errorMsg,
      duration: 4
    })
  } finally {
    loading.value = false
  }
}

async function quickLogin(role: string) {
  const usernames: Record<string, string> = {
    manufacturer: 'manufacturer_user',
    automaker: 'automaker_user',
    aftersale: 'aftersale_user'
  }
  
  loginForm.username = usernames[role]
  loginForm.password = '123456'
  
  await handleLogin()
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  position: relative;
  overflow: hidden;
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

.background-decoration {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: hidden;
  pointer-events: none;
}

.circle {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.08);
}

.circle-1 {
  width: 600px;
  height: 600px;
  top: -200px;
  left: -200px;
}

.circle-2 {
  width: 400px;
  height: 400px;
  bottom: -100px;
  right: -100px;
}

.circle-3 {
  width: 300px;
  height: 300px;
  top: 40%;
  left: 30%;
}

.grid-pattern {
  position: absolute;
  width: 100%;
  height: 100%;
  background-image: 
    linear-gradient(rgba(255, 255, 255, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.03) 1px, transparent 1px);
  background-size: 50px 50px;
}

.login-wrapper {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  max-width: 1200px;
  width: 100%;
  padding: 40px;
  z-index: 1;
}

.left-panel {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: #fff;
}

.brand-section {
  margin-bottom: 60px;
}

.brand-icon {
  width: 72px;
  height: 72px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 36px;
  margin-bottom: 24px;
  backdrop-filter: blur(10px);
}

.brand-title {
  font-size: 32px;
  font-weight: 700;
  margin: 0 0 12px;
  line-height: 1.3;
}

.brand-subtitle {
  font-size: 16px;
  opacity: 0.85;
  margin: 0;
}

.features-section {
  display: flex;
  flex-direction: column;
  gap: 28px;
  margin-bottom: 60px;
}

.feature-item {
  display: flex;
  align-items: flex-start;
  gap: 16px;
}

.feature-icon {
  width: 48px;
  height: 48px;
  background: rgba(255, 255, 255, 0.15);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  flex-shrink: 0;
}

.feature-content h3 {
  font-size: 18px;
  font-weight: 600;
  margin: 0 0 6px;
}

.feature-content p {
  font-size: 14px;
  opacity: 0.8;
  margin: 0;
}

.tech-stack {
  padding-top: 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.15);
}

.tech-label {
  font-size: 13px;
  opacity: 0.7;
  margin-bottom: 12px;
  display: block;
}

.tech-icons {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

.tech-item {
  padding: 6px 14px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 20px;
  font-size: 13px;
  backdrop-filter: blur(5px);
}

.right-panel {
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-card {
  width: 100%;
  max-width: 460px;
  border-radius: 20px;
  box-shadow: 0 25px 80px rgba(0, 0, 0, 0.3);
  position: relative;
  background: #fff;
  padding: 40px;
}

.theme-toggle {
  position: absolute;
  top: 20px;
  right: 20px;
  width: 42px;
  height: 42px;
  border: none;
  border-radius: 12px;
  background: #f3f4f6;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
}

.theme-toggle:hover {
  background: #e5e7eb;
  color: #6366f1;
}

.theme-icon {
  width: 20px;
  height: 20px;
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.header-icon {
  width: 56px;
  height: 56px;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 26px;
  color: #fff;
  margin: 0 auto 20px;
}

.login-header h2 {
  font-size: 24px;
  font-weight: 700;
  color: #1f2937;
  margin: 0 0 8px;
}

.login-header p {
  font-size: 14px;
  color: #6b7280;
  margin: 0;
}

.mode-switch {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
  margin-bottom: 32px;
  padding: 6px;
  background: #f3f4f6;
  border-radius: 14px;
}

.mode-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 14px 20px;
  border: none;
  border-radius: 10px;
  background: transparent;
  color: #6b7280;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.mode-btn:hover {
  color: #6366f1;
}

.mode-btn.active {
  background: #fff;
  color: #6366f1;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.login-form,
.register-form {
  margin-top: 0;
}

.input-wrapper {
  width: 100%;
}

.input-label {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 10px;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
}

.label-icon {
  color: #6366f1;
  font-size: 14px;
}

.input-wrapper :deep(.ant-input),
.input-wrapper :deep(.ant-input-affix-wrapper) {
  width: 100% !important;
  height: 48px !important;
  line-height: 48px !important;
  border-radius: 12px !important;
  border: 2px solid #e5e7eb !important;
  font-size: 15px !important;
  padding: 0 16px !important;
  background: #fff !important;
  color: #374151 !important;
  transition: all 0.3s ease !important;
  box-sizing: border-box !important;
}

.input-wrapper :deep(.ant-input-affix-wrapper) {
  display: flex !important;
  align-items: center !important;
}

.input-wrapper :deep(.ant-input-affix-wrapper .ant-input) {
  height: 44px !important;
  line-height: 44px !important;
  border: none !important;
  border-radius: 0 !important;
  padding: 0 !important;
  margin: 0 !important;
  background: transparent !important;
  font-size: 15px !important;
  color: #374151 !important;
  flex: 1 !important;
}

.input-wrapper :deep(.ant-input::placeholder),
.input-wrapper :deep(.ant-input-affix-wrapper .ant-input::placeholder) {
  color: #9ca3af !important;
}

.input-wrapper :deep(.ant-input:hover),
.input-wrapper :deep(.ant-input-affix-wrapper:hover) {
  border-color: #d1d5db !important;
}

.input-wrapper :deep(.ant-input:focus),
.input-wrapper :deep(.ant-input-focused),
.input-wrapper :deep(.ant-input-affix-wrapper-focused),
.input-wrapper :deep(.ant-input-affix-wrapper:focus-within) {
  border-color: #6366f1 !important;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.15) !important;
  outline: none !important;
}

.input-wrapper :deep(.ant-input-suffix),
.input-wrapper :deep(.ant-input-password-icon) {
  color: #9ca3af !important;
}

.input-wrapper :deep(.ant-input-password-icon:hover) {
  color: #6366f1 !important;
}

.input-wrapper :deep(.ant-select) {
  width: 100% !important;
}

.input-wrapper :deep(.ant-select-selector) {
  width: 100% !important;
  height: 48px !important;
  line-height: 48px !important;
  border-radius: 12px !important;
  border: 2px solid #e5e7eb !important;
  padding: 0 16px !important;
  background: #fff !important;
  transition: all 0.3s ease !important;
  box-sizing: border-box !important;
  display: flex !important;
  align-items: center !important;
}

.input-wrapper :deep(.ant-select-selection-search) {
  display: flex !important;
  align-items: center !important;
}

.input-wrapper :deep(.ant-select-selection-search-input) {
  height: 44px !important;
}

.input-wrapper :deep(.ant-select-selection-item) {
  line-height: 44px !important;
  font-size: 15px !important;
  color: #374151 !important;
  display: flex !important;
  align-items: center !important;
}

.input-wrapper :deep(.ant-select-selection-placeholder) {
  line-height: 44px !important;
  font-size: 15px !important;
  color: #9ca3af !important;
  display: flex !important;
  align-items: center !important;
}

.input-wrapper :deep(.ant-select:hover .ant-select-selector) {
  border-color: #d1d5db !important;
}

.input-wrapper :deep(.ant-select-focused .ant-select-selector) {
  border-color: #6366f1 !important;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.15) !important;
}

.input-wrapper :deep(.ant-select-arrow) {
  color: #9ca3af !important;
}

.input-wrapper :deep(.ant-select-dropdown) {
  border-radius: 12px !important;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.12) !important;
  border: 1px solid #f0f0f0 !important;
  padding: 8px 0 !important;
}

.input-wrapper :deep(.ant-select-item) {
  padding: 10px 12px !important;
  font-size: 14px !important;
  color: #374151 !important;
  border-radius: 8px !important;
  transition: all 0.2s ease !important;
  margin: 2px 0 !important;
}

.input-wrapper :deep(.ant-select-item:hover) {
  background: #f5f3ff !important;
  color: #6366f1 !important;
}

.input-wrapper :deep(.ant-select-item-option-selected) {
  background: #ede9fe !important;
  color: #6366f1 !important;
  font-weight: 500 !important;
}

.field-hint {
  display: block;
  font-size: 12px;
  color: #9ca3af;
  margin-top: 6px;
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.forgot-link {
  color: #6366f1;
  font-size: 14px;
}

.forgot-link:hover {
  text-decoration: underline;
}

.submit-btn {
  height: 50px;
  border-radius: 12px;
  background: linear-gradient(135deg, #6366f1 0%, #4f46e5 100%);
  border: none;
  font-size: 16px;
  font-weight: 600;
  box-shadow: 0 4px 15px rgba(99, 102, 241, 0.35);
}

.submit-btn:hover {
  background: linear-gradient(135deg, #4f46e5 0%, #4338ca 100%);
  box-shadow: 0 6px 20px rgba(99, 102, 241, 0.45);
}

.quick-login {
  margin-top: 28px;
}

.divider {
  text-align: center;
  position: relative;
  margin-bottom: 20px;
}

.divider::before {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  top: 50%;
  height: 1px;
  background: #e5e7eb;
}

.divider span {
  background: #fff;
  padding: 0 16px;
  position: relative;
  color: #9ca3af;
  font-size: 13px;
}

.quick-buttons {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
}

.quick-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px 12px;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  background: #fff;
  cursor: pointer;
  transition: all 0.3s ease;
}

.quick-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.quick-btn.manufacturer:hover {
  border-color: #3b82f6;
  color: #3b82f6;
}

.quick-btn.automaker:hover {
  border-color: #10b981;
  color: #10b981;
}

.quick-btn.aftersale:hover {
  border-color: #f59e0b;
  color: #f59e0b;
}

.quick-icon {
  font-size: 22px;
  color: #6b7280;
}

.quick-btn:hover .quick-icon {
  color: inherit;
}

.quick-btn span:last-child {
  font-size: 13px;
  font-weight: 500;
  color: #374151;
}

.login-footer {
  text-align: center;
  margin-top: 28px;
  padding-top: 20px;
  border-top: 1px solid #f3f4f6;
}

.login-footer p {
  font-size: 13px;
  color: #9ca3af;
  margin: 0;
}

[data-theme='dark'] .login-container {
  background: linear-gradient(135deg, #1e1b4b 0%, #312e81 100%);
}

[data-theme='dark'] .login-card {
  background: #1f2937;
}

[data-theme='dark'] .theme-toggle {
  background: #374151;
  color: #9ca3af;
}

[data-theme='dark'] .theme-toggle:hover {
  background: #4b5563;
  color: #a78bfa;
}

[data-theme='dark'] .login-header h2 {
  color: #f3f4f6;
}

[data-theme='dark'] .login-header p {
  color: #9ca3af;
}

[data-theme='dark'] .mode-switch {
  background: #111827;
}

[data-theme='dark'] .mode-btn {
  color: #9ca3af;
}

[data-theme='dark'] .mode-btn:hover {
  color: #a78bfa;
}

[data-theme='dark'] .mode-btn.active {
  background: #1f2937;
  color: #a78bfa;
}

[data-theme='dark'] .input-label {
  color: #e5e7eb !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-input),
[data-theme='dark'] .input-wrapper :deep(.ant-input-affix-wrapper) {
  background: #111827 !important;
  border-color: #374151 !important;
  color: #f3f4f6 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-input-affix-wrapper .ant-input) {
  background: transparent !important;
  color: #f3f4f6 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-input::placeholder),
[data-theme='dark'] .input-wrapper :deep(.ant-input-affix-wrapper .ant-input::placeholder) {
  color: #6b7280 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-input:hover),
[data-theme='dark'] .input-wrapper :deep(.ant-input-affix-wrapper:hover) {
  border-color: #4b5563 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-input:focus),
[data-theme='dark'] .input-wrapper :deep(.ant-input-focused),
[data-theme='dark'] .input-wrapper :deep(.ant-input-affix-wrapper-focused),
[data-theme='dark'] .input-wrapper :deep(.ant-input-affix-wrapper:focus-within) {
  border-color: #8b5cf6 !important;
  box-shadow: 0 0 0 3px rgba(139, 92, 246, 0.15) !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-input-suffix),
[data-theme='dark'] .input-wrapper :deep(.ant-input-password-icon) {
  color: #6b7280 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-input-password-icon:hover) {
  color: #a78bfa !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-selector) {
  background: #111827 !important;
  border-color: #374151 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-selection-item) {
  color: #f3f4f6 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-selection-placeholder) {
  color: #6b7280 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select:hover .ant-select-selector) {
  border-color: #4b5563 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-focused .ant-select-selector) {
  border-color: #8b5cf6 !important;
  box-shadow: 0 0 0 3px rgba(139, 92, 246, 0.15) !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-arrow) {
  color: #6b7280 !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-dropdown) {
  background: #1f2937 !important;
  border-color: #374151 !important;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.4) !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-item) {
  color: #e5e7eb !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-item:hover) {
  background: rgba(129, 140, 248, 0.15) !important;
  color: #a5b4fc !important;
}

[data-theme='dark'] .input-wrapper :deep(.ant-select-item-option-selected) {
  background: rgba(129, 140, 248, 0.2) !important;
  color: #a5b4fc !important;
}

[data-theme='dark'] .forgot-link {
  color: #a78bfa;
}

[data-theme='dark'] .submit-btn {
  background: linear-gradient(135deg, #8b5cf6 0%, #7c3aed 100%);
}

[data-theme='dark'] .divider::before {
  background: #374151;
}

[data-theme='dark'] .divider span {
  background: #1f2937;
  color: #6b7280;
}

[data-theme='dark'] .quick-btn {
  background: #111827;
  border-color: #374151;
}

[data-theme='dark'] .quick-btn:hover {
  background: #1f2937;
}

[data-theme='dark'] .quick-icon {
  color: #9ca3af;
}

[data-theme='dark'] .quick-btn span:last-child {
  color: #e5e7eb;
}

[data-theme='dark'] .login-footer {
  border-top-color: #374151;
}

[data-theme='dark'] .login-footer p {
  color: #6b7280;
}
</style>
