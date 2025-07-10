<script setup lang="ts">
import { toast } from 'vue-sonner'

import ItemsPaginator from '@/core/components/paginator/ItemsPaginator.vue'
import CreateEditRoleModal from './components/RoleFormDialog.vue'
import RolesTable from './components/RolesTable.vue'
import { RoleService } from './services/roleService'
import { useParams } from '@/core/hooks/useParams'

const { params, setParams } = useParams()
const { data, refetch } = RoleService.useGetAll(params)
const { mutateAsync } = RoleService.useUpdate()

const handleStatus = async (id: string, status: boolean) => {
  try {
    await mutateAsync({ id, data: { isActive: status } })
    await refetch()
    toast.success('Role status updated successfully')
  } catch (err) {
    console.log(err)
    toast.error('Error updating role status')
  }
}
</script>
<template>
  <div>
    <CreateEditRoleModal />
  </div>
  <RolesTable :roles="data?.items" @click:status="handleStatus" />

  <ItemsPaginator
    :items-per-page="params.perPage"
    :total="data?.totalItems || 0"
    @update:page="
      async (page) => {
        setParams({ page })
        await refetch()
      }
    "
  />
</template>
