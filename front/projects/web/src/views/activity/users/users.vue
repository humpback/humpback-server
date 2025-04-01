<script lang="ts" setup>
import { QueryActivityUserInfo } from "./common.ts"
import { TableHeight } from "@/utils"

const { t } = useI18n()
const route = useRoute()
const userStore = useUserStore()

const tableHeight = computed(() => TableHeight(346))

const isLoading = ref(false)
const filterInfo = ref<QueryActivityUserInfo>(new QueryActivityUserInfo(route.params))

const tableList = ref({
  total: 0,
  data: [] as Array<any>
})

const setRowClass = ({ row }) => {
  return !row.oldContent && !row.newContent ? "hide-expand-icon" : ""
}

function search() {}
</script>

<template>
  <div class="d-flex gap-3 flex-wrap">
    <v-date-time-range
      v-model:end-time="filterInfo.filter.endAt"
      v-model:start-time="filterInfo.filter.startAt"
      :end-placeholder="t('label.endTime')"
      :out-label="t('label.timeRange')"
      :range-separator="t('label.to')"
      :start-placeholder="t('label.startTime')"
      format="YYYY-MM-DD HH:mm"
      out-label-width="120px"
      time-format="HH:mm"
      @change="search()" />
    <div style="width: 280px">
      <v-activity-action-select v-model="filterInfo.filter.action" mode="user" show-out-label />
    </div>
    <div v-if="userStore.isAdmin" style="width: 280px">
      <v-users-select v-model="filterInfo.filter.user" :multiple="false" :out-label="t('label.operator')" clearable show-out-label />
    </div>
    <el-button type="primary">
      <template #icon>
        <el-icon :size="20">
          <IconMdiFilterVariant />
        </el-icon>
      </template>
      {{ t("btn.filter") }}
    </el-button>
  </div>

  <v-table
    v-loading="isLoading"
    v-model:page-info="filterInfo.pageInfo"
    :data="tableList.data"
    :max-height="tableHeight"
    :row-class-name="setRowClass"
    :total="tableList.total"
    class="mt-5"
    row-key="activeId"
    @page-change="search">
    <el-table-column align="left" class-name="expand-column" type="expand" width="24">
      <template #default="scope">
        <div class="px-5">
          <v-monaco-diff-editor
            v-if="scope.row.hasContent"
            :new-data="JSON.stringify(scope.row.oldContent, null, 4)"
            :old-data="JSON.stringify(scope.row.newContent, null, 4)"
            language="json" />
        </div>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.user')" min-width="140" prop="config">
      <template #default="scope">
        <v-table-column-none :text="scope.row.config" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.action')" min-width="140" prop="action">
      <template #default="scope">
        <v-table-column-none :text="scope.row.action" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.description')" min-width="200" prop="description">
      <template #default="scope">
        <v-table-column-none :text="scope.row.description" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.operator')" min-width="140" prop="user">
      <template #default="scope">
        <v-table-column-none :text="scope.row.user" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.date')" width="140">
      <template #default="scope">
        <v-date-view :timestamp="scope.row.createdAt" />
      </template>
    </el-table-column>
  </v-table>
</template>

<style lang="scss" scoped>
:deep(.hide-expand-icon) {
  .expand-column .cell {
    padding-top: 4px;
    display: none;
  }
}
</style>
