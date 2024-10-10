<template>
  <button @click="toggleTheme" class="p-2 rounded-full bg-white dark:bg-gray-800 shadow-md hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors duration-200" :title="t('themeToggle.toggleTheme')">
    <SunIcon v-if="isDark" class="h-6 w-6 text-yellow-400" />
    <MoonIcon v-else class="h-6 w-6 text-gray-700 dark:text-gray-300" />
  </button>
</template>

<script>
import { ref, onMounted, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { SunIcon, MoonIcon } from '@heroicons/vue/24/solid'

export default {
  components: {
    SunIcon,
    MoonIcon
  },
  setup() {
    const { t } = useI18n()
    const isDark = ref(true) // default to true

    const toggleTheme = () => {
      isDark.value = !isDark.value
      localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
      updateTheme()
    }

    const updateTheme = () => {
      if (isDark.value) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    }

    onMounted(() => {
      const savedTheme = localStorage.getItem('theme')
      isDark.value = savedTheme === null ? true : savedTheme === 'dark'
      updateTheme()
    })

    watch(isDark, updateTheme)

    return {
      isDark,
      toggleTheme,
      t
    }
  }
}
</script>