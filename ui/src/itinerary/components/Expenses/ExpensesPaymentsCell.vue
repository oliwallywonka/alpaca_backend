<script setup lang="ts">
import { PlusIcon } from 'lucide-vue-next'
import { Badge } from '@/core/components/ui/badge';
import { Button } from '@/core/components/ui/button'
import ExpensesPaymentFormDialog from './ExpensePaymentFormDialog.vue'
</script>
<template>
  <div class="w-full h-full grid">
    <Button size="sm"><PlusIcon /></Button>
    <ExpensesPaymentFormDialog />
    <Button size="sm" variant="ghost">123.0 usd.</Button>
    <Button size="sm" variant="ghost">... all </Button>
    <Badge variant="red" class="w-full">135.0 bol. left</Badge>
  </div>
</template>
