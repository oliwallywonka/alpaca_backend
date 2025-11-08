<script setup lang="ts">
import { computed, h } from 'vue'
import { AgGridVue } from 'ag-grid-vue3'
import {
  ClientSideRowModelModule,
  ColumnAutoSizeModule,
  ModuleRegistry,
  RowSelectionModule,
  ValidationModule,
  type GridOptions,
  type ICellRendererParams,
  type SelectionChangedEvent,
} from 'ag-grid-community'

import type { Destination } from '@/destination/interfaces/destination'
import { Switch } from '@/core/components/ui/switch'
import { Badge } from '@/core/components/ui/badge'
import DestinationFormDialog from './DestinationFormDialog.vue'

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  ColumnAutoSizeModule,
  RowSelectionModule,
  ValidationModule,
])

const props = defineProps<{
  destinations: Destination[]
  showSelect?: boolean
}>()

const emit = defineEmits<{
  (e: 'click:status', id: string, status: boolean): void
  (e: 'select:destination', destination: Destination | undefined): void
}>()

const ParentCell = {
  setup({ params }: { params: ICellRendererParams<Destination> }) {
    return () => h(Badge, () => [params.value.parent?.name?.en ?? ''])
  },
}

const LocationCell = {
  setup({ params }: { params: ICellRendererParams<Destination> }) {
    return () => h(Badge, () => [`${params.data?.location.lat} | ${params.data?.location.lon}`])
  },
}

const StatusCell = {
  setup({ params }: { params: ICellRendererParams<Destination> }) {
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

const ActionsCell = {
  setup({ params }: { params: ICellRendererParams<Destination> }) {
    return () => h(DestinationFormDialog, { destinationID: params.data?.id }, () => ['Edit'])
  },
}

const selectOptions = computed<Partial<GridOptions>>(() => {
  if (!props.showSelect) return {}
  return {
    rowSelection: {
      mode: 'singleRow',
      checkboxes: true,
    },
    selectionColumnDef: {
      pinned: 'left',
    },
    onSelectionChanged: (event: SelectionChangedEvent<Destination>) => {
      if (!event.selectedNodes || event.selectedNodes?.length === 0) {
        emit('select:destination', undefined)
        return
      }
      emit('select:destination', event.selectedNodes[0].data)
    },
  }
})

const gridOptions: GridOptions<Destination> = {
  columnDefs: [
    {
      field: 'name.en',
      headerName: 'Name',
      minWidth: 250,
      pinned: 'left',
    },
    {
      field: 'description.en',
      headerName: 'Description',
      minWidth: 350,
      maxWidth: 450,
    },
    {
      field: 'expand',
      headerName: 'Parent',
      minWidth: 100,
      cellRenderer: ParentCell,
    },
    {
      field: 'location',
      minWidth: 100,
      cellRenderer: LocationCell,
    },
    {
      field: 'isActive',
      headerName: 'Status',
      minWidth: 120,
      cellRenderer: StatusCell,
    },
    {
      headerName: 'Actions',
      pinned: 'right',
      cellRenderer: ActionsCell,
    },
  ],
  ...selectOptions.value,
  domLayout: 'autoHeight',
  autoSizeStrategy: {
    type: 'fitCellContents',
  },
}

const handleStatus = (id: string, status: boolean) => {
  emit('click:status', id, status)
}
</script>
<template>
  <AgGridVue :gridOptions="gridOptions" :rowData="props.destinations" />
</template>
