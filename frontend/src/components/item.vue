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

interface ItemProp {
  item: {
    name: string,
    id: number,
  };
}

const props = defineProps<ItemProp>();
const emit = defineEmits(['deleted'])

const removeItem = () => {
  // TODO: handle failures
  fetch(`http://localhost:8080/v1/item/${props.item.id}`, { method: 'DELETE'}).then(() => {
    emit('deleted');
  });
}
</script>

<template>
  <Card class="w-[250px] my-4">
    <CardHeader>
      <CardTitle>{{ props.item.name }}</CardTitle>
    </CardHeader>
    <CardContent />
    <CardFooter class="flex justify-between px-6 pb-6">
      <Button variant="outline">
        Edit
      </Button>
      <Button variant="destructive" @click="removeItem">
        Remove
      </Button>
    </CardFooter>
  </Card>
</template>