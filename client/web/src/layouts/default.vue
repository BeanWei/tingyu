<script setup lang="ts">
import Header from './components/Header.vue'
import NavBox from './components/NavBox.vue'
import UserBox from './components/UserBox.vue'
</script>

<template>
  <Header />
  <main class="flex max-w-1240px relative w-full py-5 m-y-0 m-x-a text-gray-700 dark:text-gray-200">
    <div class="mr-20px w-180px">
      <NavBox />
    </div>
    <div class="w-720px max-w-full flex-shrink-0 m-r-20px">
      <RouterView />
    </div>
    <div class="w-260px flex-1">
      <UserBox />
    </div>
  </main>
</template>
