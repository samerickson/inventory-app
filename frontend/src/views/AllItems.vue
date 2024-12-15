<script setup lang="ts">
import Grid from '@/components/grid.vue';
import Item from '@/components/item.vue';
import {onMounted, ref, watch} from 'vue';
import { useRoute, useRouter } from 'vue-router';

const router = useRouter();
const route = useRoute();

const items = ref();

const getData = () => {

  if (!route.query.query) {
    fetch(`http://localhost:8080/v1/item/`, {cache: 'no-cache'}).then(async response => {
      items.value = await response.json();
    });

    return;
  }

  fetch(`http://localhost:8080/v1/item/search?query=${route.query.query}`, {cache: 'no-cache'}).then(async response => {
		items.value = await response.json();
	});
};

// TODO: Might need to debounce this
watch(() => route.query, getData, { immediate: true })

onMounted(async () => {
  await router.isReady();
	getData();
});
</script>

<template>
  <Grid>
    <Item
      v-for="item in items"
      :key="item.id"
      :item="item"
    />
  </Grid>
</template>
