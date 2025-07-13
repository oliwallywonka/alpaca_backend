<script setup lang="ts">
import { Card, CardHeader, CardTitle, CardContent } from '@/core/components/ui/card'
import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableCell,
  TableHead,
} from '@/core/components/ui/table'
import { Badge } from '@/core/components/ui/badge'
import PDFVoucher from '@/trip/components/PDFVoucher.vue'
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Payments</CardTitle>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Date</TableHead>
              <TableHead>Registered By</TableHead>
              <TableHead>Customer</TableHead>
              <TableHead>Amount</TableHead>
              <TableHead>Method</TableHead>
              <TableHead>Voucher</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell>2022-01-01</TableCell>
              <TableCell>John Doe</TableCell>
              <TableCell>John Doe</TableCell>
              <TableCell><Badge variant="green">$100</Badge></TableCell>
              <TableCell><Badge>Paypal</Badge></TableCell>
              <TableCell><PDFVoucher /></TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </CardHeader>
  </Card>
</template>
