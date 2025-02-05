<script lang="ts" setup>
import { RegistryInfo } from "@/types"
import { TableHeight } from "@/utils"
import { Action } from "@/models"
import RegistryEdit from "./registry-edit.vue"
import RegistryDelete from "./registry-delete.vue"
import RegistryView from "./registry-view.vue"
import { QueryRegistryInfo } from "./common.ts"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()

const tableHeight = computed(() => TableHeight(252))

const isLoading = ref(false)
const queryInfo = ref<QueryRegistryInfo>(new QueryRegistryInfo(route.query))

const tableList = ref({
  total: 0,
  data: [] as Array<RegistryInfo>
})

const registryEditRef = useTemplateRef<InstanceType<typeof RegistryEdit>>("registryEditRef")
const registryDeleteRef = useTemplateRef<InstanceType<typeof RegistryDelete>>("registryDeleteRef")
const registryViewRef = useTemplateRef<InstanceType<typeof RegistryView>>("registryViewRef")

async function search() {
  await router.replace(queryInfo.value.urlQuery())
  isLoading.value = true
  return await registryService
    .query(queryInfo.value.searchParams())
    .then(res => {
      tableList.value.data = res.list
      tableList.value.total = res.total
    })
    .finally(() => (isLoading.value = false))
}

function openAction(action: string, info?: RegistryInfo) {
  switch (action) {
    case Action.Add:
    case Action.Edit:
      registryEditRef.value?.open(info)
      break
    case Action.Delete:
      registryDeleteRef.value?.open(info!)
      break
    case Action.View:
      registryViewRef.value?.open(info!)
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
              {{ t("btn.addRegistry") }}
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
      <el-table-column :label="t('label.registry')" fixed="left" min-width="200" prop="registryName" sortable="custom" />
      <el-table-column :label="t('label.url')" min-width="200" prop="url" />
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
      <el-table-column :label="t('label.authentication')" align="center" width="140">
        <template #default="scope">
          <el-button v-if="scope.row?.hasAuth" link type="primary" @click="openAction(Action.View, scope.row)">{{ t("btn.view") }}</el-button>
          <span v-else>--</span>
        </template>
      </el-table-column>
      <el-table-column :label="t('label.isDefault')" align="center" width="140">
        <template #default="scope">
          <el-tag v-if="scope.row.isDefault" effect="dark" round size="small" type="warning">
            {{ t("label.default") }}
          </el-tag>
          <span v-else>--</span>
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
  <registry-delete ref="registryDeleteRef" @refresh="search()" />

  <registry-edit ref="registryEditRef" @refresh="search()" />

  <registry-view ref="registryViewRef" />
</template>

<style lang="scss" scoped></style>
