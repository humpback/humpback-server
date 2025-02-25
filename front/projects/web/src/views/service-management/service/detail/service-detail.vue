<script lang="ts" setup>
const { t } = useI18n()
const route = useRoute()
const isNew = computed(() => !!route.meta.isNew)
</script>

<template>
  <h1> Service Info </h1>
</template>

<style lang="scss" scoped></style>
