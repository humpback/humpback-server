<script lang="ts" setup>
import { FormInstance, FormRules } from "element-plus"
import { SetWebTitle } from "@/utils"
import { PageGroupDetail, RuleLength, ServiceDeployType } from "@/models"
import { filter, map, omit, toLower, uniq, uniqWith } from "lodash-es"
import { groupService } from "services/group-service.ts"
import VCronInput from "@/components/business/v-corn/VCornInput.vue"
import cronstrue from "cronstrue"

interface ServiceValidDeploymentInfo extends ServiceDeploymentInfo {
  hasPlacements: boolean
  validPlacements: Array<{ id: string; mode: string; key: string; value: string; isEqual: boolean }>
  enableTimeout: boolean
  enableSchedules: boolean
}

const { t, locale } = useI18n()
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
  hasPlacements: false,
  validPlacements: [],
  enableTimeout: false,
  enableSchedules: false,
  ...NewServiceDeploymentInfo()
})

const formRef = useTemplateRef<FormInstance>("formRef")
const cronInputRef = useTemplateRef<InstanceType<typeof VCronInput>>("cronInputRef")
const rules = ref<FormRules>({
  timeout: [{ validator: checkTimeout, trigger: "blur" }],
  placementKey: [{ required: true, validator: RuleCannotBeEmpty, trigger: "change" }],
  placementValue: [{ required: true, validator: RuleCannotBeEmpty, trigger: "change" }]
})

const labelList = computed(() => {
  const result: Array<{ key: string; value: string }> = []
  map(groupNodes.value, x => {
    result.push(...map(Object.keys(x.labels), l => ({ key: l, value: x.labels[l] })))
  })
  return uniqWith(result, (a, b) => a.key === b.key && a.value === b.value)
})

function checkTimeout(rule: any, value: any, callback: any) {
  const v = value ? (value as string) : ""
  if (deploymentInfo.value.enableSchedules && deploymentInfo.value.enableTimeout && v && !/^-?\d+(\.\d+)?(ns|us|µs|ms|s|m|h)+$/.test(v)) {
    return callback(new Error(t("rules.formatErrTimeout")))
  }
  callback()
}

function cancel() {
  router.push({ name: "groupDetail", params: { groupId: groupId.value, mode: PageGroupDetail.Services } })
}

function replicatedNumChange(v: number | undefined) {
  deploymentInfo.value.replicas = v || 1
}

function changePlacementMode(index: number) {
  if (deploymentInfo.value.validPlacements[index].mode === ServicePlacementMode.PlacementModeIP) {
    deploymentInfo.value.validPlacements[index].key = "IP"
  } else {
    deploymentInfo.value.validPlacements[index].key = ""
  }
  deploymentInfo.value.validPlacements[index].value = ""
}

function addPlacementConstraint() {
  deploymentInfo.value.validPlacements.push({
    id: GenerateUUID(),
    mode: ServicePlacementMode.PlacementModeIP,
    key: "IP",
    value: "",
    isEqual: true
  })
}

function removePlacementConstraint(index: number) {
  deploymentInfo.value.validPlacements.splice(index, 1)
}

function editSchedule(index: number) {
  if (index !== -1) {
    cronInputRef.value?.openDialog(deploymentInfo.value.schedule.rules[index], index)
  } else {
    cronInputRef.value?.openDialog("", index)
  }
}

function changeSchedule(corn: string, index: number) {
  if (index !== -1) {
    deploymentInfo.value.schedule.rules[index] = corn
  } else {
    deploymentInfo.value.schedule.rules.push(corn)
  }
}

function removeSchedule(index: number) {
  deploymentInfo.value.schedule.rules.splice(index, 1)
}

function parseCronToText(corn: string) {
  try {
    return cronstrue.toString(corn, {
      use24HourTimeFormat: true,
      throwExceptionOnParseError: true,
      verbose: false,
      locale: toLower(locale.value) === "zh-cn" ? "zh_CN" : "en"
    })
  } catch (error) {
    console.error(error)
    return t("tips.invalidCorn")
  }
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
      hasPlacements: tempInfo.placements.length > 0,
      enableSchedules: tempInfo.schedule.rules.length > 0,
      enableTimeout: !!tempInfo.schedule.timeout,
      validPlacements: map(tempInfo.placements, x => Object.assign({ id: GenerateUUID() }, x)),
      ...tempInfo
    }
  })
}

async function search(init?: boolean) {
  isLoading.value = true
  await Promise.all([init ? getGroupNodes() : undefined, getGroupInfo(), getServiceInfo()]).finally(() => (isLoading.value = false))
}

async function save() {
  if (!(await formRef.value?.validate())) {
    return
  }

  const body = {
    serviceId: serviceId.value,
    type: "deployment",
    data: {
      type: deploymentInfo.value.schedule.rules.length > 0 ? ServiceDeployType.DeployTypeSchedule : ServiceDeployType.DeployTypeBackground,
      mode: deploymentInfo.value.mode,
      replicas: deploymentInfo.value.mode === ServiceDeployMode.DeployModeGlobal ? 1 : deploymentInfo.value.replicas,
      placements:
        deploymentInfo.value.hasPlacements && deploymentInfo.value.validPlacements.length > 0
          ? map(deploymentInfo.value.validPlacements, x => omit(x, ["id"]))
          : [],
      schedule: {
        timeout:
          deploymentInfo.value.enableSchedules && deploymentInfo.value.enableTimeout && deploymentInfo.value.schedule.rules.length > 0
            ? deploymentInfo.value.schedule.timeout
            : "",
        rules: deploymentInfo.value.enableTimeout ? deploymentInfo.value.schedule.rules : []
      }
    }
  }
  isAction.value = true
  await serviceService.update(groupId.value, body).finally(() => (isAction.value = false))
  ShowSuccessMsg(t("message.saveSuccess"))
  await search()
}

onMounted(async () => {
  await search(true)
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)
})
</script>

<template>
  <el-form ref="formRef" v-loading="isLoading" :model="deploymentInfo" :rules="rules" class="form-box" label-position="top" label-width="auto">
    <el-form-item :label="t('label.dispatchMode')" prop="mode">
      <div class="d-flex gap-5 mt-3">
        <el-radio-group v-model="deploymentInfo.mode">
          <el-radio :value="ServiceDeployMode.DeployModeGlobal">{{ t("label.global") }}</el-radio>
          <el-radio :value="ServiceDeployMode.DeployModeReplicate">{{ t("label.replicated") }}</el-radio>
        </el-radio-group>
        <div v-if="deploymentInfo.mode === ServiceDeployMode.DeployModeReplicate" class="flex-1 d-flex instances-box">
          <div class="instances-prefix">
            <el-text>{{ t("label.instanceNum") }}</el-text>
          </div>
          <v-input-number
            :max="RuleLength.InstanceNum.Max"
            :min="RuleLength.InstanceNum.Min"
            :model-value="deploymentInfo.replicas"
            @update:model-value="replicatedNumChange" />
        </div>
      </div>
    </el-form-item>
    <el-form-item>
      <v-tips v-if="deploymentInfo.mode === ServiceDeployMode.DeployModeGlobal">{{ t("tips.globalTips") }}</v-tips>
      <v-tips v-if="deploymentInfo.mode === ServiceDeployMode.DeployModeReplicate">{{ t("tips.replicatedTips") }}</v-tips>
    </el-form-item>

    <el-form-item class="mt-3">
      <el-checkbox v-model="deploymentInfo.hasPlacements">
        <strong>
          <el-text size="small">{{ t("label.placementConstraints") }}</el-text>
        </strong>
      </el-checkbox>
    </el-form-item>

    <div v-if="deploymentInfo.hasPlacements" class="content-box">
      <div v-for="(item, index) in deploymentInfo.validPlacements" :key="item.id" class="d-flex gap-3 flex-wrap" style="align-items: start">
        <el-form-item style="margin: 0">
          <el-radio-group v-model="deploymentInfo.validPlacements[index].mode" @change="changePlacementMode(index)">
            <el-radio :label="t('label.ip')" :value="ServicePlacementMode.PlacementModeIP" />
            <el-radio :label="t('label.label')" :value="ServicePlacementMode.PlacementModeLabel" />
          </el-radio-group>
        </el-form-item>

        <div class="d-flex gap-3 flex-1 ml-5" style="max-width: 800px; min-width: 500px">
          <el-form-item :prop="`validPlacements.${index}.key`" :rules="rules.placementKey" class="flex-1">
            <v-input
              v-if="deploymentInfo.validPlacements[index].mode === ServicePlacementMode.PlacementModeIP"
              v-model="deploymentInfo.validPlacements[index].key"
              disabled />

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

          <el-form-item :prop="`validPlacements.${index}.value`" :rules="rules.placementValue" class="flex-1">
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

    <el-form-item :label="t('label.schedulesInfo')" class="mt-5">
      <el-checkbox v-model="deploymentInfo.enableSchedules" class="mt-3">
        <strong>
          <el-text size="small">{{ t("label.setSchedules") }}</el-text>
        </strong>
      </el-checkbox>
    </el-form-item>

    <div v-if="deploymentInfo.enableSchedules" class="content-box">
      <div class="mb-3">
        <v-tips>{{ t("tips.scheduleTips") }}</v-tips>
      </div>
      <div v-for="(corn, index) in deploymentInfo.schedule.rules" :key="index" class="mb-3 cron-line">
        <div class="d-flex gap-3">
          <strong>
            <el-text> {{ t("label.cron") }}</el-text>
          </strong>
          <el-text>{{ corn }}</el-text>
        </div>
        <el-divider direction="vertical" />
        <div class="flex-1 cron-text">
          <el-text> {{ parseCronToText(corn) }}</el-text>
        </div>
        <div style="width: 100px">
          <el-button link plain type="primary" @click="editSchedule(index)">
            <el-icon :size="20">
              <IconMdiSquareEditOutline />
            </el-icon>
          </el-button>
          <el-button link plain type="danger" @click="removeSchedule(index)">
            <el-icon :size="20">
              <IconMdiClose />
            </el-icon>
          </el-button>
        </div>
      </div>

      <el-button size="small" type="info" @click="editSchedule(-1)">
        <template #icon>
          <el-icon :size="20">
            <IconMdiAdd />
          </el-icon>
        </template>
        {{ t("btn.addSchedule") }}
      </el-button>
    </div>

    <el-form-item v-if="deploymentInfo.enableSchedules" class="mt-3">
      <el-checkbox v-model="deploymentInfo.enableTimeout">
        <strong>
          <el-text size="small">{{ t("label.enableTimeout") }}</el-text>
        </strong>
      </el-checkbox>
    </el-form-item>
    <el-form-item v-if="deploymentInfo.enableTimeout" :rules="rules.timeout" prop="schedule.timeout">
      <v-input v-model="deploymentInfo.schedule.timeout" :placeholder="t('placeholder.timeoutExample')" />
    </el-form-item>
  </el-form>

  <div class="text-align-right pt-3">
    <el-button @click="cancel()">{{ t("btn.cancel") }}</el-button>
    <el-button :loading="isAction" type="primary" @click="save">{{ t("btn.save") }}</el-button>
  </div>

  <v-corn-input ref="cronInputRef" @change="changeSchedule" />
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

.content-box {
  //border: 1px solid var(--el-border-color);
  padding: 16px;
  border-radius: 4px;
  box-sizing: border-box;
  background-color: #ecf0f5;
}

.cron-line {
  padding: 4px 12px;
  border-radius: 4px;
  display: flex;
  align-items: center;

  &:hover {
    background-color: #d6dde7;
  }

  .cron-text {
    word-break: break-word;
  }
}
</style>
