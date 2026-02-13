<!--
  系统设置页面
-->

<template>
  <div class="system-settings">
    <div class="page-header">
      <h2>系统设置</h2>
    </div>
    
    <a-card title="主题设置" :bordered="false">
      <a-form :label-col="{ span: 4 }">
        <a-form-item label="当前主题">
          <a-radio-group v-model:value="theme" @change="handleThemeChange">
            <a-radio-button value="light">浅色主题</a-radio-button>
            <a-radio-button value="dark">深色主题</a-radio-button>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-card>
    
    <a-card title="API配置" :bordered="false" style="margin-top: 16px">
      <a-form :label-col="{ span: 4 }">
        <a-form-item label="API地址">
          <a-input value="http://localhost:8080" disabled />
        </a-form-item>
        <a-form-item label="Fabric通道">
          <a-input value="mychannel" disabled />
        </a-form-item>
        <a-form-item label="链码名称">
          <a-input value="parts" disabled />
        </a-form-item>
      </a-form>
    </a-card>
    
    <a-card title="关于系统" :bordered="false" style="margin-top: 16px">
      <a-descriptions :column="1">
        <a-descriptions-item label="系统名称">汽车零部件区块链管理系统</a-descriptions-item>
        <a-descriptions-item label="版本号">v1.0.0</a-descriptions-item>
        <a-descriptions-item label="技术栈">Vue 3 + TypeScript + Ant Design Vue</a-descriptions-item>
        <a-descriptions-item label="区块链">Hyperledger Fabric</a-descriptions-item>
      </a-descriptions>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAppStore } from '../../stores'

const appStore = useAppStore()
const theme = ref(appStore.theme)

function handleThemeChange() {
  appStore.setTheme(theme.value)
}
</script>

<style scoped>
.system-settings { padding: 0; }
.page-header { margin-bottom: 24px; }
.page-header h2 { font-size: 20px; font-weight: 600; margin: 0; }
</style>
