<script lang="ts" setup>
import { SetWebTitle } from "@/utils"
import { DynamicIntervalTimer } from "./instances.ts"

const { t } = useI18n()
const route = useRoute()
const stateStore = useStateStore()

const isLoading = ref(false)

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = ref<ServiceInfo>(NewServiceEmptyInfo())

const timer = ref<DynamicIntervalTimer>(new DynamicIntervalTimer())

async function startIncreasingTimer(initialDelay: number, increment: number) {
  let delay = initialDelay

  async function executeTask() {
    setTimeout(async () => {
      await executeTask()
    }, delay)
    delay += increment
  }

  await executeTask()
}

// 启动定时器：初始间隔 3 秒，每次增加 3 秒
startIncreasingTimer(3000, 3000)

async function getGroupInfo() {
  return await groupService.info(groupId.value).then(info => {
    stateStore.setGroup(groupId.value, info)
  })
}

async function getServiceInfo() {
  return await serviceService.info(groupId.value, serviceId.value).then(info => {
    serviceInfo.value = info
    stateStore.setService(serviceId.value, info)
    serviceInfo.value.containers = [
      {
        containerId: "FSJLFJLFJPQJT03QWGHQOGLANG;AHNGAHNG;KLANG",
        containerName: "humpback-losjfljsofwoj-lsfjlsj-jfslf",
        nodeId: "SLFJOPQW2JRFO",
        status: "running",
        ip: "10.16.15.12",
        statusInfo: "",
        errorMsg: "exit code -1",
        image: "docker.io/nginx:latest",
        command: "nginx -g daemon off",
        network: "",
        createAt: 1740818013735,
        startAt: 1740818013735,
        nextAt: 1740818013735,
        lastHeartbeat: 1740818013735,
        labels: { test: "skyler" },
        env: ["test=true", "amd=yesLJOQGOJQPGJOQP;JGO;EQJGPOQJGEJGNQ;LEJQG;LJGL;EWJGL;NA;GLNMA;LGMNPOQJGQP;GJ;QLJGL;ENMQG;LNMG;GQNMMGOQPHGJPQGOQHNGOQHNGO;"],
        mounts: [
          {
            source: "/var/lib/docker",
            destination: "/var/lib/docker"
          }
        ],
        ports: [
          {
            bindIP: "0.0.0.0",
            privatePort: 80,
            publicPort: 800,
            type: "TCP"
          }
        ]
      }
    ]
  })
}

async function search() {
  isLoading.value = true
  await Promise.all([getGroupInfo(), getServiceInfo()]).finally(() => (isLoading.value = false))
}

onMounted(async () => {
  await search()
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)
  timer.value.start(search)
})

onUnmounted(() => {
  timer.value.stop()
})
</script>

<template>
  <div class="d-flex gap-3">
    <strong>
      <el-text>{{ t("label.instanceOverview") }}</el-text>
    </strong>
    <div>
      <el-button plain size="small" type="success" @click="search()">{{ t("btn.refresh") }}</el-button>
      <el-button plain size="small" type="primary">{{ t("btn.viewMonitor") }}</el-button>
    </div>
  </div>
  <v-table :data="serviceInfo.containers" border class="mt-5" row-key="containerId">
    <el-table-column type="expand">
      <template #default="scope">
        <div class="expand-content">
          <div v-if="scope.row.errorMsg" class="mb-5 d-flex gap-1">
            <el-icon :size="16" color="var(--el-color-danger)">
              <IconMdiWarningCircleOutline />
            </el-icon>
            <el-text type="danger">
              {{ scope.row.errorMsg }}
            </el-text>
          </div>
          <el-form label-position="top" label-width="auto">
            <el-row :gutter="12">
              <el-col :span="12">
                <el-form-item>
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiClockTimeFourOutline />
                      </el-icon>
                      {{ t("label.createTime") }}
                    </el-text>
                  </template>
                  <v-date-view :format="-1" :timestamp="scope.row.createAt" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item>
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiClockTimeFourOutline />
                      </el-icon>
                      {{ t("label.startTime") }}
                    </el-text>
                  </template>
                  <v-date-view :format="-1" :timestamp="scope.row.startAt" />
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item>
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiClockTimeThreeOutline />
                      </el-icon>
                      {{ t("label.nextTime") }}
                    </el-text>
                  </template>
                  <v-date-view :format="-1" :timestamp="scope.row.nextAt" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item>
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiAppleKeyboardCommand />
                      </el-icon>
                      {{ t("label.command") }}
                    </el-text>
                  </template>
                  <span>{{ scope.row.command || "--" }}</span>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item :label="t('label.image')">
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiAlphaCBoxOutline />
                      </el-icon>
                      {{ t("label.image") }}
                    </el-text>
                  </template>
                  <span>{{ scope.row.image || "--" }}</span>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item :label="t('label.volumes')">
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiFileOutline />
                      </el-icon>
                      {{ t("label.volumes") }}
                    </el-text>
                  </template>
                  <div v-if="scope.row.mounts?.length > 0">
                    <div v-for="(item, index) in scope.row.mounts" :key="index" class="form-line">
                      <div class="line-prefix">-</div>
                      <div>
                        {{ `${item.source}:${item.destination}` }}
                      </div>
                    </div>
                  </div>
                  <span v-else>--</span>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item :label="t('label.ports')">
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiAlphaPCircleOutline />
                      </el-icon>
                      {{ t("label.ports") }}
                    </el-text>
                  </template>
                  <div v-if="scope.row.ports?.length > 0">
                    <div v-for="(item, index) in scope.row.ports" :key="index" class="form-line">
                      <div class="line-prefix">-</div>
                      <div>
                        {{ item.type }}
                        <el-text type="primary">{{ `${item.privatePort}:${item.publicPort}` }}</el-text>
                      </div>
                    </div>
                  </div>
                  <span v-else>--</span>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item :label="t('label.environments')">
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiMapMarkerPath />
                      </el-icon>
                      {{ t("label.environments") }}
                    </el-text>
                  </template>
                  <div v-if="scope.row.env?.length > 0">
                    <div v-for="(item, index) in scope.row.env" :key="index" class="form-line">
                      <div class="line-prefix">-</div>
                      <div> {{ item }}</div>
                    </div>
                  </div>
                  <span v-else>--</span>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item :label="t('label.labels')">
                  <template #label>
                    <el-text size="small" type="success">
                      <el-icon :size="14">
                        <IconMdiTagTextOutline />
                      </el-icon>
                      {{ t("label.labels") }}
                    </el-text>
                  </template>
                  <div v-if="Object.keys(scope.row.labels)?.length > 0">
                    <div v-for="(key, index) in Object.keys(scope.row.labels)" :key="index" class="form-line">
                      <div class="line-prefix">-</div>
                      <div> {{ `${key}:${scope.row.labels[key]}` }}</div>
                    </div>
                  </div>
                  <span v-else>--</span>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </div>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.name')" min-width="300" prop="containerName" sortable />
    <el-table-column :label="t('label.ip')" min-width="160">
      <template #default="scope">
        <el-text type="primary">{{ scope.row.ip }}</el-text>
      </template>
    </el-table-column>
    <el-table-column :label="t('label.status')" min-width="160">
      <template #default="scope">
        <div class="d-flex gap-3">
          <v-container-status :status="scope.row.status" size="small" />
          <v-tooltip v-if="scope.row.errorMsg">
            <template #content>
              <el-text type="danger">{{ scope.row.errorMsg }}</el-text>
            </template>
            <el-icon :size="20" color="var(--el-color-danger)">
              <IconMdiWarningCircleOutline />
            </el-icon>
          </v-tooltip>
        </div>
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
</template>

<style lang="scss" scoped>
.expand-content {
  box-sizing: border-box;
  padding: 20px 60px;

  :deep(.el-form-item__label) {
    margin-bottom: 4px;
  }

  :deep(.el-form-item__content) {
    padding-left: 12px;
    font-size: 12px;
    line-height: 20px;
    word-break: break-all;
  }
}

.form-line {
  display: flex;
  align-items: start;
  gap: 4px;

  .line-prefix {
    color: var(--el-color-success);
  }
}
</style>
