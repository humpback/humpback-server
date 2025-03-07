<script lang="ts" setup>
const props = withDefaults(defineProps<{ size?: number }>(), { size: 14 })
</script>

<template>
  <div :style="{ width: `${props.size}px`, height: `${props.size}px` }" class="spinner" />
</template>

<style lang="scss" scoped>
.spinner {
  width: 14px;
  height: 14px;
  border: 2px solid #e3e2e2; /* 外圈颜色 */
  border-top: 2px solid var(--el-color-primary); /* 内圈颜色 */
  border-radius: 50%; /* 圆形 */
  animation: spin 1s linear infinite; /* 旋转动画 */
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
