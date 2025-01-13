<script lang="ts" setup>
import { find } from "lodash-es"

const props = defineProps<{ role: number }>()
const { t } = useI18n()

const str = computed(() =>
  find(
    [
      { label: "role.user", value: UserRole.User },
      { label: "role.admin", value: UserRole.Admin },
      { label: "role.supperAdmin", value: UserRole.SupperAdmin }
    ],
    x => x.value === props.role
  )
)
</script>

<template>{{ str ? t(str.label) : "--" }}</template>

<style lang="scss" scoped></style>
