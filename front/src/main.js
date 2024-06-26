import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
const app=createApp(App).use(router).mount('#app')
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}