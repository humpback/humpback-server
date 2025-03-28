<script lang="ts" setup>
import { QueryActivityGroupInfo } from "./common.ts"

const { t } = useI18n()
const route = useRoute()

const filterInfo = ref<QueryActivityGroupInfo>(new QueryActivityGroupInfo(route.params))

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
    <div>
      <v-group-select v-model="filterInfo.filter.group" :placeholder="t('placeholder.all')" show-out-label />
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
</template>

<style lang="scss" scoped></style>
