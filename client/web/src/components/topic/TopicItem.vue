<script setup lang="ts">
const { data } = defineProps<{
  data: AnyObject
}>()
</script>

<template>
  <div class="bg-#fff border-rd-1 p-x-5 relative">
    <NListItem class="items-center">
      <template #prefix>
        <NAvatar v-if="data.icon" :src="data.icon" :size="48" />
        <NAvatar
          v-else
          :size="48"
          :style="{
            'background-color': 'rgba(24, 160, 88, 0.16)',
            'color': '#18a058',
          }"
        >
          {{ data.title[0] }}
        </NAvatar>
      </template>
      <NThing
        :title="data.title"
        :description="data.description"
      />
      <template #suffix>
        <NButton text type="primary">
          <template #icon>
            <n-icon>
              <ICarbonAdd />
            </n-icon>
          </template>
          关注
        </NButton>
      </template>
    </NListItem>
  </div>
</template>
