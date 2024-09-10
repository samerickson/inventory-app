<script setup lang="ts">
import { useRouter } from 'vue-router';

const props = defineProps(['box'])
const router = useRouter();
</script>

<template>
  <fieldset>
    <p>Name: {{ props.box.name }}</p>
    <p>Location: {{  props.box.location }}</p>
    <button @click="() => {
      router.push({ name: 'boxContents', params: { id: box.id }})
    }">Open</button>
  </fieldset>
</template>