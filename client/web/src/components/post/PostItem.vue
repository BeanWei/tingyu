<script setup lang="ts">
const { data } = defineProps<{
  data: AnyObject
}>()
</script>

<template>
  <div class="bg-#fff border-rd-4px p-20px relative">
    <NThing content-indented>
      <template #avatar>
        <UserAvatar
          :size="38"
          :src="data.edges.user?.avatar"
          class="flex relative border-neutral-200/70"
        />
      </template>
      <template #header>
        <PostHeaderUser :data="data" />
      </template>
      <!-- <template #header-extra>

      </template> -->
      <template #description>
        <div class="m-t-1">
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
