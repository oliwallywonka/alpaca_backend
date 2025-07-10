<script setup lang="ts">
import { ref } from 'vue'
import { useSortable } from '@vueuse/integrations/useSortable'

import { Button } from '@/core/components/ui/button'
import { Plus, X } from 'lucide-vue-next'
import { Badge } from '@/core/components/ui/badge'
import PreviewImage from './PreviewImage.vue'
import ImageModal from '@/tour/components/ImageModal.vue'
import { useVModel } from '@vueuse/core'

const props = withDefaults(
  defineProps<{
    modelValue?: {
      newFiles: File[]
      deletedFiles: string[]
      uploadedFiles: string[]
    }
  }>(),
  {
    modelValue: () => ({
      newFiles: [],
      deletedFiles: [],
      uploadedFiles: [],
    }),
  },
)

const emit = defineEmits(['update:modelValue'])
const data = useVModel(props, 'modelValue', emit, { passive: true })

// UPLOADED FILES
const uploadedFilesContainer = ref<HTMLElement>()
const hasRemovedFiles = (name: string) => {
  return data.value.deletedFiles.includes(name)
}
const removeUploadedFile = (name: string) => {
  data.value.deletedFiles.push(name)
}
const restoreUploadedFile = (name: string) => {
  data.value.deletedFiles = data.value.deletedFiles.filter((file) => file !== name)
}

// NEW FILES
const newFilesContainer = ref<HTMLElement>()
const removeNewFile = (index: number) => {
  data.value.newFiles.splice(index, 1)
}
// DRAG AND DROP
const fileInput = ref<HTMLInputElement>()
const isDragOver = ref(false)
const handleDrop = (e: DragEvent) => {
  e.preventDefault()
  const files = e.dataTransfer?.files || []
  data.value.newFiles.push(...files)
  isDragOver.value = false
}

useSortable(newFilesContainer, data.value.newFiles)
useSortable(uploadedFilesContainer, data.value.uploadedFiles)
</script>
<template>
  <div
    @drop="handleDrop"
    @dragover.prevent="isDragOver = true"
    @dragleave="isDragOver = false"
    class="flex flex-col gap-2"
  >
    <ul ref="uploadedFilesContainer" class="flex flex-col gap-2">
      <li
        v-for="(file, i) in data.uploadedFiles"
        :key="file + i"
        :class="{ 'opacity-50': hasRemovedFiles(file) }"
        class="flex justify-between items-center cursor-pointer"
      >
        <div class="flex gap-2 items-center">
          <div class="size-10">
            <ImageModal :imageURL="file" />
          </div>
          <a
            :href="file"
            :class="{ 'line-through': hasRemovedFiles(file) }"
            target="_blank"
            class="whitespace-nowrap overflow-hidden w-60 text-ellipsis text-blue-500"
          >
            {{ file }}
          </a>
        </div>
        <Button
          v-if="!hasRemovedFiles(file)"
          @click="removeUploadedFile(file)"
          variant="ghost"
          type="button"
        >
          <X class="w-4 h-4" />
        </Button>
        <Button
          v-if="hasRemovedFiles(file)"
          @click="restoreUploadedFile(file)"
          variant="destructive"
          type="button"
        >
          Restore
        </Button>
      </li>
    </ul>
    <ul class="grid gap-2" ref="newFilesContainer">
      <li
        v-for="(file, i) in data.newFiles"
        v-memo="file.name"
        :key="file.name + i"
        class="flex justify-between items-center cursor-pointer"
      >
        <figure class="flex gap-2 items-center border rounded-md">
          <PreviewImage :file="file" />
          <Badge variant="blue">New</Badge>
          <span>{{ file.name }}</span>
        </figure>
        <Button variant="ghost" type="button" @click="removeNewFile(i)">
          <X class="w-4 h-4"/>
        </Button>
      </li>
    </ul>
    <input
      ref="fileInput"
      @change="() => data.newFiles.push(...(fileInput?.files || []))"
      type="file"
      class="hidden"
      multiple
    />
    <Button type="button" variant="outline" @click="fileInput?.click()" class="mt-auto">
      <Plus /> Add
    </Button>
  </div>
</template>
