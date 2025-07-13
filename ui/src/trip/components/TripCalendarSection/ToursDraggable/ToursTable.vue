<script setup lang="ts">
import type { TourVariant } from '@/tour/interfaces/tourVariant'
import { Draggable } from '@fullcalendar/interaction/index.js'
import {
  ClientSideRowModelModule,
  ModuleRegistry,
  ValidationModule,
  type GridOptions,
  type ICellRendererParams,
} from 'ag-grid-community'
import { RowGroupingModule } from 'ag-grid-enterprise'
import { AgGridVue } from 'ag-grid-vue3'
import { h, onMounted } from 'vue'

const props = defineProps<{
  data: TourVariant[]
}>()

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  RowGroupingModule,
  ValidationModule,
])

const draggableCell = {
  setup({params}: { params: ICellRendererParams<TourVariant>}) {
    return () =>
      h(
        'div',
        {
          class: 'variant-draggable',
        },
        params.value,
      )
  }
}

const gridOptions: GridOptions<TourVariant> = {
  groupDisplayType: 'groupRows',
  domLayout: 'autoHeight',
  defaultColDef: {
    flex: 1,
    minWidth: 100,
  },
  columnDefs: [
    {
      field: 'expand.tour.code',
      headerName: 'Code',
      cellRenderer: draggableCell,
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
      field: 'totalPrice',
      headerName: 'Total Price',
    },
  ],
}

onMounted(() => {
  new Draggable(document.getElementById('variant-draggable')!, {
    itemSelector: '.variant-draggable',
    eventData: function (element) {
      const data = JSON.parse(element.getAttribute('props') ?? '{}')
      return {
        title: 'event 1',
        date: '2025-04-04',
        backgroundColor: '#e5f1fe',
        //borderColor: '#d8e6f7',
        borderColor: '#005bbf',
        textColor: '#005bbf',
        ...data,
      }
    },
  })
})
</script>

<template>
  <AgGridVue id="variant-draggable" :gridOptions="gridOptions" :rowData="props.data" />
</template>
