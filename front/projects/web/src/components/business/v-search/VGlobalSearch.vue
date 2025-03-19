<script lang="ts" setup>
import { PageServiceDetail } from "@/models"
import { onClickOutside } from "@vueuse/core"
import { ElTooltip, ElInput } from "element-plus"
import { trim } from "lodash-es"

const { t } = useI18n()
const name = ref("")
const isLoading = ref(false)
const isSearch = ref(false)
const isFocus = ref(false)

const visible = ref(false)

const contentRef = useTemplateRef<HTMLDivElement>("contentRef")
const inputRef = useTemplateRef<InstanceType<typeof ElInput>>("inputRef")
const tooltipRef = useTemplateRef<InstanceType<typeof ElTooltip>>("tooltipRef")

const result = ref<{
  [key: string]: Array<{ groupId: string; groupName: string; serviceId?: string; serviceName?: string }>
}>({ groups: [], services: [] })

const searchGroupService = CreateCancelRequest(commonService.searchGroupServiceByName)

function focusInput() {
  isFocus.value = true
  if (name.value != "" && isSearch.value) {
    tooltipRef.value?.onOpen()
  }
}

function clickLink(isHref?: boolean) {
  if (!isHref) {
    hidePopover()
  }
}

function hidePopover(event?: any) {
  if (!event || event?.target?.id !== "search-input") {
    visible.value = false
    tooltipRef.value?.onClose()
    inputRef.value?.blur()
    isFocus.value = false
  }
}

function closePopover() {
  isFocus.value = false
}

function search() {
  name.value = trim(name.value)
  if (name.value === "") {
    return
  }
  tooltipRef.value?.onOpen()
  // visible.value = true
  isLoading.value = true
  searchGroupService(name.value)
    .then(data => {
      result.value = data
      isSearch.value = true
    })
    .finally(() => (isLoading.value = false))
}

onClickOutside(contentRef, hidePopover)
</script>

<template>
  <div :class="['global-search', isFocus && 'is-focused']">
    <el-tooltip ref="tooltipRef" :hide-after="0" :visible="visible" effect="light" placement="bottom-start" @close="closePopover()">
      <!--      <template #reference>-->
      <el-input
        id="search-input"
        ref="inputRef"
        v-model="name"
        :placeholder="t('placeholder.searchGroupService')"
        size="small"
        @focus="focusInput()"
        @keydown.enter="search()">
        <template #prefix>
          <el-icon :size="16">
            <IconMdiSearch />
          </el-icon>
        </template>
      </el-input>
      <!--      </template>-->
      <template #content>
        <div ref="contentRef" v-loading="isLoading" class="global-search-content">
          <div>
            <strong>
              <el-text> {{ t("label.groups") }}</el-text>
            </strong>
            <el-divider style="margin: 8px 0 16px 0; border-color: var(--el-color-info-light-9)" />
          </div>
          <div v-if="result.groups && result.groups.length > 0" class="pl-5">
            <div v-for="(item, index) in result.groups" :key="index" class="content">
              -
              <v-router-link :href="`/ws/group/${item.groupId}/services`" :text="item.groupName" @click-route="clickLink" />
            </div>
          </div>
          <div v-else class="pl-5 d-flex gap-1">
            <el-text type="warning">
              <el-icon :size="15">
                <IconMdiWarningCircleOutline />
              </el-icon>
            </el-text>
            <el-text size="small" type="warning"> {{ t("tips.noGroupFound") }}</el-text>
          </div>

          <div class="mt-5">
            <strong>
              <el-text> {{ t("label.services") }}</el-text>
            </strong>
            <el-divider style="margin: 8px 0 16px 0; border-color: var(--el-color-info-light-9)" />
            <div v-if="result.services && result.services.length > 0" class="pl-5">
              <div v-for="(item, index) in result.services" :key="index" class="content">
                <el-text size="small" type="info">{{ item.groupName }}</el-text>
                <br />
                -
                <v-router-link
                  :href="`/ws/group/${item.groupId}/service/${item.serviceId}/${PageServiceDetail.BasicInfo}`"
                  :text="item.serviceName!"
                  @click-route="clickLink" />
              </div>
            </div>
            <div v-else class="pl-5 d-flex gap-1">
              <el-text type="warning">
                <el-icon :size="15">
                  <IconMdiWarningCircleOutline />
                </el-icon>
              </el-text>
              <el-text size="small" type="warning">{{ t("tips.noServiceFound") }}</el-text>
            </div>
          </div>
        </div>
      </template>
    </el-tooltip>
  </div>
</template>

<style lang="scss" scoped>
.global-search {
  :deep(.el-input) {
    .el-input__wrapper {
      border-radius: 16px;
    }
  }

  width: 200px;
  transition: width 0.3s ease;

  &.is-focused {
    width: 400px;
  }
}
</style>

<style lang="scss">
.global-search-content {
  margin: -12px;
  padding: 16px;
  width: 380px;

  .content {
    text-wrap: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    font-size: 14px;
  }
}
</style>
