<script lang="ts" setup>
import type { Sort, TableProps } from "element-plus"
import { omit } from "lodash-es"

type Props = Partial<
  TableProps<any> & {
    pageInfo: PageInfo
    total: number
    sortInfo: SortInfo
    selectedData: any[]
  }
>

const props = withDefaults(defineProps<Props>(), {
  data: [],
  stripe: false,
  fit: true,
  showHeader: true,
  selectOnIndeterminate: true,
  allowDragLastColumn: true
})

const emits = defineEmits<{
  (e: "update:pageInfo", pageInfo: PageInfo): void
  (e: "update:sortInfo", sortInfo: SortInfo): void
  (e: "update:selectedData", selectedData: any[]): void
  (e: "pageChange"): void
  (e: "sortChange"): void
}>()

const slots = useSlots()
const pageStore = usePageStore()

const pageSizeOptions = [10, 20, 30, 50, 100]

const defaultSort = ref<Sort | undefined>(
  props.sortInfo
    ? {
        prop: props.sortInfo?.field || "",
        order: props.sortInfo?.order === "desc" ? "descending" : "ascending"
      }
    : undefined
)
const tableAttrs = computed(() => {
  const attrs = omit(props, ["pageInfo", "total"])
  attrs.defaultSort = defaultSort.value
  defaultSort.value = props.sortInfo
    ? {
        prop: props.sortInfo.field,
        order: props.sortInfo.order === "desc" ? "descending" : "ascending"
      }
    : undefined
  return Object.keys(omit(props, ["pageInfo", "total"])).reduce((acc, key) => {
    if (typeof attrs[key] !== "undefined") {
      acc[key] = props[key]
    }
    return acc
  }, {})
})

function pageChangeEvent(index: number, size: number) {
  if (props.pageInfo) {
    emits("update:pageInfo", { index: index, size: size } as PageInfo)
    emits("pageChange")
  }
}

function sortChangeEvent(sort: any) {
  if (props.sortInfo) {
    emits("update:sortInfo", { field: sort.prop, order: sort.order === "descending" ? "desc" : "asc" } as SortInfo)
    emits("sortChange")
  }
}

function selectionChangeEvent(selectedData: any[]) {
  emits("update:selectedData", selectedData)
}
</script>

<template>
  <el-table v-bind="tableAttrs" @select="selectionChangeEvent($event.selection)" @sort-change="sortChangeEvent" @select-all="selectionChangeEvent">
    <template v-if="!!slots.default" #default>
      <slot name="default" />
    </template>
    <template v-if="!!slots.append" #append>
      <slot name="append" />
    </template>
    <template v-if="!!slots.empty" #empty>
      <slot name="empty" />
    </template>
  </el-table>
  <div v-if="props.pageInfo" class="mt-5 text-align-right">
    <el-pagination
      :background="true"
      :current-page="props.pageInfo.index"
      :layout="pageStore.isSmallScreen ? 'total, prev, pager, next' : 'total, sizes, prev, pager, next, jumper'"
      :page-size="props.pageInfo.size"
      :page-sizes="pageSizeOptions"
      :pager-count="pageStore.isSmallScreen ? 3 : 5"
      :total="props.total"
      @currentChange="pageChangeEvent($event, props.pageInfo.size)"
      @size-change="pageChangeEvent(props.pageInfo.index, $event)" />
  </div>
</template>

<style lang="scss" scoped></style>
