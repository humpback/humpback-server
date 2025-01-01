<script lang="ts" setup>
import type { DividerProps } from "element-plus"
import { omit } from "lodash"

type Props = Partial<
  DividerProps & {
    margin: string
  }
>

const props = withDefaults(defineProps<Props>(), {
  direction: "vertical",
  borderStyle: "solid",
  contentPosition: "center"
})

const slots = useSlots()

const style = computed(() => (props.direction === "vertical" && props.margin ? { margin: props.margin } : {}))
</script>

<template>
  <el-divider :style="style" v-bind="{ ...omit(props, ['margin']) }">
    <template v-if="!!slots.default" #default>
      <slot name="default"></slot>
    </template>
  </el-divider>
</template>

<style scoped></style>
