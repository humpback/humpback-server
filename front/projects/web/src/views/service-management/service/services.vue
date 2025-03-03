<script lang="ts" setup>
import { ServiceInfo } from "@/types"
import { SetWebTitle, TableHeight } from "@/utils"
import { Action } from "@/models"
import { QueryServicesInfo } from "./common.ts"
import { serviceService } from "services/service-client.ts"
import ServiceCreate from "./action/service-create.vue"
import ServiceDelete from "./action/service-delete.vue"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const tableHeight = computed(() => TableHeight(252))

const groupId = ref(route.params?.groupId as string)

const isLoading = ref(false)
const queryInfo = ref<QueryServicesInfo>(new QueryServicesInfo(route.query))

const createRef = useTemplateRef<InstanceType<typeof ServiceCreate>>("createRef")
const deleteRef = useTemplateRef<InstanceType<typeof ServiceDelete>>("deleteRef")

const tableList = ref({
  total: 0,
  data: [] as Array<ServiceInfo>
})

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
    :total="tableList.total"
    @page-change="search"
    @sort-change="search">
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
    <el-table-column :label="t('label.status')" min-width="200" prop="description">
      <template #default="scope">
        <v-table-column-none :text="scope.row.status" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.image')" min-width="200" prop="description">
      <template #default="scope">
        <v-table-column-none :text="scope.row.meta?.image" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.deployMode')" min-width="200" prop="description">
      <template #default="scope">
        <v-table-column-none :text="scope.row.deployment?.type" />
      </template>
    </el-table-column>
    <el-table-column :label="t('label.ports')" min-width="200" prop="description">
      <template #default="scope">
        <v-table-column-none :text="scope.row.description" />
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

  <service-create ref="createRef" @refresh="search()" />
</template>

<style lang="scss" scoped></style>
