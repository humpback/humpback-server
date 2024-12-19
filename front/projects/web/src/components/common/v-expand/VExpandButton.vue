<script lang="ts" setup>
const props = withDefaults(
  defineProps<{
    showType?: "icon" | "text" | "all"
    text?: string
    size?: "large" | "default" | "small"
    type?: "primary" | "success" | "warning" | "danger" | "info" | "text"
  }>(),
  {
    showType: "all",
    size: "default"
  }
)

const { t } = useI18n()
const slots = useSlots()
const isExpanded = defineModel<boolean>()

const iconSize = computed(() => {
  switch (props.size) {
    case "small":
      return 20
    case "large":
      return 28
    default:
      return 24
  }
})
</script>

<template>
  <el-button :size="props.size" :type="props.type" link @click="isExpanded = !isExpanded">
    <div class="d-flex">
      <slot v-if="!!slots.text" name="text" />
      <span v-else-if="props.text">{{ props.text }}</span>
      <span v-else-if="props.showType === 'text' || props.showType === 'all'" style="margin-left: 0">
        {{ isExpanded ? t("label.collapse") : t("label.expand") }}
      </span>
      <el-icon v-if="props.showType === 'icon' || props.showType === 'all'" :size="iconSize" style="margin-left: -2px">
        <IconMdiChevronDown :class="{ 'is-active': isExpanded }" class="arrow-icon" />
      </el-icon>
    </div>
  </el-button>
</template>

<style lang="scss" scoped>
.arrow-icon {
  transition: transform ease 150ms;

  &.is-active {
    transform: scaleY(-1);
  }
}
</style>
