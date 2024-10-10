import { createApp } from 'vue'
import { createMetaManager } from 'vue-meta'
import App from './App.vue'
import router from './router'
import store from './store'
import './styles/tailwind.css'
import i18n from './i18n'

// Toast
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

// Font Awesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'

library.add(fas)

const app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon)
app.use(router).use(store).use(createMetaManager()).use(i18n)

// Use toast
app.use(Toast, {
  transition: "Vue-Toastification__bounce",
  maxToasts: 20,
  newestOnTop: true
});

app.mount('#app')