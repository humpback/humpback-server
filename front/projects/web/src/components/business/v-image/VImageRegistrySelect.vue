<script lang="ts" setup>
import { RegistryInfo } from "@/types"

const props = withDefaults(
  defineProps<{
    image: string
    size?: "large" | "default" | "small"
    clearable?: boolean
    placeholder?: string
    options?: RegistryInfo[]
  }>(),
  { size: "default", clearable: true, placeholder: "" }
)

const emits = defineEmits<{
  (e: "change", v: string): void
}>()

const { t } = useI18n()
const registry = defineModel<string>()
const registries = ref<RegistryInfo[]>([])

const options = computed(() => props.options || registries.value)

async function getRegistryList() {
  return await registryService.list().then(list => {
    registries.value = list
  })
}

function change() {
  emits("change", registry.value!)
}

onMounted(async () => {
  if (!props.options) {
    await getRegistryList()
  }
})
</script>

<template>
  <v-select v-model="registry" :clearable="props.clearable" :loading-text="t('message.loading')" :placeholder="props.placeholder" @change="change()">
    <el-option v-for="item in options" :key="item.registryId" :label="item.registryName" :value="item.registryId" />
  </v-select>
</template>

<style lang="scss" scoped></style>
