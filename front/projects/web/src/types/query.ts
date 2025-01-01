import { SortType } from "@/models"

export interface QueryInfo {
  keywords: string
  mode: number
  filter?: any
  onlyTotal?: boolean
  IncludeDeleted?: boolean
  pageInfo: PageInfo
  sortBy: SortInfo
}

export interface PageInfo {
  index: number
  size: number
}

export interface SortInfo {
  field: string
  order: SortType | string
}

export interface QueryList<T> {
  total: number
  pageInfo: PageInfo
  objects: T[]
}
