<script lang="ts" setup>
const props = defineProps<{
  showOutLabel?: boolean
  outLabelWidth?: string
  outLabel?: string
}>()

const { t } = useI18n()

const startTime = defineModel<number>("startTime")
const endTime = defineModel<number>("endTime")

const dateTimeRange = computed<any>({
  get() {
    if (startTime.value && endTime.value) {
      return [startTime.value, endTime.value]
    }
    return []
  },
  set(value) {
    if (Array.isArray(value) && value.length == 2) {
      startTime.value = value[0]
      endTime.value = value[1]
    } else {
      startTime.value = 0
      endTime.value = 0
    }
  }
})

const shortcuts = ref([
  {
    text: t("label.last10Minutes"),
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setMinutes(start.getMinutes() - 10)
      return [start, end]
    }
  },
  {
    text: t("label.lastHour"),
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setHours(start.getHours() - 1)
      return [start, end]
    }
  },
  {
    text: t("label.lastDay"),
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setDate(start.getDay() - 1)
      return [start, end]
    }
  }
])
</script>

<template>
  <div class="date-picker">
    <div v-if="props.showOutLabel" :style="{ width: props.outLabelWidth }" class="out-label">{{ props.outLabel }}</div>
    <el-date-picker
      v-model="dateTimeRange"
      :end-placeholder="t('label.endTime')"
      :range-separator="t('label.to')"
      :shortcuts="shortcuts"
      :start-placeholder="t('label.startTime')"
      type="datetimerange"
      value-format="x" />
  </div>
</template>

<style lang="scss" scoped>
.date-picker {
  display: flex;
  align-items: center;

  .out-label {
    background-color: var(--el-fill-color-light);
    line-height: 24px;
    min-height: 32px;
    font-size: 14px;
    color: var(--el-text-color-regular);
    border: 1px solid var(--el-border-color);
    box-sizing: border-box;
    border-top-left-radius: 4px;
    border-bottom-left-radius: 4px;
    border-right: none;
    padding: 0 12px;
    display: flex;
    align-items: center;
  }

  :deep(.el-date-editor) {
    border-bottom-left-radius: unset;
    border-top-left-radius: unset;
  }
}
</style>
