<script setup lang="ts">
import { h } from 'vue'
import {
  ClientSideRowModelModule,
  ColumnApiModule,
  ModuleRegistry,
  type ColDef,
  type ICellRendererParams,
} from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue3'
import type { Resource } from '../interfaces/resource'

import Switch from '@/core/components/ui/switch/Switch.vue'
import ResourceFormDialog from './ResourceFormDialog.vue'
import Badge from '@/core/components/ui/badge/Badge.vue'

ModuleRegistry.registerModules([ClientSideRowModelModule, ColumnApiModule])
const props = withDefaults(
  defineProps<{
    services?: Resource[]
  }>(),
  { services: () => [] },
)
const emit = defineEmits<{
  (e: 'click:status', id: string, status: boolean): void
}>()

const handleStatus = (id: string, status: boolean) => {
  emit('click:status', id, status)
}

const StatusCell = {
  setup({ params }: { params: ICellRendererParams<Resource> }) {
    return () =>
      h('div', { class: 'w-full h-full flex gap-2 items-center' }, [
        h(
          Badge,
          { variant: params.data?.isActive ? 'green' : 'red' },
          { default: () => (params.data?.isActive ? 'Active' : 'Inactive') },
        ),
        h(Switch, {
          modelValue: params.data?.isActive,
          'onUpdate:model-value': (value: boolean) => handleStatus(params.data?.id || '', value),
        }),
      ])
  },
}

const TypesCell = {
  setup({ params }: { params: ICellRendererParams<Resource> }) {
    return () =>
      h(
        'div',
        { class: 'w-full h-full flex gap-2 items-center' },
        params.data?.types.map((type) => h(Badge, {}, { default: () => type })),
      )
  },
}

const ActionCell = {
  setup({ params }: { params: ICellRendererParams<Resource> }) {
    return () =>
      h(
        ResourceFormDialog,
        {
          resourceID: params.data?.id,
        },
        () => ['Edit'],
      )
  },
}

const columns: ColDef<Resource>[] = [
  {
    width: 300,
    field: 'name.en',
    pinned: 'left',
    headerName: 'Service',
    wrapText: true,
  },
  {
    width: 300,
    field: 'description.en',
    headerName: 'Description',
  },
  {
    width: 300,
    field: 'types',
    headerName: 'Types',
    cellRenderer: TypesCell,
  },
  {
    flex: 1,
    minWidth: 150,
    field: 'isActive',
    headerName: 'Status',
    cellRenderer: StatusCell,
  },
  {
    headerName: 'Actions',
    pinned: 'right',
    cellRenderer: ActionCell,
  },
]
</script>

<template>
  <AgGridVue
    :rowAutoHeight="true"
    :rowData="props.services"
    :columnDefs="columns"
    domLayout="autoHeight"
  />
</template>
