<script setup lang="ts">
const router = useRouter()
</script>

<template>
  <main p="x4 y10">
    <NResult status="404" title="404" description="你似乎来到了一片荒原">
      <template #footer>
        <NButton type="primary" ghost @click="router.push('/')">
          回到首页
        </NButton>
      </template>
    </NResult>
  </main>
</template>
