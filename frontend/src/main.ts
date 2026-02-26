/**
 * 应用主入口文件
 * 配置Vue应用、插件、全局组件等
 */

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import Antd from 'ant-design-vue'
import { ConfigProvider } from 'ant-design-vue'

// 导入样式
import './styles/global.css'
import './styles/date-picker.css'

// 导入Ant Design Vue图标
import * as Icons from '@ant-design/icons-vue'

// 创建Vue应用实例
const app = createApp(App)

// 创建Pinia实例
const pinia = createPinia()

// 注册所有图标组件
Object.keys(Icons).forEach((key) => {
  if (key !== 'default') {
    app.component(key, (Icons as any)[key])
  }
})

// 使用插件
app.use(router)
app.use(pinia)
app.use(Antd)
app.use(ConfigProvider)

// 挂载应用
app.mount('#app')
