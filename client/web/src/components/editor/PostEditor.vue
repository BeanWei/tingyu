<template>
  <div class="bg-inherit">
    <div>
      <div class="relative min-h-88px">
        123
      </div>
    </div>
    <n-space justify="space-between">
      <div class="flex relative">
        <n-button quaternary circle>
          <template #icon>
            <n-icon><i-carbon-image /></n-icon>
          </template>
        </n-button>
        <n-button quaternary circle>
          <template #icon>
            <n-icon><i-carbon-face-add /></n-icon>
          </template>
        </n-button>
      </div>
      <n-button
        strong secondary round type="primary"
      >
        发布
      </n-button>
    </n-space>
  </div>
</template>
