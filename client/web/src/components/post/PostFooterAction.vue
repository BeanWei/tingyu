<script setup lang="ts">
const { data } = defineProps<{
  data: Record<string, any>
}>()

const router = useRouter()
</script>

<template>
  <NSpace justify="space-between">
    <div class="flex items-center cursor-pointer">
      <NIcon size="18">
        <ICarbonBookmarkAdd />
      </NIcon>
      <span class="m-l-0.3em text-12px font-medium">收藏</span>
    </div>
    <div
      class="flex items-center cursor-pointer"
      @click="router.push(`/post/${data.id}`)"
    >
      <NIcon size="18">
        <ICarbonChat />
      </NIcon>
      <span class="m-l-0.3em text-12px font-medium">{{ data.comment_count || '评论' }}</span>
    </div>
    <div class="flex items-center cursor-pointer">
      <NIcon size="18">
        <ICarbonArrowUpRight />
      </NIcon>
      <span class="m-l-0.3em text-12px font-medium">分享</span>
    </div>
  </NSpace>
</template>
