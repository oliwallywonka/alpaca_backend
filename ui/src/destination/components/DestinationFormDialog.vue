<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

import { Button } from '@/core/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/core/components/ui/dialog'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/core/components/ui/select'
import { toast } from 'vue-sonner'
import LanguageInput from '@/core/components/fields/LanguageInput.vue'
import MapInput from '@/core/components/fields/MapInput.vue'
import { useForm } from 'vee-validate'
import { computed, toRef, watch } from 'vue'
import { DestinationService } from '../services/DestinationService'
import { useQueryClient } from '@tanstack/vue-query'

const isOpen = defineModel({ default: false })
const props = defineProps<{
  destinationID?: string
}>()

const startFetching = computed(() => !!props.destinationID && isOpen.value)

const queryClient = useQueryClient()
const getAllQuery = DestinationService.useGetAll(toRef({ perPage: 100 }))
const getOneQuery = DestinationService.useGetOne(toRef(props.destinationID!), startFetching)
const createQuery = DestinationService.useCreate()
const updateQuery = DestinationService.useUpdate()

const formSchema = toTypedSchema(
  z.object({
    name: z.record(z.string(), z.string()),
    description: z.record(z.string(), z.string()),
    parent: z.string(),
    location: z.object({
      lat: z.number(),
      lon: z.number(),
    }),
  }),
)

const { handleSubmit, values, setFieldValue, resetForm, setValues } = useForm({
  validationSchema: formSchema,
  initialValues: {
    name: { en: '' },
    description: { en: '' },
    parent: '',
    location: { lat: 0, lon: 0 },
  },
})

const onSubmit = handleSubmit(async (values) => {
  try {
    if (!props.destinationID) {
      await createQuery.mutateAsync(values)
      toast.success('Destination created successfully')
    }
    if (props.destinationID) {
      await updateQuery.mutateAsync({ id: props.destinationID!, data: values })
      toast.success('Destination updated successfully')
    }
    await queryClient.invalidateQueries({ queryKey: ['destinations'] })
  } catch (err) {
    console.log(err)
    toast.error('Error creating destination')
  } finally {
    isOpen.value = false
  }
})

function addLanguage(lang: string) {
  if (values.description && !Object.keys(values.description).includes(lang)) {
    setFieldValue('description', { ...values.description, [lang]: '' })
  }
  if (values.name && !Object.keys(values.name).includes(lang)) {
    setFieldValue('name', { ...values.name, [lang]: '' })
  }
}

watch(isOpen, (newValue) => {
  if (newValue) {
    resetForm()
  }
})

watch([isOpen, getOneQuery.data], ([newIsOpen, newData]) => {
  if (!newIsOpen || !newData) return

  setValues({
    name: newData.name,
    description: newData.description,
    parent: newData.parent,
    location: newData.location,
  })
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline">
        <slot> open </slot>
      </Button>
    </DialogTrigger>
    <DialogContent class="sm:max-w-4xl grid-rows-[auto_minmax(0,1fr)_auto] max-h-[90dvh]">
      <DialogHeader>
        <DialogTitle>{{ props.destinationID ? 'Edit Destination' : 'New Destination' }}</DialogTitle>
        <DialogDescription>
          Create a new tour by filling out the form below. Make sure to provide all necessary
          details. {{ updateQuery.isPending }} {{ createQuery.isPending }}
        </DialogDescription>
      </DialogHeader>
      <div class="overflow-y-auto">
        <Select @update:model-value="(value) => addLanguage(value as string)">
          <SelectTrigger id="add-language" class="w-[180px]">
            <SelectValue placeholder="Add Language" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem v-for="lang in ['en', 'es']" :key="lang" :value="lang">
                {{ lang.toUpperCase() }}
              </SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
        <form id="destination-form" @submit.prevent="onSubmit" class="grid gap-4">
          <FormField v-slot="{ componentField }" name="name">
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <LanguageInput
                  v-model="componentField.modelValue"
                  @update:model-value="setFieldValue('name', $event)"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="description">
            <FormItem>
              <FormLabel>Description</FormLabel>
              <FormControl>
                <LanguageInput
                  v-model="componentField.modelValue"
                  @update:model-value="setFieldValue('description', $event)"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="parent">
            <FormItem>
              <FormLabel>Parent Destination</FormLabel>
              <Select v-bind="componentField">
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a parent destination" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectGroup>
                    <SelectItem
                      v-for="item of getAllQuery.data.value?.items"
                      :key="item.id"
                      :value="item.id"
                    >
                      {{ item.name?.en }}
                    </SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="location">
            <FormItem>
              <FormLabel>Location</FormLabel>
              <FormControl>
                <MapInput
                  v-model="componentField.modelValue"
                  @update:model-value="setFieldValue('location', $event)"
                />
              </FormControl>
            </FormItem>
          </FormField>
        </form>
      </div>
      <DialogFooter>
        <Button type="submit" form="destination-form"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
