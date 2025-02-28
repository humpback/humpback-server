<script lang="ts" setup>
import { ServiceVolumeType } from "@/models"
import { FormInstance, FormRules } from "element-plus"
import { GenerateUUID, RulePleaseEnter } from "@/utils"
import { filter } from "lodash-es"

const props = defineProps<{ hasValid?: boolean }>()
const emits = defineEmits<{
  (e: "check"): void
}>()

const { t } = useI18n()

const volumes = defineModel<
  Array<{
    id: string
    type: ServiceVolumeType.VolumeTypeBind | ServiceVolumeType.VolumeTypeVolume
    target: string
    source: string
    "readonly": boolean
  }>
>()

const formRef = useTemplateRef<FormInstance>("formRef")
const rules = ref<FormRules>({
  hostVolume: [{ required: true, validator: RulePleaseEnter("label.hostVolume"), trigger: "blur" }],
  containerVolume: [
    { required: true, validator: RulePleaseEnter("label.containerVolume"), trigger: "blur" },
    { required: true, validator: checkContainerVolume, trigger: "blur" }
  ]
})

function checkContainerVolume(rule: any, value: any, callback: any) {
  const path = value as string
  if (filter(volumes.value, x => x.target === path).length > 1) {
    return callback(new Error(`${t("rules.duplicate")} ${t("label.containerVolume")}`))
  }
  return callback()
}

function addVolume() {
  volumes.value!.push({
    id: GenerateUUID(),
    type: ServiceVolumeType.VolumeTypeBind,
    target: "",
    source: "",
    readonly: false
  })
}

function removeVolume(index: number) {
  volumes.value!.splice(index, 1)
  emits("check")
}

async function validate() {
  return await formRef.value?.validate()
}

onMounted(() => {
  if (props.hasValid) {
    validate()
  }
})

defineExpose({ validate })
</script>

<template>
  <div>
    <div class="mb-3">
      <v-tips>{{ t("tips.volumeTips") }}</v-tips>
    </div>
    <el-form ref="formRef" :model="volumes" :rules="rules" label-position="top" label-width="auto">
      <div v-for="(volume, index) in volumes" :key="volume.id" class="d-flex gap-3">
        <el-form-item>
          <el-radio-group v-model="volumes![index].readonly" class="volume-type" fill="var(--el-color-info-light-3)" text-color="#ffffff">
            <el-radio :label="t('label.writable')" :value="false" />
            <el-radio :label="t('label.readonly')" :value="true" />
          </el-radio-group>
        </el-form-item>
        <el-form-item :prop="`${index}.target`" :rules="rules.containerVolume" class="flex-1">
          <v-input v-model="volumes![index].target" :placeholder="t('placeholder.containerVolume')" @blur="emits('check')">
            <template #prepend>{{ t("label.containerVolume") }}</template>
          </v-input>
        </el-form-item>
        <el-form-item :prop="`${index}.source`" :rules="rules.hostVolume" class="flex-1">
          <v-input v-model="volumes![index].source" :placeholder="t('placeholder.hostVolume')" @blur="emits('check')">
            <template #prepend>{{ t("label.hostVolume") }}</template>
          </v-input>
        </el-form-item>
        <el-form-item>
          <el-button plain style="padding: 4px 12px" text type="danger" @click="removeVolume(index)">
            <el-icon :size="26">
              <IconMdiClose />
            </el-icon>
          </el-button>
        </el-form-item>
      </div>
    </el-form>
    <el-button size="small" type="info" @click="addVolume()">
      <template #icon>
        <el-icon :size="20">
          <IconMdiAdd />
        </el-icon>
      </template>
      {{ t("btn.addVolume") }}
    </el-button>
  </div>
</template>

<style lang="scss" scoped></style>
