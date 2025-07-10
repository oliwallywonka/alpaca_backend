<script setup lang="ts">
import { ref } from 'vue'
import {
  Dialog,
  DialogContent,
  DialogTitle,
  DialogTrigger,
  DialogHeader,
  DialogDescription,
} from '@/core/components/ui/dialog'

const props = defineProps<{
  imageURL: string
}>()

// TODO: Fix Image Modal size
const imgSrc = ref(props.imageURL)
</script>

<template>
  <Dialog>
    <DialogTrigger as-child>
      <button class="w-full h-full rounded-md overflow-clip">
        <img
          :src="imgSrc"
          alt="Tour Image"
          class="w-full h-full object-cover"
          :width="100"
          :height="100"
          @error="
            () => {
              imgSrc = 'default-image.jpg'
            }
          "
        />
      </button>
    </DialogTrigger>

    <DialogContent class="w-full h-[90vh] mx-auto p-2">
      <DialogHeader class="hidden">
        <DialogTitle>Images</DialogTitle>
        <DialogDescription> A image modal </DialogDescription>
      </DialogHeader>
      <img :src="imgSrc" alt="Tour Image" class="h-[88vh] w-full object-cover rounded-sm" />
    </DialogContent>
  </Dialog>
</template>
