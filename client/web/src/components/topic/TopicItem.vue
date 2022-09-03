<script setup lang="ts">
import { useAxios } from '@vueuse/integrations/useAxios'
import type { Result } from '~/api'
import { instance, url } from '~/api'

const { data, isFollowed } = defineProps<{
  data: AnyObject
  isFollowed: boolean
}>()

const submitting = ref(false)
const followed = ref(isFollowed)
const followBtnEle = ref<HTMLDivElement>()
const followBtnIsHovered = useElementHover(followBtnEle)

const userStore = useUserStore()

const changeFollowState = async () => {
  submitting.value = true
  const { isFinished, error } = await useAxios<Result<AnyObject[]>>(followed.value ? url.unFollowTopic : url.followTopic, {
    params: {
      id: data.id,
    },
  }, instance)
  if (isFinished) {
    if (!error.value)
      followed.value = !followed.value
    submitting.value = false
  }
}
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
        <NButton
          v-if="userStore.info?.id && !followed"
          text
          type="primary"
          :loading="submitting"
          @click="changeFollowState"
        >
          <template #icon>
            <n-icon>
              <ICarbonAdd />
            </n-icon>
          </template>
          关注
        </NButton>
        <NButton
          v-if="userStore.info?.id && followed"
          ref="followBtnEle"
          text
          :loading="submitting"
          @click="changeFollowState"
        >
          <template #icon>
            <n-icon v-if="followBtnIsHovered">
              <ICarbonClose />
            </n-icon>
            <n-icon v-else>
              <ICarbonCheckmark />
            </n-icon>
          </template>
          {{ followBtnIsHovered ? '取消关注' : '已关注' }}
        </NButton>
      </template>
    </NListItem>
  </div>
</template>
