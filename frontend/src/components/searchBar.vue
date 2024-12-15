<script setup lang="ts">
import { Input } from '@/components/ui/input';
import { ToggleGroup, ToggleGroupItem } from '@/components/ui/toggle-group';
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const router = useRouter();
const route = useRoute();

const toggleState = ref<'item' | 'box'>('box');
const searchQuery = ref('');

onMounted(async () => {
  await router.isReady();
  searchQuery.value = route.query.query as string;

  if (route.path === '/') {
    // TODO: read the state of the toggle from browser local storage
    // and set that accordingly here
    router.push({ name: toggleState.value, query: { query: searchQuery.value }})
  }
})

// Watch for changes from item to box and vice versa
watch(() => toggleState.value, (name) => {
  router.push({ name, query: { query: searchQuery.value } })
})

// Watch for changes in search query
watch(() => searchQuery.value, (query) => {
  if (toggleState.value === 'item') {
    router.push({ name: 'item', query: { query }})
  }
  else if (toggleState.value === 'box') {
    router.push({ name: 'box', query: { query }})
  } 
})
</script>

<template>
  <div class="flex">
    <Input
      v-model="searchQuery"
      class="max-w-96 p-4"
      placeholder="Search"
    />

    <ToggleGroup
      v-model="toggleState"
      type="single"
      class="mx-4"
    >
      <ToggleGroupItem value="item">
        Item
      </ToggleGroupItem>
      <ToggleGroupItem value="box">
        Box
      </ToggleGroupItem>
    </ToggleGroup>
  </div>
</template>
