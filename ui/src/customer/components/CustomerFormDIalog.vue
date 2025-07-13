<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { useForm } from 'vee-validate'
import { useQueryClient } from '@tanstack/vue-query'
import { toTypedSchema } from '@vee-validate/zod'
import { toast } from 'vue-sonner'
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
import { Input } from '@/core/components/ui/input'

import { CustomerService } from '../services/CustomerService'
import { ContactField } from '@/core/components/fields'

const isOpen = defineModel({ default: false })
const props = defineProps<{
  customerID?: string
}>()

const startFetching = computed(() => !!props.customerID && isOpen.value)

const queryClient = useQueryClient()
const getOneQuery = CustomerService.useGetOne(toRef(props.customerID!), startFetching)
const createQuery = CustomerService.useCreate()
const updateQuery = CustomerService.useUpdate()

const formSchema = toTypedSchema(
  z.object({
    title: z.enum(['Mr', 'Mrs', 'Miss', 'Ms', 'Dr']),
    firstName: z.string(),
    middleName: z.string().optional(),
    lastName: z.string(),
    dateOfBirth: z.string(),
    contacts: z.array(z.object({ type: z.string(), value: z.string() })),
  }),
)

const { handleSubmit, values, setFieldValue, resetForm, setValues } = useForm({
  validationSchema: formSchema,
  initialValues: {
    title: getOneQuery.data.value?.title || 'Mr',
    firstName: getOneQuery.data.value?.firstName || '',
    middleName: getOneQuery.data.value?.middleName || '',
    lastName: getOneQuery.data.value?.lastName || '',
    dateOfBirth: getOneQuery.data.value?.dateOfBirth || '',
    contacts: getOneQuery.data.value?.contacts || [],
  },
})

const onSubmit = handleSubmit(async (values) => {
  try {
    if (!props.customerID) {
      await createQuery.mutateAsync(values)
      toast.success('Customer created successfully')
    }
    if (props.customerID) {
      await updateQuery.mutateAsync({ id: props.customerID!, data: values })
      toast.success('Customer updated successfully')
    }
    await queryClient.invalidateQueries({ queryKey: ['customers'] })
  } catch (err) {
    console.log(err)
    toast.error('Error creating destination')
  } finally {
    isOpen.value = false
  }
})

watch(isOpen, (newValue) => {
  if (newValue) {
    resetForm()
  }
})

watch([isOpen, getOneQuery.data], ([newIsOpen, newData]) => {
  if (!newIsOpen || !newData) return

  setValues({
    title: newData.title,
    firstName: newData.firstName,
    middleName: newData.middleName,
    lastName: newData.lastName,
    dateOfBirth: newData.dateOfBirth,
    contacts: newData.contacts,
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
        <DialogTitle>{{ props.customerID ? 'Edit Destination' : 'New Destination' }}</DialogTitle>
        <DialogDescription>
          Create a new tour by filling out the form below. Make sure to provide all necessary
          details. {{ updateQuery.isPending }} {{ createQuery.isPending }} {{ values }}
        </DialogDescription>
      </DialogHeader>
      <div class="overflow-y-auto">
        <form id="customer-form" @submit.prevent="onSubmit" class="grid gap-4">
          <FormField v-slot="{ componentField }" name="title">
            <FormItem>
              <FormLabel>Title</FormLabel>
              <Select v-bind="componentField">
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a title" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectGroup>
                    <SelectItem
                      v-for="item in ['Mr', 'Mrs', 'Miss', 'Ms', 'Dr']"
                      :key="item"
                      :value="item"
                    >
                      {{ item }}
                    </SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="firstName">
            <FormItem>
              <FormLabel>First Name</FormLabel>
              <FormControl>
                <Input v-bind="componentField" placeholder="Enter first name" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="middleName">
            <FormItem>
              <FormLabel>Middle Name</FormLabel>
              <FormControl>
                <Input v-bind="componentField" placeholder="Enter middleName name" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="lastName">
            <FormItem>
              <FormLabel>Last Name</FormLabel>
              <FormControl>
                <Input v-bind="componentField" placeholder="Enter last name" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="dateOfBirth">
            <FormItem>
              <FormLabel>Date of Birth</FormLabel>
              <FormControl>
                <Input v-bind="componentField" placeholder="Enter date of birth" type="date" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="contacts">
            <FormItem>
              <FormLabel>Contacts</FormLabel>
              <FormControl>
                <ContactField
                  v-model="componentField.modelValue"
                  @update:model-value="setFieldValue('contacts', $event)"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
        </form>
      </div>
      <DialogFooter>
        <Button type="submit" form="customer-form"> Save changes </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
