<script lang="ts" setup>
const props = defineProps<{ data?: any[] }>()
const { t } = useI18n()

const tableList = computed(() => props.data || [])
</script>

<template>
  <v-card>
    <div class="title">{{ t("header.recentActivities") }}</div>
    <div>
      <v-table v-if="tableList.length > 0" :data="tableList" hide-header-bg-color max-height="360px" minHeight="360px">
        <el-table-column :label="t('label.description')">
          <template #default>aaa</template>
        </el-table-column>
        <el-table-column :label="t('label.date')" width="120px">
          <template #default>aaa</template>
        </el-table-column>
      </v-table>
      <div v-else class="empty-content">
        <el-empty>
          <template #image>
            <el-icon :size="160">
              <IconMdiAccountFileText />
            </el-icon>
          </template>
          <template #description>
            <span>{{ t("tips.noRecentActivities") }}</span>
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
  color: gray;
  font-size: 14px;
}
</style>
