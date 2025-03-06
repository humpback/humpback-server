<script lang="ts" setup>
import { ServiceInfo } from "@/types"
import { SetWebTitle, TableHeight } from "@/utils"
import { Action } from "@/models"
import { QueryServicesInfo } from "./common.ts"
import { serviceService } from "services/service-client.ts"
import ServiceCreate from "./action/service-create.vue"
import ServiceDelete from "./action/service-delete.vue"
import VServiceStatusTag from "@/components/business/v-service/VServiceStatusTag.vue"
import { capitalize } from "lodash-es"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const tableHeight = computed(() => TableHeight(352))

const groupId = ref(route.params?.groupId as string)

const isLoading = ref(false)
const queryInfo = ref<QueryServicesInfo>(new QueryServicesInfo(route.query))

const createRef = useTemplateRef<InstanceType<typeof ServiceCreate>>("createRef")
const deleteRef = useTemplateRef<InstanceType<typeof ServiceDelete>>("deleteRef")

const tableList = ref({
  total: 0,
  data: [] as Array<ServiceInfo>
})

const setRowClass = ({ row }) => {
  return row.containers.length === 0 ? "hide-expand-icon" : ""
}

async function getGroupInfo() {
  return await groupService.info(groupId.value).then(info => {
    stateStore.setGroup(groupId.value, info)
  })
}

async function getServiceTotal() {
  return await serviceService.total(groupId.value).then(total => {
    stateStore.setGroupTotal(groupId.value, total)
  })
}

async function getServices() {
  return await serviceService.query(groupId.value, queryInfo.value.searchParams()).then(info => {
    tableList.value.data = info.list
    tableList.value.total = info.total
  })
}

async function search() {
  await router.replace(queryInfo.value.urlQuery())
  isLoading.value = true
  await Promise.all([getGroupInfo(), getServiceTotal(), getServices()]).finally(() => (isLoading.value = false))
}

function openAction(action: string, info?: ServiceInfo) {
  switch (action) {
    case Action.Add: {
      createRef.value?.open()
      break
    }
    case Action.Delete: {
      deleteRef.value?.open(info!)
    }
  }
}

onMounted(async () => {
  await search()
  SetWebTitle(`${t("webTitle.services")} - ${stateStore.getGroup()?.groupName}`)
})
</script>

<template>
  <el-form @submit.prevent="search">
    <el-form-item>
      <div class="d-flex gap-3 w-100 flex-wrap">
        <div style="width: 220px">
          <v-service-status-query-select v-model="queryInfo.filter.status" :placeholder="t('placeholder.all')" @change="search" />
        </div>
        <div style="width: 200px">
          <v-service-schedule-query-select v-model="queryInfo.filter.schedule" :placeholder="t('placeholder.all')" @change="search" />
        </div>
        <div class="flex-1" style="min-width: 300px">
          <v-input v-model="queryInfo.keywords">
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
            {{ t("btn.addService") }}
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
    :row-class-name="setRowClass"
    :total="tableList.total"
    @page-change="search"
    @sort-change="search">
    <el-table-column fixed="left" type="expand">
      <template #default="scope">
        <div class="pa-5">
          <v-table :data="scope.row.containers" :max-height="500" border headerCellClassName="">
            <el-table-column :label="t('label.instanceName')" min-width="200">
              <template #default="cscope">
                <el-text>{{ cscope.row.contianerName }}</el-text>
              </template>
            </el-table-column>
            <el-table-column :label="t('label.status')" min-width="160">
              <template #default="cscope">
                <div class="d-flex gap-3">
                  <v-container-status :status="cscope.row.status" size="small" />
                  <v-tooltip v-if="cscope.row.errorMsg">
                    <template #content>
                      <el-text type="danger">{{ cscope.row.errorMsg }}</el-text>
                    </template>
                    <el-icon :size="20" color="var(--el-color-danger)">
                      <IconMdiWarningCircleOutline />
                    </el-icon>
                  </v-tooltip>
                </div>
              </template>
            </el-table-column>
            <el-table-column :label="t('label.ip')" min-width="160">
              <template #default="cscope">
                <el-text type="primary">{{ cscope.row.ip }}</el-text>
              </template>
            </el-table-column>
            <el-table-column :label="t('label.createTime')" min-width="140">
              <template #default="scope">
                <v-date-view :timestamp="scope.row.createAt" />
              </template>
            </el-table-column>
            <el-table-column :label="t('label.startTime')" min-width="140">
              <template #default="scope">
                <v-date-view :timestamp="scope.row.startDate" />
              </template>
            </el-table-column>
            <el-table-column :label="t('label.action')" width="180">
              <template #default>
                <el-button :title="t('label.restart')" link type="success">
                  <el-icon :size="16">
                    <IconMdiRestart />
                  </el-icon>
                </el-button>
                <el-button :title="t('label.start')" link type="success">
                  <el-icon :size="16">
                    <IconMdiPlay />
                  </el-icon>
                </el-button>
                <el-button :title="t('label.stop')" link type="danger">
                  <el-icon :size="16">
                    <IconMdiSquare />
                  </el-icon>
                </el-button>
                <el-button :title="t('label.log')" link type="primary">
                  <el-icon :size="16">
                    <IconMdiNoteText />
                  </el-icon>
                </el-button>
              </template>
            </el-table-column>
          </v-table>
        </div>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.service')" fixed="left" min-width="200" prop="serviceName" sortable="custom">
      <template #default="scope">
        <v-router-link :href="`/ws/group/${groupId}/service/${scope.row.serviceId}/basic-info`" :text="scope.row.serviceName" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.description')" min-width="200" prop="description">
      <template #default="scope">
        <v-table-column-none :text="scope.row.description" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.status')" min-width="130" prop="description">
      <template #default="scope">
        <v-service-status-tag :is-enabled="scope.row.isEnabled" :status="scope.row.status" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.image')" min-width="200" prop="image">
      <template #default="scope">
        <v-table-column-none :text="scope.row.meta?.image" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.deployMode')" min-width="200">
      <template #default="scope">
        <div v-if="scope.row.deployment">
          <el-text>{{ capitalize(scope.row.deployment.mode) }}</el-text>
          <el-text v-if="scope.row.deployment.mode === ServiceDeployMode.DeployModeReplicate" type="primary">
            {{ ` (${scope.row.deployment.replicas})` }}
          </el-text>
        </div>
        <span v-else>--</span>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.updateDate')" min-width="140" prop="updatedAt" sortable="custom">
      <template #default="scope">
        <v-date-view :timestamp="scope.row.updatedAt" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.action')" align="right" fixed="right" header-align="center" width="130">
      <template #default="scope">
        <el-button link type="primary" @click="openAction(Action.Edit, scope.row)">{{ t("btn.action") }}</el-button>
        <el-button link type="danger" @click="openAction(Action.Delete, scope.row)">{{ t("btn.delete") }}</el-button>
      </template>
    </el-table-column>
  </v-table>

  <service-delete ref="deleteRef" @refresh="search()" />

  <service-create ref="createRef" />
</template>

<style lang="scss" scoped>
:deep(.hide-expand-icon) .el-table__expand-icon {
  display: none;
}
</style>
