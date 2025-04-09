<script lang="ts" setup>
const props = defineProps<{ data?: any[]; enabledServices?: number }>()
const { t } = useI18n()

const tableList = computed(() => props.data || [])
</script>

<template>
  <v-card>
    <div class="title">{{ t("header.abnormalNodes") }}</div>
    <div>
      <v-table v-if="tableList.length > 0" :data="tableList" hide-header-bg-color max-height="360px" minHeight="360px">
        <el-table-column :label="t('label.ip')">
          <template #default>aaa</template>
        </el-table-column>
        <el-table-column :label="t('label.status')" width="100px">
          <template #default>aaa</template>
        </el-table-column>
      </v-table>
      <div v-else :class="['empty-content', props.enabledServices && 'no-abnormal']">
        <el-empty>
          <template #image>
            <el-icon :size="160">
              <IconMdiServer />
            </el-icon>
          </template>
          <template #description>
            <span>{{ props.enabledServices ? t("tips.noAbnormalNodes") : t("tips.noEnabledNodes") }}</span>
          </template>
        </el-empty>
      </div>
    </div>
  </v-card>
</template>

<style lang="scss" scoped>
.title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 20px;
  padding-left: 4px;
}

.empty-content {
  height: 360px;
  font-size: 14px;
  color: gray;
}

.no-abnormal {
  color: #13b913;

  :deep(.el-icon) svg {
    color: #d8f2d8;
  }
}
</style>
