<script lang="ts" setup>
import { find, findIndex } from "lodash-es"

const { t } = useI18n()
const capabilities = defineModel<string[]>()

const options = ref<Array<{ label: string; i18nTips: string }>>([
  { label: "AUDIT_CONTROL", i18nTips: "tips.AUDIT_CONTROL" },
  { label: "AUDIT_WRITE", i18nTips: "tips.AUDIT_WRITE" },
  { label: "BLOCK_SUSPEND", i18nTips: "tips.BLOCK_SUSPEND" },
  { label: "CHOWN", i18nTips: "tips.CHOWN" },
  { label: "DAC_OVERRIDE", i18nTips: "tips.DAC_OVERRIDE" },
  { label: "DAC_READ_SEARCH", i18nTips: "tips.DAC_READ_SEARCH" },
  { label: "FOWNER", i18nTips: "tips.FOWNER" },
  { label: "FSETID", i18nTips: "tips.FSETID" },
  { label: "IPC_LOCK", i18nTips: "tips.IPC_LOCK" },
  { label: "IPC_OWNER", i18nTips: "tips.IPC_OWNER" },
  { label: "KILL", i18nTips: "tips.KILL" },
  { label: "LEASE", i18nTips: "tips.LEASE" },
  { label: "LINUX_IMMUTABLE", i18nTips: "tips.LINUX_IMMUTABLE" },
  { label: "MAC_ADMIN", i18nTips: "tips.MAC_ADMIN" },
  { label: "MAC_OVERRIDE", i18nTips: "tips.MAC_OVERRIDE" },
  { label: "MKNOD", i18nTips: "tips.MKNOD" },
  { label: "NET_ADMIN", i18nTips: "tips.NET_ADMIN" },
  { label: "NET_BIND_SERVICE", i18nTips: "tips.NET_BIND_SERVICE" },
  { label: "NET_BROADCAST", i18nTips: "tips.NET_BROADCAST" },
  { label: "NET_RAW", i18nTips: "tips.NET_RAW" },
  { label: "SETFCAP", i18nTips: "tips.SETFCAP" },
  { label: "SETGID", i18nTips: "tips.SETGID" },
  { label: "SETPCAP", i18nTips: "tips.SETPCAP" },
  { label: "SETUID", i18nTips: "tips.SETUID" },
  { label: "SYSLOG", i18nTips: "tips.SYSLOG" },
  { label: "SYS_ADMIN", i18nTips: "tips.SYS_ADMIN" },
  { label: "SYS_BOOT", i18nTips: "tips.SYS_BOOT" },
  { label: "SYS_CHROOT", i18nTips: "tips.SYS_CHROOT" },
  { label: "SYS_MODULE", i18nTips: "tips.SYS_MODULE" },
  { label: "SYS_NICE", i18nTips: "tips.SYS_NICE" },
  { label: "SYS_PACCT", i18nTips: "tips.SYS_PACCT" },
  { label: "SYS_PTRACE", i18nTips: "tips.SYS_PTRACE" },
  { label: "SYS_RAWIO", i18nTips: "tips.SYS_RAWIO" },
  { label: "SYS_RESOURCE", i18nTips: "tips.SYS_RESOURCE" },
  { label: "SYS_TIME", i18nTips: "tips.SYS_TIME" },
  { label: "SYS_TTY_CONFIG", i18nTips: "tips.SYS_TTY_CONFIG" },
  { label: "WAKE_ALARM", i18nTips: "tips.WAKE_ALARM" }
])

function changeValue(label: string) {
  const index = findIndex(capabilities.value, x => x === label)
  if (index === -1) {
    capabilities.value?.push(label)
  } else {
    capabilities.value?.splice(index, 1)
  }
}

function getValue(label: string) {
  return !!find(capabilities.value, x => x === label)
}
</script>

<template>
  <el-row :gutter="12">
    <el-col v-for="item in options" :key="item.label" :span="8" class="mb-1">
      <div class="d-flex">
        <div class="d-flex gap-1" style="flex: 7">
          <strong>
            <el-text>{{ item.label }}</el-text>
          </strong>
          <v-tooltip :content="t(item.i18nTips)" effect="dark" max-width="300px" placement="top-start">
            <el-icon :size="16">
              <IconMdiHelpCircleOutline />
            </el-icon>
          </v-tooltip>
        </div>
        <div style="flex: 3">
          <el-switch :model-value="getValue(item.label)" @click="changeValue(item.label)" />
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<style lang="scss" scoped></style>
