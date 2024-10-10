<template>
  <div class="flex flex-col min-h-screen bg-primary-50 dark:bg-neutral-900 text-neutral-800 dark:text-neutral-200">
    <nav class="sticky top-0 z-10 bg-white dark:bg-neutral-800 shadow-md">
      <div class="w-full px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center cursor-pointer" @click="navigateHome">
            <BookmarkIcon class="h-8 w-8 text-primary-600 dark:text-primary-400" />
            <span class="ml-2 text-xl font-semibold text-primary-900 dark:text-primary-100">{{ t('nav.siteTitle') }}</span>
          </div>
          <div class="flex items-center space-x-4">
            <div v-if="isLoggedIn" class="relative">
              <input type="text" :placeholder="t('search')" 
                     class="w-64 px-4 py-2 rounded-lg border border-primary-300 dark:border-neutral-600 bg-white dark:bg-neutral-700 text-neutral-900 dark:text-neutral-200 focus:outline-none focus:ring-2 focus:ring-primary-500">
            </div>
            <div class="relative" v-if="isLoggedIn">
              <button @click="toggleUserMenu" class="flex items-center space-x-2 focus:outline-none">
                <img class="h-8 w-8 rounded-full" src="https://via.placeholder.com/32" alt="User avatar">
                <span class="text-neutral-700 dark:text-neutral-300">{{ username }}</span>
              </button>
              <div v-if="showUserMenu" class="absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white dark:bg-neutral-800 ring-1 ring-black ring-opacity-5">
                <router-link to="/settings" class="block px-4 py-2 text-sm text-neutral-700 dark:text-neutral-300 hover:bg-primary-100 dark:hover:bg-neutral-700">{{ t('nav.settings') }}</router-link>
                <a @click="logout" class="block px-4 py-2 text-sm text-neutral-700 dark:text-neutral-300 hover:bg-primary-100 dark:hover:bg-neutral-700 cursor-pointer">{{ t('nav.logout') }}</a>
              </div>
            </div>
            <template v-else>
              <router-link to="/home" class="nav-link">{{ t('nav.home') }}</router-link>
              <router-link to="/login" class="nav-link">{{ t('nav.login') }}</router-link>
              <router-link to="/register" class="nav-link">{{ t('nav.register') }}</router-link>
            </template>
          </div>
        </div>
      </div>
    </nav>

    <main class="flex-grow overflow-hidden">
      <router-view></router-view>
    </main>

    <!-- 主題切換按鈕 -->
    <div class="fixed bottom-4 right-4 space-y-4">
      <ThemeToggle />
      <button v-if="isLoggedIn" @click="openAddBookmarkModal" 
              class="p-2 rounded-full bg-primary-500 text-white shadow-md hover:bg-primary-600 transition-colors duration-200">
        <PlusIcon class="h-6 w-6" />
      </button>
    </div>

    <!-- 新增書籤Modal -->
    <AddBookmarkModal v-if="showAddBookmarkModal" @close="closeAddBookmarkModal" @bookmarkAdded="refreshData" />

    <!-- 語言切換下拉選單 -->
    <div class="fixed bottom-4 left-4">
      <select v-model="currentLocale" @change="changeLanguage" 
              class="px-3 py-1 bg-primary-500 text-white text-sm rounded-md hover:bg-primary-600 focus:outline-none focus:ring-2 focus:ring-primary-300">
        <option value="en">English</option>
        <option value="zh-TW">繁體中文</option>
        <option value="zh-CN">简体中文</option>
      </select>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted, watch } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { useMeta } from 'vue-meta'
import { useI18n } from 'vue-i18n'
import ThemeToggle from '@/components/ThemeToggle.vue'
import BookmarkIcon from '@/components/BookmarkIcon.vue'
import AddBookmarkModal from '@/components/AddBookmarkModal.vue'
import { PlusIcon } from '@heroicons/vue/24/solid'

export default {
  name: 'App',
  components: {
    ThemeToggle,
    BookmarkIcon,
    AddBookmarkModal,
    PlusIcon
  },
  setup() {
    const store = useStore()
    const router = useRouter()
    const showUserMenu = ref(false)
    const showAddBookmarkModal = ref(false)

    const isLoggedIn = computed(() => store.getters.isLoggedIn)
    const username = computed(() => store.state.user?.username || '')

    useMeta({
      title: 'Go Navigator'
    })

    const toggleUserMenu = () => {
      showUserMenu.value = !showUserMenu.value
    }

    const logout = () => {
      store.dispatch('logout')
      router.push('/')
      showUserMenu.value = false
    }

    const openAddBookmarkModal = async () => {
      if (store.state.categories.length === 0) {
        await store.dispatch('fetchCategories')
      }
      showAddBookmarkModal.value = true
    }

    const closeAddBookmarkModal = () => {
      showAddBookmarkModal.value = false
    }

    const refreshData = async () => {
      await store.dispatch('fetchBookmarks')
      await store.dispatch('fetchCategories')
    }

    const { t, locale } = useI18n()
    const currentLocale = ref(locale.value)

    const changeLanguage = () => {
      locale.value = currentLocale.value
    }

    watch(locale, (newLocale) => {
      document.querySelector('html').setAttribute('lang', newLocale)
    })

    const initializeApp = async () => {
      const token = localStorage.getItem('token')
      if (token) {
        store.commit('setToken', token)
        try {
          await store.dispatch('fetchUser')
          if (isLoggedIn.value) {
            await store.dispatch('fetchCategories')
          }
        } catch (error) {
          console.error('Error fetching user data:', error)
          store.dispatch('logout')
        }
      }
    }

    onMounted(async () => {
      await initializeApp()
      
      const savedTheme = localStorage.getItem('theme')
      if (savedTheme === null || savedTheme === 'dark') {
        document.documentElement.classList.add('dark')
      }
      
      // Set initial language
      const savedLanguage = localStorage.getItem('language')
      if (savedLanguage) {
        locale.value = savedLanguage
        currentLocale.value = savedLanguage
      } else {
        locale.value = 'en'
        currentLocale.value = 'en'
      }
    })

    // Watch for language changes and save to localStorage
    watch(currentLocale, (newLocale) => {
      localStorage.setItem('language', newLocale)
    })

    // Watch for route changes to ensure user menu is closed
    watch(() => router.currentRoute.value, () => {
      showUserMenu.value = false
    })

    const navigateHome = () => {
      if (isLoggedIn.value) {
        router.push('/home')
      } else {
        router.push('/')
      }
    }

    return {
      isLoggedIn,
      username,
      showUserMenu,
      showAddBookmarkModal,
      toggleUserMenu,
      logout,
      openAddBookmarkModal,
      closeAddBookmarkModal,
      refreshData,
      t,
      currentLocale,
      changeLanguage,
      navigateHome
    }
  }
}
</script>

<style>
/* Other styles remain unchanged */

/* Add the following styles */
html, body {
  height: 100%;
  overflow: hidden;
}

#app {
  height: 100%;
  display: flex;
  flex-direction: column;
}
</style>