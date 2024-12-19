<script lang="ts" setup>
const props = defineProps<{ title?: string; defaultExpanded?: boolean; type?: "primary" | "success" | "warning" | "danger" | "info" }>()

const slots = useSlots()

const isExpanded = ref<boolean>(!!props.defaultExpanded)

function toggleCard() {
  isExpanded.value = !isExpanded.value
}
</script>

<template>
  <div class="card-container">
    <div class="card-header d-flex gap-1">
      <el-button :type="props.type" link @click="toggleCard">
        <el-icon :size="24">
          <IconMdiChevronDown :class="{ 'is-active': isExpanded }" class="arrow-icon" />
        </el-icon>
      </el-button>
      <div class="flex-1">
        <slot v-if="!!slots.title" name="title" />
        <h3 v-else-if="props.title">{{ props.title }}</h3>
      </div>
    </div>

    <v-expand>
      <div v-if="isExpanded" class="card-content">
        <slot name="default" />
      </div>
    </v-expand>
  </div>
</template>

<style lang="scss" scoped>
.card-container {
  border: 1px solid var(--el-border-color);
  overflow: hidden;
  border-radius: 4px;
  width: 100%;

  .card-header {
    background-color: var(--card-header-background-color);
    padding: 6px 10px;
    cursor: pointer;

    .arrow-icon {
      transition: transform ease 150ms;

      &.is-active {
        transform: rotateZ(-180deg);
      }
    }
  }

  .card-content {
    will-change: height;
    overflow-y: hidden;
    padding: 16px;
    background-color: #ffffff;
    //border-top: 1px solid var(--card-header-background-color);
  }
}
</style>
