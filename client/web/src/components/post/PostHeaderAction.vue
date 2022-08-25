<script setup lang="ts">
import type { DropdownOption } from 'naive-ui'

const options: DropdownOption[] = [
  {
    label: '编辑',
    key: 'edit',
  },
  {
    label: '删除',
    key: 'delete',
  },
]
</script>

<template>
  <n-dropdown
    trigger="hover"
    :options="options"
    :show-arrow="true"
    placement="bottom-end"
  >
    <n-button quaternary circle>
      <template #icon>
        <n-icon><i-carbon-overflow-menu-vertical /></n-icon>
      </template>
    </n-button>
  </n-dropdown>
</template>
