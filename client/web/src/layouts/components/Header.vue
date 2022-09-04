<script setup lang="ts">
import type { MenuOption } from 'naive-ui'
import { menuLabelRender } from '~/utils/ui'

const menuOptions: MenuOption[] = [
  {
    label: '首页',
    path: '/',
    key: 'index',
  },
  {
    label: '关于',
    key: 'about',
  },
]

const userStore = useUserStore()
const miscStore = useMiscStore()
const route = useRoute()
const router = useRouter()

const activeKey = ref(route.name as string)
const searchWord = ref('')

const handleSearch = () => {
  if (searchWord.value)
    router.push(`/?keyword=${encodeURIComponent(searchWord.value)}`)
  else
    router.push('/')
}

watch(route, () => {
  activeKey.value = route.name as string
})
</script>

<template>
  <div class="relative h-54px">
    <header class="border-b-1 border-b-#f1f1f1 color-#909090 z-250 fixed top-0 left-0 right-0 bg-#fff">
      <div class="p-y-6px m-auto max-w-1240px flex items-center h-full w-full relative">
        <nav class="h-full flex-auto">
          <NSpace justify="space-between" align="center">
            <NMenu
              :value="activeKey"
              mode="horizontal"
              :options="menuOptions"
              :render-label="menuLabelRender"
            />
            <NSpace align="center">
              <NInput
                v-model:value="searchWord"
                clearable
                round
                placeholder="搜索"
                @keyup.enter.prevent="handleSearch"
              >
                <template #suffix>
                  <ICarbonSearch />
                </template>
              </NInput>
              <NSpace v-if="!!userStore.info">
                <NButton quaternary circle>
                  <template #icon>
                    <NIcon><ICarbonNotification /></NIcon>
                  </template>
                </NButton>
                <UserAvatar :src="userStore.info?.avatar" :size="32" />
              </NSpace>
              <NButton
                v-else
                strong
                secondary
                round
                type="primary"
                @click="miscStore.setAuthModalVisible(true)"
              >
                登录/注册
              </NButton>
            </NSpace>
          </NSpace>
        </nav>
      </div>
    </header>
  </div>
</template>
