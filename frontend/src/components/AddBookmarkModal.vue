<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center" id="my-modal">
    <div :class="['bg-white dark:bg-gray-800 p-5 rounded-lg', isFullScreen ? 'w-full h-full' : 'w-96']">
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white">{{ t('bookmark.add') }}</h3>
        <button @click="toggleFullScreen" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
          <font-awesome-icon :icon="isFullScreen ? 'compress' : 'expand'" />
        </button>
      </div>
      <div class="mt-3 text-center">
        <div class="mt-2 px-7 py-3">
          <form @submit.prevent="confirmAddBookmark">
            <input type="text" v-model="title" :placeholder="t('bookmark.title')" 
                   class="mb-3 px-3 py-2 border rounded-md w-full dark:bg-gray-700 dark:text-white" required>
            <div class="mb-3">
              <input type="url" v-model="url" :placeholder="t('bookmark.url')" 
                     @blur="validateUrl"
                     class="mb-1 px-3 py-2 border rounded-md w-full dark:bg-gray-700 dark:text-white" required>
              <p v-if="urlError" class="text-red-500 text-sm">{{ urlError }}</p>
            </div>
            <textarea v-model="description" :placeholder="t('bookmark.description')" maxlength="50"
                      class="mb-3 px-3 py-2 border rounded-md w-full dark:bg-gray-700 dark:text-white"
                      rows="3"></textarea>
            <div class="mb-3">
              <select v-model.number="categoryId" class="px-3 py-2 border rounded-md w-full dark:bg-gray-700 dark:text-white">
                <option :value="null">{{ t('common.select') }}</option>
                <option v-for="category in categories" :key="category.ID" :value="category.ID">
                  {{ category.Name }}
                </option>
              </select>
              <button @click.prevent="showAddCategoryModal = true" 
                      class="mt-2 px-3 py-1 bg-green-500 text-white text-sm rounded-md hover:bg-green-600">
                {{ t('category.add') }}
              </button>
            </div>
            <div class="mb-3">
              <label for="weight" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1 text-left">
                {{ t('bookmark.weight') }} <small>{{ t('common.weightDesc') }}</small>
              </label>
              <input type="number" id="weight" v-model="weight" 
                     class="px-3 py-2 border rounded-md w-full dark:bg-gray-700 dark:text-white">
            </div>
            <div class="mb-3 flex items-center">
              <input type="checkbox" id="isPrivate" v-model="isPrivate" class="mr-2">
              <label for="isPrivate" class="text-gray-700 dark:text-gray-300">{{ t('bookmark.isPrivate') }}</label>
            </div>
            <button type="submit" 
                    class="px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-300">
              {{ t('common.add') }}
            </button>
          </form>
        </div>
      </div>
      <div class="items-center px-7 py-3">
        <button @click="$emit('close')" 
                class="px-4 py-2 bg-gray-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-300">
          {{ t('common.close') }}
        </button>
      </div>
    </div>

    <!-- 新增分類的Modal -->
    <div v-if="showAddCategoryModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center">
      <div class="p-5 border w-96 shadow-lg rounded-md bg-white dark:bg-gray-800">
        <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white mb-4">{{ t('category.add') }}</h3>
        <form @submit.prevent="confirmAddCategory">
          <div class="mb-3">
            <label for="categoryName" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              {{ t('category.name') }}
            </label>
            <input type="text" id="categoryName" v-model="newCategory.name" required
                   class="px-3 py-2 border rounded-md w-full dark:bg-gray-700 dark:text-white">
          </div>
          <div class="mb-3">
            <label for="categoryIcon" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              {{ t('category.icon') }}
            </label>
            <div class="flex items-center">
              <input type="text" id="categoryIcon" v-model="newCategory.icon"
                     class="px-3 py-2 border rounded-md w-full dark:bg-gray-700 dark:text-white mr-2"
                     placeholder="圖標類名" readonly>
              <IconPicker @select="selectIcon" />
            </div>
            <div v-if="newCategory.icon" class="mt-2">
              <i :class="newCategory.icon" class="text-2xl text-gray-700 dark:text-gray-300"></i>
            </div>
          </div>
          <div class="mb-3">
            <label for="categoryWeight" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
              {{ t('category.weight') }} {{ t('common.weightDesc') }}
            </label>
            <input type="number" id="categoryWeight" v-model="newCategory.weight"
                   class="px-3 py-2 border rounded-md w-full dark:bg-gray-700 dark:text-white">
          </div>
          <div class="mb-3 flex items-center">
            <input type="checkbox" id="categoryIsPrivate" v-model="newCategory.isPrivate" class="mr-2">
            <label for="categoryIsPrivate" class="text-gray-700 dark:text-gray-300">{{ t('category.isPrivate') }}</label>
          </div>
          <button type="submit" 
                  class="mb-2 px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-300">
            {{ t('category.add') }}
          </button>
          <button @click.prevent="closeAddCategoryModal" 
                  class="px-4 py-2 bg-gray-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-300">
            {{ t('common.close') }}
          </button>
        </form>
      </div>
    </div>

    <!-- 確認新增書籤的Modal -->
    <div v-if="showConfirmBookmarkModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center">
      <div class="p-5 border w-96 shadow-lg rounded-md bg-white dark:bg-gray-800">
        <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white mb-4">{{ t('common.confirmAdd') }}</h3>
        <div class="mb-4">
          <p><strong>{{ t('bookmark.title') }}:</strong> {{ title }}</p>
          <p><strong>{{ t('bookmark.url') }}:</strong> {{ url }}</p>
          <p><strong>{{ t('bookmark.description') }}:</strong> {{ description }}</p>
          <p><strong>{{ t('bookmark.category') }}:</strong> {{ selectedCategoryName }}</p>
          <p><strong>{{ t('bookmark.weight') }}:</strong> {{ weight }}</p>
          <p><strong>{{ t('bookmark.isPrivate') }}:</strong> {{ isPrivate ? t('common.yes') : t('common.no') }}</p>
        </div>
        <button @click="addBookmark" 
                class="mb-2 px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-300">
          {{ t('common.confirm') }}
        </button>
        <button @click="showConfirmBookmarkModal = false" 
                class="px-4 py-2 bg-gray-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-300">
          {{ t('common.cancel') }}
        </button>
      </div>
    </div>

    <!-- 確認新增分類的Modal -->
    <div v-if="showConfirmCategoryModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center">
      <div class="p-5 border w-96 shadow-lg rounded-md bg-white dark:bg-gray-800">
        <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white mb-4">{{ t('common.confirmAdd') }}</h3>
        <div class="mb-4">
          <p><strong>{{ t('category.name') }}:</strong> {{ newCategory.name }}</p>
          <p><strong>{{ t('category.icon') }}:</strong> {{ newCategory.icon || t('common.none') }}</p>
          <p><strong>{{ t('category.weight') }}:</strong> {{ newCategory.weight }}</p>
          <p><strong>{{ t('category.isPrivate') }}:</strong> {{ newCategory.isPrivate ? t('common.yes') : t('common.no') }}</p>
        </div>
        <button @click="addCategory" 
                class="mb-2 px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-300">
          {{ t('common.confirm') }}
        </button>
        <button @click="showConfirmCategoryModal = false" 
                class="px-4 py-2 bg-gray-500 text-white text-base font-medium rounded-md w-full shadow-sm hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-300">
          {{ t('common.cancel') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'
import IconPicker from './IconPicker.vue'
import { useToast } from "vue-toastification"

export default {
  name: 'AddBookmarkModal',
  components: {
    IconPicker
  },
  emits: ['close', 'bookmarkAdded'],  // Add bookmarkAdded event
  setup(props, { emit }) {
    const { t } = useI18n()
    const store = useStore()
    const toast = useToast()
    const title = ref('')
    const url = ref('')
    const description = ref('')
    const isPrivate = ref(false)
    const categoryId = ref(null)
    const weight = ref(0)
    const showAddCategoryModal = ref(false)
    const showConfirmBookmarkModal = ref(false)
    const showConfirmCategoryModal = ref(false)
    const newCategory = ref({
      name: '',
      icon: '',
      isPrivate: false,
      weight: 0
    })
    const isFullScreen = ref(false)
    const urlError = ref('')

    const categories = computed(() => store.state.categories)

    const selectedCategoryName = computed(() => {
      if (categoryId.value === null) return t('common.select')
      const category = categories.value.find(c => c.ID === categoryId.value)
      return category ? category.Name : t('common.select')
    })

    onMounted(async () => {
      if (categories.value.length === 0) {
        await store.dispatch('fetchCategories')
      }
    })

    const getCategoryName = (id) => {
      const category = categories.value.find(c => c.id === id)
      return category ? category.name : t('common.select')
    }

    const validateUrl = () => {
      const urlPattern = /^https?:\/\//i
      if (!urlPattern.test(url.value)) {
        urlError.value = t('bookmark.urlNotValid')
        return false
      }
      urlError.value = ''
      return true
    }

    const confirmAddBookmark = async () => {
      if (await validateUrl()) {
        showConfirmBookmarkModal.value = true
      }
    }

    const addBookmark = async () => {
      try {
        await store.dispatch('createBookmark', {
          title: title.value,
          url: url.value,
          description: description.value.replace(/\n/g, '\\n'), // Convert newline characters to \n
          isPrivate: isPrivate.value,
          categoryId: categoryId.value,
          weight: weight.value
        })
        showConfirmBookmarkModal.value = false
        emit('bookmarkAdded')
        toast.success(t('bookmark.addSuccess'))
        emit('close')
      } catch (error) {
        console.error('Error adding bookmark:', error)
        toast.error(t('bookmark.addError'))
      }
    }

    const confirmAddCategory = (event) => {
      event.preventDefault()
      showConfirmCategoryModal.value = true
    }

    const closeAddCategoryModal = () => {
      showAddCategoryModal.value = false
      // 重置新分類表單
      newCategory.value = {
        name: '',
        icon: '',
        isPrivate: false,
        weight: 0
      }
    }

    const addCategory = async () => {
      try {
        const createdCategory = await store.dispatch('createCategory', newCategory.value)
        categoryId.value = createdCategory.id
        showConfirmCategoryModal.value = false
        closeAddCategoryModal()
        toast.success(t('category.addSuccess'))
      } catch (error) {
        console.error('Error adding category:', error)
        toast.error(t('category.addError'))
      }
    }

    const selectIcon = (icon) => {
      newCategory.value.icon = icon
    }

    const toggleFullScreen = () => {
      isFullScreen.value = !isFullScreen.value
    }

    return {
      title,
      url,
      description,
      isPrivate,
      categoryId,
      weight,
      categories,
      showAddCategoryModal,
      showConfirmBookmarkModal,
      showConfirmCategoryModal,
      newCategory,
      confirmAddBookmark,
      addBookmark,
      confirmAddCategory,
      closeAddCategoryModal,
      addCategory,
      getCategoryName,
      selectIcon,
      isFullScreen,
      toggleFullScreen,
      urlError,
      validateUrl,
      selectedCategoryName,
      t
    }
  }
}
</script>