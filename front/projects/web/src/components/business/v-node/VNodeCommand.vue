<script lang="ts" setup>
import { CopyToClipboard } from "@/utils"

const props = defineProps<{ command: string }>()

const { t } = useI18n()

function copy() {
  CopyToClipboard(props.command)
}
</script>

<template>
  <div class="command-box">
    <div class="command-body">
      <div class="copy-btn">
        <v-tooltip :content="t('btn.copy')" :teleported="false" placement="top-end">
          <el-button link type="primary" @click="copy()">
            <el-icon :size="16">
              <IconMdiContentCopy />
            </el-icon>
          </el-button>
        </v-tooltip>
      </div>
      <div class="command-content">
        <pre>{{ props.command }}</pre>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.command-box {
  width: 100%;
  border-radius: 4px;
  padding: 4px;
  background-color: #f5f5f5;
  box-sizing: border-box;

  .command-body {
    border-radius: 4px;
    background-color: #ffffff;
    border: 1px dashed #adadad;

    .copy-btn {
      text-align: left;
      padding: 4px 0 0 4px;
    }

    .command-content {
      padding: 0 16px;
      min-height: 50px;
      font-size: 14px;
      font-weight: 500;
      line-height: 20px;
    }
  }
}
</style>
