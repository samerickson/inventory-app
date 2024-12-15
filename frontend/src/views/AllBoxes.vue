<script setup lang="ts">
import BoxEntry from '@/components/boxEntry.vue';
import CreateBox from '@/components/createBox.vue';
import Grid from '@/components/grid.vue';
import {onMounted, ref, watch} from 'vue';
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();

const boxes = ref();

const getData = () => {
  if (!route.query.query) {
    // TODO: handle 404 or other errors.
    fetch('http://localhost:8080/v1/box/', {cache: 'no-cache'}).then(async response => {
      boxes.value = await response.json();
    });

    return;
  }

  
  fetch(`http://localhost:8080/v1/box/search?query=${route.query.query}`, {cache: 'no-cache'}).then(async response => {
		boxes.value = await response.json();
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
  <!-- TODO: These should be displayed in a flexbox grid. -->
  <Grid>
    <BoxEntry
      v-for="box in boxes"
      :key="box.id"
      class="my-4"
      :box="box"
      @updated="getData"
    />
  </Grid>
  <CreateBox @created="getData" />
</template>
