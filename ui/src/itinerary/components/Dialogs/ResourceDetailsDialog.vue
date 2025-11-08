<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { useQueryClient } from '@tanstack/vue-query'
import { z } from 'zod'
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import { toast } from 'vue-sonner'
import { Trash } from 'lucide-vue-next'
import { useParams } from '@/core/hooks/useParams'

import { Button } from '@/core/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/core/components/ui/dialog'
import { FormField, FormItem, FormControl, FormLabel, FormMessage } from '@/core/components/ui/form'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/core/components/ui/select'
import { Badge } from '@/core/components/ui/badge'
import PriceListField from '@/core/components/fields/PriceListField.vue'
import ConfirmDeleteDialog from '@/core/components/Dialogs/ConfirmDeleteDialog.vue'

import { ItineraryResourceService } from '@/itinerary/services/itineraryResources'
import { ResourceProviderService } from '@/providerResource/providerResourcesService'
import type { ResourceProvider } from '@/providerResource/ResourceProvider'

/* TODO: Fix Reactivity, it stop working after changing view */
const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  itineraryResourceID: string
}>()

const query = useQueryClient()

const { params } = useParams({
  expand: 'resourceProvider.resource, ',
})
const { data: itineraryResource, isLoading: irIsLoading } = ItineraryResourceService.useGetOne(
  toRef(props.itineraryResourceID),
  toRef(true),
  params,
)
const update = ItineraryResourceService.useUpdate()
const remove = ItineraryResourceService.useDelete()

const filter = computed(() => {
  if (!itineraryResource.value) return undefined
  return `resource="${itineraryResource.value?.expand.resourceProvider.resource}"`
})
const startFetching = computed(() => {
  if (!itineraryResource.value || !filter.value) return false
  return true
})

const rpParams = computed(() => ({
  filter: filter.value,
  expand: 'provider,user',
}))

const { data: resourceProviders, isLoading: rpIsLoading } = ResourceProviderService.useGetAll(
  rpParams,
  startFetching,
)

const schema = toTypedSchema(
  z.object({
    resourceProvider: z.string(),
    price: z.object({
      minPersons: z.number(),
      maxPersons: z.number(),
      price: z.number(),
      currency: z.string(),
    }),
  }),
)

const initialValues = {
  resourceProvider: itineraryResource.value?.expand.resourceProvider.id || '',
  price: {
    price: itineraryResource.value?.originalCost || 0.0,
    minPersons: itineraryResource.value?.minPersons || 1,
    maxPersons: itineraryResource.value?.maxPersons || 1,
    currency: itineraryResource.value?.currency || 'USD',
  },
}

const { values, handleSubmit, setValues, resetForm } = useForm({
  validationSchema: schema,
  initialValues,
})

const currentResourceProvider = computed(() => {
  if (!resourceProviders.value) return undefined
  return resourceProviders.value.items.find((rp) => rp.id === values.resourceProvider)
})

const onSubmit = handleSubmit(async ({ price, resourceProvider }) => {
  try {
    await update.mutateAsync({
      id: props.itineraryResourceID,
      data: {
        originalCost: price.price,
        minPersons: price.minPersons,
        maxPersons: price.maxPersons,
        currency: price.currency,
        resourceProvider: resourceProvider,
      },
    })
    await query.invalidateQueries({ queryKey: ['itineraryResources'] })
    await query.invalidateQueries({ queryKey: ['itineraries'] })
    toast.success("Resource's price updated successfully")
  } catch (err) {
    console.log(err)
    toast.error('Error updating resource')
  } finally {
    isOpen.value = false
  }
})

async function handleDelete() {
  if (!props.itineraryResourceID) return
  try {
    await remove.mutateAsync(props.itineraryResourceID)
    query.invalidateQueries({ queryKey: ['itineraryResources'] })
    toast.success('Resource deleted successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error deleting Resource')
  } finally {
    isOpen.value = false
  }
}

function isUser(resourceProvider?: ResourceProvider) {
  if (!resourceProvider) return false
  return !!resourceProvider.user
}

function isProvider(resourceProvider?: ResourceProvider) {
  if (!resourceProvider) return false
  return !!resourceProvider.provider
}

watch(isOpen, (newIsOpen) => {
  if (!newIsOpen) return
  resetForm()
})

watch([isOpen, itineraryResource], ([newIsOpen, newIR]) => {
  if (!newIsOpen || !newIR) return
  setValues({
    resourceProvider: newIR.expand.resourceProvider.id,
    price: {
      minPersons: newIR.minPersons,
      maxPersons: newIR.maxPersons,
      price: newIR.originalCost,
      currency: newIR.currency,
    },
  })
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogContent class="sm:max-w-4xl">
      <p v-if="irIsLoading" class="text-muted-foreground text-sm">Loading...</p>
      <DialogHeader v-if="!irIsLoading">
        <DialogTitle>
          Resource:
          {{ itineraryResource?.expand.resourceProvider.expand.resource.name.en }}
        </DialogTitle>
        <DialogDescription> Resources Details </DialogDescription>
      </DialogHeader>
      <form
        v-if="!rpIsLoading && !irIsLoading"
        id="itinerary-resource-form"
        @submit.prevent="onSubmit"
        class="grid grid-cols-2 gap-4 place-items-start"
      >
        <FormField v-slot="{ componentField }" name="resourceProvider">
          <FormItem>
            <FormLabel>Providers</FormLabel>

            <Select v-bind="componentField">
              <FormControl>
                <SelectTrigger>
                  <SelectValue placeholder="Select a Provider">
                    <Badge>{{ isProvider(currentResourceProvider) ? 'Provider' : 'User' }}</Badge>
                    <span v-if="isProvider(currentResourceProvider)">
                      {{ currentResourceProvider?.expand.provider.fullName }}
                    </span>
                    <span v-if="isUser(currentResourceProvider)">
                      {{ currentResourceProvider?.expand.user.name }}
                    </span>
                  </SelectValue>
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                <SelectItem
                  v-for="resourceProvider in resourceProviders!.items"
                  :value="resourceProvider.id"
                  :key="resourceProvider.id"
                >
                  <div v-if="isProvider(resourceProvider)" class="flex items-center gap-2 w-full">
                    <Badge> Provider: </Badge>
                    <Badge> {{ resourceProvider.type[0] }} </Badge>
                    {{ resourceProvider.expand.provider.fullName }}
                  </div>
                  <div v-if="isUser(resourceProvider)" class="flex items-center gap-2 w-full">
                    <Badge> User: </Badge>
                    <Badge> {{ resourceProvider.type[0] }} </Badge>
                    {{ resourceProvider.expand.user.name }}
                  </div>
                </SelectItem>
              </SelectContent>
            </Select>
            <FormMessage />
          </FormItem>
        </FormField>
        <FormField name="price" v-slot="{ componentField }">
          <FormItem>
            <FormLabel>Prices List</FormLabel>
            <!-- TODO fix price list field -->
            <PriceListField
              v-model="componentField.modelValue"
              @update:modelValue="
                setValues({
                  price: {
                    minPersons: componentField.modelValue?.minPersons,
                    maxPersons: componentField.modelValue?.maxPersons,
                    price: componentField.modelValue?.price,
                    currency: componentField.modelValue?.currency,
                  },
                })
              "
              :priceList="currentResourceProvider?.refPrices || []"
            />
            <FormMessage />
          </FormItem>
        </FormField>
      </form>
      <DialogFooter>
        <ConfirmDeleteDialog @confirm:delete="handleDelete">
          <Button variant="destructive"><Trash /> Delete Resource</Button>
        </ConfirmDeleteDialog>
        <Button type="submit" form="itinerary-resource-form"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
