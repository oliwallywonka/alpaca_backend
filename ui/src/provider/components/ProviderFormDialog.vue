<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'
import { toTypedSchema } from '@vee-validate/zod'
import { useForm } from 'vee-validate'
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
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
  SelectGroup,
} from '@/core/components/ui/select'
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/core/components/ui/form'
import { Input } from '@/core/components/ui/input'
import ContactField from '@/core/components/fields/ContactField.vue'
import { ProviderService } from '../services/providerService'
import { ProviderTypes } from '@/core/constants/providerTypes'

const isOpen = defineModel<boolean>({ default: false })
const props = defineProps<{
  providerID?: string
}>()

const startFetching = computed(() => !!props.providerID && isOpen.value)
const { mutateAsync: createMutate, isPending: isPendingCreate } = ProviderService.useCreate()
const { mutateAsync: updateMutate, isPending: isPendingUpdate } = ProviderService.useUpdate()
const { data } = ProviderService.useGetOne(toRef(props.providerID!), startFetching)
const queryClient = useQueryClient()

const formSchema = toTypedSchema(
  z.object({
    fullName: z.string().min(2).max(50),
    description: z.string().optional(),
    contacts: z.array(z.object({ type: z.string(), value: z.string() })),
    types: z.string(),
  }),
)

const initialValues = {
  fullName: data.value?.fullName || '',
  description: data.value?.description || '',
  contacts: data.value?.contacts || [],
  types: data.value?.types[0] || '',
}

const { handleSubmit, setFieldValue, resetForm, setValues } = useForm({
  validationSchema: formSchema,
  initialValues,
})

const onSubmit = handleSubmit(async (values) => {
  const { types: _, ...rest } = values
  try {
    if (props.providerID) {
      await updateMutate({ id: props.providerID!, data: { ...rest, types: [_] } })
      toast.success('Provider updated successfully')
      queryClient.invalidateQueries({ queryKey: ['providers'] })
    }
    if (!props.providerID) {
      await createMutate({ ...rest, types: [_] })
      toast.success('Provider created successfully')
      queryClient.invalidateQueries({ queryKey: ['providers'] })
    }
  } catch (err) {
    console.log(err)
    toast.error('Error creating provider')
  } finally {
    isOpen.value = false
  }
})

// TODO: Fix form empty context after reopen dialog!!!! This happends on most Dialogs
// this Could work if on the previous watch i use setValues and de newData watch value
watch(isOpen, (newIsOpen) => {
  if (!newIsOpen) return
  resetForm()
})

watch([data, isOpen], ([newData, newIsOpen]) => {
    if (!newData || !newIsOpen) return
    setValues({
      fullName: newData.fullName,
      description: newData.description,
      contacts: newData.contacts,
      types: newData.types[0] || ''
    })
  },
)
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline">
        <slot> Open </slot>
      </Button>
    </DialogTrigger>
    <DialogContent class="sm:max-w-2xl">
      <DialogHeader>
        <DialogTitle>{{ props.providerID ? 'Edit' : 'Create' }} Provider</DialogTitle>
        <DialogDescription>
          Make changes to the provider selected here. Click save when you're done.
        </DialogDescription>
      </DialogHeader>

      <form id="provider-form" @submit.prevent="onSubmit" class="grid gap-3">
        <FormField v-slot="{ componentField }" name="fullName">
          <FormItem>
            <FormLabel>Name</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Provider Name" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="description">
          <FormItem>
            <FormLabel>Description</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Provider Description" v-bind="componentField" />
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
                  <SelectItem v-for="type of ProviderTypes" :key="type.value" :value="type.value">
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
        <Button type="submit" form="provider-form" :disabled="isPendingCreate || isPendingUpdate">
          Save changes
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
