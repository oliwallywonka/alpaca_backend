<script setup lang="ts">
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'

import TourForm from '@/tour/components/TourForm.vue'
import ImagesForm from '../components/ImagesForm.vue'
import ItineraryList from '@/itinerary/components/List/ItineraryList.vue'
</script>

<template>
  <Tabs default-value="itineraries">
    <TabsList class="grid w-full grid-cols-3">
      <TabsTrigger value="tourForm"> Basic Info </TabsTrigger>
      <TabsTrigger value="images"> Images </TabsTrigger>
      <TabsTrigger value="itineraries"> Itineraries </TabsTrigger>
    </TabsList>
    <TabsContent value="tourForm">
      <TourForm />
    </TabsContent>
    <TabsContent value="images">
      <ImagesForm />
    </TabsContent>
    <TabsContent value="itineraries">
      <ItineraryList />
    </TabsContent>
  </Tabs>
</template>
