<script lang="ts" setup>
import { CopyToClipboard } from "@/utils/index.ts"

const props = defineProps<{ text: string; size?: "large" | "small" | "default" }>()

const iconSize = computed(() => {
  switch (props.size) {
    case "large":
      return 16
    case "small":
      return 12
    default:
      return 14
  }
})
</script>

<template>
  <el-icon v-if="props.text" :size="iconSize" class="copy-icon" @click="CopyToClipboard(props.text)">
    <IconMdiContentCopy />
  </el-icon>
</template>

<style lang="scss" scoped>
.copy-icon:hover {
  opacity: 0.7; /* 设置透明度 */
  cursor: pointer;
  color: var(--el-color-primary);
}
</style>
