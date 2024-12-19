<script lang="ts" setup>
import { TimeExpiration, TimestampToTime } from "@/utils"

const props = withDefaults(defineProps<{ timestamp?: number; checkNearExpiration?: boolean; format?: number }>(), {
  timestamp: 0,
  checkNearExpiration: false,
  format: 3
})
const textColor = computed(() => {
  if (!props.checkNearExpiration || !props.timestamp) {
    return ""
  }
  if (TimeExpiration(props.timestamp)) {
    return "danger"
  }
  return ""
})
</script>

<template>
  <el-text :type="textColor">
    {{ TimestampToTime(props.timestamp, props.format) }}
  </el-text>
</template>

<style lang="scss" scoped></style>
