<script setup lang="ts">
import { ref, watch } from 'vue'
import { animations } from '@formkit/drag-and-drop'
import { dragAndDrop } from '@formkit/drag-and-drop/vue'

import { GripVertical, TrashIcon } from 'lucide-vue-next'
import { Checkbox } from '@/core/components/ui/checkbox'
import { Badge } from '@/core/components/ui/badge'

const items = ref([
  'Depeche Mode',
  'Duran Duran',
  'Pet Shop Boys',
  'Kraftwerk',
  'Tears for Fears',
  'Spandau Ballet',
])
const parent = ref<HTMLElement | null>(null)

watch(parent, (newVal) => {
  if (!newVal) return
  dragAndDrop({
    parent: newVal,
    values: items,
    group: 'itinerary-day',
    plugins: [animations()],
    onDragend: (data) => {
      console.log('Drag ended', data)
    },
  })
})
</script>

<template>
  <div ref="parent" class="grid gap-4">
    <div
      v-for="(tape, index) in items"
      :key="tape"
      class="grid grid-cols-[auto_1fr_auto] gap-2 hover:cursor-grab"
    >
      <div class="flex items-center">
        <GripVertical />
      </div>
      <div class="bg-secondary rounded-md flex items-start gap-2">
        <span class="grid w-8 place-items-center -ml-4">
          <svg
            data-prefix="fas"
            data-icon="location-pin"
            role="img"
            viewBox="0 0 384 512"
            aria-hidden="true"
            class="w-8 h-8 col-start-1 row-start-1"
          >
            <path
              fill="currentColor"
              d="M192 0C86 0 0 84.4 0 188.6 0 307.9 120.2 450.9 170.4 505.4 182.2 518.2 201.8 518.2 213.6 505.4 263.8 450.9 384 307.9 384 188.6 384 84.4 298 0 192 0z"
            ></path>
          </svg>
          <span class="col-start-1 row-start-1 text-white font-bold mb-1.5">
            {{ index + 1 }}
          </span>
        </span>
        <div class="flex flex-col justify-between">
          <h1 class="font-bold">Lake Titikaka {{ tape }}</h1>
          <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti illo deserunt ea,
            commodi sit esse, nostrum eligendi quod accusamus excepturi dolore voluptatum non
            eveniet totam numquam ipsam architecto quae dolorem?
          </p>
          <div class="flex gap-2">
            <Badge variant="blue">12:30 PM - 14:00 PM</Badge>
            <Badge>142.00 us</Badge>
          </div>
        </div>
      </div>
      <div class="flex items-center gap-2">
        <img
          src="https://images.unsplash.com/photo-1506744038136-46273834b3fb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8M3x8dHJhdmVsJTIwcGljdHVyZXN8ZW58MHx8MHx8fDA%3D&w=1000&q=80"
          alt="Travel Image"
          class="aspect-video h-24 object-cover rounded-md"
        />
        <TrashIcon />
      </div>
    </div>
  </div>
</template>
