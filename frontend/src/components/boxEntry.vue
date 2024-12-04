<script setup lang="ts">
import {ref} from 'vue';
import {useRouter} from 'vue-router';
import {Button} from './ui/button';

// See: https://www.shadcn-vue.com/docs/components/card.html
import {
	Card,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from '@/components/ui/card';

import {
	Popover,
	PopoverContent,
	PopoverTrigger,
} from '@/components/ui/popover';

import {Label} from '@/components/ui/label';
import {Input} from '@/components/ui/input';

const props = defineProps(['box']);
const router = useRouter();

const name = ref<string>(props.box.name);
const location = ref<string>(props.box.location);
const isOpen = ref<boolean>(false);

const emit = defineEmits(['updated']);

const saveEdit = () => {
	fetch(`http://localhost:8080/v1/box/${props.box.id}`, {
		method: 'PUT',
		body: JSON.stringify({name: name.value, location: location.value}),
	}).then(() => {
		emit('updated');
		isOpen.value = false; // Close the popup on success.
	}).catch(error => {
		console.error(error);
	});
};
</script>

<template>
  <Card class="w-[250px]">
    <CardHeader>
      <CardTitle>{{ props.box.name }}</CardTitle>
      <CardDescription>Location: {{ props.box.location }}</CardDescription>
    </CardHeader>
    <CardFooter class="flex justify-end px-6 pb-6 space-x-1">
      <Popover v-model:open="isOpen">
        <PopoverTrigger as-child>
          <Button variant="outline">
            Edit
          </Button>
        </PopoverTrigger>
        <PopoverContent class="w-80">
          <div class="grid gap-4">
            <div class="space-y-2">
              <h4 class="font-medium leading-none">
                Modify box details.
              </h4>
              <p class="text-sm text-muted-foreground">
                Enter the changes you want to make then click save.
              </p>
            </div>
            <div class="grid gap-2">
              <div class="grid grid-cols-3 items-center gap-4">
                <Label for="width">Name</Label>
                <Input
                  id="width"
                  v-model:model-value="name"
                  type="text"
                  :default-value="props.box.name"
                  class="col-span-2 h-8"
                />
              </div>
              <div class="grid grid-cols-3 items-center gap-4">
                <Label for="maxWidth">Location</Label>
                <Input
                  id="maxWidth"
                  v-model:model-value="location"
                  type="text"
                  :default-value="props.box.location"
                  class="col-span-2 h-8"
                />
              </div>
            </div>
            <div class="flex flex-row-reverse">
              <Button @click="saveEdit()">
                Save
              </Button>
            </div>
          </div>
        </PopoverContent>
      </Popover>
      <Button
        @click="() => {
          router.push({ name: 'boxContents', params: { id: box.id }})
        }"
      >
        Open
      </Button>
    </CardFooter>
  </Card>
</template>
