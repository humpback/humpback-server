<script lang="ts" setup>
import { menuI18nPrefix, MenuInfo } from "./types"

const props = defineProps<{ menuInfo: MenuInfo }>()
const { t } = useI18n()
const router = useRouter()

function navigateToRoute(event: MouseEvent, href: string) {
  if (event.ctrlKey) {
    window.open(href, "_blank")
  } else {
    router.push(href)
  }
}
</script>

<template>
  <router-link v-slot="{ href }" custom to="">
    <div @click="navigateToRoute($event, href)">
      <el-menu-item :index="props.menuInfo.name" :route="{ name: props.menuInfo.name }">
        <template #title>
          {{ t(`${menuI18nPrefix}.${props.menuInfo.name}`) }}
        </template>
        <el-icon v-if="props.menuInfo.icon" :size="20">
          <component :is="props.menuInfo.icon" />
        </el-icon>
      </el-menu-item>
    </div>
  </router-link>
</template>

<style lang="scss" scoped></style>
