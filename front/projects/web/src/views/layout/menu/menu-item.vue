<script lang="ts" setup>
import { menuI18nPrefix, MenuInfo } from "./types"

const props = defineProps<{ menuInfo: MenuInfo }>()
const { t } = useI18n()
</script>

<template>
  <el-menu-item
    :index="props.menuInfo?.routeName || props.menuInfo.name"
    :route="{ name: props.menuInfo?.routeName || props.menuInfo.name, params: props.menuInfo.params }">
    <template #title>
      {{ t(`${menuI18nPrefix}.${props.menuInfo.name}`) }}
    </template>
    <el-icon v-if="props.menuInfo.icon" :size="20">
      <component :is="props.menuInfo.icon" />
    </el-icon>
  </el-menu-item>
</template>

<style lang="scss" scoped></style>
