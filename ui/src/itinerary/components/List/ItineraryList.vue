<script setup lang="ts">
import { ref, watch } from 'vue'
import { Plus, EllipsisVertical, Trash, Copy, Edit, RefreshCcwIcon } from 'lucide-vue-next'
import { useRoute } from 'vue-router'

import { Badge } from '@/core/components/ui/badge'
import { Button } from '@/core/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/core/components/ui/dropdown-menu'
import { Label } from '@/core/components/ui/label'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'
import ItineraryFormDialog from '@/itinerary/components/List/ItineraryFormDialog.vue'
import { ToggleGroup, ToggleGroupItem } from '@/core/components/ui/toggle-group'
import ItineraryCalendar from '@/itinerary/components/Calendar/ItineraryCalendar.vue'
import { ItineraryService } from '@/itinerary/services/itineraryService'
import { useParams } from '@/core/hooks/useParams'
import ItineraryDeleteDialog from './ItineraryDeleteDialog.vue'
import ExpensesSection from '../Expenses/ExpensesSection.vue'

const route = useRoute()
const tourID = route.params.tourID as string
const { params } = useParams({
  filter: `tour="${tourID}" && isActive=true && isTemplate=true`,
})

const { data, isLoading, refetch } = ItineraryService.useGetAll(params)
const currentItineraryID = ref<string>()

watch(data, (newData) => {
  if (!newData || !newData.items || newData.items.length === 0) return
  currentItineraryID.value = newData.items[0].id
})
</script>

<template>
  <div class="grid gap-4">
    <p v-if="isLoading" class="text-gray-500">Loading...</p>
    <p v-if="!data || data.items.length === 0" class="text-gray-500">
      Itineraries have not been found
    </p>
    <div class="grid xl:grid-cols-[auto_1fr] gap-2">
      <aside class="flex flex-row xl:flex-col gap-2">
        <Label>Itineraries</Label>
        <ToggleGroup
          :default-value="tourID"
          v-model:modelValue="currentItineraryID"
          class="h-auto flex flex-wrap xl:flex-col gap-2 items-start justify-center"
          variant="outline"
          type="single"
        >
          <div class="flex flex-wrap gap-2">
            <ItineraryFormDialog :tourID="tourID">
              <Button variant="outline" class="grow"> <Plus /> Itinerary </Button>
            </ItineraryFormDialog>
            <Button @click="refetch" :disabled="isLoading">
              <RefreshCcwIcon :class="{ 'animate-spin': isLoading }" />
            </Button>
          </div>
          <ToggleGroupItem
            v-for="variant in data?.items"
            :key="variant.id"
            :value="variant.id"
            class="rounded-md flex gap-2 justify-between"
          >
            <p class="font-semibold">
              MinMax Pax: {{ variant.minPersons }}-{{ variant.maxPersons }}
            </p>
            <div class="flex justify-between">
              <Badge>Price: ${{ variant.finalPrice.toFixed(2) }}</Badge>

              <DropdownMenu>
                <DropdownMenuTrigger @click.stop><EllipsisVertical /></DropdownMenuTrigger>
                <DropdownMenuContent>
                  {{ /** TODO: Duplicate variant */ '' }}
                  {{ /** TODO: Fix click propagation on dropdown and items */ '' }}
                  <DropdownMenuItem @click.stop><Copy /> Duplicate </DropdownMenuItem>

                  <DropdownMenuItem>
                    <ItineraryFormDialog :variantID="variant.id">
                      <button @click.stop class="w-full flex gap-2"><Edit /> Edit</button>
                    </ItineraryFormDialog>
                  </DropdownMenuItem>

                  <DropdownMenuItem>
                    <ItineraryDeleteDialog :variantID="variant.id">
                      <button class="w-full flex gap-2"><Trash /> Delete</button>
                    </ItineraryDeleteDialog>
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>
          </ToggleGroupItem>
        </ToggleGroup>
      </aside>
      <Label v-if="!currentItineraryID">Select a itinerary</Label>
      <Tabs v-if="currentItineraryID" default-value="expenses">
        <TabsList class="grid w-full grid-cols-2">
          <TabsTrigger value="itinerary"> Itinerary </TabsTrigger>
          <TabsTrigger value="expenses"> Operative Expenses </TabsTrigger>
        </TabsList>
        <TabsContent value="itinerary">
          <ItineraryCalendar :itineraryID="currentItineraryID" :tourID="tourID" />
        </TabsContent>
        <TabsContent value="expenses">
          <ExpensesSection :itineraryID="currentItineraryID" />
        </TabsContent>
      </Tabs>
    </div>
  </div>
</template>
