<script lang="ts" setup>
import { cloneDeep } from "lodash-es"

const props = defineProps<{ params?: any }>()
const { t } = useI18n()
const route = useRoute()
const stateStore = useStateStore()

const breadcrumbs = computed(() => {
  const list = cloneDeep(route.meta.breadcrumb) || []
  for (const b of list) {
    if (!b.isLink) {
      switch (b.customName) {
        case "service":
          b.name = stateStore.getService()?.serviceName
          break
        case "group":
          b.name = stateStore.getGroup()?.groupName
          break
      }
    } else {
      switch (b.customName) {
        case "service":
          b.routeParams = Object.assign({}, b.routeParams, { serviceId: stateStore.getService()?.serviceId }, props.params)
          b.name = stateStore.getService()?.serviceName
          break
        case "group":
          b.routeParams = Object.assign({}, b.routeParams, { groupId: stateStore.getGroup()?.groupId }, props.params)
          b.name = stateStore.getGroup()?.groupName
          break
      }
    }
  }
  return list
})
</script>

<template>
  <el-breadcrumb v-if="breadcrumbs.length > 0" separator="/">
    <el-breadcrumb-item v-for="(item, index) in breadcrumbs" :key="index" :to="item.isLink ? { name: item.routeName, params: item.routeParams } : undefined">
      {{ item.name || t(item.i18nLabel) }}
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<style lang="scss" scoped></style>
