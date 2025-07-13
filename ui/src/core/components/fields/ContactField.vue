<script setup lang="ts">
import { Input } from '@/core/components/ui/input'
import { Button } from '@/core/components/ui/button'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/core/components/ui/select'
import type { ContactField } from '@/core/interfaces/fields'
import { useVModel } from '@vueuse/core'
import { PlusIcon, X } from 'lucide-vue-next'

const contactTypes = ['phone', 'email', 'address', 'url']

const props = withDefaults(
  defineProps<{
    modelValue?: ContactField[]
  }>(),
  {
    modelValue: () => [],
  },
)
const emit = defineEmits<{
  (e: 'update:modelValue', contacts: ContactField[]): void
}>()

const data = useVModel(props, 'modelValue', emit, {
  passive: true,
})

const addContact = () => {
  data.value.push({
    type: '',
    value: '',
  })
}

const removeContact = (index: number) => {
  data.value.splice(index, 1)
}
</script>

<template>
  <div class="flex flex-col gap-2">
    <Button variant="outline" @click.stop="addContact" type="button">
      <PlusIcon /> Add Contact
    </Button>
    <div v-for="(_, index) in data" :key="index" class="flex gap-2">
      <Select v-model:model-value="data[index].type">
        <SelectTrigger class="w-full">
          <SelectValue placeholder="Select a contact type" />
        </SelectTrigger>
        <SelectContent>
          <SelectGroup>
            <SelectItem
              v-for="(contact, index) of contactTypes"
              :key="`${contact}-${index}`"
              :value="contact"
              class="font-semibold"
            >
              {{ contact }}
            </SelectItem>
          </SelectGroup>
        </SelectContent>
      </Select>
      <Input v-model="data[index].value" />
      <Button variant="outline" @click="() => removeContact(index)"><X class="w-4 h-4" /></Button>
    </div>
  </div>
</template>
