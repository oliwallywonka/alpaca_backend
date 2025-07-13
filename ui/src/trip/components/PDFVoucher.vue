<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import type { TDocumentDefinitions } from 'pdfmake/interfaces'
import * as PdfPrinter from 'pdfmake/build/pdfmake'
import 'pdfmake/build/vfs_fonts'
import { logoBase64 } from './base64'

const voucher: TDocumentDefinitions = {
  // a string or { width: number, height: number }
  pageSize: 'LETTER',

  // by default we use portrait, you can change it to landscape if you wish
  pageOrientation: 'portrait',

  // [left, top, right, bottom] or [horizontal, vertical] or just a number for equal margins
  pageMargins: [40, 30, 40, 30],
  content: [
    {
      columns: [
        {
          width: 90,
          height: 75,
          fontSize: 9,
          margin: [22, 0, 0, 0],
          image: logoBase64
        },
        [
          {
            text: 'Voucher',

            color: '#333333',
            width: 800,
            fontSize: 20,
            bold: true,
            alignment: 'center',
            margin: [0, 0, 0, 0],
          },
        ],
      ],
    },
    {
      columns: [
        {
          width: 150,
          text: [
            'CEL:78857480 - 78857480\n',
            'Urb. 21 de octubre "A" Dist. 7 Ex Tranca San Roque Av. Panamericana Nro 3174 Entre calle Gral. juan Jose Torrez\n',
            'El Alto-La Paz-Bolivia\n',
          ],
          fontSize: 8,
          alignment: 'center',
          margin: [0, 0, 0, 0],
        },
        [
          '\n',
          {
            text: 'Date: 12 jul 2025 11:48:08AM',

            color: '#333333',
            width: '*',
            fontSize: 11,
            bold: true,
            alignment: 'rigth',
            margin: [19, 0, 0, 0],
          },
        ],
        [
          '\n',
          {
            text: 'N: xxxxxxx',

            color: '#333333',
            width: '*',
            fontSize: 11,
            bold: true,
            alignment: 'rigth',
            margin: [19, 0, 0, 0],
          },
        ],
      ],
    },
    {
      alignment: 'justify',
      columns: [
        {
          width: 400,

          //nombre del cliente
          text: ['Customer: Harold Martinez\n'],
          fontSize: 11,
          alignment: 'center',
          margin: [0, 18, 0, 0],
        },
        [
          {
            text: 'Email: harold@martinez.com',

            color: '#333333',
            width: '*',
            fontSize: 11,
            bold: true,
            alignment: 'rigth',
            margin: [0, 18, 0, 5],
          },
        ],
      ],
    },
    { canvas: [{ type: 'line', x1: 0, y1: 0, x2: 530, y2: 0, lineWidth: 3 }] },
    {
      style: 'tableExample',
      table: {
        widths: [48, '*', '*', 80],
        headerRows: 4,
        body: [
          [
            { text: '' },
            { text: 'SUMMARY', alignment: 'center', style: 'tableHeader', bold: true, colSpan: 2 },
            {},
            { text: 'TOTAL', alignment: 'center', bold: true },
          ],
          [
            { text: 'TOUR', alignment: 'center' },
            { text: 'Uyuni full tour 2 days 6 people', fontSize: 10, alignment: 'left', colSpan: 2 },
            {},
            { text: '325 USD', alignment: 'right' },
          ],
          [
            { text: 'PAYMENT', fontSize: 10, alignment: 'center' },
            { },
            {},
            { text: '75 USD', alignment: 'right' },
          ],
          [
            {
              text: 'PAYMENT LEFT',
              alignment: 'left',
              fontSize: 11,
              bold: true,
              colSpan: 3,
            },
            {},
            {},
            { text: '250 USD', alignment: 'right', bold: true },
          ],
        ],
      },
      layout: 'lightHorizontalLines',
    },
    ' \n\n\n',
    {
      text: [
        'DEPARTURE DATE: 27 JUL 2025 11:48:08AM BOLIVIAN TIME',
      ],
      fontSize: 12,
      bold: true,
      alignment: 'center',
    },
    '\n\n\n\n\n',
    //aca segu
  ],

  styles: {
    notesTitle: {
      fontSize: 8,
      bold: true,
      margin: [0, 50, 0, 3],
    },
    notesText: {
      fontSize: 8,
    },
  },
  tableHeader: {
    bold: true,
    fontSize: 12,
    color: 'black',
  },
  defaultStyle: {
    columnGap: 20,
    //font: 'Quicksand',
  },
}

function generateVoucher() {
  PdfPrinter.createPdf(voucher).open()
}
</script>

<template>
  <Button variant="outline" size="sm" class="w-24" @click="generateVoucher"> Voucher </Button>
</template>
