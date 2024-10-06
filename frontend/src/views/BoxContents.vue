<script setup lang="ts">
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import Item from '@/components/item.vue';

import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import Grid from '@/components/grid.vue';

const route = useRoute();
const router = useRouter();

const boxId = Number(route.params.id);
const box = ref();
const newItemName = ref();

const loadContents = () => {
  // TODO: handle 404 or other errors.
  fetch(`http://localhost:8080/v1/box/${boxId}`).then(async (response) => {
    box.value = await response.json();
  });
}

const addNewItem = () => {
  fetch(`http://localhost:8080/v1/item/`, {
    method: 'POST',
    headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      name: newItemName.value,
      box: boxId,
    })
  }).then(async () => {
    loadContents();
  });
}

loadContents();
</script>

<template>
  <!-- Only show the details once we get them -->
  <div v-if="box">
    <h2 class="text-2xl">
      📦 {{ box.name }}
    </h2>
    <h3 class=" text-xl text-muted-foreground">
      🗺️ Location: {{ box.location }}
    </h3>
    <Dialog>
      <Button
        class="mr-2"
        variant="outline"
        @click="router.push({ name: 'allBoxes'})"
      >
        Go Back
      </Button>
      <DialogTrigger as-child>
        <Button variant="outline">
          Add Item
        </Button>
      </DialogTrigger>
      <Grid>
        <Item
          v-for="item in box.items"
          :key="item.id"
          :item="item"
          @deleted="loadContents()"
          @updated="loadContents()"
        />
      </Grid>
      <DialogContent class="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Add Item</DialogTitle>
          <DialogDescription>
            Add a new item, it will automatically be added to this box.
          </DialogDescription>
        </DialogHeader>
        <div class="grid gap-4 py-4">
          <div class="grid grid-cols-4 items-center gap-4">
            <Label
              for="name"
              class="text-right"
            >
              Name
            </Label>
            <Input
              id="name"
              v-model="newItemName"
              class="col-span-3"
            />
          </div>
        </div>
        <DialogFooter>
          <DialogClose as-child>
            <Button
              type="button"
              @click="addNewItem()"
            >
              Add Item
            </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </div>
  <!-- TODO: Show a skeleton until we get the box details from our API call. -->
</template>