<script setup lang="ts">
import { ref, useAttrs } from 'vue'
import { useVModel } from '@vueuse/core'
import { X } from 'lucide-vue-next'

import { Label } from '@/core/components/ui/label'
import { Badge } from '@/core/components/ui/badge'
import { Textarea } from '@/core/components/ui/textarea'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/core/components/ui/select'
import { Button } from '@/core/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/core/components/ui/dialog'

const attrs = useAttrs()

const props = defineProps<{
  modelValue: Record<string, string>
  showAddLanguage?: boolean
}>()
const emit = defineEmits(['update:modelValue'])
const data = useVModel(props, 'modelValue', emit)

const laguagesCodes = ['en', 'es']
const languageToRemove = ref<string>()
const isDialogOpen = ref(false)

const openDialog = (lang: string) => {
  isDialogOpen.value = true
  languageToRemove.value = lang
}

const removeLanguage = () => {
  if (languageToRemove.value && data.value) {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const { [languageToRemove.value]: _, ...rest } = data.value
    data.value = rest
    isDialogOpen.value = false
    languageToRemove.value = undefined
  }
}

const onChangeLanguage = (lang: string, value: string) => {
  if (lang in data.value) {
    return
  }
  data.value[lang] = value
}
</script>
<template>
  <div v-for="(_, key) in data" :key="key" class="flex items-center gap-2">
    <Badge>{{ key }}</Badge>
    <Textarea v-model="data[key]" rows="4" v-bind="key === Object.keys(data)[0] ? attrs : {}" />
    <Button v-if="key !== 'en'" @click="openDialog(key)" variant="ghost">
      <X class="h-4 w-4" />
    </Button>
  </div>
  <div v-if="showAddLanguage" class="flex items-center gap-2">
    <Label htmlFor="add-language"> Add: </Label>

    <Select @update:model-value="(value) => onChangeLanguage(value as string, '')">
      <SelectTrigger id="add-language" class="w-[180px]">
        <SelectValue placeholder="Select a language" />
      </SelectTrigger>
      <SelectContent>
        <SelectItem v-for="lang in laguagesCodes" :key="lang" :value="lang">
          {{ lang.toUpperCase() }}
        </SelectItem>
      </SelectContent>
    </Select>

    <Dialog v-bind:open="isDialogOpen">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Remove Language</DialogTitle>
          <DialogDescription>
            Are you sure you want to remove the
            {{ languageToRemove?.toUpperCase() }} language input? This action cannot be undone.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button variant="outline" @click="isDialogOpen = false"> Cancel </Button>
          <Button @click="removeLanguage"> Remove </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  </div>
</template>
