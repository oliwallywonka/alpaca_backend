<script setup lang="ts">
import Badge from '@/core/components/ui/badge/Badge.vue'
import type { LanguageField } from '@/core/interfaces/fields'

const props = withDefaults(defineProps<{ data: LanguageField }>(), {})
</script>

<template>
  <p v-for="(value, key) of props.data" :key="key" class="">
    <Badge>{{ key }}</Badge>
    <span>{{ value }}</span>
  </p>
</template>
