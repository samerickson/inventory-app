<script setup lang="ts">
import CreateBox from '@/components/createBox.vue';
</script>

<template>
  <main>
    <h1>This is an about page</h1>
    <router-view />
    <CreateBox />
  </main>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>