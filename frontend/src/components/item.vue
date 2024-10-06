<script setup lang="ts">
import {
  Button
} from '@/components/ui/button';
import {
  Card,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';

import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'

import { ref } from 'vue';

interface ItemProp {
  item: {
    name: string,
    id: number,
  };
}

const props = defineProps<ItemProp>();
const emit = defineEmits(['deleted', 'updated'])

const name = ref<string>(props.item.name);
const isOpen = ref<boolean>(false);

const removeItem = () => {
  // TODO: handle failures
  fetch(`http://localhost:8080/v1/item/${props.item.id}`, { method: 'DELETE'}).then(() => {
    emit('deleted');
  });
}

const save = () => {
  fetch(`http://localhost:8080/v1/item/${props.item.id}`, {
    method: "PUT",
    body: JSON.stringify({ name: name.value })
  }).then(() => {
    emit('updated')
    isOpen.value = false; // Close the popup on success.
  }).catch(error => {
    console.error(error)
  })
}
</script>

<template>
  <Card class="w-[250px] my-4">
    <CardHeader>
      <CardTitle>{{ props.item.name }}</CardTitle>
    </CardHeader>
    <CardContent />
    <CardFooter class="flex justify-between px-6 pb-6">
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
                Modify Item.
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
                  :default-value="props.item.name"
                  class="col-span-2 h-8"
                />
              </div>
            </div>
            <div class="flex flex-row-reverse">
              <Button @click="save()">
                Save
              </Button>
            </div>
          </div>
        </PopoverContent>
      </Popover>
      <Button variant="destructive" @click="removeItem">
        Remove
      </Button>
    </CardFooter>
  </Card>
</template>