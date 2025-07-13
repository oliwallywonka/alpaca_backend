<script setup lang="ts">
import { Card, CardContent, CardHeader, CardTitle } from '@/core/components/ui/card'
import { Button } from '@/core/components/ui/button'
</script>

<template>
  <Card>
    <CardHeader>
      <CardTitle>Client Details</CardTitle>
    </CardHeader>
    <CardContent>
      <p class="font-semibold">Name: Pedro Duarte</p>
      <p class="font-semibold">Email: pedro@duarte.com</p>
      <p class="font-semibold">Phone: pedro@duarte.com</p>
      <p class="font-semibold">Country: Texax, USA</p>
      <Button>Edit</Button>
    </CardContent>
  </Card>
</template>
