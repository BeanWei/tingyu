<script setup lang="ts">
const { data } = defineProps<{
  data: AnyObject
}>()
</script>

<template>
  <div class="bg-#fff border-rd-1 p-y-4 p-x-5 relative">
    <NThing content-indented>
      <template #avatar>
        <UserAvatar
          :size="38"
          :src="data.user?.avatar"
          class="flex relative border-neutral-200/70"
        />
      </template>
      <template #header>
        <div class="inline">
          <a class="font-medium text-16px color-#252933 decoration-none cursor-pointer">
            {{ data.user?.nickname }}
            <span v-if="data.user?.headline" class="color-#8a919f text-14px whitespace-nowrap text-ellipsis before:content-[，]">{{ data.user?.headline }}</span>
          </a>
        </div>
        <div class="flex items-center text-14px color-#8a919f whitespace-nowrap">
          <CreationInfo :time="data.created_at" :location="data.ip_loc" />
        </div>
      </template>
      <template #header-extra>
        <slot name="header-extra" />
      </template>
      <template #description>
        <div class="m-t-2">
          <Editor :read-only="true" :initial-state="data.content" />
        </div>
      </template>
      <!-- <template #footer>

      </template> -->
      <template #action>
        <PostFooterAction :data="data" />
      </template>
    </NThing>
  </div>
</template>
