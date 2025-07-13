<script setup lang="ts">
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'
import TripsSummarySection from './components/TripsSummarySection.vue'
import TripsTable from './components/TripsTable.vue'
import TripsCalendarSection from './components/TripsCalendarSection.vue';
</script>

<template>
  <TripsSummarySection />

  <Tabs default-value="tripsTable">
    <TabsList>
      <TabsTrigger value="tripsTable"> Trips List </TabsTrigger>
      <TabsTrigger value="tripsCalendar"> Trips Calendar </TabsTrigger>
    </TabsList>
    <TabsContent value="tripsTable">
      <TripsTable />
    </TabsContent>
    <TabsContent value="tripsCalendar" class="min-h-[80vh]">
      <TripsCalendarSection />
    </TabsContent>
  </Tabs>
</template>
