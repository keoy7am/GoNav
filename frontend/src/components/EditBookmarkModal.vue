<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center" id="edit-modal">
    <div class="bg-white dark:bg-gray-800 p-5 rounded-lg w-96">
      <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white mb-4">{{ t('bookmark.edit') }}</h3>
      <form @submit.prevent="updateBookmark">
        <div class="mb-3">
          <label for="title" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('bookmark.title') }}</label>
          <input type="text" id="title" v-model="editedBookmark.Title" required
                 class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white">
        </div>
        <div class="mb-3">
          <label for="url" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('bookmark.url') }}</label>
          <input type="url" id="url" v-model="editedBookmark.URL" required
                 class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white">
        </div>
        <div class="mb-3">
          <label for="description" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('bookmark.description') }}</label>
          <textarea id="description" v-model="editedBookmark.Description"
                    class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white"
                    rows="3"></textarea>
        </div>
        <div class="mb-3">
          <label for="category" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('bookmark.category') }}</label>
          <select id="category" v-model="editedBookmark.CategoryID"
                  class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white">
            <option v-for="category in categories" :key="category.ID" :value="category.ID">{{ category.Name }}</option>
          </select>
        </div>
        <div class="mb-3">
          <label for="weight" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('bookmark.weight') }}</label>
          <input type="number" id="weight" v-model="editedBookmark.Weight"
                 class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white">
        </div>
        <div class="mb-3 flex items-center">
          <input type="checkbox" id="isPrivate" v-model="editedBookmark.IsPrivate" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
          <label for="isPrivate" class="ml-2 block text-sm text-gray-900 dark:text-gray-300">{{ t('bookmark.isPrivate') }}</label>
        </div>
        <div class="mt-4 flex justify-end space-x-2">
          <button type="button" @click="$emit('close')" class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400">{{ t('common.cancel') }}</button>
          <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">{{ t('common.save') }}</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'

export default {
  name: 'EditBookmarkModal',
  props: {
    bookmark: {
      type: Object,
      required: true
    }
  },
  emits: ['close', 'update'],
  setup(props, { emit }) {
    const store = useStore()
    const { t } = useI18n()
    const editedBookmark = ref({ ...props.bookmark })
    const categories = computed(() => store.state.categories)

    const updateBookmark = () => {
      const updatedBookmark = {
        ...editedBookmark.value,
        Description: editedBookmark.value.Description.replace(/\n/g, '\\n')
      }
      emit('update', updatedBookmark)
    }

    return {
      editedBookmark,
      categories,
      updateBookmark,
      t
    }
  }
}
</script>