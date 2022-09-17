<script setup lang="ts">
const props = defineProps<{ id: string }>()
</script>

<template>
  <div class="bg-#fff box-border border-rd-4px relative p-20px">
    {{ props.id }}
  </div>
</template>
