import './assets/main.css'

import { createApp, provide } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'

import router from './router'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)

app.provide('hostBacked', 'http://127.0.0.1:8080')

app.mount('#app')
