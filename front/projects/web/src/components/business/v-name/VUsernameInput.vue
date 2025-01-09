<script lang="ts" setup>
const props = withDefaults(
  defineProps<{
    size?: "large" | "default" | "small"
    placeholder?: string
    disabled?: boolean
    clearable?: boolean
    showWordLimit?: boolean
  }>(),
  {
    clearable: true,
    showWordLimit: true,
    placeholder: ""
  }
)

const name = defineModel<string>()
</script>

<template>
  <v-input
    v-model="name"
    :clearable="props.clearable"
    :disabled="props.disabled"
    :maxlength="LimitUserName.Max"
    :placeholder="props.placeholder"
    :show-word-limit="props.showWordLimit"
    :size="props.size" />
</template>

<style lang="scss" scoped></style>
