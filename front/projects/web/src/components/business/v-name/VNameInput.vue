<script lang="ts" setup>
const props = withDefaults(
  defineProps<{
    modelValue?: string
    size?: "large" | "default" | "small"
    placeholder?: string
    disabled?: boolean
    clearable?: boolean
    showWordLimit?: boolean
    showTitle?: boolean
  }>(),
  {
    modelValue: "",
    clearable: true,
    showWordLimit: true,
    showTitle: false,
    placeholder: ""
  }
)
const emits = defineEmits<{
  (e: "update:model-value", data: string): void
}>()

const { t } = useI18n()
const name = computed({
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
    v-model="name"
    :clearable="props.clearable"
    :disabled="props.disabled"
    :placeholder="props.placeholder"
    :show-word-limit="props.showWordLimit"
    :size="props.size">
    <template v-if="props.showTitle" #prefix>
      <el-text :size="props.size">{{ t("label.name") }}</el-text>
      <v-divider />
    </template>
  </v-input>
</template>

<style lang="scss" scoped></style>
