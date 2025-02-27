<script lang="ts" setup>
import { PageGroupDetail, RuleLength } from "@/models"
import { FormInstance, FormRules } from "element-plus"
import { RulePleaseEnter, SetWebTitle } from "@/utils"

const { t } = useI18n()
const route = useRoute()
const router = useRouter()
const stateStore = useStateStore()

const isLoading = ref(false)
const isAction = ref(false)

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = ref<ServiceInfo>(NewServiceEmptyInfo())

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

async function getGroupInfo() {
  return await groupService.info(groupId.value).then(info => {
    stateStore.setGroup(groupId.value, info)
  })
}

async function getServiceInfo() {
  return await serviceService.info(groupId.value, serviceId.value).then(info => {
    serviceInfo.value = info
    stateStore.setService(serviceId.value, info)
  })
}

async function search() {
  isLoading.value = true
  await Promise.all([getGroupInfo(), getServiceInfo()]).finally(() => (isLoading.value = false))
}

async function save() {
  if (!(await formRef.value?.validate())) {
    return
  }

  // isAction.value = true
  // serviceService
  //   .create({
  //     serviceName: dialogInfo.value.info.serviceName,
  //     description: dialogInfo.value.info.description,
  //     groupId: groupId.value
  //   })
  //   .then(() => {
  //     ShowSuccessMsg(t("message.addSuccess"))
  //     dialogInfo.value.show = false
  //     emits("refresh")
  //   })
  //   .finally(() => (isAction.value = false))
}

onMounted(async () => {
  await search()
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)
})
</script>

<template>
  <el-form ref="formRef" v-loading="isLoading" :model="serviceInfo" :rules="rules" class="form-box" label-position="top" label-width="auto">
    <el-form-item :label="t('label.name')" prop="serviceName">
      <v-input :maxlength="RuleLength.ServiceName.Max" :model-value="serviceInfo.serviceName" clearable disabled show-word-limit />
    </el-form-item>
    <el-form-item :label="t('label.description')" prop="description">
      <v-description-input v-model="serviceInfo.description" />
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
</style>
