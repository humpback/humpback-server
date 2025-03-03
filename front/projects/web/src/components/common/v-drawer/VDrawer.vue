<script lang="ts" setup>
import { omit } from "lodash-es"
import { DrawerProps } from "element-plus"

type Props = Partial<DrawerProps>

const props = withDefaults(defineProps<Props>(), {
  showClose: true,
  destroyOnClose: true,
  withHeader: true,
  closeOnPressEscape: true,
  modal: true,
  lockScroll: true,
  appendToBody: true
})
const emits = defineEmits<{
  (e: "update:model-value", data: boolean): void
  (e: "closed"): void
}>()

const pageStore = usePageStore()
const slots = useSlots()

const attrs = computed(() => {
  return omit(props, ["size"])
})

const size = computed(() => {
  if (!pageStore.isBigScreen) {
    return "100%"
  }
  return props.size
})

function changeModelValue(v: boolean) {
  emits("update:model-value", v)
}

function closedDrawer() {
  emits("closed")
}
</script>

<template>
  <el-drawer :size="size" class="custom-drawer" v-bind="{ ...attrs }" @closed="closedDrawer()" @update:modelValue="changeModelValue">
    <template v-if="!!slots.header" #header>
      <slot name="header" />
    </template>
    <template v-if="!!slots.default" #default>
      <slot />
    </template>
    <template v-if="!!slots.footer" #footer>
      <slot name="footer" />
    </template>
  </el-drawer>
</template>

<style scoped></style>
