<script lang="ts" setup>
const { t } = useI18n()
</script>

<template>
  <div class="d-flex gap-3">
    <div class="header-icon">
      <el-icon :size="18">
        <IconMdiAccountFileText />
      </el-icon>
    </div>
    <el-text class="f-bold" size="large">{{ t("label.activity") }}</el-text>
  </div>
</template>

<style lang="scss" scoped>
.header-icon {
  color: var(--el-color-primary);
  background-color: var(--el-color-primary-light-9);
  border-radius: 50%;
  padding: 8px;
  width: 18px;
  height: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
