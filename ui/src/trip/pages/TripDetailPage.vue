<script setup lang="ts">
import { Card, CardContent, CardHeader, CardTitle } from '@/core/components/ui/card'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/core/components/ui/form'
import { Button, buttonVariants } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import { Badge } from '@/core/components/ui/badge'
import { cn } from '@/core/lib/utils'

import ExpensesCard from '@/trip/components/ExpensesCard.vue'
import PaymentsCard from '@/trip/components/PaymentsCard.vue'
import TripCalendar from '@/trip/components/TripCalendar.vue'
import ItineraryBuilder from '@/trip/components/ItineraryBuilder.vue'
import SummarySection from '@/trip/components/SummarySection.vue'
</script>

<template>
  <SummarySection />
  <!-- <PDFVoucher /> -->
  <Card>
    <CardHeader>
      <CardTitle>Client Details</CardTitle>
    </CardHeader>
    <CardContent>
      <p class="font-semibold">Name: Pedro Duarte</p>
      <p class="font-semibold">Email: pedro@duarte.com</p>
      <p class="font-semibold">Phone: pedro@duarte.com</p>
      <p class="font-semibold">Country: Texax, USA</p>
      <Button>Edit</Button>
    </CardContent>
  </Card>
  <Card>
    <CardHeader>
      <CardTitle>Trip Basic Information</CardTitle>
    </CardHeader>
    <CardContent class="grid grid-cols-2 gap-4">
      <div class="grid gap-4">
        <div class="travel grid grid-cols-[150px_1fr] border border-gray-200 rounded-md">
          <figure class="h-full">
            <img
              class="object-cover object-center w-full h-full rounded-l-md"
              src="https://images.unsplash.com/photo-1554629947-334ff61d85dc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=387&q=80"
              alt="Tour Image"
            />
          </figure>
          <div class="grid gap-2 p-4 bg-muted/50">
            <ul class="grid gap-1">
              <li class="text-lg font-bold">Tour Name</li>
              <li class="text-sm text-muted-foreground">Slug:</li>
              <li class="text-sm text-muted-foreground flex flex-wrap gap-0.5">
                Destinations: <Badge v-for="(_, index) in 8" :key="index">La Paz</Badge>
              </li>
              <li class="text-sm text-muted-foreground">Short Description:</li>
              <li class="text-sm text-muted-foreground">Long Description:</li>
              <li class="text-sm text-muted-foreground">Ref price pp:</li>
              <li class="text-sm text-muted-foreground">Discount:</li>
              <li class="text-sm text-muted-foreground">Days:</li>
              <li class="text-sm text-muted-foreground">Group Size:</li>
              <li class="text-sm text-muted-foreground">Transportation:</li>
              <li class="text-sm text-muted-foreground">Accommodation:</li>
              <li class="text-sm text-muted-foreground">Team:</li>
              <li class="text-sm text-muted-foreground">Public:</li>
            </ul>
            <div class="flex flex-wrap gap-2">
              <Button>Itinerary</Button>
              <RouterLink to="/tours/1" :class="cn(buttonVariants())">Edit</RouterLink>
              <Button>Delete</Button>
            </div>
          </div>
        </div>
      </div>

      <form>
        <FormField name="" v-slot="{ componentField }">
          <FormItem>
            <FormLabel >Price per person</FormLabel>
            <FormControl>
              <Input v-bind="componentField" />
            </FormControl>
            <FormDescription />
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="" v-slot="{ componentField }">
          <FormItem>
            <FormLabel>Price per person</FormLabel>
            <FormControl>
              <Input v-bind="componentField" />
            </FormControl>
            <FormDescription />
            <FormMessage />
          </FormItem>
        </FormField>
      </form>
    </CardContent>
  </Card>
  <ExpensesCard />
  <PaymentsCard />
  <Card>
    <CardContent>
      <Tabs default-value="account">
        <TabsList>
          <TabsTrigger value="account"> Travel Calendar </TabsTrigger>
          <TabsTrigger value="password"> Itinerary Builder </TabsTrigger>
        </TabsList>
        <TabsContent value="account" class="grid grid-cols-2 min-h-[85vh] gap-4">
          <TripCalendar />
        </TabsContent>
        <TabsContent value="password" class="grid grid-cols-3 min-h-[85vh] gap-4">
          <ItineraryBuilder />
        </TabsContent>
      </Tabs>
    </CardContent>
  </Card>
</template>
