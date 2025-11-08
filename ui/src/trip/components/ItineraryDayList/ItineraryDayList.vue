<script setup lang="ts">
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/core/components/ui/accordion'
import ItineraryDay from './ItineraryDay.vue'

const dates = [
  'friday, december 12, 2025',
  'saturday, december 13, 2025',
  'sunday, december 14, 2025',
  'monday, december 15, 2025',
  'tuesday, december 16, 2025',
  'wednesday, december 17, 2025',
]
</script>

<template>
  <div class="px-4">
    <Accordion type="multiple" class="w-full" collapsible :default-value="dates[0]">
      <AccordionItem v-for="date in dates" :key="date" :value="date">
        <AccordionTrigger class="items-center">
          <div>
            <p class="text-2xl font-bold">Friday, December 12, 2025</p>
            <p class="text-muted-foreground">Lake Titikaka • Lake Titikaka • Lake Titikaka</p>
          </div>
        </AccordionTrigger>
        <AccordionContent>
            <ItineraryDay /> 
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  </div>
</template>
