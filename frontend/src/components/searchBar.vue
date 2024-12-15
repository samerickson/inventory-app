<script setup lang="ts">
import { Input } from '@/components/ui/input';
import { ToggleGroup, ToggleGroupItem } from '@/components/ui/toggle-group';
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const router = useRouter();
const route = useRoute();

const toggleState = ref<'item' | 'box'>('item');
const searchQuery = ref('');

onMounted(async () => {
  await router.isReady();

  if (!route.query.query) {
    router.push({ name: 'allBoxes' })
  }

console.log(route.query)

  searchQuery.value = route.query.query as string;
})

watch(() => searchQuery.value, (val) => {

  if (val === '') {
    router.push({ name: 'allBoxes' })
  }
  else if (toggleState.value === 'item') {
    router.push({ name: 'searchItems', query: { query: val}})
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
