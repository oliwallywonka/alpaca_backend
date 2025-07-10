<script setup lang="ts">
import { permissionData } from './constants/permisionList'
import { Switch } from '@/core/components/ui/switch'
import { useVModel } from '@vueuse/core'

const props = withDefaults(
  defineProps<{
    modelValue?: string[]
  }>(),
  {
    modelValue: () => [],
  },
)
const emit = defineEmits<{
  (e: 'update:modelValue', permissions: string[]): void
}>()
const data = useVModel(props, 'modelValue', emit, {
  // I dont know why but this is needed to make the switch work if is not into a form
  passive: true,
})

const permissionDataLength = permissionData.flatMap((item) =>
  item.permissions.map((p) => p.value),
).length

const togglePermission = (permission: string) => {
  if (data.value.includes(permission)) {
    data.value = data.value.filter((i) => i !== permission)
  } else {
    data.value = [...data.value, permission]
  }
}

const toggleAllPermissions = () => {
  if (data.value.length === permissionDataLength) {
    data.value = []
  } else {
    data.value = permissionData.flatMap((item) =>
      item.permissions.map((p) => p.value),
    )
  }
}
</script>

<template>
  <ul>
    <li class="flex items-center gap-2 font-medium">
      <Switch
        :model-value="data.length === permissionDataLength"
        @update:model-value="() => toggleAllPermissions()"
      />
      <p>Select All</p>
    </li>
    <li v-for="item in permissionData" :key="item.title">
      <p class="font-medium">{{ item.title }}</p>
      <ul>
        <li
          v-for="permission in item.permissions"
          :key="permission.value"
          class="flex items-center gap-2"
        >
          <Switch
            :model-value="data.includes(permission.value)"
            @update:model-value="() => togglePermission(permission.value)"
          /> 
          <p>{{ permission.name }}</p>
        </li>
      </ul>
    </li>
  </ul>
</template>
