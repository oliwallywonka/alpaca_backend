<script setup lang="ts">
import { ref } from 'vue'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'

import type { TourResource } from '@/tour/interfaces/tourResources'
import type { ResourceProvider } from '@/providerResource/ResourceProvider'
import { TourResourcesService } from '@/tour/services/tourResources'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/core/components/ui/dialog'
import ResourceProviderPicker from '../ResourceProviderPicker/ResourceProviderPicker.vue'
import { Badge } from '@/core/components/ui/badge'
import { Button } from '@/core/components/ui/button'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  tourResource: Partial<TourResource>
}>()
const currentResource = ref<ResourceProvider | undefined>()
const create = TourResourcesService.useCreate()
const query = useQueryClient()

const handleSubmit = async () => {
  if (!currentResource.value) return
  try {
    await create.mutateAsync({
      ...props.tourResource,
      resourceProvider: currentResource.value.id,
      originalCost: currentResource.value.refPrices[0].price,
      profitPercent: 0,
      currency: currentResource.value.refPrices[0].currency,
      dollarChangeRate: 10,
      quantity: 1,
      isVisible: true,
    })
    query.invalidateQueries({ queryKey: ['tourResources'] })
    toast.success('Resource created successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error creating resource')
  } finally {
    isOpen.value = false
  }
}
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Resources and providers</DialogTitle>
        <DialogDescription>
          Select the resource and provider of the itinerary "this will be showed on the summary
          section"
        </DialogDescription>
      </DialogHeader>

      <p>
        <span class="font-semibold text-sm">* Selected resource: </span>
        <span v-if="!currentResource" class="text-muted-foreground text-sm">
          No resource selected
        </span>
      </p>
      <div v-if="currentResource" class="flex flex-wrap gap-1">
        <Badge>{{ currentResource.expand.resource.name.en }}</Badge>
        <Badge>{{
          currentResource.expand.user?.name || currentResource.expand.provider?.fullName
        }}</Badge>
        <Badge>{{
          `
            ${currentResource.refPrices[0].minPersons} |
            ${currentResource.refPrices[0].maxPersons} |
            ${currentResource.refPrices[0].price} ${currentResource.refPrices[0].currency}
          `
        }}</Badge>
      </div>

      <ResourceProviderPicker @select:resource-provider="currentResource = $event" />

      <DialogFooter>
        <Button @click="handleSubmit">Save</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
