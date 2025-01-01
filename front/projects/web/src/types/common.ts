export interface ColorBtn {
  title: string
  color: string
  value: number
  disabled?: boolean
}

export interface TableSortEvent {
  field: string
  order: "asc" | "desc"
  property: string
}
