<script setup lang="ts">
import type { HTMLAttributes } from 'vue'
import { cn } from '@/core/lib/utils'

const props = defineProps<{
  class?: HTMLAttributes['class']
}>()
</script>

<template>
  <table data-slot="table" :class="cn('w-full caption-bottom text-sm', props.class)">
    <slot />
  </table>
</template>
