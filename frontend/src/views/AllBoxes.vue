<script setup lang="ts">
import BoxEntry from '@/components/boxEntry.vue';
import { ref } from 'vue';

const boxes = ref();

// TODO: handle 404 or other errors.
fetch("http://localhost:8080/v1/box/", {cache: "no-cache"}).then(async (response) => {
  boxes.value = await response.json();
});
</script>

<template>
  <BoxEntry v-for="box in boxes" :box="box" :key="box.id"/>
</template>