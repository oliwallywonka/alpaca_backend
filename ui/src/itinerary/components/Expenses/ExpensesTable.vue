<script setup lang="ts">
import { h, ref } from 'vue'
import { AgGridVue } from 'ag-grid-vue3'
import {
  CellStyleModule,
  ClientSideRowModelModule,
  ColumnApiModule,
  ColumnAutoSizeModule,
  ModuleRegistry,
  NumberEditorModule,
  PinnedRowModule,
  RowAutoHeightModule,
  TextEditorModule,
  ValidationModule,
  ClientSideRowModelApiModule,
  HighlightChangesModule,
  type GridOptions,
  type IAggFuncParams,
  type ICellRendererParams,
  type GetRowIdFunc,
  type GetRowIdParams,
  RenderApiModule,
} from 'ag-grid-community'
import { RowGroupingModule } from 'ag-grid-enterprise'

import Badge from '@/core/components/ui/badge/Badge.vue'
import type { ItineraryResource } from '@/itinerary/interfaces/itineraryResource'
import type { ResourceProvider } from '@/providerResource/ResourceProvider'
import ResourceDetailsDialog from '@/itinerary/components/Dialogs/ResourceDetailsDialog.vue'
import ExpensesPaymentsCell from './ExpensesPaymentsCell.vue'

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  NumberEditorModule,
  TextEditorModule,
  ColumnApiModule,
  HighlightChangesModule,
  CellStyleModule,
  RowGroupingModule,
  ClientSideRowModelApiModule,
  RowAutoHeightModule,
  PinnedRowModule,
  ColumnAutoSizeModule,
  RenderApiModule,
  ValidationModule,
])

const props = defineProps<{
  resources: ItineraryResource[]
}>()

const emit = defineEmits<{
  (e: 'change:itineraryResource', itineraryResource: Partial<ItineraryResource>): void
}>()

const itineraryResourceID = ref<string>()
const showDialog = ref(false)

const DescriptionCell = {
  setup({ params }: { params: ICellRendererParams<ItineraryResource, ResourceProvider> }) {
    if (params.node.level === -1 && params.node.footer) {
      return () => h('span', { class: 'font-bold' }, 'Total')
    }
    return () =>
      h(
        'div',
        {
          class: 'flex flex-col items-start px-2 pb-2 gap-1 hover:cursor-pointer',
          onClick: () => {
            showDialog.value = true
            itineraryResourceID.value = params.data?.id
          },
        },
        [
          params.value?.expand.resource.name.en,
          h(
            'p',
            {
              class: 'flex flex-wrap max-w-24 gap-1',
            },
            [
              h(Badge, {}, () => [params.value?.provider ? 'Provider' : 'User']),
              h(Badge, {}, () => [
                params.value?.expand.provider?.fullName || params.value?.expand.user?.name,
              ]),
              h(Badge, {}, () => [
                `Persons Min: ${params.data?.minPersons} Max: ${params.data?.maxPersons}`,
              ]),
            ],
          ),
        ],
      )
  },
}

const CurrencyCell = {
  setup({ params }: { params: ICellRendererParams }) {
    return () => h(Badge, () => [params.value])
  },
}

const ActionsCell = {
  setup({ params }: { params: ICellRendererParams }) {
    if (params.node.level === -1 && params.node.footer) {
      return () => {}
    }
    //return () => h(DeleteExpenseDialog, { itineraryResourceID: params.data?.id })
    return () => h(ExpensesPaymentsCell)
  },
}

// Helper function to round the number to two decimal places
const roundHelper = (number: number) => Math.round(number * 100) / 100

const sumHelper = (params: IAggFuncParams) => {
  const sum = params.values.reduce((acc, val) => acc + val, 0)
  return roundHelper(sum)
}

const getRowId = ref<GetRowIdFunc>((params: GetRowIdParams) => String(params.data.id))

const gridOptions: GridOptions<ItineraryResource> = {
  defaultColDef: {
    minWidth: 90,
    enableCellChangeFlash: true,
    type: 'rightAligned',
    cellStyle: (params) => {
      return {
        backgroundColor: params.colDef.editable ? '#f5f5f5' : '#f9fafb',
      }
    },
  },
  rowModelType: 'clientSide',
  suppressAggFuncInHeader: true,
  autoSizeStrategy: {
    type: 'fitCellContents',
  },
  grandTotalRow: 'pinnedBottom',
  enableRowPinning: true,
  domLayout: 'autoHeight',
  columnDefs: [
    {
      headerName: 'Description',
      field: 'expand.resourceProvider',
      pinned: 'left',
      type: 'leftAligned',
      wrapText: true,
      autoHeight: true,
      width: 300,
      cellRenderer: DescriptionCell,
    },
    {
      headerName: 'Original Cost',
      field: 'originalCost',
      editable: true,
      onCellValueChanged: (params) => {
        emit('change:itineraryResource', {
          id: params.data?.id,
          originalCost: params.newValue,
        })
      },
      aggFunc: sumHelper,
    },
    {
      field: 'currency',
      cellRenderer: CurrencyCell,
    },
    {
      field: 'quantity',
      editable: true,
      onCellValueChanged: (params) => {
        emit('change:itineraryResource', {
          id: params.data?.id,
          quantity: params.newValue,
        })
      },
      aggFunc: sumHelper,
    },
    {
      headerName: 'SubTotal Cost',
      enableValue: false,
      valueFormatter: (params) =>
        `${params.value} ${params.data?.currency.toLocaleLowerCase() || ''}.`,
      valueGetter: (params) => {
        const originalCost = params.data?.originalCost || 0
        const quantity = params.data?.quantity || 1
        return roundHelper(originalCost * quantity)
      },
      aggFunc: () => '',
    },
    {
      field: 'dollarChangeRate',
      headerName: 'USD Change Rate',
      editable: true,
      onCellValueChanged: (params) => {
        emit('change:itineraryResource', {
          id: params.data?.id,
          dollarChangeRate: params.newValue,
        })
      },
      aggFunc: () => '',
    },
    {
      colId: 'SUBTOTALCOST',
      headerName: 'SubTotal USD',
      valueGetter: (params) => {
        const costUSD =
          ((params.data?.originalCost || 0) * (params.data?.quantity || 1)) /
          (params.data?.dollarChangeRate || 0)
        return roundHelper(costUSD)
      },
      valueFormatter: (params) => `${params.value} usd.`,
      aggFunc: sumHelper,
    },
    {
      field: 'profitPercent',
      valueFormatter: (params) => `${params.value}%`,
      aggFunc: () => '',
      editable: true,
      onCellValueChanged: (params) => {
        emit('change:itineraryResource', {
          id: params.data?.id,
          profitPercent: params.newValue,
        })
      },
    },
    {
      colId: 'PROFIT',
      headerName: 'Profit',
      valueGetter: (params) => {
        const subTotalCost = params.getValue('SUBTOTALCOST') ?? 1
        const profitPercent = params.data?.profitPercent || 0
        const profit = roundHelper((profitPercent / 100) * subTotalCost)
        return profit
      },
      valueFormatter: (params) => `${params.value} usd.`,
      aggFunc: sumHelper,
    },
    {
      headerName: 'SubTotal Price',
      valueGetter: (params) => {
        const subTotalCost = params.getValue('SUBTOTALCOST') ?? 1
        const profit = params.getValue('PROFIT') ?? 0
        return roundHelper(subTotalCost + profit)
      },
      valueFormatter: (params) => `${params.value} usd.`,
      aggFunc: sumHelper,
      pinned: 'right',
    },
    {
      headerName: 'Payments',
      pinned: 'right',
      type: 'leftAligned',
      autoHeight: true,
      cellRenderer: ActionsCell,
    },
  ],
}
</script>

<template>
  <ResourceDetailsDialog
    v-if="itineraryResourceID && showDialog"
    v-model="showDialog"
    :itineraryResourceID="itineraryResourceID"
  />
  <AgGridVue :gridOptions="gridOptions" :rowData="props.resources" :getRowId="getRowId" />
</template>
