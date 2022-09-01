<script setup lang="ts">
const props = defineProps<{
  time: number
  location: string
}>()
</script>

<template>
  <div class="flex items-center text-14px color-#8a919f whitespace-nowrap">
    <a class="decoration-none color-#909090">
      <NTime unix type="relative" :time="props.time" />
    </a>
    <div class="m-x-0.5em">
      ·
    </div>
    <div class="max-w-24rem truncate">
      {{ props.location }}
    </div>
  </div>
</template>
