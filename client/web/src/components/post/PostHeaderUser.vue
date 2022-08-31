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
    <div class="max-w-24rem truncate">
      {{ data.ip_loc }}
    </div>
    <div class="m-x-0.5em">
      ·
    </div>
    <a class="decoration-none color-#909090">
      <NTime unix type="relative" :time="data.created_at" />
    </a>
  </div>
</template>
