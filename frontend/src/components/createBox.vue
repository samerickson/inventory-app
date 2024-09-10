<script setup lang="ts">
import { ref } from 'vue';

const name = ref();
const location = ref();

const submit = () => {
fetch("http://localhost:8080/v1/box/", {
  method: "POST",
  body: JSON.stringify({ name: name.value, location: location.value })
}).catch(error => {
  console.error(error)
})
}
</script>

<template>
  <h2>Create New Box</h2>
  <label for="name">Name:</label>
  <input id="name" type="text" v-model="name"/>
  <br>
  <label for="name">Location:</label>
  <input id="location" type="text" v-model="location"/>
  <br>
  <button @click="submit">Submit</button>
</template>