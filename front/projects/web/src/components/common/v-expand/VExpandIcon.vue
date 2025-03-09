<script lang="ts" setup>
const props = withDefaults(defineProps<{ iconSize?: number; expand?: boolean; disabled?: boolean }>(), { iconSize: 16 })
const emits = defineEmits<{
  (e: "expand"): void
}>()

function expandClick() {
  if (props.disabled) {
    return
  }
  emits("expand")
}
</script>

<template>
  <el-icon :class="['expand-icon', props.expand && !props.disabled && 'expanded', props.disabled && 'disabled']" :size="props.iconSize" @click="expandClick()">
    <IconMdiChevronRight />
  </el-icon>
</template>

<style lang="scss" scoped>
.expand-icon {
  display: inline-block;
  color: var(--el-disabled-text-color);

  &:not(.disabled) {
    color: var(--el-text-color);
    transition: transform 0.3s ease;
    cursor: pointer;
  }
}

.expanded {
  transform: rotate(90deg);
}
</style>
