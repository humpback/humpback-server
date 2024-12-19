<script lang="ts" setup>
import SubMenu from "./sub-menu.vue"
import MenuItem from "./menu-item.vue"
import { MenuInfo } from "./types"

const props = defineProps<{ menuInfo: MenuInfo }>()

const hasChild = computed(() => Array.isArray(props.menuInfo.children) && props.menuInfo.children.length > 0)

const menuComponent = computed(() => {
  return hasChild.value ? SubMenu : MenuItem
})
</script>

<template>
  <component :is="menuComponent" :menuInfo="menuInfo">
    <template v-if="hasChild">
      <menu-container v-for="item in menuInfo.children" :key="item.name" :menuInfo="item" />
    </template>
  </component>
</template>

<style lang="scss" scoped></style>
