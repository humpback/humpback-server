<script lang="ts" setup>
import { NodeInfo } from "@/types"
import { TableHeight } from "@/utils"
import { Action } from "@/models"
import NodeEdit from "./node-edit.vue"
import NodeDelete from "./node-delete.vue"
import NodeView from "./node-view.vue"
import { QueryNodesInfo, statusOptions, modeOptions } from "./common.ts"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()

const tableHeight = computed(() => TableHeight(252))

const isLoading = ref(false)
const queryInfo = ref<QueryNodesInfo>(new QueryNodesInfo(route.query, []))

const tableList = ref({
  total: 0,
  data: [] as Array<NodeInfo>
})

const editRef = useTemplateRef<InstanceType<typeof NodeEdit>>("editRef")
const deleteRef = useTemplateRef<InstanceType<typeof NodeDelete>>("deleteRef")
const viewValueRef = useTemplateRef<InstanceType<typeof NodeView>>("viewValueRef")

async function search() {
  await router.replace(queryInfo.value.urlQuery())
  tableList.value.data.push({
    nodeId: "wr",
    name: "e11dbts01.buyabs.corp",
    ipAddress: "172.16.171.52",
    port: 8566,
    status: "Online",
    isEnable: true,
    cpuUsage: 0,
    memoryUsage: 0,
    labels: {},
    createdAt: 0,
    updatedAt: 0
  })
  // isLoading.value = true
  // return await nodeService
  //   .query(queryInfo.value.getSearch())
  //   .then(res => {
  //     tableList.value.data = res.list
  //     tableList.value.total = res.total
  //   })
  //   .finally(() => (isLoading.value = false))
}

function openAction(action: string, info?: NodeInfo) {
  // switch (action) {
  //   case Action.Add:
  //   case Action.Edit:
  //     editRef.value?.open(info)
  //     break
  //   case Action.Delete:
  //     deleteRef.value?.open(info!)
  //     break
  //   case Action.View:
  //     viewValueRef.value?.open(info!)
  //     break
  // }
}

onMounted(() => search())
</script>

<template>
  <v-card>
    <el-form @submit.prevent="search">
      <el-form-item>
        <div class="d-flex gap-3 w-100 flex-wrap">
          <div style="width: 220px">
            <v-select
              v-model="queryInfo.filter.status"
              :out-label="t('label.status')"
              :placeholder="t('placeholder.all')"
              clearable
              out-label-width="80px"
              show-out-label>
              <el-option v-for="(item, index) in statusOptions" :key="index" :label="t(item.label)" :value="item.value" />
            </v-select>
          </div>
          <div class="flex-1" style="min-width: 300px">
            <v-input
              v-model="queryInfo.keywords"
              :placeholder="queryInfo.mode === 'keywords' ? t('placeholder.enterIpOrHostname') : t('placeholder.enterLabelKey')">
              <template #prepend>
                <el-select v-model="queryInfo.mode" placeholder="" style="width: 120px">
                  <el-option v-for="item in modeOptions" :key="item.value" :label="t(item.label)" :value="item.value" />
                </el-select>
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
              {{ t("btn.addNode") }}
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
      <el-table-column :label="t('label.ip')" fixed="left" min-width="160" prop="nodeName" sortable="custom">
        <template #default="scope">
          <div class="custom-column">
            <div class="d-flex gap-2">
              <v-node-enable-tag :enabled="scope.row.isEnable" />
              <el-link type="primary">{{ scope.row.ipAddress }}</el-link>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column :label="t('label.hostname')" min-width="140" prop="description">
        <template #default="scope">
          <div class="custom-column">
            <v-table-column-none :text="scope.row.name" />
          </div>
        </template>
      </el-table-column>
      <el-table-column :label="t('label.status')" min-width="220">
        <template #default="scope">
          <v-node-status-tag :status="scope.row.status" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.resources')" min-width="200">
        <template #default="scope"></template>
      </el-table-column>
      <el-table-column :label="t('label.labels')" min-width="200" />
      <!--      <el-table-column :label="t('label.updateDate')" min-width="140" prop="updatedAt" sortable="custom">-->
      <!--        <template #default="scope">-->
      <!--          <v-date-view :timestamp="scope.row.updatedAt" />-->
      <!--        </template>-->
      <!--      </el-table-column>-->
      <!--      <el-table-column :label="t('label.createDate')" min-width="140" prop="createdAt" sortable="custom">-->
      <!--        <template #default="scope">-->
      <!--          <v-date-view :timestamp="scope.row.createdAt" />-->
      <!--        </template>-->
      <!--      </el-table-column>-->
      <el-table-column :label="t('label.action')" align="right" fixed="right" header-align="center" width="130">
        <template #default="scope">
          <el-button link type="primary" @click="openAction(Action.Edit, scope.row)">{{ t("btn.edit") }}</el-button>
          <el-button link type="danger" @click="openAction(Action.Delete, scope.row)">{{ t("btn.delete") }}</el-button>
        </template>
      </el-table-column>
    </v-table>
  </v-card>
  <node-delete ref="deleteRef" @refresh="search()" />

  <node-edit ref="editRef" @refresh="search()" />

  <node-view ref="viewValueRef" />
</template>

<style lang="scss" scoped>
.custom-column {
  min-height: 60px;
}
</style>
