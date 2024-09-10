<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const box = ref();

// TODO: handle 404 or other errors.
fetch(`http://localhost:8080/v1/box/${route.params.id}`).then(async (response) => {
  console.log(response)
  box.value = await response.json();
});


console.log(route.params.id)
</script>

<template>
  <h1>Box contents</h1>
  {{ box  }}
</template>