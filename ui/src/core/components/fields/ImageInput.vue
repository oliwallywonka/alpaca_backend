<script setup lang="ts">
import { ref, useId } from 'vue'

import { PlusIcon, EditIcon } from 'lucide-vue-next'
import { useVModel } from '@vueuse/core'

const props = withDefaults(
  defineProps<{
    modelValue?: {
      newFile: File | null
      uploadedFile: string
    }
  }>(),
  {
    modelValue: () => ({
      newFile: null,
      uploadedFile: '',
    }),
  },
)
const emit = defineEmits(['update:modelValue'])
const data = useVModel(props, 'modelValue', emit, { passive: true })

const inputRef = ref<HTMLInputElement>()
const id = useId()

const handleImageUpload = (e: Event) => {
  const target = e.target as HTMLInputElement
  const file = target.files?.[0]
  if (!file) return
  data.value.newFile = file
  const reader = new FileReader()

  reader.onload = function (e) {
    data.value.uploadedFile = e.target?.result as string
  }
  reader.readAsDataURL(file)
}

const triggerFileInput = () => {
  inputRef.value?.click()
}
</script>

<template>
  <div
    class="w-full min-h-[25vh] max-h-[50vh] relative bg-gray-200 flex items-center justify-center overflow-hidden rounded-md"
  >
    <div
      v-if="data?.uploadedFile"
      class="w-full h-full relative cursor-pointer group"
      @click="triggerFileInput"
    >
      <img
        :src="data.uploadedFile"
        alt="Tour Image"
        class="w-full h-full relative bg-gray-200 object-cover rounded-md"
      />
      <div
        class="absolute inset-0 bg-black/50 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300"
      >
        <EditIcon class="w-8 h-8 text-white" />
      </div>
    </div>

    <label v-else :htmlFor="id" class="cursor-pointer flex flex-col items-center">
      <PlusIcon class="w-8 h-8 text-gray-500 mb-2" />
      <span class="text-gray-500">Upload an image</span>
    </label>

    <input
      ref="inputRef"
      type="file"
      :id="id"
      accept="image/*"
      name="image"
      className="hidden"
      @change="handleImageUpload"
    />
  </div>
</template>
