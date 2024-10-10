<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-2xl font-bold mb-6">{{ t('settings.title') }}</h1>
    
    <div class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6 mb-6">
      <h2 class="text-xl font-semibold mb-4">{{ t('settings.publicBookmarks') }}</h2>
      <div class="flex items-center mb-4">
        <label class="inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="settings.EnablePublicBookmarks" class="sr-only peer" @change="updateSettings">
          <div class="relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
          <span class="ml-3 text-sm font-medium text-gray-900 dark:text-gray-300">{{ t('settings.enablePublicBookmarks') }}</span>
        </label>
      </div>
      <div v-if="settings.EnablePublicBookmarks" class="mt-4">
        <p class="mb-2">{{ t('settings.publicBookmarksUrl') }}</p>
        <div class="flex items-center">
          <input type="text" :value="publicBookmarksUrl" readonly class="flex-grow px-3 py-2 border rounded-l-md dark:bg-gray-700 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500">
          <button @click="copyPublicUrl" class="px-4 py-2 bg-blue-500 text-white rounded-r-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-300">
            {{ t('common.copyUrl') }}
          </button>
        </div>
      </div>
    </div>

    <div class="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6">
      <h2 class="text-xl font-semibold mb-4">{{ t('settings.accessCode') }}</h2>
      <div class="flex items-center mb-4">
        <label class="inline-flex items-center cursor-pointer">
          <input type="checkbox" v-model="settings.EnableAccessCode" class="sr-only peer" @change="updateSettings">
          <div class="relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
          <span class="ml-3 text-sm font-medium text-gray-900 dark:text-gray-300">{{ t('settings.enableAccessCode') }}</span>
        </label>
      </div>
      <div v-if="settings.EnableAccessCode" class="mt-4">
        <div class="mb-3">
          <label for="accessCode" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('settings.accessCode') }}</label>
          <input type="text" id="accessCode" v-model="settings.AccessCode"
                 class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white">
        </div>
        <button @click="updateSettings" class="mt-2 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">
          {{ t('common.save') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'
import { useToast } from "vue-toastification"

export default {
  name: 'Settings',
  setup() {
    const store = useStore()
    const { t } = useI18n()
    const settings = ref({
      EnablePublicBookmarks: false,
      PublicID: '',
      EnableAccessCode: false,
      AccessCode: ''
    })

    const publicBookmarksUrl = computed(() => {
      if (settings.value.EnablePublicBookmarks && settings.value.PublicID) {
        return `${window.location.origin}/share/${settings.value.PublicID}`
      }
      return ''
    })

    const fetchSettings = async () => {
      try {
        const response = await store.dispatch('fetchSettings')
        settings.value = response
      } catch (error) {
        console.error('Error fetching settings:', error)
      }
    }

    const updateSettings = async () => {
      try {
        if (settings.value.EnableAccessCode && !settings.value.AccessCode) {
          toast.warning(t('settings.accessCodeRequired'))
          return
        }

        const updatedSettings = await store.dispatch('updateSettings', settings.value)
        settings.value = updatedSettings
        toast.success(t('settings.settingsSaved'))
      } catch (error) {
        console.error('Error updating settings:', error)
        toast.error(t('settings.errorSavingSettings'))
      }
    }

    const copyPublicUrl = () => {
      navigator.clipboard.writeText(publicBookmarksUrl.value)
        .then(() => toast.success(t('settings.urlCopied')))
        .catch(err => {
          console.error('Error copying text: ', err)
          toast.error(t('settings.urlCopyError'))
        })
    }

    fetchSettings()

    const toast = useToast()

    return {
      settings,
      publicBookmarksUrl,
      updateSettings,
      copyPublicUrl,
      t
    }
  }
}
</script>