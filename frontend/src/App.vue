<script setup lang="ts">
import {RouterView} from 'vue-router';
import ThemeToggle from '@/components/themeToggle.vue';
import {Input} from '@/components/ui/input';
</script>

<template>
  <div class="container">
    <header class="flex justify-between my-4">
      <h1 class="text-3xl">
        📦 Inventory App
      </h1>
      <ThemeToggle />
    </header>

    <Input
      class="max-w-96"
      placeholder="Search"
    />
    <RouterView />
  </div>
</template>
