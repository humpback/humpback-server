<script lang="ts" setup>
import PageHeader from "./page-header.vue"
import PageAside from "./page-aside.vue"

const route = useRoute()
const pageStore = usePageStore()

const leftWidth = computed(() => (pageStore.menuIsCollapse ? "var(--hp-aside-collapse-width)" : "var(--hp-aside-width)"))
</script>

<template>
  <div id="page-container" :style="{ paddingLeft: leftWidth }">
    <div id="page-header" :style="{ paddingLeft: leftWidth }">
      <page-header />
    </div>
    <div id="page-aside" :style="{ width: leftWidth }">
      <page-aside />
    </div>
    <div id="page-main">
      <router-view :key="route.name as string" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
#page-container {
  box-sizing: border-box;
  width: 100%;
  min-width: 320px;
  height: 100%;
  min-height: 500px;

  #page-header {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000;
    width: 100%;
    min-width: 320px;
    height: var(--hp-header-height);
    box-sizing: border-box;
    border-bottom: 1px solid var(--el-border-color);
    //box-shadow: 0 1px 1px var(--el-border-color);
    background-color: var(--hp-header-bg-color);
  }

  #page-aside {
    position: fixed;
    z-index: 1001;
    left: 0;
    top: 0;
    height: 100%;
    background-color: var(--hp-aside-bg-color);
  }

  #page-main {
    box-sizing: border-box;
    padding: calc(var(--hp-header-height) + 12px) 12px 0 12px;
    max-width: 100%;
  }
}
</style>
