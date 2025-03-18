<script lang="ts" setup>
import { cloneDeep } from "lodash-es"

const props = defineProps<{ params?: any }>()
const { t } = useI18n()
const route = useRoute()
const stateStore = useStateStore()

const breadcrumbs = computed(() => {
  const list = cloneDeep(route.meta.breadcrumb) || []
  for (const b of list) {
    switch (b.customName) {
      case "service":
        b.routeParams = b.isLink ? Object.assign({}, b.routeParams, { serviceId: stateStore.getService()?.serviceId }, props.params) : b.routeParams
        b.name = stateStore.getService()?.serviceName
        break
      case "group":
        b.routeParams = b.isLink ? Object.assign({}, b.routeParams, { groupId: stateStore.getGroup()?.groupId }, props.params) : b.routeParams
        b.name = stateStore.getGroup()?.groupName
        break
    }
  }
  return list
})
</script>

<template>
  <el-breadcrumb v-if="breadcrumbs.length > 0" separator="/">
    <el-breadcrumb-item v-for="(item, index) in breadcrumbs" :key="index">
      <router-link v-if="item.isLink" v-slot="{ href }" :to="{ name: item?.routeName, params: item?.routeParams }" custom>
        <a :href="href">{{ item?.name ? item.name : item?.i18nLabel ? t(item?.i18nLabel) : "" }}</a>
      </router-link>
      <span v-else>{{ item?.name ? item.name : item?.i18nLabel ? t(item?.i18nLabel) : "" }}</span>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<style lang="scss" scoped></style>
