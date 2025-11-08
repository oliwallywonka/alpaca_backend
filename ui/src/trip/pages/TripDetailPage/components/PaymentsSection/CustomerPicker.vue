<script setup lang="ts">
import { Badge } from '@/core/components/ui/badge'
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput,
} from '@/core/components/ui/input-group'

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/core/components/ui/dialog'
import { Button } from '@/core/components/ui/button'
import { computed, ref } from 'vue'
import CustomersPage from '@/customer/pages/CustomersPage.vue'
import type { Customer } from '@/customer/interfaces/customer'
import { useVModel } from '@vueuse/core'

const props = withDefaults(defineProps<{ modelValue?: Customer }>(), {
  modelValue: undefined,
})
const emit = defineEmits<{
  (e: 'update:modelValue', value: Customer): void
}>()
const selectedCustomer = useVModel(props, 'modelValue', emit, { passive: true })

const isOpen = ref(false)
const customer = computed(() => {
  if (!selectedCustomer.value) return
  return `${selectedCustomer.value.title} ${selectedCustomer.value.firstName} ${selectedCustomer.value.middleName} ${selectedCustomer.value.lastName}`
})
</script>
<template>
  <div>
    <InputGroup>
      <InputGroupAddon align="inline-start">
        <Badge v-if="selectedCustomer">{{ customer }}</Badge>
      </InputGroupAddon>
      <InputGroupInput disabled :placeholder="!customer ? 'No customer selected' : ''" />
      <InputGroupAddon align="inline-end">
        <InputGroupButton variant="secondary" @click="isOpen = true"> Search </InputGroupButton>
      </InputGroupAddon>
    </InputGroup>

    <Dialog v-model:open="isOpen">
      <DialogContent class="sm:max-w-4xl">
        <DialogHeader>
          <DialogTitle>Select Customer</DialogTitle>
          <DialogDescription> Select a customer from the list below. </DialogDescription>
          <DialogDescription v-if="!selectedCustomer"> No customer Selected </DialogDescription>
          <Badge v-if="selectedCustomer">{{ customer }}</Badge>
        </DialogHeader>
        <CustomersPage :showSelect="true" @select:customer="selectedCustomer = $event" />
        <DialogFooter>
          <Button type="submit" @click="isOpen = false"> Save changes </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </div>
</template>
