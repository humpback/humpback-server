<script lang="ts" setup>
import { NodeInfo } from "@/types"
import { BytesToGB, TableHeight } from "@/utils"
import { Action } from "@/models"
import NodeAdd from "./node-add.vue"
import NodeDelete from "./node-delete.vue"
import NodeEnable from "./node-enable.vue"
import NodeEditLabel from "./node-edit-label.vue"
import NodeViewCommand from "./node-view-command.vue"
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

const addRef = useTemplateRef<InstanceType<typeof NodeAdd>>("addRef")
const editLabelRef = useTemplateRef<InstanceType<typeof NodeEditLabel>>("editLabelRef")
const viewValueRef = useTemplateRef<InstanceType<typeof NodeViewCommand>>("viewValueRef")
const enableRef = useTemplateRef<InstanceType<typeof NodeEnable>>("enableRef")
const deleteRef = useTemplateRef<InstanceType<typeof NodeDelete>>("deleteRef")

async function search() {
  await router.replace(queryInfo.value.urlQuery())
  isLoading.value = true
  return await nodeService
    .query(queryInfo.value.searchParams())
    .then(res => {
      tableList.value.data = res.list
      tableList.value.total = res.total
    })
    .finally(() => (isLoading.value = false))

  // tableList.value.data.push(
  //   {
  //     nodeId: "wr",
  //     name: "e11dbts01.buyabs.corp",
  //     ipAddress: "172.16.171.52",
  //     port: 8566,
  //     status: "Online",
  //     isEnable: true,
  //     cpuUsage: 0,
  //     cpu: 8,
  //     memoryUsage: 0.36,
  //     memoryTotal: 62500000200,
  //     memoryUsed: 1343130000,
  //     labels: {
  //       test: "skyler",
  //       nice: "true",
  //       "bts-zk-node": "true",
  //       "eggkeeper-api": "true",
  //       "app": "true",
  //       "newegg.platform": "true",
  //       "etcd": "true",
  //       "skyler": "true"
  //     },
  //     createdAt: 0,
  //     updatedAt: 0
  //   },
  //   {
  //     nodeId: "wrsf2rf2",
  //     name: "e11dbts02.buyabs.corp",
  //     ipAddress: "172.16.171.53",
  //     port: 8566,
  //     status: "Offline",
  //     isEnable: true,
  //     cpuUsage: 0,
  //     cpu: 0,
  //     memoryUsage: 0,
  //     memoryTotal: 0,
  //     memoryUsed: 0,
  //     labels: {
  //       test: "skyler",
  //       nice: "true",
  //       "bts-zk-node": "true",
  //       "eggkeeper-api": "true",
  //       "app": "true",
  //       "newegg.platform": "true",
  //       "etcd": "true",
  //       "skyler": "true"
  //     },
  //     createdAt: 0,
  //     updatedAt: 0
  //   },
  //   {
  //     nodeId: "wrsf2rf2fss",
  //     name: "e11dbts03.buyabs.corp",
  //     ipAddress: "172.16.171.54",
  //     port: 8566,
  //     status: "Offline",
  //     isEnable: false,
  //     cpuUsage: 0,
  //     cpu: 0,
  //     memoryUsage: 0,
  //     memoryTotal: 0,
  //     memoryUsed: 0,
  //     labels: {
  //       test: "skyler",
  //       nice: "true",
  //       "bts-zk-node": "true",
  //       "eggkeeper-api": "true"
  //     },
  //     createdAt: 0,
  //     updatedAt: 0
  //   }
  // )
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
  switch (action) {
    case Action.Add:
      addRef.value?.open()
      break
    case Action.EditLabel:
      editLabelRef.value?.open(info!)
      break
    case Action.View:
      viewValueRef.value?.open(info!)
      break
    case Action.Delete:
      deleteRef.value?.open(info!)
      break
    case Action.Enable:
      enableRef.value?.open(info!)
  }
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
              show-out-label
              @change="search">
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
              {{ t("btn.addNodes") }}
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
          <div class="d-flex gap-2">
            <v-node-enable-tag :enabled="scope.row.isEnable" />
            <el-link :type="scope.row.isEnable ? 'primary' : 'info'">{{ scope.row.ipAddress }}</el-link>
          </div>
        </template>
      </el-table-column>
      <el-table-column :label="t('label.hostname')" min-width="200" prop="description">
        <template #default="scope">
          <v-table-column-none :text="scope.row.name" />
        </template>
      </el-table-column>
      <el-table-column :label="t('label.status')" min-width="400">
        <template #default="scope">
          <div class="custom-column">
            <div class="status">
              <div class="status-content">
                <div class="status-cpu">
                  <el-text size="small" type="info">
                    <strong>{{ t("label.cpu") }}</strong>
                    <div>{{ scope.row.cpu || "--/--" }} {{ t("label.core") }}</div>
                  </el-text>
                </div>
                <div class="status-memory">
                  <el-text size="small" type="info">
                    <strong>{{ t("label.memoryUsed") }}</strong>
                    <el-progress v-if="scope.row.memoryTotal" :percentage="Math.trunc(scope.row.memoryUsage * 100)" />
                    <div v-if="scope.row.memoryTotal">
                      {{ `${BytesToGB(scope.row.memoryUsed)} ${t("label.gb")} / ${BytesToGB(scope.row.memoryTotal)} ${t("label.gb")}` }}
                    </div>
                    <div v-else>
                      {{ `--/-- ${t("label.gb")}` }}
                    </div>
                  </el-text>
                </div>
              </div>
              <div class="status-tag">
                <v-node-status-tag v-if="scope.row.isEnable" :status="scope.row.status" />
              </div>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column :label="t('label.labels')" min-width="240">
        <template #default="scope">
          <div class="custom-column">
            <v-label-table-view :labels="scope.row.labels" :line="4" />
          </div>
        </template>
      </el-table-column>
      <el-table-column :label="t('label.action')" align="center" fixed="right" header-align="center" width="130">
        <template #default="scope">
          <el-button link type="primary" @click="openAction(Action.View, scope.row)">
            {{ t("btn.command") }}
          </el-button>
          <el-dropdown placement="bottom-end" @command="openAction($event, scope.row)">
            <el-button link type="primary">
              <el-icon :size="20">
                <IconMdiMoreHoriz />
              </el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item :command="Action.EditLabel">
                  <el-link :underline="false" type="primary">{{ t("btn.editLabel") }}</el-link>
                </el-dropdown-item>
                <el-dropdown-item :command="Action.Enable">
                  <el-link :type="scope.row.isEnable ? 'info' : 'success'" :underline="false">
                    {{ scope.row.isEnable ? t("btn.disable") : t("btn.enable") }}
                  </el-link>
                </el-dropdown-item>
                <el-dropdown-item :command="Action.Delete">
                  <el-link :underline="false" type="danger">{{ t("btn.delete") }}</el-link>
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </el-table-column>
    </v-table>
  </v-card>

  <node-add ref="addRef" @refresh="search()" />

  <node-view-command ref="viewValueRef" />

  <node-edit-label ref="editLabelRef" @refresh="search()" />

  <node-enable ref="enableRef" @refresh="search()" />

  <node-delete ref="deleteRef" @refresh="search()" />
</template>

<style lang="scss" scoped>
.custom-column {
  min-height: 60px;
}

.status {
  display: flex;
  align-items: start;
  gap: 20px;
  width: 100%;

  .status-content {
    display: flex;
    align-items: start;
    gap: 20px;
    flex: 1;

    .status-cpu {
      flex: 3;
      min-width: 100px;
    }

    .status-memory {
      flex: 7;
      min-width: 180px;
    }
  }

  .status-tag {
    width: 100px;
    text-align: right;
    padding-right: 20px;
  }
}
</style>
