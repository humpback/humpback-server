<script lang="ts" setup>
import PageHeader from "./page-header.vue"
import PageAside from "./page-aside.vue"

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
      <router-view v-slot="{ Component, route }">
        <Transition mode="out-in" name="drawer">
          <component :is="Component" :key="route.name" />
        </Transition>
      </router-view>
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
    padding: calc(var(--hp-header-height) + 8px) 10px 0 10px;
    max-width: 100%;
  }
}

.drawer-enter-active {
  transition:
    transform 0.3s ease,
    opacity 0.3s ease 0.2s;
}

.drawer-leave-active {
  transition: opacity 0.3s ease;
}

.drawer-enter-from {
  transform: translateX(-100%);
  opacity: 0;
}

.drawer-enter-to {
  transform: translateX(0);
  opacity: 1;
}

.drawer-leave-to {
  opacity: 0;
}
</style>
