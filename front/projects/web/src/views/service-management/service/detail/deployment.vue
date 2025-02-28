<script lang="ts" setup>
import { FormInstance, FormRules } from "element-plus"
import { RulePleaseEnter, SetWebTitle } from "@/utils"
import { PageGroupDetail, RuleLength } from "@/models"

interface ServiceValidDeploymentInfo extends ServiceDeploymentInfo {
  hasPlacementConstraints: boolean
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

const deploymentInfo = ref<ServiceValidDeploymentInfo>({ hasPlacementConstraints: false, ...NewServiceDeploymentInfo() })

const formRef = useTemplateRef<FormInstance>("formRef")
const rules = ref<FormRules>({
  serviceName: [
    { required: true, validator: RulePleaseEnter("label.name"), trigger: "blur" },
    { required: true, validator: RuleLimitRange(RuleLength.ServiceName.Min, RuleLength.ServiceName.Max), trigger: "blur" }
  ],
  description: [{ validator: RuleLimitMax(RuleLength.Description.Max), trigger: "blur" }]
})

function cancel() {
  router.push({ name: "groupDetail", params: { groupId: groupId.value, mode: PageGroupDetail.Services } })
}

function replicatedNumChange(v: number | undefined) {
  deploymentInfo.value.replicas = v || 1
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
    deploymentInfo.value = { hasPlacementConstraints: tempInfo.placements.length > 0, ...tempInfo }
  })
}

async function search() {
  isLoading.value = true
  await Promise.all([getGroupInfo(), getServiceInfo()]).finally(() => (isLoading.value = false))
}

async function save() {}

onMounted(async () => {
  await search()
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
      <v-tips v-if="deploymentInfo.mode === ServiceDeployMode.DeployModeReplicate">{{ t("tips.replicatedTips") }}</v-tips>
    </el-form-item>
    <el-form-item>
      <el-checkbox v-model="deploymentInfo.hasPlacementConstraints">
        <strong>
          <el-text size="small">{{ t("label.placementConstraints") }}</el-text>
        </strong>
      </el-checkbox>
    </el-form-item>
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
</style>
