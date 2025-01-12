<script lang="ts" setup>
import { CopyToClipboard } from "./copy.ts"

const props = withDefaults(
  defineProps<{
    msg: string
    width?: number
    disabled?: boolean
    placement?:
      | "top"
      | "top-start"
      | "top-end"
      | "bottom"
      | "bottom-start"
      | "bottom-end"
      | "left"
      | "left-start"
      | "left-end"
      | "right"
      | "right-start"
      | "right-end"
  }>(),
  { placement: "top-start" }
)

const { t } = useI18n()
const slots = useSlots()
</script>

<template>
  <el-popover :disabled="props.disabled" :placement="props.placement" :width="props.width" trigger="hover">
    <template v-if="!!slots.reference" #reference>
      <slot name="reference"></slot>
    </template>
    <div v-if="props.msg">
      <slot v-if="!!slots.message" name="message"></slot>
      <div v-else>{{ props.msg }}</div>
      <el-divider border-style="dashed" class="divider" />
      <el-button link type="success" @click="CopyToClipboard(props.msg)">
        <el-icon size="14">
          <IconMdiContentCopy />
        </el-icon>
        {{ t("btn.copy") }}
      </el-button>
    </div>
  </el-popover>
</template>

<style lang="scss" scoped>
.divider {
  margin: 8px 0 4px 0;
}
</style>
