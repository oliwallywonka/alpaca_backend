<script setup lang="ts">
import { ref } from 'vue'
import { useDebounceFn } from '@vueuse/core'

import { Loader, Search, X } from 'lucide-vue-next'
import { Input } from '@/core/components/ui/input'
import { Button } from '@/core/components/ui/button'
const props = defineProps<{
  placeholder?: string
  isLoading?: boolean
  debounceMilis?: number
}>()
const emit = defineEmits<{
  (e: 'change:debounce', value: string): void
}>()
const input = ref('')
const debouncedFn = useDebounceFn(() => {
  emit('change:debounce', input.value)
}, props.debounceMilis || 400)

function handleClose() {
  input.value = ''
  emit('change:debounce', '')
}
</script>
<template>
  <div class="relative w-full items-center">
    <Input
      id="search"
      type="text"
      autocomplete="off"
      class="pl-10"
      :placeholder="props.placeholder || 'Search...'"
      @input="debouncedFn"
      v-model="input"
    />
    <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
      <Search class="size-6 text-muted-foreground" />
    </span>
    <span class="absolute right-1 top-2">
      <Loader v-if="isLoading" class="animate-spin" />
    </span>
    <span v-if="input" class="absolute right-1 top-0">
      <Button @click="handleClose" variant="ghost">
        <X class="size-4" />
      </Button>
    </span>
  </div>
</template>
