<script lang="ts" setup>
import { GenerateUUID, RulePleaseEnter, SetWebTitle } from "@/utils"
import { PageGroupDetail, RuleLength, ServiceNetworkMode, ServiceNetworkProtocol } from "@/models"
import { FormInstance, FormRules } from "element-plus"
import { NewServiceMetaDockerEmptyInfo } from "@/types"
import { find, map } from "lodash-es"

interface ServiceApplicationInfo extends ServiceMetaDockerInfo {
  imageDomain: string
  imageName: string
  ports: Array<{
    id: string
    hostPort: number
    containerPort: number
    protocol: string
  }>
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
const registries = ref<RegistryInfo[]>([])

const metaInfo = ref<ServiceApplicationInfo>({ imageDomain: "", imageName: "", ports: [], ...NewServiceMetaDockerEmptyInfo() })

const formRef = useTemplateRef<FormInstance>("formRef")
const rules = ref<FormRules>({
  imageName: [
    { required: true, validator: RulePleaseEnter("label.image"), trigger: "blur" },
    { required: true, validator: RuleLimitRange(RuleLength.ServiceName.Min, RuleLength.ServiceName.Max), trigger: "blur" }
  ]
})

function cancel() {
  router.push({ name: "groupDetail", params: { groupId: groupId.value, mode: PageGroupDetail.Services } })
}

function parseMetaInfo() {
  const defaultImage = find(registries.value, x => x.isDefault)
  const domain = defaultImage ? defaultImage.url : registries.value.length > 0 ? registries.value[0].url : ""
  let meta = serviceInfo.value.meta ? serviceInfo.value.meta : NewServiceMetaDockerEmptyInfo()
  const imageSplit = meta.image.indexOf("/")
  metaInfo.value = {
    imageDomain: imageSplit > 0 ? meta.image.slice(0, imageSplit) : domain,
    imageName: imageSplit > 0 ? meta.image.slice(imageSplit) : "",
    ports: map(meta.network.ports, x => ({ id: GenerateUUID(), ...x })),
    ...meta
  }
}

function addPort() {
  metaInfo.value.ports.push({ id: GenerateUUID(), containerPort: 0, protocol: ServiceNetworkProtocol.NetworkProtocolTCP, hostPort: 0 })
}

function removePort(index: number) {
  metaInfo.value.ports.splice(index, 1)
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

async function getRegistryList() {
  return await registryService.list().then(list => {
    registries.value = list
  })
}

async function search(init?: boolean) {
  isLoading.value = true
  await Promise.all([getGroupInfo(), getServiceInfo(), init ? getRegistryList() : undefined])
    .then(() => {
      parseMetaInfo()
    })
    .finally(() => (isLoading.value = false))
}

async function save() {}

onMounted(async () => {
  await search(true)
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)
})
</script>

<template>
  <el-form ref="formRef" v-loading="isLoading" :model="serviceInfo" :rules="rules" class="form-box" label-position="top" label-width="auto">
    <el-row :gutter="12">
      <el-col :span="24">
        <el-form-item :label="t('label.image')" prop="imageName">
          <v-input v-model="metaInfo.imageName" :maxlength="RuleLength.ImageName?.Max" clearable show-word-limit>
            <template #prepend>
              <el-select v-model="metaInfo.imageDomain" placeholder="" style="width: auto; min-width: 200px">
                <el-option v-for="item in registries" :key="item.registryId" :label="item.url" :value="item.url" />
              </el-select>
            </template>
          </v-input>
        </el-form-item>
      </el-col>
      <el-col :span="24">
        <el-form-item :label="t('label.command')">
          <v-description-input v-model="metaInfo.command" />
        </el-form-item>
      </el-col>

      <el-col :span="24">
        <el-form-item :label="t('label.network')">
          <el-select v-model="metaInfo.network.mode">
            <el-option :value="ServiceNetworkMode.NetworkModeHost" label="Host" />
            <el-option :value="ServiceNetworkMode.NetworkModeBridge" label="Bridge" />
            <el-option :value="ServiceNetworkMode.NetworkModeCustom" label="Custom" />
          </el-select>
        </el-form-item>
      </el-col>

      <el-col v-if="metaInfo.network.mode !== ServiceNetworkMode.NetworkModeHost">
        <div class="network-box">
          <el-row :gutter="12">
            <el-col v-if="metaInfo.network.mode === ServiceNetworkMode.NetworkModeCustom" :span="12">
              <el-form-item :label="t('label.networkName')" prop="network.networkName">
                <v-input v-model="metaInfo.network.networkName" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item :label="t('label.hostname')" prop="network.hostname">
                <v-input v-model="metaInfo.network.hostname" />
              </el-form-item>
            </el-col>
            <el-col :span="12" />
            <el-col v-for="(portInfo, index) in metaInfo.ports" :key="index" :span="24">
              <div class="d-flex gap-2">
                <el-form-item :prop="`ports.${index}.containerPort`">
                  <v-input-number v-model="metaInfo.ports[index].containerPort" :controls="false" :min="0" :placeholder="t('placeholder.containerPort')" />
                </el-form-item>
                <el-form-item :prop="`ports.${index}.protocol`">
                  <el-select v-model="metaInfo.ports[index].protocol" :placeholder="t('placeholder.protocol')" style="width: 100px">
                    <el-option :value="ServiceNetworkProtocol.NetworkProtocolTCP" label="TCP" />
                    <el-option :value="ServiceNetworkProtocol.NetworkProtocolUDP" label="UDP" />
                  </el-select>
                </el-form-item>
                <el-form-item :prop="`ports.${index}.hostPort`">
                  <div class="d-flex gap-1">
                    <v-input-number v-model="metaInfo.ports[index].hostPort" :controls="false" :placeholder="t('placeholder.hostPort')" />
                    <el-button plain style="padding: 4px 12px" text type="danger" @click="removePort(index)">
                      <el-icon :size="26">
                        <IconMdiClose />
                      </el-icon>
                    </el-button>
                  </div>
                </el-form-item>
              </div>
            </el-col>
            <el-col>
              <el-button size="small" type="info" @click="addPort">
                <template #icon>
                  <el-icon :size="20">
                    <IconMdiAdd />
                  </el-icon>
                </template>
                {{ t("btn.addPort") }}
              </el-button>
            </el-col>
          </el-row>
        </div>
      </el-col>
    </el-row>
  </el-form>
  <div class="text-align-right pt-3">
    <el-button @click="cancel()">{{ t("btn.cancel") }}</el-button>
    <el-button :loading="isAction" type="primary" @click="save">{{ t("btn.save") }}</el-button>
  </div>
</template>

<style lang="scss" scoped>
.form-box {
  :deep(.el-form-item__label) {
    //font-weight: 500;
    //font-size: 12px;
    //margin-bottom: 2px;
  }
}

.network-box {
  border: 1px solid var(--el-border-color);
  padding: 16px;
  border-radius: 4px;
  box-sizing: border-box;
  background-color: #ecf0f5;
}
</style>
