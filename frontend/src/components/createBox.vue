<script setup lang="ts">
import {ref} from 'vue';

const name = ref();
const location = ref();

const emit = defineEmits(['created']);

const submit = () => {
	fetch('http://localhost:8080/v1/box/', {
		method: 'POST',
		body: JSON.stringify({name: name.value, location: location.value}),
	}).then(() => {
		emit('created');
	}).catch(error => {
		console.error(error);
	}).finally(() => {
		clearInputs();
	});
};

const clearInputs = () => {
	name.value = '';
	location.value = '';
};

import {
	Card,
	CardTitle,
	CardHeader,
	CardContent,
} from '@/components/ui/card';
import {Button} from '@/components/ui/button';
import {Input} from '@/components/ui/input';
</script>

<template>
  <Card class="w-[450px]">
    <CardHeader>
      <CardTitle>Create new Box</CardTitle>
    </CardHeader>
    <CardContent>
      <!-- TODO: Make this a proper form that you can submit with enter -->
      <div class="grid w-full max-w-sm items-center gap-1.5">
        <Label for="name">Name</Label>
        <Input
          id="name"
          v-model="name"
          type="text"
          placeholder="Box 1"
        />
        <Label for="location">Location</Label>
        <Input
          id="location"
          v-model="location"
          type="text"
          placeholder="Spare bedroom"
        />
      </div>
    </CardContent>
    <CardFooter class="flex justify-between px-6 pb-6">
      <Button
        variant="outline"
        @click="clearInputs"
      >
        Clear
      </Button>
      <Button @click="submit">
        Create
      </Button>
    </CardFooter>
  </Card>
</template>
