<script setup lang="ts">
const { data } = defineProps<{
  data: Record<string, any>
}>()

const router = useRouter()
</script>

<template>
  <NSpace justify="space-between" align="center">
    <NSpace>
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
          <ICarbonFaceAdd />
        </NIcon>
      </div>
    </NSpace>
    <CreationInfo :time="data.created_at" :location="data.ip_loc" />
  </NSpace>
</template>
