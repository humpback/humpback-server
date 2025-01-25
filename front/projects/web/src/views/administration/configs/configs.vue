<script lang="ts" setup>
import { ConfigInfo } from "@/types"
import { TableHeight } from "@/utils"
import { Action, ConfigType } from "@/models"
import ConfigEdit from "./config-edit.vue"
import ConfigDelete from "./config-delete.vue"
import ConfigView from "./config-view.vue"
import { QueryConfigsInfo } from "./common.ts"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()

const tableHeight = computed(() => TableHeight(252))

const isLoading = ref(false)
const queryInfo = ref<QueryConfigsInfo>(new QueryConfigsInfo(route.query))

const tableList = ref({
  total: 0,
  data: [] as Array<ConfigInfo>
})

const editRef = useTemplateRef<InstanceType<typeof ConfigEdit>>("editRef")
const deleteRef = useTemplateRef<InstanceType<typeof ConfigDelete>>("deleteRef")
const viewValueRef = useTemplateRef<InstanceType<typeof ConfigView>>("viewValueRef")

async function search() {
  await router.replace(queryInfo.value.getQuery())
  isLoading.value = true
  return await configService
    .query(queryInfo.value.getSearch())
    .then(res => {
      tableList.value.data = res.list
      tableList.value.total = res.total
    })
    .finally(() => (isLoading.value = false))
}

function openAction(action: string, info?: ConfigInfo) {
  switch (action) {
    case Action.Add:
    case Action.Edit:
      editRef.value?.open(info)
      break
    case Action.Delete:
      deleteRef.value?.open(info!)
      break
    case Action.View:
      viewValueRef.value?.open(info!)
      break
  }
}

onMounted(() => search())
</script>

<template>
  <v-card>
    <el-form @submit.prevent="search">
      <el-form-item>
        <div class="d-flex gap-3 w-100 flex-wrap">
          <div>
            <v-config-type-query-select v-model="queryInfo.filter.configType" :placeholder="t('placeholder.all')" @change="search()" />
          </div>
          <div class="flex-1" style="min-width: 300px">
            <v-input v-model="queryInfo.keywords">
              <template #prepend>
                <span>{{ t("label.name") }}</span>
              </template>
            </v-input>
          </div>
          <div>
            <el-button native-type="submit" type="primary">{{ t("btn.search") }}</el-button>
            <el-button plain type="primary" @click="openAction(Action.Add)">
              <template #icon>
                <el-icon :size="20">
                  <IconMdiAdd />
                </el-icon>
              </template>
              {{ t("btn.addConfig") }}
            </el-button>
          </div>
        </div>
      </el-form-item>
    </el-form>

    <v-table
      v-loading="isLoading"
      v-model:page-info="queryInfo.pageInfo"
      v-model:sort-info="queryInfo.sortInfo"
      :data="tableList.data"
      :max-height="tableHeight"
      :total="tableList.total"
      @page-change="search"
      @sort-change="search">
      <el-table-column :label="t('label.config')" fixed="left" min-width="160" prop="configName" sortable="custom" />
      <el-table-column :label="t('label.description')" min-width="140" prop="description">
        <template #default="scope">
          <v-table-column-none :text="scope.row.description" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.type')" min-width="120">
        <template #default="scope">
          <v-config-type-view :configType="scope.row.configType" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.value')" min-width="200">
        <template #default="scope">
          <span v-if="scope.row.configType === ConfigType.Static">{{ scope.row.configValue }}</span>
          <el-button v-else link type="primary" @click="openAction(Action.View, scope.row)">{{ t("btn.view") }}</el-button>
        </template>
      </el-table-column>
      <el-table-column :label="t('label.updateDate')" min-width="140" prop="updatedAt" sortable="custom">
        <template #default="scope">
          <v-date-view :timestamp="scope.row.updatedAt" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.createDate')" min-width="140" prop="createdAt" sortable="custom">
        <template #default="scope">
          <v-date-view :timestamp="scope.row.createdAt" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.action')" align="right" fixed="right" header-align="center" width="130">
        <template #default="scope">
          <el-button link type="primary" @click="openAction(Action.Edit, scope.row)">{{ t("btn.edit") }}</el-button>
          <el-button link type="danger" @click="openAction(Action.Delete, scope.row)">{{ t("btn.delete") }}</el-button>
        </template>
      </el-table-column>
    </v-table>
  </v-card>
  <config-delete ref="deleteRef" @refresh="search()" />

  <config-edit ref="editRef" @refresh="search()" />

  <config-view ref="viewValueRef" />
</template>

<style lang="scss" scoped></style>
