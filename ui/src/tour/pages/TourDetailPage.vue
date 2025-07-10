<script setup lang="ts">
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/core/components/ui/tabs'

import TourForm from '@/tour/components/TourForm.vue'
import ImagesForm from '../components/ImagesForm.vue'
import VariantsList from '../components/TourVariants/VariantsList.vue'
</script>

<template>
  <Tabs default-value="variants">
    <TabsList class="grid w-full grid-cols-3">
      <TabsTrigger value="tourForm"> Basic Info </TabsTrigger>
      <TabsTrigger value="images"> Images </TabsTrigger>
      <TabsTrigger value="variants"> Variants </TabsTrigger>
    </TabsList>
    <TabsContent value="tourForm">
      <TourForm />
    </TabsContent>
    <TabsContent value="images">
      <ImagesForm />
    </TabsContent>
    <TabsContent value="variants">
      <VariantsList />
    </TabsContent>
  </Tabs>
</template>
