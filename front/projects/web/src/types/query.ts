import { ParseNumber } from "@/utils"

export interface QueryInfo {
  keywords: string
  mode: string
  filter: {
    [key: string]: any
  }
  pageInfo?: PageInfo
  sortInfo?: SortInfo
}

export default class Query {
  keywords: string
  mode: string
  filter: {
    [key: string]: any
  }
  pageInfo: PageInfo
  sortInfo: SortInfo

  constructor(queryInfo: any) {
    this.keywords = queryInfo["keywords"] ? queryInfo["keywords"] : ""
    this.mode = queryInfo["mode"] ? queryInfo["mode"] : ""
    this.filter = {}
    this.pageInfo = {
      index: ParseNumber(queryInfo["pageIndex"], 1)!,
      size: ParseNumber(queryInfo["pageIndex"], 20)!
    }
    this.sortInfo = {
      field: queryInfo["field"] ? queryInfo["field"] : "",
      order: ParseQueryOrder(queryInfo["order"])
    }
  }
}

export interface PageInfo {
  index: number
  size: number
}

export function NewPageInfo(index: number, size: number): PageInfo {
  return {
    index: index,
    size: size
  }
}

export interface SortInfo {
  field: string
  order: "asc" | "desc" | ""
}

export function NewSortInfo(field: string, order: "asc" | "desc" | ""): SortInfo {
  return {
    field: field,
    order: order
  }
}

export interface QueryList<T> {
  total: number
  pageInfo: PageInfo
  objects: T[]
}

function ParseQueryOrder(value: any, defaultValue: "desc" | "asc" = "asc") {
  const result = toString(value).trim()
  if (["desc", "asc"].indexOf(result) === -1) {
    return defaultValue
  }
  return result
}
