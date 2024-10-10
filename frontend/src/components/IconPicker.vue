<template>
  <div>
    <button @click="toggleIconPicker" class="px-3 py-1 bg-blue-500 text-white text-sm rounded-md hover:bg-blue-600">
      {{ t('iconPicker.selectIcon') }}
    </button>
    <div v-if="showIconPicker" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center">
      <div :class="['bg-white dark:bg-gray-800 p-5 rounded-lg', isFullScreen ? 'w-full h-full' : 'w-[32rem] max-h-[80vh]']">
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-medium text-gray-900 dark:text-white">{{ t('iconPicker.selectIcon') }}</h3>
          <button @click="toggleFullScreen" class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
            <font-awesome-icon :icon="isFullScreen ? 'compress' : 'expand'" />
          </button>
        </div>
        <div class="mb-4 flex">
          <input v-model="searchQuery" type="text" :placeholder="t('iconPicker.searchIcon')" 
                 class="flex-grow px-3 py-2 border rounded-l-md dark:bg-gray-700 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500" />
          <button @click="clearSearch" class="px-3 py-2 bg-gray-200 dark:bg-gray-600 rounded-r-md hover:bg-gray-300 dark:hover:bg-gray-500">
            <font-awesome-icon icon="times" />
          </button>
        </div>
        <div class="mb-4 flex flex-wrap">
          <button v-for="category in categories" :key="category" 
                  @click="selectCategory(category)"
                  :class="['mr-2 mb-2 px-3 py-1 rounded-md text-sm', 
                           selectedCategory === category ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-gray-300']">
            {{ category }}
          </button>
        </div>
        <div class="flex-grow overflow-y-auto" :style="{ height: isFullScreen ? 'calc(100vh - 200px)' : '300px' }">
          <div class="grid grid-cols-8 gap-2">
            <button v-for="icon in filteredIcons" :key="icon" @click="selectIcon(icon)" 
                    class="p-2 hover:bg-gray-200 dark:hover:bg-gray-700 rounded flex items-center justify-center">
              <font-awesome-icon :icon="['fas', icon]" class="text-xl text-gray-700 dark:text-gray-300" />
            </button>
          </div>
        </div>
        <div class="mt-4 flex justify-between items-center">
          <span class="text-sm text-gray-500 dark:text-gray-400">
            {{ filteredIcons.length }} {{ t('iconPicker.icons') }}
          </span>
          <button @click="closeIconPicker" class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600">
            {{ t('common.close') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'

export default {
  name: 'IconPicker',
  emits: ['select'],
  setup(props, { emit }) {
    const { t } = useI18n()
    const showIconPicker = ref(false)
    const isFullScreen = ref(false)
    const searchQuery = ref('')
    const selectedCategory = ref('All')

    const categories = ['All', 'Common', 'Weather', 'Transportation', 'Sports', 'Technology', 'Food', 'Others']

    const iconCategories = {
      Common: ['home', 'user', 'cog', 'envelope', 'bell', 'calendar', 'folder', 'file', 'heart', 'star', 'search', 'plus', 'minus', 'times', 'check'],
      Weather: ['sun', 'moon', 'cloud', 'umbrella', 'snowflake', 'wind', 'rainbow', 'temperature-high', 'temperature-low'],
      Transportation: ['car', 'plane', 'ship', 'train', 'bicycle', 'bus', 'taxi', 'truck', 'motorcycle'],
      Sports: ['football-ball', 'basketball-ball', 'baseball-ball', 'volleyball-ball', 'golf-ball', 'table-tennis', 'running', 'swimming-pool', 'dumbbell'],
      Technology: ['laptop', 'desktop', 'mobile-alt', 'keyboard', 'mouse', 'print', 'camera', 'wifi', 'microchip', 'battery-full', 'usb'],
      Food: ['apple-alt', 'carrot', 'coffee', 'pizza-slice', 'hamburger', 'ice-cream', 'birthday-cake', 'beer', 'utensils', 'wine-glass'],
      Others: ['trophy', 'medal', 'crown', 'gem', 'lightbulb', 'fire', 'bolt', 'book', 'map', 'gift', 'music', 'paint-brush', 'camera-retro', 'headphones']
    }

    const allIcons = Object.values(iconCategories).flat()

    const filteredIcons = computed(() => {
      let icons = selectedCategory.value === 'All' ? allIcons : iconCategories[selectedCategory.value]
      if (searchQuery.value) {
        icons = icons.filter(icon => icon.toLowerCase().includes(searchQuery.value.toLowerCase()))
      }
      return icons
    })

    const toggleIconPicker = () => {
      showIconPicker.value = !showIconPicker.value
    }

    const closeIconPicker = () => {
      showIconPicker.value = false
    }

    const selectIcon = (icon) => {
      emit('select', icon)
      closeIconPicker()
    }

    const toggleFullScreen = () => {
      isFullScreen.value = !isFullScreen.value
    }

    const selectCategory = (category) => {
      selectedCategory.value = category
    }

    const clearSearch = () => {
      searchQuery.value = ''
    }

    return {
      showIconPicker,
      isFullScreen,
      searchQuery,
      categories,
      selectedCategory,
      filteredIcons,
      toggleIconPicker,
      closeIconPicker,
      selectIcon,
      toggleFullScreen,
      selectCategory,
      clearSearch,
      t
    }
  }
}
</script>