<template>
  <div class="flex h-full bg-primary-50 dark:bg-neutral-900">
    <!-- 左側導航欄 -->
    <div class="w-64 bg-white dark:bg-neutral-800 shadow-md overflow-y-auto">
      <!-- 左側導航欄內容保持不變 -->
    </div>

    <!-- 主要內容區 -->
    <div class="flex-1 overflow-hidden flex flex-col">
      <div class="p-8">
        <div class="mb-6 flex justify-between items-center">
          <!-- 搜索欄和顯示模式切換保持不變 -->
        </div>
      </div>
      <div class="flex-1 overflow-y-auto px-8 pb-8">
        <!-- 書籤列表內容 -->
        <div v-if="filteredBookmarks.length === 0" class="text-center text-gray-500 dark:text-gray-400">
          {{ t('bookmark.noBookmarks') }}
        </div>
        <div v-else :class="{'grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6': displayMode === 'block', 'space-y-4': displayMode === 'list'}">
          <!-- 書籤卡片內容保持不變 -->
        </div>
      </div>
    </div>

    <!-- Modal保持不變 -->
  </div>
</template>

<script>
// script 部分保持不變
</script>

<style scoped>
/* 可以刪除之前的樣式，因為我們現在使用 Tailwind 的類來控制佈局 */
</style>