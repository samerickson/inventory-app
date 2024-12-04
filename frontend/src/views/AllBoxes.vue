<script setup lang="ts">
import BoxEntry from '@/components/boxEntry.vue';
import CreateBox from '@/components/createBox.vue';
import Grid from '@/components/grid.vue';
import {onMounted, ref} from 'vue';

const boxes = ref();

const getData = () => {
	// TODO: handle 404 or other errors.
	fetch('http://localhost:8080/v1/box/', {cache: 'no-cache'}).then(async response => {
		boxes.value = await response.json();
	});
};

onMounted(() => {
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
