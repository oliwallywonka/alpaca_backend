<script setup lang="ts">
import { Card, CardContent } from '@/core/components/ui/card'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'

import ExpensesCard from '@/trip/components/ExpensesCard.vue'
import PaymentsCard from '@/trip/components/PaymentsCard.vue'
import ItineraryBuilder from '@/trip/components/ItineraryBuilder.vue'
import SummarySection from '@/trip/components/SummarySection.vue'
import TripCustomersTable from '@/trip/components/TripCustomersTable.vue'
import TripCalendarSection from '@/trip/components/TripCalendarSection/TripCalendarSection.vue'
</script>

<template>
  <SummarySection />
  <Card>
    <CardContent>
      <Tabs default-value="tripCalendar">
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
        <TabsContent value="itinerary" class="grid grid-cols-3 gap-4">
          <ItineraryBuilder />
        </TabsContent>
        <TabsContent value="expenses">
          <ExpensesCard />
        </TabsContent>
        <TabsContent value="payments">
          <PaymentsCard />
        </TabsContent>
        <TabsContent value="customers">
          <TripCustomersTable />
        </TabsContent>
      </Tabs>
    </CardContent>
  </Card>
</template>
