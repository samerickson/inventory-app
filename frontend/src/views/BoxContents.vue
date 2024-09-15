<script setup lang="ts">
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { Button } from '@/components/ui/button';
import Grid from '@/components/grid.vue';
import Card from '@/components/ui/card/Card.vue';
import Item from '@/components/item.vue';

const route = useRoute();
const router = useRouter();
const box = ref();

const loadContents = () => {
  // TODO: handle 404 or other errors.
  fetch(`http://localhost:8080/v1/box/${route.params.id}`).then(async (response) => {
    box.value = await response.json();
  });
}

loadContents();
</script>

<template>
  <!-- Only show the details once we get them -->
  <div v-if="box">
    <h2 class="text-2xl">📦 {{ box.name }}</h2>
    <h3 class=" text-xl text-muted-foreground">🗺️ Location: {{ box.location }}</h3>
    <Button variant="outline" @click="router.push({ name: 'allBoxes'})">Go Back</Button>
    <Item :item="item" v-for="item in box.items" :key="item.id"></Item>
  </div>
  <!-- TODO: Show a skeleton until we get the box details from our API call. -->
</template>