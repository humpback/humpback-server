<script lang="ts" setup>
const props = defineProps<{ keywords?: string; showAddBtn?: boolean; keywordsPlaceholder?: string }>()
const emits = defineEmits<{
  (e: "update:keywords", keywords: string): void
  (e: "search"): void
  (e: "add"): void
}>()

const { t } = useI18n()
const slots = useSlots()

function handleKeywordsChange(v: string) {
  emits("update:keywords", v)
}

function search() {
  emits("search")
}

function add() {
  emits("add")
}
</script>

<template>
  <el-form @submit.prevent="search">
    <el-form-item class="search">
      <div class="d-flex w-100 flex-wrap gap-3">
        <slot v-if="!!slots.prefix" name="prefix" />
        <slot v-if="!!slots.input" name="input" />
        <v-input v-else :model-value="props.keywords" :placeholder="props.keywordsPlaceholder" class="search-input" @update:model-value="handleKeywordsChange">
          <template #prepend>
            <slot v-if="!!slots.keywordsPrepend" name="keywordsPrepend"></slot>
            <span v-else style="color: var(--el-text-color-regular)">
              {{ t("label.keywords") }}
            </span>
          </template>
        </v-input>
        <el-button native-type="submit" type="primary">
          <el-icon :size="16">
            <IconMdiSearch />
          </el-icon>
          {{ t("btn.search") }}
        </el-button>
        <slot v-if="!!slots.append" name="append" />
        <el-button v-if="props.showAddBtn" type="primary" @click="add()">
          <el-icon :size="18">
            <IconMdiAdd />
          </el-icon>
          {{ t("btn.add") }}
        </el-button>
      </div>
    </el-form-item>
  </el-form>
</template>

<style lang="scss" scoped>
.search {
  .search-input {
    flex: 1;
    min-width: 200px;
  }

  .el-button {
    margin: 0;
  }
}
</style>
