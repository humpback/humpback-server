<script lang="ts" setup>
import LogoSVG from "@/assets/logo.png"

const props = defineProps<{
  isHorizontal?: boolean
  showText?: boolean
  enableClick?: boolean
}>()

const router = useRouter()
const pageStore = usePageStore()

const logoName = "Humpback"
const logoImg = LogoSVG

function clickLogo() {
  if (props.enableClick) {
    router.push({ name: "manage" })
  }
}
</script>

<template>
  <div v-if="props.isHorizontal" :class="{ 'logo-click': props.enableClick, 'logo-box': true }" @click="clickLogo()">
    <el-avatar :size="32" :src="logoImg" class="color-light" fit="fill" shape="square" />
    <span v-if="!pageStore.menuIsCollapse || props.showText" class="gradient-text"> {{ logoName }} </span>
  </div>
  <div v-else class="w-100 text-align-center mb-2">
    <el-avatar :size="80" :src="logoImg" class="color-transparent" fit="fill" shape="square" />
    <h3 class="gradient-text m-none"> {{ logoName }} </h3>
  </div>
</template>

<style lang="scss" scoped>
.logo-box {
  height: 100%;
  width: 100%;
  display: flex;
  gap: 12px;
  align-items: center;
  padding: 0 0 0 12px;
  justify-content: left;
  box-sizing: border-box;

  .color-light {
    background-color: #ffffff;
  }
}

.logo-click:hover {
  cursor: pointer;
}

.el-avatar {
  border-radius: 50%;
}

.color-transparent {
  background-color: rgba(0, 0, 0, 0);
}

.gradient-text {
  font-size: 20px;
  font-weight: bold;
  background-image: linear-gradient(45deg, #ffc700, #ff3d00);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
}
</style>
