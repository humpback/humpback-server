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
  <div class="copy-box" @click="CopyToClipboard(props.text)">
    <div class="overflow_div">
      <el-text :size="props.size">{{ props.text }}</el-text>
    </div>
    <div class="d-flex copy-icon">
      <el-icon :size="iconSize">
        <IconMdiContentCopy />
      </el-icon>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.copy-box {
  display: flex;
  align-items: center;
  gap: 4px;

  &:hover {
    opacity: 0.7; /* 设置透明度 */
    cursor: pointer;

    .copy-icon {
      color: var(--el-color-primary);
    }
  }
}
</style>
