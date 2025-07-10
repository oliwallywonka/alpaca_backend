<script setup lang="ts">
import { FileIcon } from 'lucide-vue-next';
import { onMounted, ref } from 'vue'

const props = defineProps<{
  file: File
}>()

const previewURL = ref<string>()

function generateThumb(file: File, width = 100, height = 100) {
  return new Promise((resolve) => {
    const reader = new FileReader()

    reader.onload = function (e) {
      const img = new Image()

      img.onload = function () {
        const canvas = document.createElement('canvas')
        const ctx = canvas.getContext('2d')
        const imgWidth = img.width
        const imgHeight = img.height

        canvas.width = width
        canvas.height = height

        ctx?.drawImage(
          img,
          imgWidth > imgHeight ? (imgWidth - imgHeight) / 2 : 0,
          0, // top aligned
          // imgHeight > imgWidth ? (imgHeight - imgWidth) / 2 : 0,
          imgWidth > imgHeight ? imgHeight : imgWidth,
          imgWidth > imgHeight ? imgHeight : imgWidth,
          0,
          0,
          width,
          height,
        )

        return resolve(canvas.toDataURL(file.type))
      }

      img.src = e.target?.result as string
    }

    reader.readAsDataURL(file)
  })
}

function loadPreviewURL() {
  generateThumb(props.file)
    .then((url) => {
      previewURL.value = url as string
    })
    .catch((err) => {
      previewURL.value = ''
      console.error(err)
    })
}

onMounted(() => {
  loadPreviewURL()
})
</script>

<template>
  <img v-if="previewURL" :src="previewURL" class="rounded-sm size-10 bg-gray-200"/>
  <div v-else class="bg-gray-200">
    <FileIcon  class="rounded-sm size-10"/>
  </div>
</template>
