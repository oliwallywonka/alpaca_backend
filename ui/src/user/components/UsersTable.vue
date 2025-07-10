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
import { Avatar, AvatarFallback, AvatarImage } from '@/core/components/ui/avatar'
import { Badge } from '@/core/components/ui/badge'
import type { User } from '../interfaces/user'
import { buttonVariants } from '@/core/components/ui/button'
import UserFormDialog from './UserFormDialog.vue'

interface Props {
  users?: User[]
}

const props = withDefaults(defineProps<Props>(), {
  users: () => [],
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
    <TableCaption>User List</TableCaption>
    <TableHeader>
      <TableRow>
        <TableHead>Name</TableHead>
        <TableHead>Role</TableHead>
        <TableHead>Email</TableHead>
        <TableHead>Contacts</TableHead>
        <TableHead>Status</TableHead>
        <TableHead>Actions</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow v-for="user of props.users" :key="user.id">
        <TableCell class="flex items-center gap-2">
          <Avatar>
            <AvatarImage src="https://github.com/unovues.png" alt="@unovue" />
            <AvatarFallback>UN</AvatarFallback>
          </Avatar>
          <div>
            <p class="text-sm font-medium uppercase">{{ user.name }}</p>
            <p class="text-sm text-muted-foreground">{{ user.expand.role.description }}</p>
          </div>
        </TableCell>
        <TableCell>
          <Badge>{{ user.expand.role.name }}</Badge>
        </TableCell>
        <TableCell>{{ user.email }}</TableCell>
        <TableCell>
          <ul class="grid gap-1">
            <li v-for="contact of user.contacts" :key="JSON.stringify(contact)">
              {{ contact.type }}: {{ contact.value }}
            </li>
          </ul>
        </TableCell>
        <TableCell>
          <Badge v-if="!user.isActive" variant="red">Inactive</Badge>
          <Badge v-else variant="green">Active</Badge>
          <Switch
            :model-value="user.isActive"
            @update:model-value="(value) => handleStatus(user.id, value)"
            @click.stop
          />
        </TableCell>
        <TableCell class="flex gap-2">
          <RouterLink :to="`/users/${user.id}/resources`" :class="buttonVariants()">
            Resources
          </RouterLink>
          <UserFormDialog :userID="user.id" />
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>
