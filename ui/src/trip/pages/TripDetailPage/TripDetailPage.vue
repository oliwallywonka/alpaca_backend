<script setup lang="ts">
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'

import ExpensesCard from '@/trip/components/ExpensesCard.vue'
import ItineraryBuilder from '@/trip/components/ItineraryBuilder.vue'
import SummarySection from '@/trip/components/SummarySection.vue'
import TripCustomersTable from '@/trip/components/TripCustomersTable.vue'
import TripCalendarSection from '@/trip/components/TripCalendarSection/TripCalendarSection.vue'
import PaymensSection from '@/trip/pages/TripDetailPage/components/PaymentsSection/PaymensSection.vue'
</script>

<template>
  <SummarySection />

  <Tabs default-value="itinerary">
    <TabsList>
      <TabsTrigger value="tripCalendar"> Trip Calendar </TabsTrigger>
      <TabsTrigger value="itinerary"> Itinerary Builder </TabsTrigger>
      <TabsTrigger value="expenses"> Expenses </TabsTrigger>
      <TabsTrigger value="payments"> Payments </TabsTrigger>
      <TabsTrigger value="customers"> Customers </TabsTrigger>
    </TabsList>
    <TabsContent value="tripCalendar">
      <TripCalendarSection />
    </TabsContent>
    <TabsContent value="itinerary">
      <ItineraryBuilder />
    </TabsContent>
    <TabsContent value="expenses">
      <ExpensesCard />
    </TabsContent>
    <TabsContent value="payments">
      <PaymensSection />
    </TabsContent>
    <TabsContent value="customers">
      <TripCustomersTable />
    </TabsContent>
  </Tabs>
</template>
