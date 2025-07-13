<script setup lang="ts">
import { h } from 'vue'
import { RouterLink } from 'vue-router'

import { AgGridVue } from 'ag-grid-vue3'
import {
  ClientSideRowModelModule,
  ColumnApiModule,
  ColumnAutoSizeModule,
  ModuleRegistry,
  RowAutoHeightModule,
  ValidationModule,
  type GridOptions,
  type ICellRendererParams,
} from 'ag-grid-community'
import ImageModal from '../components/ImageModal.vue'
import { buttonVariants } from '@/core/components/ui/button'
import type { Tour } from '../interfaces/tour'
import { Badge } from '@/core/components/ui/badge'
import { Switch } from '@/core/components/ui/switch'
import { getFileURL } from '@/core/lib/fileURL'

ModuleRegistry.registerModules([
  ClientSideRowModelModule,
  ColumnApiModule,
  RowAutoHeightModule,
  ColumnAutoSizeModule,
  ValidationModule,
])

const props = defineProps<{
  tours: Tour[]
}>()

const ImageCell = {
  setup({ params }: { params: ICellRendererParams<Tour> }) {
    return () =>
      h(ImageModal, {
        imageURL: getFileURL({
          collectionName: 'tours',
          recordID: params.data?.id || '',
          fileName: params.value,
        }),
      })
  },
}
const URLCell = {
  setup({ params }: { params: ICellRendererParams }) {
    return () =>
      h(
        'a',
        {
          href: params.value,
          target: '_blank',
          class: 'text-blue-500 hover:underline',
        },
        params.value,
      )
  },
}

const PublicCell = {
  setup({ params }: { params: ICellRendererParams<Tour> }) {
    return () =>
      h('div', { class: 'w-full h-full flex gap-2 items-center' }, [
        h(
          Badge,
          { variant: params.data?.isPublic ? 'green' : 'red' },
          { default: () => (params.data?.isPublic ? 'Public' : 'Hidden') },
        ),
        h(Switch, {
          modelValue: params.data?.isPublic,
          'onUpdate:model-value': () => {},
        }),
      ])
  },
}

const ActiveCell = {
  setup({ params }: { params: ICellRendererParams<Tour> }) {
    return () =>
      h('div', { class: 'w-full h-full flex gap-2 items-center' }, [
        h(
          Badge,
          { variant: params.data?.isActive ? 'green' : 'red' },
          { default: () => (params.data?.isActive ? 'Active' : 'Inactive') },
        ),
        h(Switch, {
          modelValue: params.data?.isActive,
          'onUpdate:model-value': () => {},
        }),
      ])
  },
}

const ActionCell = {
  setup({ params }: { params: ICellRendererParams<Tour> }) {
    return () =>
      h(
        RouterLink,
        {
          to: `/tours/${params.data?.id}`,
          class: buttonVariants({
            variant: 'outline',
          }),
        },
        () => 'Edit',
      )
  },
}

const gridOptions: GridOptions<Tour> = {
  columnDefs: [
    { field: 'code', minWidth: 100, pinned: 'left' },
    { field: 'banner', width: 100, cellRenderer: ImageCell },
    { field: 'name.en', minWidth: 300, flex: 3 },
    { field: 'slug.en', minWidth: 200, flex: 2, cellRenderer: URLCell },
    { field: 'days', minWidth: 100, flex: 1 },
    { field: 'groupSize', headerName: 'Group Size', minWidth: 150, flex: 1 },
    {
      field: 'isPublic',
      headerName: 'Plublished',
      minWidth: 150,
      flex: 1,
      cellRenderer: PublicCell,
    },
    { field: 'isActive', headerName: 'Active', minWidth: 150, flex: 1, cellRenderer: ActiveCell },
    { headerName: 'Actions', pinned: 'right', cellRenderer: ActionCell },
  ],
  domLayout: 'autoHeight',
  autoSizeStrategy: {
    type: 'fitCellContents',
  },
}

</script>

<template>
  <AgGridVue :gridOptions="gridOptions" :rowData="props.tours" />
</template>
