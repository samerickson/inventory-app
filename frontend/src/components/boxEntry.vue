<script setup lang="ts">
import { useRouter } from 'vue-router';
import { Button } from './ui/button';

// See: https://www.shadcn-vue.com/docs/components/card.html
import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'

const props = defineProps(['box'])
const router = useRouter();
</script>

<template>
  <Card class="w-[250px]">
    <CardHeader>
      <CardTitle>{{ props.box.name }}</CardTitle>
      <CardDescription>Location: {{ props.box.location }}</CardDescription>
    </CardHeader>
    <CardFooter class="flex justify-end px-6 pb-6">
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