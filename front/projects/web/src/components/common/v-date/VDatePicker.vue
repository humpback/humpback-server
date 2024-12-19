<script lang="ts" setup>
import type { DatePickerProps } from "element-plus"
import { CircleClose, Calendar } from "@element-plus/icons-vue"

type Props = Partial<
  DatePickerProps & {
    modelValue: number
  }
>

const props = withDefaults(defineProps<Props>(), {
  editable: true,
  clearable: true,
  clearIcon: CircleClose,
  rangeSeparator: "-",
  popperOptions: { persistent: true },
  validateEvent: true,
  teleported: true,
  type: "date",
  format: "MM/DD/YYYY",
  placeholder: "",
  prefixIcon: Calendar,
  valueFormat: "x"
})

const emits = defineEmits<{
  (e: "update:modelValue", data: number): void
  (e: "change"): void
}>()

function updateModel(v: number | undefined) {
  emits("update:modelValue", v || 0)
}

function change() {
  emits("change")
}
</script>
<template>
  <el-date-picker style="width: 100%" v-bind="{ ...props }" @change="change()" @update:modelValue="updateModel" />
</template>

<style lang="scss" scoped></style>
