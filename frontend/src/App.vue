<script setup lang="ts">
import { RouterView } from 'vue-router'
import ThemeToggle from './components/themeToggle.vue';
</script>

<template>
  <div class="container">
    <header class="flex justify-between my-4">
      <h1 class="text-3xl">
        Inventory App
      </h1>
      <ThemeToggle />
    </header>
    <RouterView />
  </div>
</template>
