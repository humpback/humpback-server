<script lang="ts" setup>
import { ChangeLanguage, GetCurrentLanguageName, GetLanguageOptions } from "@/locales"

const props = withDefaults(defineProps<{ showType?: "icon" | "text" | "all"; trigger?: "hover" | "click" | "contextmenu" }>(), {
  trigger: "click",
  showType: "all"
})

const languageOptions = computed(() => GetLanguageOptions())

const style = computed(() => {
  return props.showType === "text" ? {} : { "margin-left": "2px" }
})
</script>

<template>
  <el-dropdown :show-timeout="0" :trigger="props.trigger" @command="ChangeLanguage">
    <el-button link>
      <el-icon v-if="props.showType === 'icon' || props.showType === 'all'" :size="16">
        <IconMdiLanguage />
      </el-icon>
      <span v-if="props.showType === 'text' || props.showType === 'all'" :style="style">
        {{ GetCurrentLanguageName() }}
      </span>
    </el-button>
    <template #dropdown>
      <el-dropdown-menu>
        <template v-for="item in languageOptions" :key="item.value">
          <el-dropdown-item v-if="!item.disabled" :command="item.value">
            <div> {{ item.name }}</div>
          </el-dropdown-item>
        </template>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<style lang="scss" scoped></style>
