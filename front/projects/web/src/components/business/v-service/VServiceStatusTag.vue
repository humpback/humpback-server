<script lang="ts" setup>
import { toLower } from "lodash-es"

const props = defineProps<{ isEnabled?: boolean; status?: string }>()

const { t } = useI18n()

const statusInfo = computed<{ type: "primary" | "success" | "info" | "warning" | "danger"; i18nText: string }>(() => {
  if (!props.isEnabled) {
    return {
      type: "info",
      i18nText: "label.disabled"
    }
  }
  switch (toLower(props.status)) {
    case toLower(ServiceStatus.ServiceStatusNotReady):
      return {
        type: "warning",
        i18nText: "label.notReady"
      }
    case toLower(ServiceStatus.ServiceStatusFailed):
      return {
        type: "danger",
        i18nText: "label.failed"
      }
    default:
      return {
        type: "success",
        i18nText: "label.running"
      }
  }
})
</script>

<template>
  <el-tag :type="statusInfo.type" effect="dark">{{ t(statusInfo.i18nText) }}</el-tag>
</template>

<style lang="scss" scoped></style>
