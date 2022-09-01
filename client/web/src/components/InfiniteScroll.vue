<script setup lang="ts">
const emits = defineEmits(['intersect'])
let observer: IntersectionObserver

const root = ref<HTMLDivElement>()
const isIntersecting = ref(false)
const isComplete = ref(false)

onMounted(() => {
  observer = new IntersectionObserver(([entry]) => {
    if (entry && entry.isIntersecting && root.value) {
      isIntersecting.value = true
      observer.unobserve(root.value)
      emits('intersect', {
        loaded() {
          isIntersecting.value = false
          observer.observe(root.value!)
        },
        complete() {
          observer?.disconnect()
          isIntersecting.value = false
          isComplete.value = true
        },
      })
    }
  })
  if (root.value)
    observer.observe(root.value)
})

onUnmounted(() => {
  observer?.disconnect()
})
</script>

<template>
  <div ref="root">
    <slot v-if="isIntersecting" name="placeholder">
      <Skeleton />
    </slot>
    <slot v-if="isComplete" name="no-more">
      <NDivider>
        没有了哦
      </NDivider>
    </slot>
  </div>
</template>
