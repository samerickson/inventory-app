<script setup lang="ts">
import {RouterView} from 'vue-router';
import ThemeToggle from '@/components/themeToggle.vue';
import SearchBar from './components/searchBar.vue';
</script>

<template>
  <div class="container">
    <header class="flex justify-between my-4">
      <h1 class="text-3xl">
        📦 Inventory App
      </h1>
      <ThemeToggle />
    </header>

    <SearchBar />

    <RouterView />
  </div>
</template>
