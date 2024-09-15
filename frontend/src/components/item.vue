<script setup lang="ts">
import {
  Button
} from '@/components/ui/button';
import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'

interface ItemProp {
  item: {
    name: string
  };
}

const props = defineProps<ItemProp>()
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
      <Button variant="destructive">
        Remove
      </Button>
    </CardFooter>
  </Card>
</template>