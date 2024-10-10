<template>
  <div v-if="needAccessCode" class="flex items-center justify-center h-screen bg-gray-100 dark:bg-gray-900">
    <div class="bg-white dark:bg-gray-800 p-8 rounded-lg shadow-md">
      <h2 class="text-2xl font-bold mb-4 text-gray-800 dark:text-white">{{ t('publicBookmarks.enterAccessCode') }}</h2>
      <input v-model="accessCode" type="text" :placeholder="t('publicBookmarks.accessCode')" 
             class="w-full px-3 py-2 border rounded-md mb-4 bg-white dark:bg-gray-700 text-gray-900 dark:text-white placeholder-gray-500 dark:placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500">
      <button @click="submitAccessCode" class="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600 transition duration-200">
        {{ t('common.submit') }}
      </button>
    </div>
  </div>
  
  <!-- 原有的書籤顯示邏輯 -->
  <div v-else class="flex h-screen bg-primary-50 dark:bg-neutral-900">
    <!-- 左側導航欄 -->
    <div class="w-64 bg-white dark:bg-neutral-800 shadow-md overflow-y-auto">
      <div class="p-4">
        <h2 class="text-lg font-semibold text-primary-800 dark:text-primary-200 mb-4">{{ t('category.title') }}</h2>
        <ul class="space-y-2">
          <li @click="selectCategory(null)" 
              class="cursor-pointer text-neutral-600 dark:text-neutral-400 hover:text-primary-800 dark:hover:text-primary-200 flex items-center justify-between p-2 rounded hover:bg-primary-100 dark:hover:bg-neutral-700"
              :class="{ 'bg-primary-100 dark:bg-neutral-700': selectedCategoryId === null }">
            <span>{{ t('category.all') }}</span>
            <span class="text-sm text-neutral-500">{{ getCategoryBookmarkCount(null) }}</span>
          </li>
          <li v-for="category in categories" :key="category.ID" 
              @click="selectCategory(category.ID)"
              class="cursor-pointer text-neutral-600 dark:text-neutral-400 hover:text-primary-800 dark:hover:text-primary-200 flex items-center justify-between p-2 rounded hover:bg-primary-100 dark:hover:bg-neutral-700"
              :class="{ 'bg-primary-100 dark:bg-neutral-700': selectedCategoryId === category.ID }">
            <div class="flex items-center">
              <i v-if="category.Icon" :class="category.Icon" class="mr-2"></i>
              <span>{{ category.Name }}</span>
            </div>
            <span class="text-sm text-neutral-500">{{ getCategoryBookmarkCount(category.ID) }}</span>
          </li>
        </ul>
      </div>
    </div>

    <!-- 主要內容區 -->
    <div class="flex-1 overflow-auto p-8 bg-primary-50 dark:bg-neutral-900">
      <div class="mb-6 flex justify-between items-center">
        <input type="text" v-model="searchQuery" :placeholder="t('search')" 
               class="w-full max-w-md px-4 py-2 rounded-lg border border-primary-300 dark:border-neutral-600 bg-white dark:bg-neutral-800 text-neutral-800 dark:text-neutral-200 focus:outline-none focus:ring-2 focus:ring-primary-500">
        <div class="flex items-center space-x-4">
          <div class="flex bg-white dark:bg-neutral-800 rounded-md overflow-hidden">
            <button @click="setDisplayMode('block')" 
                    :class="['px-3 py-2 transition-colors duration-200', 
                             displayMode === 'block' ? 'bg-primary-500 text-white' : 'text-neutral-600 dark:text-neutral-400 hover:bg-primary-100 dark:hover:bg-neutral-700']">
              <font-awesome-icon icon="th-large" />
            </button>
            <button @click="setDisplayMode('list')" 
                    :class="['px-3 py-2 transition-colors duration-200', 
                             displayMode === 'list' ? 'bg-primary-500 text-white' : 'text-neutral-600 dark:text-neutral-400 hover:bg-primary-100 dark:hover:bg-neutral-700']">
              <font-awesome-icon icon="th-list" />
            </button>
          </div>
        </div>
      </div>
      <div v-if="filteredBookmarks.length === 0" class="text-center text-gray-500 dark:text-gray-400">
        {{ t('bookmark.noBookmarks') }}
      </div>
      <div v-else :class="{'grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6': displayMode === 'block', 'space-y-4': displayMode === 'list'}">
        <div v-for="bookmark in filteredBookmarks" :key="bookmark.ID" 
             :class="{'bg-white dark:bg-neutral-800 rounded-lg shadow-md p-4 flex flex-col items-start space-y-4 hover:shadow-lg transition-shadow duration-200 relative': displayMode === 'block',
                      'bg-white dark:bg-neutral-800 rounded-lg shadow-md p-4 hover:shadow-lg transition-shadow duration-200 relative': displayMode === 'list'}">
          <div :class="{'flex items-start space-x-4 w-full': displayMode === 'block', 'flex items-center space-x-4 w-full': displayMode === 'list'}">
            <div class="flex-shrink-0">
              <img :src="getIconUrl(bookmark.URL)" alt="Site icon" class="w-8 h-8 rounded">
            </div>
            <div class="flex-grow overflow-hidden">
              <h3 class="text-lg font-semibold text-primary-800 dark:text-primary-200 truncate">{{ bookmark.Title }}</h3>
              <a :href="bookmark.URL" target="_blank" class="text-sm text-primary-600 hover:text-primary-800 dark:text-primary-400 dark:hover:text-primary-300 break-all">
                {{ bookmark.URL }}
              </a>
              <p v-if="displayMode === 'block' || bookmark.expanded" class="mt-1 text-sm text-neutral-600 dark:text-neutral-400 whitespace-pre-wrap">{{ formatDescription(bookmark.Description) }}</p>
            </div>
            <button v-if="displayMode === 'list' && bookmark.Description" @click="toggleExpand(bookmark)" class="text-primary-500 hover:text-primary-600 dark:text-primary-400 dark:hover:text-primary-300">
              <font-awesome-icon :icon="bookmark.expanded ? 'chevron-up' : 'chevron-down'" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'  // 添加 onMounted
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import api from '@/services/api'

export default {
  name: 'PublicBookmarks',
  setup() {
    const { t } = useI18n()
    const route = useRoute()
    const publicId = route.params.publicId

    const bookmarks = ref([])
    const categories = ref([])
    const selectedCategoryId = ref(null)
    const searchQuery = ref('')
    const displayMode = ref('block')

    const needAccessCode = ref(false)
    const accessCode = ref('')

    const fetchPublicBookmarks = async (code = null) => {
      try {
        const response = await api.getPublicBookmarks(publicId, code)
        bookmarks.value = response.data.bookmarks.map(b => ({ ...b, expanded: false }))
        categories.value = response.data.categories
        needAccessCode.value = false
      } catch (error) {
        if (error.response && error.response.status === 401) {
          needAccessCode.value = true
        } else {
          console.error('Error fetching public bookmarks:', error)
        }
      }
    }

    const selectCategory = (categoryId) => {
      selectedCategoryId.value = categoryId
    }

    const filteredBookmarks = computed(() => {
      return bookmarks.value.filter(bookmark => 
        (selectedCategoryId.value === null || bookmark.CategoryID === selectedCategoryId.value) &&
        ((bookmark.Title && bookmark.Title.toLowerCase().includes(searchQuery.value.toLowerCase())) ||
        (bookmark.Description && bookmark.Description.toLowerCase().includes(searchQuery.value.toLowerCase())) ||
        (bookmark.URL && bookmark.URL.toLowerCase().includes(searchQuery.value.toLowerCase())))
      ).sort((a, b) => a.Weight - b.Weight)
    })

    const getCategoryBookmarkCount = (categoryId) => {
      if (categoryId === null) {
        return filteredBookmarks.value.length
      }
      return filteredBookmarks.value.filter(bookmark => bookmark.CategoryID === categoryId).length
    }

    const getIconUrl = (url) => {
      return `https://www.google.com/s2/favicons?domain=${url}&sz=64`
    }

    const setDisplayMode = (mode) => {
      displayMode.value = mode
    }

    const toggleExpand = (bookmark) => {
      bookmark.expanded = !bookmark.expanded
    }

    const formatDescription = (description) => {
      return description ? description.replace(/\\n/g, '\n') : ''
    }

    const submitAccessCode = () => {
      fetchPublicBookmarks(accessCode.value)
    }

    onMounted(() => {
      fetchPublicBookmarks()
    })

    return {
      categories,
      filteredBookmarks,
      selectedCategoryId,
      selectCategory,
      searchQuery,
      displayMode,
      setDisplayMode,
      getCategoryBookmarkCount,
      getIconUrl,
      toggleExpand,
      formatDescription,
      t,
      needAccessCode,
      accessCode,
      submitAccessCode,
    }
  }
}
</script>