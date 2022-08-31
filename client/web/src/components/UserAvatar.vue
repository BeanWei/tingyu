<script setup lang="ts">
const props = defineProps<{
  src?: string
  size?: number
}>()
</script>

<template>
  <NAvatar
    v-if="props.src"
    round
    :size="props.size"
    :src="props.src"
  />
  <NAvatar
    v-else
    round
    :size="props.size"
    :style="{
      'background-color': 'rgba(24, 160, 88, 0.16)',
      'color': '#18a058',
    }"
  >
    <NIcon><ICarbonRainDrop /></NIcon>
  </NAvatar>
</template>
