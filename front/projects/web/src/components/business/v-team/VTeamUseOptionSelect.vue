<script lang="ts" setup>
import { TeamInfo } from "@/types"
import { map } from "lodash-es"

const props = withDefaults(defineProps<{ options: TeamInfo[]; filterable?: boolean; multiple?: boolean; placeholder?: string }>(), {
  filterable: true,
  multiple: true,
  placeholder: ""
})

const { t } = useI18n()
const teams = defineModel<string[]>()
const selectOptions = computed(() => map(props.options, x => ({ label: x.name, value: x.teamId })))
</script>

<template>
  <el-select-v2 v-model="teams" :filterable="props.filterable" :multiple="props.multiple" :options="selectOptions" :placeholder="props.placeholder">
    <template #footer>
      <el-link href="/ws/user-related/team" target="_blank" type="primary">
        <strong>{{ t("btn.goToAddTeam") }}</strong>
      </el-link>
    </template>
  </el-select-v2>
</template>

<style lang="scss" scoped></style>
