<script setup lang="ts">
import { AgGridVue } from 'ag-grid-vue3'
import {
  ClientSideRowModelModule,
  ColumnAutoSizeModule,
  ModuleRegistry,
  RowSelectionModule,
  ValidationModule,
  type GridOptions,
} from 'ag-grid-community'
import { RowGroupingModule } from 'ag-grid-enterprise'

import type { ResourceProvider } from '@/providerResource/ResourceProvider'
import { computed } from 'vue'

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  ColumnAutoSizeModule,
  RowSelectionModule,
  ValidationModule,
  RowGroupingModule,
])

const props = defineProps<{
  data: ResourceProvider[]
}>()
const emit = defineEmits<{
  (e: 'select:resourceProvider', providerResource: ResourceProvider | undefined): void
}>()

const flatedData = computed(() => {
  return props.data.flatMap((rp) => {
    return rp.refPrices.map((price) => ({
      ...rp,
      minPersons: price.minPersons,
      maxPersons: price.maxPersons,
      price: price.price,
      currency: price.currency,
      isPerPerson: price.isPerPerson,
    }))
  })
})

const gridOptions: GridOptions<(typeof flatedData.value)[0]> = {
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
  selectionColumnDef: {
    pinned: 'left',
  },
  onSelectionChanged: (event) => {
    if (!event.selectedNodes || event.selectedNodes?.length === 0) {
      emit('select:resourceProvider', undefined)
      return
    }
    emit('select:resourceProvider', {
      ...event.selectedNodes[0].data,
      refPrices: [
        {
          minPersons: event.selectedNodes[0].data.minPersons,
          maxPersons: event.selectedNodes[0].data.maxPersons,
          price: event.selectedNodes[0].data.price,
          currency: event.selectedNodes[0].data.currency,
          isPerPerson: event.selectedNodes[0].data.isPerPerson,
        },
      ],
    })
  },
  columnDefs: [
    {
      field: 'expand.resource.name.en',
      headerName: 'Resource',
      hide: true,
      rowGroup: true,
    },
    {
      headerName: 'Provider',
      valueGetter: (params) => {
        const data = params.data?.expand
        return data?.user?.name || data?.provider?.fullName
      },
      rowGroup: true,
      hide: true,
    },
    {
      headerName: 'Prices',
      valueGetter: (params) => {
        return `Min: ${params.data?.minPersons} Max: ${params.data?.maxPersons} Price: ${params.data?.price} ${params.data?.currency}`
      },
    },
  ],
}
</script>
<template>
  <AgGridVue :gridOptions="gridOptions" :rowData="flatedData" class="max-h-[70vh]" />
</template>
