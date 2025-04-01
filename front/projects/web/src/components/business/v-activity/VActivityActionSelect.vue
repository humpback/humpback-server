<script lang="ts" setup>
import { filter, find } from "lodash-es"

const props = defineProps<{ mode: "group" | "service" | "config" | "registry" | "node" | "user" | "team"; showOutLabel?: boolean }>()

const { t } = useI18n()

const action = defineModel<string>()

const options = computed<Array<{ i18nLabel: string; value: string; modeList: string[] }>>(() => {
  const actions = [
    { i18nLabel: "label.add", value: "add", modeList: ["group", "service", "registry", "config", "node", "user", "team"] },
    { i18nLabel: "label.update", value: "update", modeList: ["group", "registry", "config", "user", "team"] },
    { i18nLabel: "label.delete", value: "delete", modeList: ["group", "service", "registry", "config", "node", "user", "team"] },
    { i18nLabel: "label.updateLabel", value: "updateLabel", modeList: ["node"] },
    { i18nLabel: "label.addNode", value: "addNode", modeList: ["group"] },
    { i18nLabel: "label.removeNode", value: "removeNode", modeList: ["group"] },
    { i18nLabel: "label.updateBasic", value: "updateBasic", modeList: ["service"] },
    { i18nLabel: "label.updateApplication", value: "updateApplication", modeList: ["service"] },
    { i18nLabel: "label.updateDeployment", value: "updateDeployment", modeList: ["service"] },
    { i18nLabel: "label.enable", value: "enable", modeList: ["service", "node"] },
    { i18nLabel: "label.disable", value: "disable", modeList: ["service", "node"] },
    { i18nLabel: "label.start", value: "start", modeList: ["service"] },
    { i18nLabel: "label.restart", value: "restart", modeList: ["service"] },
    { i18nLabel: "label.stop", value: "stop", modeList: ["service"] }
  ]
  return filter(actions, x => !!find(x.modeList, m => m === props.mode))
})
</script>

<template>
  <v-select
    v-model="action"
    :out-label="props.showOutLabel ? t('label.action') : ''"
    :placeholder="t('placeholder.all')"
    clearable
    filterable
    out-label-width="80px">
    <el-option v-for="item in options" :key="item.value" :label="t(item.i18nLabel)" :value="item.value" />
  </v-select>
</template>

<style lang="scss" scoped></style>
