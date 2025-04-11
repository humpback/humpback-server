<script lang="ts" setup>
import { TableHeight } from "@/utils"
import { ActivityInfo } from "@/types"

const { t } = useI18n()
const route = useRoute()

const tableHeight = computed(() => TableHeight(362))

const isLoading = ref(false)
const queryInfo = ref<any>({
  filter: {
    "type": "services",
    "serviceId": route.params.serviceId as string
  },
  pageInfo: {
    size: 20,
    index: 1
  }
})

const tableList = ref({
  total: 0,
  data: [] as Array<ActivityInfo>
})

const setRowClass = ({ row }) => {
  return !row.oldContent && !row.newContent ? "hide-expand-icon" : ""
}

async function getActivities() {
  isLoading.value = true
  return await activityService
    .query(queryInfo.value)
    .then(data => {
      tableList.value.total = data.total
      tableList.value.data = data.list
    })
    .finally(() => (isLoading.value = false))
}

onMounted(() => {
  getActivities()
})
</script>

<template>
  <div class="d-flex gap-3">
    <div class="header-icon">
      <el-icon :size="18">
        <IconMdiAccountFileText />
      </el-icon>
    </div>
    <div class="d-flex gap-1">
      <el-text class="f-bold" size="large">{{ t("label.activities") }}</el-text>
      <el-button :disabled="isLoading" :title="t('label.refresh')" link type="primary" @click="getActivities()">
        <el-icon v-if="!isLoading" :size="20">
          <IconMdiRefresh />
        </el-icon>
        <v-loading v-else />
      </el-button>
    </div>
  </div>

  <v-table
    v-loading="isLoading"
    v-model:page-info="queryInfo.pageInfo"
    :data="tableList.data"
    :max-height="tableHeight"
    :row-class-name="setRowClass"
    :total="tableList.total"
    class="mt-5"
    row-key="activityId"
    @page-change="getActivities()">
    <el-table-column align="left" class-name="expand-column" type="expand" width="24">
      <template #default="scope">
        <div class="px-5">
          <v-monaco-diff-editor
            v-if="scope.row.oldContent || scope.row.newContent"
            :new-data="scope.row.newContent ? JSON.stringify(scope.row.newContent, null, 4) : ''"
            :old-data="scope.row.oldContent ? JSON.stringify(scope.row.oldContent, null, 4) : ''"
            language="json" />
        </div>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.description')" min-width="200">
      <template #default="scope">
        <span v-if="scope.row.instanceName">
          {{ t(`activity.service.${scope.row.action}Instance`, { name: scope.row.instanceName }) }}
        </span>
        <span v-else>{{ t(`activity.service.${scope.row.action}`) }} </span>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.operator')" min-width="140" prop="operator" />
    <el-table-column :label="t('label.date')" width="160">
      <template #default="scope">
        <v-date-view :timestamp="scope.row.operateAt" />
      </template>
    </el-table-column>
  </v-table>
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

:deep(.expand-column) {
  .cell {
    padding: 0 4px 0 8px;
  }
}

:deep(.hide-expand-icon) {
  .expand-column .cell {
    padding-top: 4px;
    display: none;
  }
}
</style>
