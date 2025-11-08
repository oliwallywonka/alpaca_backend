<script setup lang="ts">
import { cn } from '@/core/lib/utils'
import { ChevronDown } from 'lucide-vue-next'
import { AccordionHeader, AccordionTrigger, type AccordionTriggerProps } from 'reka-ui'
import { computed, type HTMLAttributes } from 'vue'

const props = defineProps<AccordionTriggerProps & { class?: HTMLAttributes['class'] }>()

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props

  return delegated
})
</script>

<template>
  <AccordionHeader class="flex">
    <AccordionTrigger
      data-slot="accordion-trigger"
      v-bind="delegatedProps"
      :class="
        cn(
          'focus-visible:border-ring focus-visible:ring-ring/50 flex flex-1 items-start gap-4 rounded-md py-4 text-left text-sm font-medium transition-all outline-none hover:underline focus-visible:ring-[3px] disabled:pointer-events-none disabled:opacity-50 [&[data-state=open]>svg]:rotate-180',
          props.class,
        )
      "
    >
      <slot name="icon">
        <ChevronDown
          class="pointer-events-none size-8 shrink-0 translate-y-0.5 transition-transform duration-200"
        />
      </slot>
      <slot />
    </AccordionTrigger>
  </AccordionHeader>
</template>
