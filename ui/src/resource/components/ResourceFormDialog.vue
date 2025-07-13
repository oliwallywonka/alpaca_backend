<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
import * as z from 'zod'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'

import { LanguageInput } from '@/core/components/fields'
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
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectValue,
} from '@/core/components/ui/select'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import { ResourceService } from '../services/resourceService'
import { ResourceTypes } from '@/core/constants/resourceTypes'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  resourceID?: string
}>()

const startFetching = computed(() => !!props.resourceID && isOpen.value)
const { data, isLoading } = ResourceService.useGetOne(toRef(props.resourceID!), startFetching)
const { mutateAsync } = ResourceService.useCreate()
const { mutateAsync: updateMutate } = ResourceService.useUpdate()
const queryClient = useQueryClient()

const { handleSubmit, setFieldValue, resetForm, setValues } = useForm({
  validationSchema: toTypedSchema(
    z.object({
      name: z.record(z.string(), z.string().min(1, 'Every lang must has at least 1 character')),
      description: z.record(
        z.string(),
        z.string().optional(),
      ),
      // TODO: use an array string and do a multiselect component
      types: z.string().min(1, 'Type is required'),
    }),
  ),
  initialValues: {
    name: data.value?.name || {
      en: '',
    },
    description: data.value?.description || {
      en: '',
    },
    types: data.value?.types[0] || '',
  },
})

const onSubmit = handleSubmit(async (values) => {
  try {
    const { types: _, ...rest } = values
    if (props.resourceID) {
      await updateMutate({
        id: props.resourceID,
        data: {
          ...rest,
          types: [_],
        },
      })
      toast.success('Form submitted successfully')
      queryClient.invalidateQueries({ queryKey: ['resources'] })
      return
    }

    await mutateAsync({
      ...rest,
      types: [_],
    })
    toast.success('Form submitted successfully')
    queryClient.invalidateQueries({ queryKey: ['resources'] })
  } catch (error) {
    toast.error('Error submitting form')
    console.error('Error submitting form:', error)
  } finally {
    isOpen.value = false
  }
})

watch([isOpen], (newValue) => {
  if (newValue) {
    resetForm()
  }
})

watch([data, isOpen], ([newData, newIsOpen]) => {
  if (!newData || !newIsOpen) return
  setValues({
    name: newData?.name,
    description: newData?.description,
  })
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline"> <slot> open </slot> </Button>
    </DialogTrigger>
    <DialogContent class="w-full max-w-5xl">
      <DialogHeader>
        <DialogTitle>{{ props.resourceID ? 'Edit Service' : 'Create Service' }}</DialogTitle>
        <DialogDescription> Click save when you're done. </DialogDescription>
      </DialogHeader>

      <form id="resource-form" @submit.prevent="onSubmit">
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

        <FormField v-slot="{ componentField }" name="types">
          <FormItem>
            <FormLabel>Type</FormLabel>
            <Select v-bind="componentField">
              <FormControl>
                <SelectTrigger class="w-full">
                  <SelectValue placeholder="Select a service Type" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                <SelectGroup>
                  <SelectItem v-for="type of ResourceTypes" :key="type.value" :value="type.value">
                    {{ type.label }}
                  </SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>

      <DialogFooter>
        <Button form="resource-form" type="submit" :disabled="isLoading"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
