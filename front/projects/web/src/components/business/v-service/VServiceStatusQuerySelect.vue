<script lang="ts" setup>
import { NodeSwitch, ServiceStatus } from "@/models"

const props = withDefaults(
  defineProps<{
    modelValue: string
    size?: "large" | "default" | "small"
    clearable?: boolean
    placeholder?: string
  }>(),
  { size: "default", clearable: true, placeholder: "" }
)
const emits = defineEmits<{
  (e: "update:modelValue", data: string): void
  (e: "change", data: string): void
}>()

const { t } = useI18n()

const status = computed({
  get: () => props.modelValue || "",
  set: (v: string) => {
    emits("update:modelValue", v)
  }
})

const options = computed(() => [
  { label: "label.enabled", value: NodeSwitch.Enabled },
  { label: "label.disabled", value: NodeSwitch.Disabled },
  { label: "label.notReady", value: ServiceStatus.ServiceStatusNotReady },
  { label: "label.running", value: ServiceStatus.ServiceStatusRunning },
  { label: "label.failed", value: ServiceStatus.ServiceStatusFailed }
])

function change() {
  emits("change", status.value)
}
</script>

<template>
  <v-select
    v-model="status"
    :clearable="props.clearable"
    :out-label="t('label.status')"
    :placeholder="props.placeholder"
    out-label-width="80px"
    show-out-label
    @change="change()">
    <el-option :label="t('label.all')" :value="0" />
    <el-option v-for="item in options" :key="item.value" :label="t(item.label)" :value="item.value" />
  </v-select>
</template>

<style lang="scss" scoped></style>
