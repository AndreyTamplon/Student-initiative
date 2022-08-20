import { createApp } from 'vue'
import App from './App.vue'
import { createVuestic } from 'vuestic-ui' // <-
import 'vuestic-ui/dist/vuestic-ui.css'
import router from "@/router/router"; // <-
import store from "@/store/store";


const app = createApp(App)
export const BASE_CONTENT_URL = "BASE_CONTENT_URL_PLACEHOLDER"
export const BASE_AUTH_URL = "BASE_AUTH_URL_PLACEHOLDER"
app.use(createVuestic())
app.use(router)
app.use(store)
app.mount('#app')

