<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-sm w-full space-y-8">
      <div>
        <BookmarkIcon class="mx-auto h-12 w-auto text-blue-500 dark:text-blue-400" />
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900 dark:text-white">
          {{ t('login.title') }}
        </h2>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="login">
        <input type="hidden" name="remember" value="true">
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="username" class="sr-only">{{ t('login.username') }}</label>
            <input id="username" name="username" type="text" required
                   class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 dark:border-gray-700 placeholder-gray-500 dark:placeholder-gray-400 text-gray-900 dark:text-white rounded-t-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-400 dark:focus:border-blue-400 focus:z-10 sm:text-sm bg-white dark:bg-gray-800"
                   :placeholder="t('login.username')" v-model="username">
          </div>
          <div>
            <label for="password" class="sr-only">{{ t('login.password') }}</label>
            <input id="password" name="password" type="password" autocomplete="current-password" required
                   class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 dark:border-gray-700 placeholder-gray-500 dark:placeholder-gray-400 text-gray-900 dark:text-white rounded-b-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-400 dark:focus:border-blue-400 focus:z-10 sm:text-sm bg-white dark:bg-gray-800"
                   :placeholder="t('login.password')" v-model="password">
          </div>
        </div>

        <div>
          <button type="submit"
                  class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
            {{ t('login.submit') }}
          </button>
        </div>
      </form>
      <div class="text-center">
        <router-link to="/register" class="font-medium text-blue-600 hover:text-blue-500 dark:text-blue-400 dark:hover:text-blue-300">
          {{ t('login.registerLink') }}
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import BookmarkIcon from '@/components/BookmarkIcon.vue'
import { useToast } from "vue-toastification"

export default {
  name: 'Login',
  components: {
    BookmarkIcon
  },
  setup() {
    const store = useStore()
    const router = useRouter()
    const { t } = useI18n()
    const username = ref('')
    const password = ref('')
    const toast = useToast()

    const login = async () => {
      try {
        const success = await store.dispatch('login', {
          username: username.value,
          password: password.value
        })
        if (success) {
          router.push('/')
          toast.success(t('login.success'))
        } else {
          // 這裡是關鍵：當登入不成功但沒有拋出錯誤時，我們也應該顯示一個錯誤通知
          toast.error(t('login.failed'))
        }
      } catch (error) {
        console.error('Login error:', error)
        // 這裡處理網絡錯誤或其他異常情況
        toast.error(t('login.error'))
      }
    }

    return {
      username,
      password,
      login,
      t
    }
  }
}
</script>