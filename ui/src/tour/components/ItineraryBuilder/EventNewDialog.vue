<script setup lang="ts">
import { Text, Handshake, MapPinHouse } from 'lucide-vue-next'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/core/components/ui/dialog'

const eventTypes = [
  {
    type: 'commentary',
    title: 'Commentary',
    icon: Text,
    description:
      'Images and commentaries of the itinerary "this will be showed on the summary section"',
  },
  {
    type: 'resource',
    title: 'Resource',
    icon: Handshake,
    description:
      'Services, activities, hotels, guides and providers of the itinerary "this will be showed on the operative expenses section"',
  },
  {
    type: 'destination',
    title: 'Destination',
    icon: MapPinHouse,
    description: 'The destination of the itinerary "this will be showed on the summary section"',
  },
]

const isOpen = defineModel<boolean>({ default: false })

const emit = defineEmits<{
  (e: 'click:event', type: string): void
}>()

function clickEvent(type: string) {
  emit('click:event', type)
  isOpen.value = false
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="max-w-6xl">
      <DialogHeader>
        <DialogTitle>New Event</DialogTitle>
        <DialogDescription> Choose the type of event you want to create </DialogDescription>
      </DialogHeader>

      <div class="grid md:grid-cols-3 gap-2">
        <button
          v-for="eventType in eventTypes"
          :key="eventType.type"
          @click="clickEvent(eventType.type)"
          class="flex flex-col justify-center items-center border rounded-md p-2 hover:bg-muted/40"
        >
          <component :is="eventType.icon" class="size-12" />
          <p class="font-bold text-sm">{{ eventType.title }}</p>
          <p class="text-muted-foreground text-sm">
            {{ eventType.description }}
          </p>
        </button>
      </div>

      <DialogFooter> </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
