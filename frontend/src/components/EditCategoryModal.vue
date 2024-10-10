<template>
  <div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center">
    <div class="bg-white dark:bg-gray-800 p-5 rounded-lg w-96">
      <h3 class="text-lg leading-6 font-medium text-gray-900 dark:text-white mb-4">{{ t('category.edit') }}</h3>
      <form @submit.prevent="updateCategory">
        <div class="mb-3">
          <label for="categoryName" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('category.name') }}</label>
          <input type="text" id="categoryName" v-model="editedCategory.Name" required
                 class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white">
        </div>
        <div class="mb-3">
          <label for="categoryIcon" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('category.icon') }}</label>
          <div class="flex items-center">
            <input type="text" id="categoryIcon" v-model="editedCategory.Icon"
                   class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white mr-2"
                   placeholder="Icon class" readonly>
            <IconPicker @select="selectIcon" />
          </div>
        </div>
        <div class="mb-3">
          <label for="categoryWeight" class="block text-sm font-medium text-gray-700 dark:text-gray-300">{{ t('category.weight') }}</label>
          <input type="number" id="categoryWeight" v-model="editedCategory.Weight"
                 class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 dark:bg-gray-700 dark:text-white">
        </div>
        <div class="mb-3 flex items-center">
          <input type="checkbox" id="categoryIsPrivate" v-model="editedCategory.IsPrivate" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
          <label for="categoryIsPrivate" class="ml-2 block text-sm text-gray-900 dark:text-gray-300">{{ t('category.isPrivate') }}</label>
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
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import IconPicker from './IconPicker.vue'

export default {
  name: 'EditCategoryModal',
  components: {
    IconPicker
  },
  props: {
    category: {
      type: Object,
      required: true
    }
  },
  emits: ['close', 'update'],
  setup(props, { emit }) {
    const { t } = useI18n()
    const editedCategory = ref({ ...props.category })

    const updateCategory = () => {
      emit('update', editedCategory.value)
    }

    const selectIcon = (icon) => {
      editedCategory.value.Icon = icon
    }

    return {
      editedCategory,
      updateCategory,
      selectIcon,
      t
    }
  }
}
</script>