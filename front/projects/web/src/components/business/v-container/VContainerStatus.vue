<script lang="ts" setup>
import { ContainerStatus } from "@/models"
import { toLower } from "lodash-es"

const props = defineProps<{ status?: string; size?: "large" | "default" | "small" }>()

const { t } = useI18n()

const statusInfo = computed<{ type: "primary" | "success" | "info" | "warning" | "danger"; i18nText: string }>(() => {
  switch (toLower(props.status)) {
    case toLower(ContainerStatus.ContainerStatusStarting):
      return {
        type: "primary",
        i18nText: "enum.status.starting"
      }
    case toLower(ContainerStatus.ContainerStatusCreated):
      return {
        type: "info",
        i18nText: "enum.status.created"
      }
    case toLower(ContainerStatus.ContainerStatusRunning):
      return {
        type: "success",
        i18nText: "enum.status.running"
      }
    case toLower(ContainerStatus.ContainerStatusFailed):
      return {
        type: "danger",
        i18nText: "enum.status.failed"
      }
    case toLower(ContainerStatus.ContainerStatusExited):
      return {
        type: "warning",
        i18nText: "enum.status.exited"
      }
    case toLower(ContainerStatus.ContainerStatusRemoved):
      return {
        type: "danger",
        i18nText: "enum.status.removed"
      }
    case toLower(ContainerStatus.ContainerStatusWarning):
      return {
        type: "warning",
        i18nText: "enum.status.warning"
      }
    default:
      return {
        type: "primary",
        i18nText: "enum.status.pending"
      }
  }
})
</script>

<template>
  <el-tag :size="props.size" :type="statusInfo.type" effect="dark">{{ t(statusInfo.i18nText) }}</el-tag>
</template>

<style lang="scss" scoped></style>
