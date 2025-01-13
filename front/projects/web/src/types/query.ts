export interface QueryInfo {
  keywords: string
  mode: string
  filter: {
    [key: string]: any
  }
  pageInfo?: PageInfo
  sortInfo?: SortInfo
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
