<!--
  应用根组件
  提供路由视图和全局配置
-->

<template>
  <a-config-provider :locale="zhCN">
    <router-view />
  </a-config-provider>
</template>

<script setup lang="ts">
import zhCN from 'ant-design-vue/es/locale/zh_CN'
</script>

<style>
/* 全局样式在 styles/global.css 中定义 */
</style>
