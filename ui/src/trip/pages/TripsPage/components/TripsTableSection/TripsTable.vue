<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { computed, h } from 'vue'
import { AgGridVue } from 'ag-grid-vue3'
import { cn } from '@/core/lib/utils'
import { buttonVariants } from '@/core/components/ui/button'
import { Badge } from '@/core/components/ui/badge'
import {
  ClientSideRowModelModule,
  ColumnAutoSizeModule,
  ModuleRegistry,
  ValidationModule,
  type GridOptions,
  type ICellRendererParams,
  type SelectionChangedEvent,
} from 'ag-grid-community'
import type { TripSummary } from '@/trip/interfaces/trip'
import { dateToSuggar } from '@/core/lib/date'

ModuleRegistry.registerModules([ClientSideRowModelModule, ColumnAutoSizeModule, ValidationModule])

const props = defineProps<{
  trips: TripSummary[]
  showSelect?: boolean
}>()

const emit = defineEmits<{
  (e: 'select:trip', trip: TripSummary | undefined): void
}>()

const StatusCell = {
  setup({ params }: { params: ICellRendererParams<TripSummary> }) {
    return () => h(Badge, {}, () => [params.value])
  },
}

const DetailsCell = {
  setup({ params }: { params: ICellRendererParams<TripSummary> }) {
    return () =>
      h(
        RouterLink,
        {
          to: `/trips/${params.data?.id}`,
          class: cn(buttonVariants()),
        },
        () => 'Details',
      )
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
    onSelectionChanged: (event: SelectionChangedEvent<TripSummary>) => {
      if (!event.selectedNodes || event.selectedNodes?.length === 0) {
        emit('select:trip', undefined)
        return
      }
      emit('select:trip', event.selectedNodes[0].data)
    }
  }
})

const gridOptions: GridOptions<TripSummary> = {
  domLayout: 'autoHeight',
  autoSizeStrategy: {
    type: 'fitCellContents',
  },
  ...selectOptions.value,
  defaultColDef: {
    minWidth: 150,
  },
  columnDefs: [
    {
      field: 'expand.tour.name.en',
      headerName: 'Tour',
      pinned: 'left',
    },
    {
      field: 'customerLead',
      headerName: 'Customer Lead',
    },
    {
      field: 'startDate',
      headerName: 'Start Date',
      valueFormatter: (params) => {
        return dateToSuggar(params.value)
      },
    },
    {
      field: 'endDate',
      headerName: 'End Date',
      valueFormatter: (params) => {
        return dateToSuggar(params.value)
      },
    },
    {
      field: 'state',
      headerName: 'State',
      cellRenderer: StatusCell,
    },
    {
      headerName: 'Actions',
      pinned: 'right',
      cellRenderer: DetailsCell,
    },
  ],
}
</script>
<template>
  <AgGridVue :gridOptions="gridOptions" :rowData="props.trips" />
</template>
