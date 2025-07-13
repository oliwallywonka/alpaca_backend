<script setup lang="ts">
import { ref, watch } from 'vue'
import { Plus, EllipsisVertical, Trash, Copy, Edit, RefreshCcwIcon } from 'lucide-vue-next'

import { Badge } from '@/core/components/ui/badge'
import { Button } from '@/core/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/core/components/ui/dropdown-menu'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'
import VariantFormDialog from '@/tour/components/TourVariants/VariantFormDialog.vue'
import { ToggleGroup, ToggleGroupItem } from '@/core/components/ui/toggle-group'
import ItineraryBuilder from '../ItineraryBuilder/ItineraryBuilder.vue'
import { TourVariantService } from '@/tour/services/tourVariantService'
import { useParams } from '@/core/hooks/useParams'
import { useRoute } from 'vue-router'
import VariantDeleteDialog from './VariantDeleteDialog.vue'
import ExpensesSection from '../Expenses/ExpensesSection.vue'

const route = useRoute()
const tourID = route.params.tourID as string
const { params } = useParams({
  filter: `tour="${tourID}" && isActive=true`,
})

const { data, isLoading, refetch } = TourVariantService.useGetAll(params)
const currentVariantID = ref()

watch(data, (newData) => {
  if (!newData || !newData.items.length) return
  currentVariantID.value = newData.items[0].id
})
</script>

<template>
  <div class="grid gap-4">
    <p v-if="isLoading" class="text-gray-500">Loading...</p>
    <p v-if="!data || data.items.length === 0" class="text-gray-500">
      TourVarianst have not been found
    </p>
    <ToggleGroup
      v-model:model-value="currentVariantID"
      class="w-full grid md:grid-cols-4 gap-2"
      variant="outline"
      type="single"
    >
      <div class="flex gap-2">
        <VariantFormDialog :tourID="tourID">
          <Button variant="outline" class="grow"> <Plus /> Add Variant </Button>
        </VariantFormDialog>
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
        <p class="font-semibold">MinMax Pax: {{ variant.minPersons }}-{{ variant.maxPersons }}</p>
        <div class="flex justify-between">
          <Badge>Price: ${{ variant.totalPrice }}</Badge>

          <DropdownMenu>
            <DropdownMenuTrigger @click.stop><EllipsisVertical /></DropdownMenuTrigger>
            <DropdownMenuContent>
              {{ /** TODO: Duplicate variant */ '' }}
              {{ /** TODO: Fix click propagation on dropdown and items */ '' }}
              <DropdownMenuItem @click.stop><Copy /> Duplicate </DropdownMenuItem>

              <DropdownMenuItem>
                <VariantFormDialog :variantID="variant.id">
                  <button @click.stop class="w-full flex gap-2"><Edit /> Edit</button>
                </VariantFormDialog>
              </DropdownMenuItem>

              <DropdownMenuItem>
                <VariantDeleteDialog :variantID="variant.id">
                  <button class="w-full flex gap-2"><Trash /> Delete</button>
                </VariantDeleteDialog>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </ToggleGroupItem>
    </ToggleGroup>

    <Tabs v-if="currentVariantID" default-value="expenses">
      <TabsList class="grid w-full grid-cols-2">
        <TabsTrigger value="itinerary"> Itinerary </TabsTrigger>
        <TabsTrigger value="expenses"> Operative Expenses </TabsTrigger>
      </TabsList>
      <TabsContent value="itinerary">
        <ItineraryBuilder :tourVariantID="currentVariantID" />
      </TabsContent>
      <TabsContent value="expenses">
        <ExpensesSection :tourVariantID="currentVariantID" />
      </TabsContent>
    </Tabs>
  </div>
</template>
