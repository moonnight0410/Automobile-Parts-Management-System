<!--
  登录页面
  用户登录和注册入口
-->

<template>
  <div class="login-container">
    <!-- 背景装饰 -->
    <div class="background-decoration">
      <div class="circle circle-1"></div>
      <div class="circle circle-2"></div>
      <div class="circle circle-3"></div>
    </div>
    
    <!-- 登录卡片 -->
    <a-card class="login-card" :bordered="false">
      <!-- 主题切换按钮 -->
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
        <img src="/vite.svg" alt="Logo" class="logo" />
        <h1 class="title">汽车零部件区块链管理系统</h1>
        <p class="subtitle">Automobile Parts Blockchain Management System</p>
      </div>
      
      <!-- 登录/注册切换按钮 -->
      <div class="mode-switch">
        <button 
          class="mode-btn" 
          :class="{ active: activeTab === 'login' }"
          @click="activeTab = 'login'"
        >
          <svg class="mode-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 12C14.7614 12 17 9.76142 17 7C17 4.23858 14.7614 2 12 2C9.23858 2 7 4.23858 7 7C7 9.76142 9.23858 12 12 12Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            <path d="M20.59 22C20.59 18.13 16.74 15 12 15C7.26 15 3.41 18.13 3.41 22" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
          <span>登录</span>
        </button>
        <button 
          class="mode-btn" 
          :class="{ active: activeTab === 'register' }"
          @click="activeTab = 'register'"
        >
          <svg class="mode-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M16 21V19C16 17.9391 15.5786 16.9217 14.8284 16.1716C14.0783 15.4214 13.0609 15 12 15H5C3.93913 15 2.92172 15.4214 2.17157 16.1716C1.42143 16.9217 1 17.9391 1 19V21" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            <circle cx="8.5" cy="7" r="4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            <line x1="20" y1="8" x2="20" y2="14" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            <line x1="23" y1="11" x2="17" y2="11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
          <span>注册</span>
        </button>
      </div>
      
      <!-- 登录表单 -->
      <a-form
        v-show="activeTab === 'login'"
        :model="loginForm"
        :rules="loginRules"
        @finish="handleLogin"
        layout="vertical"
        class="login-form"
      >
        <a-form-item name="username" label="用户名">
          <a-input
            v-model:value="loginForm.username"
            size="large"
            prefix-icon="user"
          >
            <template #prefix>
              <UserOutlined />
            </template>
          </a-input>
        </a-form-item>
        
        <a-form-item name="password" label="密码">
          <a-input-password
            v-model:value="loginForm.password"
            size="large"
          >
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input-password>
        </a-form-item>
        
        <a-form-item>
          <div class="form-actions">
            <a-checkbox v-model:checked="rememberMe">记住我</a-checkbox>
            <a href="#" class="forgot-password">忘记密码？</a>
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
            登录
          </a-button>
        </a-form-item>
      </a-form>
      
      <!-- 快速登录（演示用） -->
      <div v-show="activeTab === 'login'" class="quick-login">
        <a-divider>快速登录（演示）</a-divider>
        <a-space>
          <a-button @click="quickLogin('manufacturer')">制造商</a-button>
          <a-button @click="quickLogin('automaker')">车企</a-button>
          <a-button @click="quickLogin('aftersale')">售后中心</a-button>
        </a-space>
      </div>
      
      <!-- 注册表单 -->
      <a-form
        v-show="activeTab === 'register'"
        :model="registerForm"
        :rules="registerRules"
        @finish="handleRegister"
        layout="vertical"
        class="register-form"
      >
        <a-form-item name="username" label="用户名">
          <a-input
            v-model:value="registerForm.username"
            size="large"
          >
            <template #prefix>
              <UserOutlined />
            </template>
          </a-input>
          <span class="field-hint">至少3个字符</span>
        </a-form-item>
        
        <a-form-item name="password" label="密码">
          <a-input-password
            v-model:value="registerForm.password"
            size="large"
          >
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input-password>
          <span class="field-hint">至少6位字符</span>
        </a-form-item>
        
        <a-form-item name="confirmPassword" label="确认密码">
          <a-input-password
            v-model:value="registerForm.confirmPassword"
            size="large"
          >
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input-password>
          <span class="field-hint">需与密码一致</span>
        </a-form-item>
        
        <a-form-item name="org" label="组织">
          <a-select
            v-model:value="registerForm.org"
            size="large"
          >
            <a-select-option value="Org1MSP">制造商组织</a-select-option>
            <a-select-option value="Org2MSP">车企组织</a-select-option>
            <a-select-option value="Org3MSP">售后组织</a-select-option>
          </a-select>
        </a-form-item>
        
        <a-form-item name="role" label="角色">
          <a-select
            v-model:value="registerForm.role"
            size="large"
          >
            <a-select-option value="manufacturer">制造商</a-select-option>
            <a-select-option value="automaker">车企</a-select-option>
            <a-select-option value="aftersale">售后中心</a-select-option>
          </a-select>
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
            注册
          </a-button>
        </a-form-item>
      </a-form>
      
      <!-- 底部信息 -->
      <div class="login-footer">
        <p>© 2024 汽车零部件区块链管理系统 All Rights Reserved</p>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { message } from 'ant-design-vue'
import { useAuthStore } from '../stores/auth.store'
import { UserRole } from '../types'

// ==================== 组合式API ====================

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

// ==================== 响应式状态 ====================

// 当前标签页
const activeTab = ref('login')

// 加载状态
const loading = ref(false)

// 记住我
const rememberMe = ref(false)

// 主题状态
const isDark = ref(false)

// 切换主题
const toggleTheme = () => {
  isDark.value = !isDark.value
  const html = document.documentElement
  const newTheme = isDark.value ? 'dark' : 'light'
  html.setAttribute('data-theme', newTheme)
  localStorage.setItem('theme', newTheme)
}

// 初始化主题
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

// 登录表单
const loginForm = reactive({
  username: '',
  password: ''
})

// 注册表单
const registerForm = reactive({
  username: '',
  password: '',
  confirmPassword: '',
  org: '',
  role: '' as UserRole | ''
})

// 确认密码验证
const validateConfirmPassword = async (_rule: any, value: string) => {
  if (value !== registerForm.password) {
    return Promise.reject('两次输入的密码不一致')
  }
  return Promise.resolve()
}

// 登录表单验证规则
const loginRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码至少6位', trigger: 'blur' }
  ]
}

// 注册表单验证规则
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

// ==================== 方法 ====================

/**
 * 处理登录
 */
async function handleLogin() {
  loading.value = true
  
  try {
    await authStore.login(loginForm.username, loginForm.password)
    message.success('登录成功')
    
    // 跳转到之前的页面或首页
    const redirect = (route.query.redirect as string) || '/'
    router.push(redirect)
  } catch (error: any) {
    message.error(error.message || '登录失败')
  } finally {
    loading.value = false
  }
}

/**
 * 处理注册
 */
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

/**
 * 快速登录（演示用）
 */
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
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  position: relative;
  overflow: hidden;
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

.background-decoration {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.circle {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
}

.circle-1 {
  width: 400px;
  height: 400px;
  top: -100px;
  left: -100px;
}

.circle-2 {
  width: 300px;
  height: 300px;
  bottom: -50px;
  right: -50px;
}

.circle-3 {
  width: 200px;
  height: 200px;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.login-card {
  width: 450px;
  border-radius: 16px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  z-index: 1;
  position: relative;
}

.theme-toggle {
  position: absolute;
  top: 16px;
  right: 16px;
  width: 40px;
  height: 40px;
  border: 2px solid #e0e0e0;
  border-radius: 50%;
  background: #ffffff;
  color: #666;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  z-index: 10;
}

.theme-toggle:hover {
  border-color: #1890ff;
  color: #1890ff;
  transform: rotate(15deg);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.25);
}

.theme-icon {
  width: 20px;
  height: 20px;
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.logo {
  width: 64px;
  height: 64px;
  margin-bottom: 16px;
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.title {
  font-size: 24px;
  font-weight: bold;
  color: var(--text-color);
  margin-bottom: 8px;
  text-align: center;
}

.subtitle {
  font-size: 14px;
  color: var(--text-color-secondary);
  text-align: center;
}

.mode-switch {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-bottom: 24px;
}

.mode-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 32px;
  border: 2px solid #e0e0e0;
  border-radius: 12px;
  background: #ffffff;
  color: #666;
  font-size: 16px;
  font-weight: 500;
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.mode-btn:hover {
  border-color: #1890ff;
  color: #1890ff;
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(24, 144, 255, 0.25);
}

.mode-btn.active {
  border-color: #1890ff;
  background: linear-gradient(135deg, #1890ff 0%, #40a9ff 100%);
  color: #ffffff;
  box-shadow: 0 6px 20px rgba(24, 144, 255, 0.4);
}

.mode-btn.active:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(24, 144, 255, 0.5);
}

.mode-icon {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
}

.login-form,
.register-form {
  margin-top: 24px;
}

.login-form :deep(.ant-form-item),
.register-form :deep(.ant-form-item) {
  text-align: center;
  margin-bottom: 28px;
}

.login-form :deep(.ant-form-item-label),
.register-form :deep(.ant-form-item-label) {
  text-align: center;
  padding-bottom: 10px;
}

.login-form :deep(.ant-form-item-label > label),
.register-form :deep(.ant-form-item-label > label) {
  justify-content: center;
  display: flex;
  font-weight: 500;
  color: #333;
  font-size: 14px;
}

.field-hint {
  display: block;
  font-size: 12px;
  color: #999;
  margin-top: 10px;
  margin-bottom: 4px;
  text-align: center;
  line-height: 1.5;
  min-height: 18px;
}

.submit-btn {
  height: 48px;
  width: calc(60%);
  margin: 0 auto;
  display: block;
  border: none;
  border-radius: 12px;
  background: linear-gradient(135deg, #1890ff 0%, #40a9ff 100%);
  box-shadow: 0 4px 15px rgba(24, 144, 255, 0.35);
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 2px;
  transition: all 0.3s ease;
}

.submit-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(24, 144, 255, 0.5);
  background: linear-gradient(135deg, #40a9ff 0%, #69c0ff 100%);
}

.submit-btn:active {
  transform: translateY(0);
  box-shadow: 0 2px 10px rgba(24, 144, 255, 0.3);
}

.form-actions {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
}

.forgot-password {
  color: #1890ff;
  transition: color 0.3s ease;
}

.forgot-password:hover {
  color: #40a9ff;
}

.quick-login {
  margin-top: 20px;
  text-align: center;
}

.quick-login :deep(.ant-divider) {
  margin: 16px 0;
}

.quick-login :deep(.ant-divider-inner-text) {
  color: #999;
  font-size: 13px;
}

.quick-login :deep(.ant-space) {
  justify-content: center;
  width: 100%;
  flex-wrap: nowrap;
  gap: 12px;
  display: flex;
  align-items: stretch;
}

.quick-login :deep(.ant-space-item) {
  flex: 1;
  display: flex;
}

.quick-login :deep(.ant-space .ant-btn) {
  border-radius: 10px;
  border: 2px solid #e8e8e8;
  background: #fff;
  color: #666;
  font-weight: 500;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
  padding: 8px 16px;
  height: 40px;
  width: 100%;
  min-width: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.quick-login :deep(.ant-space .ant-btn:hover) {
  border-color: #1890ff;
  color: #1890ff;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.2);
}

.login-footer {
  text-align: center;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid var(--border-color-secondary);
  color: var(--text-color-secondary);
  font-size: 12px;
}

:deep(.ant-input-affix-wrapper) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  border-radius: 10px !important;
  border: 2px solid #e8e8e8 !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06) !important;
  transition: all 0.3s ease !important;
  padding: 10px 16px !important;
  background: #fff !important;
  display: flex !important;
  align-items: center !important;
}

:deep(.ant-input-affix-wrapper:hover) {
  border-color: #1890ff !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.15) !important;
}

:deep(.ant-input-affix-wrapper-focused) {
  border-color: #1890ff !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.2) !important;
}

:deep(.ant-input-affix-wrapper .ant-input) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  border: none !important;
  box-shadow: none !important;
  padding: 0 !important;
  background: transparent !important;
  text-align: center !important;
  flex: 1 !important;
}

:deep(.ant-input) {
  text-align: center !important;
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

:deep(.ant-input-affix-wrapper .ant-input-prefix) {
  margin-right: 8px !important;
  color: #999 !important;
}

:deep(.ant-input-affix-wrapper .ant-input-suffix) {
  margin-left: 8px !important;
  color: #999 !important;
}

:deep(.ant-form-item-label) {
  text-align: center !important;
}

:deep(.ant-form-item-label > label) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  font-weight: 500;
  color: #333;
}

:deep(.ant-form-item-label::after) {
  content: '' !important;
}

:deep(.ant-form-item) {
  margin-bottom: 16px !important;
}

:deep(.ant-select-selector) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  border-radius: 10px !important;
  border: 2px solid #e8e8e8 !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06) !important;
  transition: all 0.3s ease !important;
  padding: 10px 16px !important;
  height: auto !important;
  min-height: 44px !important;
  background: #fff !important;
}

:deep(.ant-select-selection-item) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  text-align: center !important;
  justify-content: center !important;
  width: 100% !important;
  font-weight: 500;
  color: #333;
}

:deep(.ant-select-selection-placeholder) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  text-align: center !important;
  justify-content: center !important;
  left: 0 !important;
  right: 0 !important;
  width: 100% !important;
  color: #999;
}

:deep(.ant-select .ant-select-selector .ant-select-selection-item) {
  text-align: center !important;
}

:deep(.ant-select .ant-select-selector) {
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

:deep(.ant-select-selection-search) {
  left: 50% !important;
  transform: translateX(-50%) !important;
  position: absolute !important;
}

:deep(.ant-select-arrow) {
  right: 12px !important;
  color: #999;
}

:deep(.ant-select-selector:hover) {
  border-color: #1890ff !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.15) !important;
}

:deep(.ant-select-focused .ant-select-selector) {
  border-color: #1890ff !important;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.2) !important;
}

:deep(.ant-select-selection-search-input) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

:deep(.ant-select) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  width: 100%;
}

:deep(.ant-select-selection-item) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

:deep(.ant-select-item) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
  padding: 10px 16px;
  text-align: center;
  transition: all 0.2s ease;
}

:deep(.ant-select-item-option-active) {
  background: #e6f7ff !important;
}

:deep(.ant-select-item-option-selected) {
  background: #1890ff !important;
  color: #fff !important;
  font-weight: 500;
}

:deep(.ant-select-dropdown) {
  border-radius: 10px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  padding: 4px 0;
}

:deep(.ant-tabs-tab) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

:deep(.ant-btn) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

:deep(.ant-checkbox-wrapper) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

:deep(.ant-checkbox-inner) {
  border-radius: 4px;
}

:deep(.ant-divider-inner-text) {
  font-family: "Microsoft YaHei", "微软雅黑", sans-serif;
}

/* 深色主题 */
[data-theme='dark'] .login-card {
  background: var(--bg-color-secondary);
}

[data-theme='dark'] .theme-toggle {
  background: #1f1f1f;
  border-color: #434343;
  color: #a0a0a0;
}

[data-theme='dark'] .theme-toggle:hover {
  border-color: #1890ff;
  color: #40a9ff;
}

[data-theme='dark'] .login-footer {
  border-top-color: var(--border-color);
}

[data-theme='dark'] .mode-btn {
  background: #1f1f1f;
  border-color: #434343;
  color: #a0a0a0;
}

[data-theme='dark'] .mode-btn:hover {
  border-color: #1890ff;
  color: #40a9ff;
}

[data-theme='dark'] .mode-btn.active {
  background: linear-gradient(135deg, #1890ff 0%, #40a9ff 100%);
  border-color: #1890ff;
  color: #ffffff;
}

[data-theme='dark'] .login-form :deep(.ant-form-item-label > label),
[data-theme='dark'] .register-form :deep(.ant-form-item-label > label) {
  color: #e0e0e0;
}

[data-theme='dark'] :deep(.ant-input-affix-wrapper) {
  background: #1f1f1f !important;
  border-color: #434343 !important;
}

[data-theme='dark'] :deep(.ant-input-affix-wrapper .ant-input) {
  background: transparent !important;
  color: #e0e0e0 !important;
}

[data-theme='dark'] :deep(.ant-input-affix-wrapper:hover) {
  border-color: #1890ff !important;
}

[data-theme='dark'] :deep(.ant-input-affix-wrapper-focused) {
  border-color: #1890ff !important;
}

[data-theme='dark'] :deep(.ant-select-selector) {
  background: #1f1f1f !important;
  border-color: #434343 !important;
  color: #e0e0e0 !important;
}

[data-theme='dark'] :deep(.ant-select-selector:hover) {
  border-color: #1890ff !important;
}

[data-theme='dark'] :deep(.ant-select-focused .ant-select-selector) {
  border-color: #1890ff !important;
}

[data-theme='dark'] :deep(.ant-select-selection-item) {
  color: #e0e0e0 !important;
}

[data-theme='dark'] :deep(.ant-select-selection-placeholder) {
  color: #666 !important;
}

[data-theme='dark'] :deep(.ant-select-arrow) {
  color: #666 !important;
}

[data-theme='dark'] :deep(.ant-select-item) {
  color: #e0e0e0;
}

[data-theme='dark'] :deep(.ant-select-item-option-active) {
  background: #2a2a2a !important;
}

[data-theme='dark'] :deep(.ant-select-item-option-selected) {
  background: #1890ff !important;
  color: #fff !important;
}

[data-theme='dark'] :deep(.ant-select-dropdown) {
  background: #1f1f1f;
  border: 1px solid #434343;
}

[data-theme='dark'] .quick-login :deep(.ant-space .ant-btn) {
  background: #1f1f1f;
  border-color: #434343;
  color: #a0a0a0;
}

[data-theme='dark'] .quick-login :deep(.ant-space .ant-btn:hover) {
  border-color: #1890ff;
  color: #40a9ff;
}

[data-theme='dark'] .field-hint {
  color: #777;
}

[data-theme='dark'] .forgot-password {
  color: #40a9ff;
}

[data-theme='dark'] .forgot-password:hover {
  color: #69c0ff;
}
</style>
