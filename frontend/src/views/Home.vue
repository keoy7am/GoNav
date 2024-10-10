<template>
  <div class="flex h-screen bg-primary-50 dark:bg-neutral-900">
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
          <li v-for="category in getSortedCategories()" :key="category.ID" 
              @click="selectCategory(category.ID)"
              class="cursor-pointer text-neutral-600 dark:text-neutral-400 hover:text-primary-800 dark:hover:text-primary-200 flex items-center justify-between p-2 rounded hover:bg-primary-100 dark:hover:bg-neutral-700"
              :class="{ 'bg-primary-100 dark:bg-neutral-700': selectedCategoryId === category.ID }">
            <div class="flex items-center">
              <i v-if="category.Icon" :class="category.Icon" class="mr-2"></i>
              <span>{{ category.Name }}</span>
            </div>
            <div class="flex items-center space-x-2">
              <span class="text-sm text-neutral-500">{{ getCategoryBookmarkCount(category.ID) }}</span>
              <div class="flex space-x-2">
                <button @click.stop="editCategory(category)" class="text-blue-500 hover:text-blue-600 dark:text-blue-400 dark:hover:text-blue-300">
                  <font-awesome-icon icon="edit" />
                </button>
                <button @click.stop="deleteCategory(category.ID)" class="text-red-500 hover:text-red-600 dark:text-red-400 dark:hover:text-red-300">
                  <font-awesome-icon icon="trash" />
                </button>
              </div>
            </div>
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
          <!-- 過濾器下拉選單 -->
          <select v-model="privacyFilter" class="px-3 py-2 bg-white dark:bg-neutral-800 border border-primary-300 dark:border-neutral-600 rounded-md text-neutral-800 dark:text-neutral-200 focus:outline-none focus:ring-2 focus:ring-primary-500">
            <option value="all">{{ t('bookmark.filterAll') }}</option>
            <option value="public">{{ t('bookmark.filterPublic') }}</option>
            <option value="private">{{ t('bookmark.filterPrivate') }}</option>
          </select>
          <!-- 顯示模式按鈕 -->
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
          <!-- 編輯和刪除按鈕 -->
          <div class="absolute top-2 right-2 flex space-x-2">
            <button @click="editBookmark(bookmark)" class="text-blue-500 hover:text-blue-600 dark:text-blue-400 dark:hover:text-blue-300">
              <font-awesome-icon icon="edit" />
            </button>
            <button @click="deleteBookmark(bookmark.ID)" class="text-red-500 hover:text-red-600 dark:text-red-400 dark:hover:text-red-300">
              <font-awesome-icon icon="trash" />
            </button>
          </div>
          
          <div :class="{'flex items-start space-x-4 w-full': displayMode === 'block', 'flex items-center space-x-4 w-full': displayMode === 'list'}">
            <div class="flex-shrink-0">
              <img :src="getIconUrl(bookmark.URL)" alt="Site icon" class="w-8 h-8 rounded">
            </div>
            <div class="flex-grow overflow-hidden">
              <div class="flex items-center">
                <h3 class="text-lg font-semibold text-primary-800 dark:text-primary-200 truncate">{{ bookmark.Title }}</h3>
                <!-- 私人書籤圖示 -->
                <font-awesome-icon v-if="bookmark.IsPrivate" icon="lock" class="ml-2 text-gray-500 dark:text-gray-400" :title="t('bookmark.isPrivate')" />
              </div>
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

    <!-- 編輯書籤的Modal -->
    <EditBookmarkModal
      v-if="showEditModal"
      :bookmark="editingBookmark"
      @close="closeEditModal"
      @update="updateBookmark"
    />

    <!-- 編輯分類的Modal -->
    <EditCategoryModal
      v-if="showEditCategoryModal"
      :category="editingCategory"
      @close="closeEditCategoryModal"
      @update="updateCategory"
    />

    <!-- 新增書籤Modal -->
    <AddBookmarkModal 
      v-if="showAddBookmarkModal" 
      @close="closeAddBookmarkModal" 
      @bookmarkAdded="handleBookmarkAdded" 
    />
  </div>
</template>

<script>
import { ref, computed, onMounted, reactive, watch } from 'vue'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'
import { useToast } from "vue-toastification"
import EditBookmarkModal from '@/components/EditBookmarkModal.vue'
import EditCategoryModal from '@/components/EditCategoryModal.vue'
import AddBookmarkModal from '@/components/AddBookmarkModal.vue'

export default {
  name: 'Home',
  components: {
    EditBookmarkModal,
    EditCategoryModal,
    AddBookmarkModal
  },
  setup() {
    const store = useStore()
    const { t } = useI18n()
    const toast = useToast()
    const searchQuery = ref('')
    const selectedCategoryId = ref(null)
    const showEditModal = ref(false)
    const editingBookmark = ref(null)
    const showEditCategoryModal = ref(false)
    const editingCategory = ref(null)
    const displayMode = ref('block')
    const privacyFilter = ref('all')

    const bookmarksMap = reactive(new Map())

    const updateBookmarksMap = () => {
      const storeBookmarks = store.state.bookmarks
      storeBookmarks.forEach(bookmark => {
        if (!bookmarksMap.has(bookmark.ID)) {
          bookmarksMap.set(bookmark.ID, reactive({ ...bookmark, expanded: false }))
        } else {
          Object.assign(bookmarksMap.get(bookmark.ID), bookmark)
        }
      })
      return Array.from(bookmarksMap.values())
    }

    const bookmarks = ref([])

    watch(() => store.state.bookmarks, () => {
      bookmarks.value = updateBookmarksMap()
    }, { immediate: true })

    const categories = computed(() => store.state.categories)

    const filteredBookmarks = computed(() => {
      return bookmarks.value.filter(bookmark => {
        const matchesSearch = (
          (bookmark.Title && bookmark.Title.toLowerCase().includes(searchQuery.value.toLowerCase())) ||
          (bookmark.Description && bookmark.Description.toLowerCase().includes(searchQuery.value.toLowerCase())) ||
          (bookmark.URL && bookmark.URL.toLowerCase().includes(searchQuery.value.toLowerCase()))
        );

        // Check privacy settings
        const matchesPrivacyFilter = 
          privacyFilter.value === 'all' || 
          (privacyFilter.value === 'public' && !bookmark.IsPrivate) || 
          (privacyFilter.value === 'private' && bookmark.IsPrivate);

        // If all categories are selected
        if (selectedCategoryId.value === null) {
          return matchesSearch && matchesPrivacyFilter;
        } else {
          // If a specific category is selected
          return matchesSearch && 
                 bookmark.CategoryID === selectedCategoryId.value && 
                 matchesPrivacyFilter;
        }
      }).sort((a, b) => a.Weight - b.Weight);
    });

    const getSortedCategories = () => {
      // Sort categories by weight
      return categories.value.sort((a, b) => a.Weight - b.Weight)
    }

    const getCategoryBookmarkCount = (categoryId) => {
      if (categoryId === null) {
        // For "All Categories", return the count of all bookmarks that match the current privacy filter
        return bookmarks.value.filter(bookmark => 
          privacyFilter.value === 'all' || 
          (privacyFilter.value === 'public' && !bookmark.IsPrivate) || 
          (privacyFilter.value === 'private' && bookmark.IsPrivate)
        ).length
      }
      // For a specific category, return the count of all bookmarks in that category, ignoring privacy filters
      return bookmarks.value.filter(bookmark => bookmark.CategoryID === categoryId).length
    }

    const getIconUrl = (url) => {
      return `https://www.google.com/s2/favicons?domain=${url}&sz=64`
    }

    const selectCategory = (categoryId) => {
      selectedCategoryId.value = categoryId
    }

    const setDisplayMode = (mode) => {
      displayMode.value = mode
    }

    const editBookmark = (bookmark) => {
      editingBookmark.value = { ...bookmark }
      showEditModal.value = true
    }

    const closeEditModal = () => {
      showEditModal.value = false
      editingBookmark.value = null
    }

    const updateBookmark = async (updatedBookmark) => {
      try {
        await store.dispatch('updateBookmark', { id: updatedBookmark.ID, bookmarkData: updatedBookmark })
        closeEditModal()
        await store.dispatch('fetchBookmarks')
        toast.success(t('bookmark.updateSuccess'))
      } catch (error) {
        console.error('Error updating bookmark:', error)
        toast.error(t('bookmark.updateError'))
      }
    }

    const deleteBookmark = async (bookmarkId) => {
      if (confirm(t('common.confirmDelete'))) {
        try {
          await store.dispatch('deleteBookmark', bookmarkId)
          toast.success(t('bookmark.deleteSuccess'))
        } catch (error) {
          console.error('Error deleting bookmark:', error)
          toast.error(t('bookmark.deleteError'))
        }
      }
    }

    const editCategory = (category) => {
      editingCategory.value = { ...category }
      showEditCategoryModal.value = true
    }

    const closeEditCategoryModal = () => {
      showEditCategoryModal.value = false
      editingCategory.value = null
    }

    const updateCategory = async (updatedCategory) => {
      try {
        await store.dispatch('updateCategory', { id: updatedCategory.ID, categoryData: updatedCategory })
        closeEditCategoryModal()
        toast.success(t('category.updateSuccess'))
      } catch (error) {
        console.error('Error updating category:', error)
        toast.error(t('category.updateError'))
      }
    }

    const deleteCategory = async (categoryId) => {
      if (confirm(t('common.confirmDelete'))) {
        try {
          await store.dispatch('deleteCategory', categoryId)
          if (selectedCategoryId.value === categoryId) {
            selectedCategoryId.value = null
          }
          toast.success(t('category.deleteSuccess'))
        } catch (error) {
          console.error('Error deleting category:', error)
          toast.error(t('category.deleteError'))
        }
      }
    }

    const toggleExpand = (bookmark) => {
      bookmark.expanded = !bookmark.expanded
    }

    const formatDescription = (description) => {
      return description ? description.replace(/\\n/g, '\n') : ''
    }

    const showAddBookmarkModal = ref(false)

    const refreshData = async () => {
      await store.dispatch('fetchBookmarks')
      await store.dispatch('fetchCategories')
      toast.success(t('common.dataRefreshed'))
    }

    const closeAddBookmarkModal = () => {
      showAddBookmarkModal.value = false
    }

    const handleBookmarkAdded = () => {
      refreshData()
      closeAddBookmarkModal()
    }

    onMounted(async () => {
      await refreshData()
    })

    return {
      searchQuery,
      categories,
      filteredBookmarks,
      bookmarks,
      getIconUrl,
      selectedCategoryId,
      selectCategory,
      editBookmark,
      deleteBookmark,
      showEditModal,
      editingBookmark,
      closeEditModal,
      updateBookmark,
      showEditCategoryModal,
      editingCategory,
      editCategory,
      closeEditCategoryModal,
      updateCategory,
      deleteCategory,
      displayMode,
      setDisplayMode,
      toggleExpand,
      privacyFilter,
      getCategoryBookmarkCount,
      getSortedCategories,
      t,
      formatDescription,
      showAddBookmarkModal,
      closeAddBookmarkModal,
      refreshData,
      handleBookmarkAdded
    }
  }
}
</script>

<style scoped>
.home {
  display: flex;
}
.sidebar {
  width: 200px;
  padding: 1rem;
  background-color: #f0f0f0;
}
.main-content {
  flex-grow: 1;
  padding: 1rem;
}
.search-bar {
  margin-bottom: 1rem;
}
.bookmark {
  margin-bottom: 1rem;
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}
.break-all {
  word-break: break-all;
}
</style>