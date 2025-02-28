<script lang="ts" setup>
import { FormInstance, FormRules } from "element-plus"
import { RulePleaseEnter, SetWebTitle } from "@/utils"
import { PageGroupDetail, RuleLength } from "@/models"
import { filter, map, uniq, uniqWith } from "lodash-es"
import { groupService } from "services/group-service.ts"

interface ServiceValidDeploymentInfo extends ServiceDeploymentInfo {
  hasPlacementConstraints: boolean
  validPlacements: Array<{ id: string; mode: string; key: string; value: string; isEqual: boolean }>
}

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const isLoading = ref(false)
const isAction = ref(false)

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = ref<ServiceInfo>(NewServiceEmptyInfo())
const groupNodes = ref<NodeInfo[]>([])

const deploymentInfo = ref<ServiceValidDeploymentInfo>({
  hasPlacementConstraints: false,
  validPlacements: [],
  ...NewServiceDeploymentInfo()
})

const formRef = useTemplateRef<FormInstance>("formRef")
const rules = ref<FormRules>({
  serviceName: [
    { required: true, validator: RulePleaseEnter("label.name"), trigger: "blur" },
    { required: true, validator: RuleLimitRange(RuleLength.ServiceName.Min, RuleLength.ServiceName.Max), trigger: "blur" }
  ],
  description: [{ validator: RuleLimitMax(RuleLength.Description.Max), trigger: "blur" }]
})

const labelList = computed(() => {
  const result: Array<{ key: string; value: string }> = []
  map(groupNodes.value, x => {
    result.push(...map(Object.keys(x.labels), l => ({ key: l, value: x.labels[l] })))
  })
  return uniqWith(result, (a, b) => a.key === b.key && a.value === b.value)
})

function cancel() {
  router.push({ name: "groupDetail", params: { groupId: groupId.value, mode: PageGroupDetail.Services } })
}

function replicatedNumChange(v: number | undefined) {
  deploymentInfo.value.replicas = v || 1
}

function changePlacementMode(index: number) {
  deploymentInfo.value.validPlacements[index].key = ""
  deploymentInfo.value.validPlacements[index].value = ""
}

function addPlacementConstraint() {
  deploymentInfo.value.validPlacements.push({
    id: GenerateUUID(),
    mode: ServicePlacementMode.PlacementModeIP,
    key: "",
    value: "",
    isEqual: true
  })
}

function removePlacementConstraint(index: number) {
  deploymentInfo.value.validPlacements.splice(index, 1)
}

async function getGroupNodes() {
  return await groupService.getNodes(groupId.value).then(nodes => {
    groupNodes.value = nodes
  })
}

async function getGroupInfo() {
  return await groupService.info(groupId.value).then(info => {
    stateStore.setGroup(groupId.value, info)
  })
}

async function getServiceInfo() {
  return await serviceService.info(groupId.value, serviceId.value).then(info => {
    serviceInfo.value = info
    stateStore.setService(serviceId.value, info)
    const tempInfo = info.deployment || NewServiceDeploymentInfo()
    deploymentInfo.value = {
      hasPlacementConstraints: tempInfo.placements.length > 0,
      validPlacements: map(tempInfo.placements, x => Object.assign({ id: GenerateUUID() }, x)),
      ...tempInfo
    }
  })
}

async function search(init?: boolean) {
  isLoading.value = true
  await Promise.all([init ? getGroupNodes() : undefined, getGroupInfo(), getServiceInfo()]).finally(() => (isLoading.value = false))
}

async function save() {}

onMounted(async () => {
  await search(true)
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)
})
</script>

<template>
  <el-form ref="formRef" v-loading="isLoading" :model="deploymentInfo" :rules="rules" class="form-box" label-position="top" label-width="auto">
    <el-form-item :label="t('label.dispatchMode')" prop="mode">
      <div class="d-flex gap-5">
        <el-radio-group v-model="deploymentInfo.mode">
          <el-radio :value="ServiceDeployMode.DeployModeGlobal">{{ t("label.global") }}</el-radio>
          <el-radio :value="ServiceDeployMode.DeployModeReplicate">{{ t("label.replicated") }}</el-radio>
        </el-radio-group>
        <div class="flex-1 d-flex instances-box">
          <div class="instances-prefix">
            <el-text>{{ t("label.instanceNum") }}</el-text>
          </div>
          <v-input-number :max="100" :min="1" :model-value="deploymentInfo.replicas" @update:model-value="replicatedNumChange" />
        </div>
      </div>
    </el-form-item>
    <el-form-item>
      <v-tips v-if="deploymentInfo.mode === ServiceDeployMode.DeployModeGlobal">{{ t("tips.globalTips") }}</v-tips>
      <v-tips v-if="deploymentInfo.mode === ServiceDeployMode.DeployModeReplicate">{{ t("tips.replicatedTips") }} </v-tips>
    </el-form-item>
    <el-form-item>
      <el-checkbox v-model="deploymentInfo.hasPlacementConstraints">
        <strong>
          <el-text size="small">{{ t("label.placementConstraints") }}</el-text>
        </strong>
      </el-checkbox>
    </el-form-item>
    <div v-if="deploymentInfo.hasPlacementConstraints" class="placement-constraints">
      <div v-for="(item, index) in deploymentInfo.validPlacements" :key="item.id" class="d-flex gap-3 flex-wrap" style="align-items: start">
        <el-form-item style="margin: 0">
          <el-radio-group v-model="deploymentInfo.validPlacements[index].mode" @change="changePlacementMode(index)">
            <el-radio :label="t('label.ip')" :value="ServicePlacementMode.PlacementModeIP" />
            <el-radio :label="t('label.label')" :value="ServicePlacementMode.PlacementModeLabel" />
          </el-radio-group>
        </el-form-item>

        <div class="d-flex gap-3 flex-1 ml-5" style="max-width: 800px; min-width: 500px">
          <el-form-item :prop="`${index}.key`" class="flex-1">
            <v-input v-if="deploymentInfo.validPlacements[index].mode === ServicePlacementMode.PlacementModeIP" disabled>
              <template #prefix>{{ t("label.ip") }}</template>
            </v-input>
            <v-select v-else v-model="deploymentInfo.validPlacements[index].key" :out-label="t('label.label')" placeholder="" show-out-label>
              <template v-if="deploymentInfo.validPlacements[index].mode === ServicePlacementMode.PlacementModeLabel">
                <el-option
                  v-for="item in uniq(
                    filter(labelList, x => !deploymentInfo.validPlacements[index].value || x.value === deploymentInfo.validPlacements[index].value)
                  )"
                  :key="item.key"
                  :label="item.key"
                  :value="item.key" />
              </template>
              <template v-else>
                <el-option label="ip" value="ip" />
              </template>
            </v-select>
          </el-form-item>

          <el-form-item>
            <el-select v-model="deploymentInfo.validPlacements[index].isEqual" style="width: 140px">
              <el-option :value="true" label="=" />
              <el-option :value="false" label="!=" />
            </el-select>
          </el-form-item>

          <el-form-item :prop="`${index}.value`" class="flex-1">
            <v-select v-model="deploymentInfo.validPlacements[index].value" :out-label="t('label.value')" placeholder="" show-out-label>
              <template v-if="deploymentInfo.validPlacements[index].mode === ServicePlacementMode.PlacementModeLabel">
                <el-option
                  v-for="item in uniq(
                    filter(labelList, x => !deploymentInfo.validPlacements[index].key || x.key === deploymentInfo.validPlacements[index].key)
                  )"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value" />
              </template>
              <template v-else>
                <el-option v-for="item in groupNodes" :key="item.ipAddress" :label="item.ipAddress" :value="item.ipAddress" />
              </template>
            </v-select>
          </el-form-item>

          <el-form-item>
            <el-button plain style="padding: 4px 12px" text type="danger" @click="removePlacementConstraint(index)">
              <el-icon :size="26">
                <IconMdiClose />
              </el-icon>
            </el-button>
          </el-form-item>
        </div>
      </div>

      <el-form-item style="margin: 0">
        <el-button size="small" type="info" @click="addPlacementConstraint">
          <template #icon>
            <el-icon :size="20">
              <IconMdiAdd />
            </el-icon>
          </template>
          {{ t("btn.addPlacementConstraint") }}
        </el-button>
      </el-form-item>
    </div>
  </el-form>
  <div class="text-align-right pt-3">
    <el-button @click="cancel()">{{ t("btn.cancel") }}</el-button>
    <el-button :loading="isAction" type="primary" @click="save">{{ t("btn.save") }}</el-button>
  </div>
</template>

<style lang="scss" scoped>
.form-box {
  :deep(.el-form-item__label) {
    font-weight: 600;
    font-size: 12px;
    margin-bottom: 4px;
  }
}

.instances-box {
  margin-left: 16px;

  .instances-prefix {
    border: 1px solid var(--el-border-color);
    line-height: 30px;
    padding: 0 8px;
    border-bottom-left-radius: 4px;
    border-top-left-radius: 4px;
    border-right: none;
    background-color: var(--el-fill-color-light);
  }

  :deep(.el-input__wrapper) {
    border-bottom-left-radius: 0;
    border-top-left-radius: 0;
  }
}

.placement-constraints {
  border: 1px solid var(--el-border-color);
  padding: 16px;
  border-radius: 4px;
  box-sizing: border-box;
  background-color: #ecf0f5;
}
</style>
