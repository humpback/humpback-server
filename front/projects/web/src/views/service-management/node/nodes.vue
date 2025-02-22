<script lang="ts" setup>
import { NodeInfo } from "@/types"
import { BytesToGB, TableHeight } from "@/utils"
import { Action } from "@/models"
import NodeAdd from "./node-add.vue"
import NodeRemove from "./node-remove.vue"
import { QueryGroupNodesInfo } from "./common.ts"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const tableHeight = computed(() => TableHeight(252))

const groupId = ref(route.params?.groupId as string)

const isLoading = ref(false)
const queryInfo = ref<QueryGroupNodesInfo>(new QueryGroupNodesInfo(route.query))

const tableList = ref({
  total: 0,
  data: [] as Array<NodeInfo>
})

const addRef = useTemplateRef<InstanceType<typeof NodeAdd>>("addRef")
const deleteRef = useTemplateRef<InstanceType<typeof NodeRemove>>("deleteRef")

async function getGroupInfo() {
  return await groupService.info(groupId.value).then(info => {
    stateStore.setGroup(groupId.value, info)
  })
}

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
}

function openAction(action: string, info?: NodeInfo) {
  switch (action) {
    case Action.Add:
      addRef.value?.open()
      break
    case Action.Delete:
      deleteRef.value?.open(info!)
      break
  }
}

onMounted(() => search())
</script>

<template>
  <el-form @submit.prevent="search">
    <el-form-item>
      <div class="d-flex gap-3 w-100 flex-wrap">
        <div class="flex-1" style="min-width: 300px">
          <v-input v-model="queryInfo.keywords" :placeholder="t('placeholder.enterIpHostNameOrLabel')">
            <template #prepend>
              <span>{{ t("label.keywords") }}</span>
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
    <el-table-column :label="t('label.ip')" fixed="left" min-width="160" prop="ipAddress" sortable="custom">
      <template #default="scope">
        <div class="d-flex gap-2">
          <v-node-enable-tag :enabled="scope.row.isEnable" />
          <v-router-link :text="scope.row.ipAddress" :type="scope.row.isEnable ? 'primary' : 'info'" href="" />
        </div>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.hostname')" min-width="200" prop="name" sortable="custom">
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
        <el-button link type="danger" @click="openAction(Action.Delete, scope.row)">
          {{ t("btn.remove") }}
        </el-button>
      </template>
    </el-table-column>
  </v-table>

  <node-add ref="addRef" @refresh="search()" />

  <node-remove ref="deleteRef" @refresh="search()" />
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
