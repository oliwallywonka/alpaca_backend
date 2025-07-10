<script setup lang="ts">
import { toast } from 'vue-sonner'
import { RefreshCwIcon } from 'lucide-vue-next'

import SearchInput from '@/core/components/fields/SearchInput.vue'
import UsersTable from '@/user/components/UsersTable.vue'
import CreateUserDialog from '@/user/components/UserFormDialog.vue'
import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import { Button } from '@/core/components/ui/button'

import { UserService } from '@/user/services/userService'
import { useParams } from '@/core/hooks/useParams'

const { params, setParams } = useParams({
  orderBy: 'user.created',
})
const { data, refetch } = UserService.useGetAll(params)
const { mutateAsync } = UserService.useUpdate()

const handleStatus = async (id: string, status: boolean) => {
  try {
    await mutateAsync({ id, data: { isActive: status } })
    await refetch()
    toast.success('User status updated successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error updating user status')
  }
}

const handleSearchInput = async (value: string) => {
  setParams({ filter: `name~"${value}" || email~"${value}" || role.description~"${value}"` })
  await refetch()
}
</script>

<template>
  <h1 class="text-xl font-bold">Users management</h1>
  <SearchInput @change:debounce="handleSearchInput" />
  <div class="flex gap-2">
    <CreateUserDialog />
    <Button @click="refetch"> <RefreshCwIcon /> </Button>
  </div>
  <UsersTable :users="data?.items" @click:status="handleStatus" />
  <ItemsPaginator
    :itemsPerPage="params.perPage"
    :total="data?.totalItems || 0"
    @update:page="
      async (page) => {
        setParams({ page })
        await refetch()
      }
    "
  />
</template>
