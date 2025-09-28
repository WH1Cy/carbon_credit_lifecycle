// @ts-ignore
import { createApp } from 'vue'
// @ts-ignore
import { createPinia } from 'pinia'
// @ts-ignore
import Antd from 'ant-design-vue'
// @ts-ignore
import 'ant-design-vue/dist/reset.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Antd)

app.mount('#app')
