<script setup lang="ts">
import PostHeaderUser from './PostHeaderUser.vue'
import PostHeaderAction from './PostHeaderAction.vue'
</script>

<template>
  <div class="bg-#fff border-rd-4px p-x-20px p-t-20px relative">
    <div class="flex items-center justify-between">
      <PostHeaderUser />
      <PostHeaderAction />
    </div>
    <div class="p-t-2">
      <div>123</div>
    </div>
    <div class="m-t-12px m-x--20px">
      <PostFooterAction />
    </div>
  </div>
</template>
