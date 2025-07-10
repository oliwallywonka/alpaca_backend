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
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Expenses</CardTitle>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Date</TableHead>
              <TableHead>Registered By</TableHead>
              <TableHead>Covered By</TableHead>
              <TableHead>Amount</TableHead>
              <TableHead>Covered</TableHead>
              <TableHead>Type</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell>2022-01-01</TableCell>
              <TableCell>John Doe</TableCell>
              <TableCell>John Doe</TableCell>
              <TableCell><Badge variant="green">$100</Badge></TableCell>
              <TableCell><Badge variant="red">NO</Badge></TableCell>
              <TableCell><Badge variant="yellow">Service</Badge></TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
    </CardHeader>
  </Card>
</template>
