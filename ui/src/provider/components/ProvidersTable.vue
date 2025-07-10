<script setup lang="ts">
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
  TableCaption,
} from '@/core/components/ui/table'
import { Switch } from '@/core/components/ui/switch'
import { Badge } from '@/core/components/ui/badge'
import type { Provider } from '../interfaces/provider'
import { parseDate } from '@/core/lib/date'
import { buttonVariants } from '@/core/components/ui/button'
import ProviderFormDialog from './ProviderFormDialog.vue'

interface Props {
  providers?: Provider[]
}

const props = withDefaults(defineProps<Props>(), {
  providers: () => [],
})

const emit = defineEmits<{
  (e: 'click:status', id: string, status: boolean): void
}>()

const handleStatus = (id: string, status: boolean) => {
  emit('click:status', id, status)
}
</script>

<template>
  <Table class="min-w-[800px] w-full">
    <TableCaption>Provider List</TableCaption>
    <TableHeader>
      <TableRow>
        <TableHead>Name</TableHead>
        <TableHead>Contacts</TableHead>
        <TableHead>Type</TableHead>
        <TableHead>Updated At</TableHead>
        <TableHead>Status</TableHead>
        <TableHead>Actions</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow v-for="provider of props.providers" :key="provider.id">
        <TableCell class="flex items-center">
          <div>
            <p class="text-sm font-medium uppercase">{{ provider.fullName }}</p>
            <p class="text-sm text-muted-foreground">{{ provider.description }}</p>
          </div>
        </TableCell>
        <TableCell>
          <ul class="grid gap-1">
            <li
              v-for="(contact, index) of provider.contacts"
              :key="JSON.stringify(contact) + index"
            >
              {{ contact.type }}: {{ contact.value }}
            </li>
          </ul>
        </TableCell>
        <TableCell>
          <Badge >{{ provider.types[0] }}</Badge>
        </TableCell>
        <TableCell>{{ parseDate(provider.created) }}</TableCell>
        <TableCell>
          <Badge v-if="!provider.isActive" variant="red">Inactive</Badge>
          <Badge v-else variant="green">Active</Badge>
          <Switch
            :model-value="provider.isActive"
            @update:model-value="(value) => handleStatus(provider.id, value)"
            @click.stop
          />
        </TableCell>
        <TableCell>
          <RouterLink :to="`/providers/${provider.id}/resources`" :class="buttonVariants()">
            Resources
          </RouterLink>
          <ProviderFormDialog :providerID="provider.id" >Edit</ProviderFormDialog>
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>
