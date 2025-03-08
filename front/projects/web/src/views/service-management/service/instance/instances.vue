<script lang="ts" setup>
import { SetWebTitle } from "@/utils"
import { refreshData } from "@/views/service-management/service/common.ts"

const { t } = useI18n()
const route = useRoute()
const stateStore = useStateStore()

const groupId = ref(route.params.groupId as string)
const serviceId = ref(route.params.serviceId as string)
const serviceInfo = computed<ServiceInfo | undefined>(() => stateStore.getService(serviceId.value))

const timer = ref<any>(null)
const interval = ref(5000)

async function resetLoopSearch() {
  if (timer.value) {
    clearInterval(timer.value)
    timer.value = null
  }
  interval.value = 5000
  loopSearch()
}

async function search() {
  await refreshData(groupId.value, serviceId.value, "instances")
}

function loopSearch() {
  timer.value = setTimeout(async () => {
    await search()
    if (serviceInfo.value?.status === ServiceStatus.ServiceStatusRunning) {
      interval.value = 1000
    }
    if (serviceInfo.value?.isEnabled) {
      loopSearch()
    }
  }, interval.value)
}

onMounted(async () => {
  await search()
  SetWebTitle(`${t("webTitle.serviceInfo")} - ${stateStore.getService()?.serviceName}`)
  loopSearch()
})

onUnmounted(() => {
  if (timer.value) {
    clearTimeout(timer.value)
    timer.value = null
  }
})

defineExpose({ resetLoopSearch })
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
  <v-table :data="serviceInfo?.containers || []" border class="mt-5" row-key="containerId">
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
              <el-col v-if="scope.row.nextAt" :span="24">
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
                  <div v-if="Object.keys(scope.row.labels || {})?.length > 0">
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
    margin-bottom: 12px;
  }

  :deep(.el-form-item__content) {
    padding-left: 12px;
    font-size: 14px;
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
