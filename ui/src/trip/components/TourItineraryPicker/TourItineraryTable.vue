<script setup lang="ts">
import {
  ClientSideRowModelModule,
  ModuleRegistry,
  RowSelectionModule,
  ValidationModule,
  type GridOptions,
} from 'ag-grid-community'
import { RowGroupingModule } from 'ag-grid-enterprise'
import { AgGridVue } from 'ag-grid-vue3'

import type { Itinerary } from '@/itinerary/interfaces/itinerary'

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  RowGroupingModule,
  RowSelectionModule,
  ValidationModule,
])

const props = defineProps<{
  data: Itinerary[]
}>()

const emit = defineEmits<{
  (e: 'select:itinerary', itinerary: Itinerary | undefined): void
}>()

const gridOptions: GridOptions<Itinerary> = {
  groupDisplayType: 'groupRows',
  domLayout: 'autoHeight',
  defaultColDef: {
    flex: 1,
    minWidth: 100,
  },
  rowSelection: {
    mode: 'singleRow',
    checkboxes: true,
  },
  onSelectionChanged: (event) => {
    if (!event.selectedNodes || event.selectedNodes?.length === 0) {
      emit('select:itinerary', undefined)
      return
    }
    emit('select:itinerary', event.selectedNodes[0].data)
  },
  columnDefs: [
    {
      field: 'expand.tour.code',
      headerName: 'Code',
    },
    {
      field: 'expand.tour.name.en',
      headerName: 'Tour',
      rowGroup: true,
      hide: true,
    },
    {
      field: 'minPersons',
      headerName: 'Min Persons',
    },
    {
      field: 'maxPersons',
      headerName: 'Max Persons',
    },
    {
      field: 'finalPrice',
      headerName: 'Total Price',
    },
  ],
}
</script>

<template>
  <AgGridVue :gridOptions="gridOptions" :rowData="props.data" />
</template>
