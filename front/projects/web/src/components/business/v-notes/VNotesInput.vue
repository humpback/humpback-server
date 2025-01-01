<script lang="ts" setup>
const props = withDefaults(
  defineProps<{
    modelValue: string
    size?: "large" | "default" | "small"
    placeholder?: string
    row?: number
    autosize?: { minRows?: number; maxRows?: number } | boolean | undefined
    resize?: "none" | "both" | "horizontal" | "vertical" | undefined
  }>(),
  {
    modelValue: "",
    row: 2,
    resize: "vertical",
    autosize: () => {
      return { minRows: 2, maxRows: 4 }
    }
  }
)
const emits = defineEmits<{
  (e: "update:model-value", data: string): void
}>()

const notes = computed({
  get() {
    return props.modelValue
  },
  set(data) {
    emits("update:model-value", data)
  }
})
</script>

<template>
  <v-input
    v-model="notes"
    :autosize="props.autosize"
    :placeholder="props.placeholder"
    :resize="props.resize"
    :row="2"
    :size="props.size"
    clearable
    show-word-limit
    type="textarea" />
</template>

<style lang="scss" scoped></style>
