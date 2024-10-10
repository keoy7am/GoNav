import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import LandingPage from '../views/LandingPage.vue'
import store from '../store'
import Settings from '../views/Settings.vue'
import PublicBookmarks from '../views/PublicBookmarks.vue'

const routes = [
  {
    path: '/',
    name: 'LandingPage',
    component: LandingPage
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
    meta: { requiresAuth: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings,
    meta: { requiresAuth: true }
  },
  {
    path: '/share/:publicId',
    name: 'PublicBookmarks',
    component: PublicBookmarks
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const isLoggedIn = store.getters.isLoggedIn
  if (to.matched.some(record => record.meta.requiresAuth) && !isLoggedIn) {
    next('/')
  } else if (to.path === '/' && isLoggedIn) {
    next('/home')
  } else {
    next()
  }
})

export default router