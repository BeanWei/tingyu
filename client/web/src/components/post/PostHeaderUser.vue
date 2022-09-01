<script setup lang="ts">
const { data } = defineProps<{
  data: Record<string, any>
}>()
</script>

<template>
  <div class="inline">
    <a class="font-medium text-16px color-#252933 decoration-none cursor-pointer"> {{ data.edges.user?.nickname }} </a>
  </div>
  <div class="flex items-center text-14px color-#8a919f whitespace-nowrap">
    我有一头小毛驴，我从来都不骑
  </div>
</template>
