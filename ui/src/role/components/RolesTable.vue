<script setup lang="ts">
/* import { ref } from 'vue' */
import {
  Table,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
  TableCaption,
} from '@/core/components/ui/table'
/* import { Checkbox } from '@/core/components/ui/checkbox' */
import { Badge } from '@/core/components/ui/badge'
import type { Role } from '../interfaces/role'
import { parseDate } from '@/core/lib/date'
import Switch from '@/core/components/ui/switch/Switch.vue'
import CreateEditRoleModal from './RoleFormDialog.vue'

interface Props {
  roles?: Role[]
}

const props = withDefaults(defineProps<Props>(), {
  roles: () => [],
})

const emit = defineEmits<{
  (e: 'click:status', id: string, status: boolean): void
  (e: 'click:row', id: string): void
  (e: 'click:selectedIDs', selectedIDs: string[]): void
}>()

const handleStatus = (id: string, status: boolean) => {
  emit('click:status', id, status)
}

const handleClickRow = (id: string) => {
  emit('click:row', id)
}
</script>

<template>
  <Table class="min-w-[800px] w-full">
    <TableCaption>Roles List</TableCaption>
    <TableHeader>
      <TableRow>
        <!-- <TableHead class="sticky left-0.5 bg-inherit min-w-5">
          <Checkbox
            :model-value="selectedIDs.length === props.roles.length"
            @update:model-value="addAllToSelectedIDs"
            class="size-5"
          />
        </TableHead> -->
        <TableHead>Name</TableHead>
        <TableHead>Created At</TableHead>
        <TableHead>Updated At</TableHead>
        <TableHead>Status</TableHead>
        <TableHead>Actions</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow
        v-for="role in props.roles"
        :key="role.id"
        class="hover:bg-gray-100 hover:cursor-pointer"
        @click="handleClickRow(role.id)"
      >
        <!-- <TableCell class="sticky left-0.5 bg-inherit min-w-5">
          <Checkbox
            :model-value="selectedIDs.includes(role.ID)"
            @update:model-value="addToSelectedIDs(role.ID)"
            class="size-5"
          />
        </TableCell> -->
        <TableCell class="flex items-center gap-2">
          <div>
            <p class="text-sm font-medium uppercase">{{ role.name }}</p>
            <p class="text-sm text-muted-foreground">{{ role.description }}</p>
          </div>
        </TableCell>
        <TableCell>{{ parseDate(role.created) }}</TableCell>
        <TableCell>{{ parseDate(role.updated) }}</TableCell>
        <TableCell>
          <Badge v-if="!role.isActive" variant="red">Inactive</Badge>
          <Badge v-else variant="green">Active</Badge>
          <Switch
            :model-value="role.isActive"
            @update:model-value="(value) => handleStatus(role.id, value)"
            @click.stop
          />
        </TableCell>
        <TableCell>
          <CreateEditRoleModal :roleID="role.id" />
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>
