<script lang="ts" setup>
import { toLower } from "lodash-es"

const props = withDefaults(
  defineProps<{
    isEnabled?: boolean
    status?: string
    isText?: boolean
    effect?: "dark" | "light" | "plain"
  }>(),
  { effect: "dark" }
)

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
  <el-text v-if="props.isText" :type="statusInfo.type">{{ t(statusInfo.i18nText) }}</el-text>
  <el-tag v-else :type="statusInfo.type" effect="dark"> {{ t(statusInfo.i18nText) }}</el-tag>
</template>

<style lang="scss" scoped></style>
