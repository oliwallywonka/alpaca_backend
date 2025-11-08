<script setup lang="ts">
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'
import TripsSummarySection from './components/TripsSummarySection.vue'
import TripsCalendarSection from './components/TripsCalendarSection.vue'
import TripsTableSection from './components/TripsTableSection/TripsTableSection.vue'
</script>

<template>
  <TripsSummarySection />
  <Tabs default-value="tripsTableSection">
    <TabsList>
      <TabsTrigger value="tripsTableSection"> Trips List </TabsTrigger>
      <TabsTrigger value="tripsCalendar"> Trips Calendar </TabsTrigger>
    </TabsList>
    <TabsContent value="tripsTableSection">
      <TripsTableSection />
    </TabsContent>
    <TabsContent value="tripsCalendar" class="min-h-[80vh]">
      <TripsCalendarSection />
    </TabsContent>
  </Tabs>
</template>
