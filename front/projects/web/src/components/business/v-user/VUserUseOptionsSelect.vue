<script lang="ts" setup>
import { UserInfo } from "@/types"
import { map } from "lodash-es"

const props = withDefaults(defineProps<{ options: UserInfo[]; filterable?: boolean; multiple?: boolean; placeholder?: string }>(), {
  filterable: true,
  multiple: true,
  placeholder: ""
})

const { t } = useI18n()
const teams = defineModel<string[]>()
const selectOptions = computed(() => map(props.options, x => ({ label: x.username, value: x.userId })))
</script>

<template>
  <el-select-v2 v-model="teams" :filterable="props.filterable" :multiple="props.multiple" :options="selectOptions" :placeholder="props.placeholder">
    <template #footer>
      <el-link href="/ws/user-related/users" target="_blank" type="primary">
        <strong>{{ t("btn.goToAddUser") }}</strong>
      </el-link>
    </template>
  </el-select-v2>
</template>

<style lang="scss" scoped></style>
