<script setup lang="ts">
import { computed, ref, toRef, watch } from 'vue'
import { toTypedSchema } from '@vee-validate/zod'
import { z } from 'zod'
import { useForm } from 'vee-validate'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from '@/core/components/ui/dialog'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import { Button } from '@/core/components/ui/button'
import { Input } from '@/core/components/ui/input'
import { ItineraryService } from '@/itinerary/services/itineraryService'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'

const props = defineProps<{
  variantID?: string
  tourID?: string
}>()

const isOpen = ref(false)

const startFetching = computed(() => !!props.variantID || isOpen.value)

const query = useQueryClient()
const { data } = ItineraryService.useGetOne(toRef(props.variantID as string), startFetching)
const update = ItineraryService.useUpdate()
const create = ItineraryService.useCreate()

const formSchema = toTypedSchema(
  z.object({
    minPersons: z.number().min(1, 'Min persons must be greater than 0'),
    maxPersons: z.number().min(1, 'Max persons must be greater than 0'),
    finalPrice: z.number().min(1, 'Final price must be greater than 0'),
  }),
)

const { handleSubmit, values, setValues, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues: {
    minPersons: data.value?.minPersons || 1,
    maxPersons: data.value?.maxPersons || 1,
    finalPrice: data.value?.finalPrice || 0,
  },
})

const onSubmit = handleSubmit(async (values) => {
  try {
    if (props.variantID) {
      await update.mutateAsync({
        id: props.variantID,
        data: { ...values, tour: props.tourID, isActive: true },
      })
      toast.success('Variant updated successfully')
      query.invalidateQueries({ queryKey: ['itineraries'] })
    }
    if (!props.variantID) {
      await create.mutateAsync({ ...values, tour: props.tourID, isActive: true, isTemplate: true })
      toast.success('Variant created successfully')
      query.invalidateQueries({ queryKey: ['itineraries'] })
    }
  } catch (err) {
    console.log(err)
    toast.error('Error creating variant')
  } finally {
    isOpen.value = false
  }
})

watch(isOpen, (newIsOpen) => {
  if (!newIsOpen) return
  resetForm()
})

watch([isOpen, data], ([newIsOpen, newData]) => {
  if (!newIsOpen || !newData) return
  setValues({
    minPersons: newData.minPersons,
    maxPersons: newData.maxPersons,
    finalPrice: newData.finalPrice,
  })
})
</script>
<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <slot>
        <Button variant="outline"> Open </Button>
      </slot>
    </DialogTrigger>
    <DialogContent class="sm:max-w-md">
      <DialogHeader>
        <DialogTitle>{{ props.variantID ? 'Edit Variant' : 'New Variant' }}</DialogTitle>
        <DialogDescription> Complete the form {{ values }}</DialogDescription>
      </DialogHeader>

      <form id="variant-form" class="grid grid-cols-2 gap-2" @submit.prevent="onSubmit">
        <FormField v-slot="{ componentField }" name="minPersons">
          <FormItem class="flex flex-col">
            <FormLabel>Min Persons</FormLabel>
            <FormControl>
              <Input type="number" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="maxPersons">
          <FormItem class="flex flex-col">
            <FormLabel>Max Persons</FormLabel>
            <FormControl>
              <Input type="number" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="finalPrice">
          <FormItem class="col-span-2 flex flex-col">
            <FormLabel>USD Price</FormLabel>
            <FormControl>
              <Input type="number" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>
      <DialogFooter>
        <DialogClose as-child>
          <Button type="button" variant="secondary"> Close </Button>
        </DialogClose>
        <Button type="submit" form="variant-form"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
