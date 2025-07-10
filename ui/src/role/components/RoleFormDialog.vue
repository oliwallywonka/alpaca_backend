<script setup lang="ts">
import { computed, toRef, watch } from 'vue'
import { toast } from 'vue-sonner'
import { useQueryClient } from '@tanstack/vue-query'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { useForm } from 'vee-validate'

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

import { Input } from '@/core/components/ui/input'
import { Button } from '@/core/components/ui/button'
import { EditIcon, PlusIcon } from 'lucide-vue-next'
import { PermissionsField } from '@/core/components/fields'
import { ScrollArea } from '@/core/components/ui/scroll-area'

import { RoleService } from '../services/roleService'

const isOpen = defineModel<boolean>({ default: false })

const props = defineProps<{
  roleID?: string
}>()

const startFetching = computed(() => !!props.roleID && isOpen.value)

const queryClient = useQueryClient()
const { data, isLoading } = RoleService.useGetOne(toRef(props.roleID!), startFetching)
const { mutateAsync: createMutate, isPending: isCreatePending } = RoleService.useCreate()
const { mutateAsync: updateMutate, isPending: isUpdatePending } = RoleService.useUpdate()

const formSchema = toTypedSchema(
  z.object({
    name: z.string().min(1, 'Name is required'),
    description: z.string().optional(),
    permissions: z.array(z.string()).optional(),
  }),
)

const { handleSubmit, setFieldValue, setValues } = useForm({
  validationSchema: formSchema,
  initialValues: {
    name: '',
    description: '',
    permissions: [],
  },
})

const onSubmit = handleSubmit(async (values) => {
  console.log('Form submitted!', values)
  try {
    if (props.roleID) {
      await updateMutate({ id: props.roleID!, data: values })
      toast.success('Role updated successfully')
    }

    if (!props.roleID) {
      await createMutate(values)
      toast.success('Role created successfully')
    }

    queryClient.invalidateQueries({ queryKey: ['roles'] })
  } catch (err) {
    console.log(err)
    toast.error('Error updating role')
  }
  isOpen.value = false
})

watch([data, isOpen], ([newData, newIsOpen]) => {
  if (!newData || !newIsOpen) return
  setValues({
    name: newData?.name,
    description: newData?.description,
    permissions: newData?.permissions,
  })
})
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="outline">
        <PlusIcon v-if="!props.roleID" class="w-4 h-4" />
        <EditIcon v-if="props.roleID" class="w-4 h-4" />
        {{ props.roleID ? 'Edit Role' : 'Create Role' }}
      </Button>
    </DialogTrigger>
    <DialogContent class="w-full max-w-5xl">
      <DialogHeader>
        <DialogTitle>{{ props.roleID ? 'Edit Role' : 'Create Role' }} </DialogTitle>
        <DialogDescription> * Are mandatory fields </DialogDescription>
      </DialogHeader>
      <form
        v-if="!isLoading"
        id="role-form"
        @submit.prevent="onSubmit"
        class="grid grid-cols-2 gap-4"
      >
        <div class="flex flex-col gap-4">
          <FormField v-slot="{ componentField }" name="name">
            <FormItem>
              <FormLabel>*Name</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Rol name" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField v-slot="{ componentField }" name="description">
            <FormItem>
              <FormLabel>Description</FormLabel>
              <FormControl>
                <Input type="text" placeholder="Rol description" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
        </div>

        <FormField v-slot="{ componentField }" name="permissions">
          <FormItem>
            <FormLabel>Permissions</FormLabel>
            <FormControl>
              <ScrollArea class="max-h-[400px]">
                <PermissionsField
                  v-model="componentField.modelValue"
                  @update:model-value="setFieldValue('permissions', $event)"
                />
              </ScrollArea>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>

      <p v-if="isLoading">Loading...</p>
      <DialogFooter>
        <Button type="submit" form="role-form" :disabled="isCreatePending || isUpdatePending">
          Save changes
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
