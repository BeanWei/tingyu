<script setup lang="ts">
const userStore = useUserStore()
</script>

<template>
  <div class="bg-#fff border-rd-4px p-24px">
    <div class="flex items-center border-b-1 border-b-#f1f1f1 pb-20px">
      {{ userStore.info?.nickname }}
    </div>
    <div class="flex pt-20px">
      <NRow>
        <NCol :span="6">
          <NStatistic label="帖子">
            {{ userStore.info?.count_post || 0 }}
          </NStatistic>
        </NCol>
        <NCol :span="6">
          <NStatistic label="话题">
            {{ userStore.info?.count_topic || 0 }}
          </NStatistic>
        </NCol>
        <NCol :span="6">
          <NStatistic label="关注">
            {{ userStore.info?.count_following || 0 }}
          </NStatistic>
        </NCol>
        <NCol :span="6">
          <NStatistic label="粉丝">
            {{ userStore.info?.count_follower || 0 }}
          </NStatistic>
        </NCol>
      </NRow>
    </div>
  </div>
</template>
