<script lang="ts" setup>
const props = withDefaults(
  defineProps<{
    modelValue: number
    size?: "large" | "default" | "small"
    clearable?: boolean
    placeholder?: string
  }>(),
  { size: "default", clearable: true, placeholder: "" }
)
const emits = defineEmits<{
  (e: "update:modelValue", data: number): void
  (e: "change", data: number): void
}>()

const { t } = useI18n()

const role = computed({
  get: () => props.modelValue || 0,
  set: (v: number) => {
    emits("update:modelValue", v || 0)
  }
})

const options = computed(() => [
  { label: "role.user", value: UserRole.User },
  { label: "role.admin", value: UserRole.Admin },
  { label: "role.supperAdmin", value: UserRole.SupperAdmin }
])

function change() {
  emits("change", role.value)
}
</script>

<template>
  <div class="d-flex role-select">
    <div class="label">{{ t("label.role") }}</div>
    <el-select v-model="role" :clearable="props.clearable" :placeholder="props.placeholder" class="role-content-select" @change="change()">
      <el-option :label="t('label.all')" :value="0" />
      <el-option v-for="item in options" :key="item.value" :label="t(item.label)" :value="item.value" />
    </el-select>
  </div>
</template>

<style lang="scss" scoped>
.role-select {
  min-width: 260px;

  .label {
    width: 100px;
    background-color: var(--el-fill-color-light);
    height: 32px;
    color: var(--el-text-color-regular);
    border: 1px solid var(--el-border-color);
    box-sizing: border-box;
    border-top-left-radius: 4px;
    border-bottom-left-radius: 4px;
    border-right: none;
    padding: 0 12px;
  }

  .role-content-select {
    :deep(.el-select__wrapper) {
      border-top-left-radius: 0;
      border-bottom-left-radius: 0;
    }
  }
}
</style>
