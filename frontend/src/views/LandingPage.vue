<template>
  <div class="min-h-screen bg-gray-100 dark:bg-gray-900 flex flex-col justify-center items-center px-4 sm:px-6 lg:px-8">
    <div class="max-w-3xl w-full space-y-8 text-center">
      <BookmarkIcon class="mx-auto h-20 w-20 text-blue-500 dark:text-blue-400" />
      <h1 class="text-4xl font-extrabold text-gray-900 dark:text-white sm:text-5xl sm:tracking-tight lg:text-6xl">
        {{ t('landing.welcome') }}
      </h1>
      <p class="mt-5 text-xl text-gray-500 dark:text-gray-300 max-w-2xl mx-auto">
        {{ t('landing.description') }}
      </p>
      <div class="mt-8 flex justify-center">
        <div class="inline-flex rounded-md shadow">
          <router-link to="/login" class="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
            {{ t('landing.getStarted') }}
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useI18n } from 'vue-i18n'
import BookmarkIcon from '@/components/BookmarkIcon.vue'

export default {
  name: 'LandingPage',
  components: {
    BookmarkIcon
  },
  setup() {
    const { t } = useI18n()

    return { t }
  }
}
</script>