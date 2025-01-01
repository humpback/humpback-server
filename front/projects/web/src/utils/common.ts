import { includes, toLower } from "lodash"

export function TableHeight(invalidHeight: number, minHeight = 500): number {
  const pageStore = usePageStore()
  return Math.max(pageStore.screenHeight - invalidHeight, minHeight)
}

export function IncludesIgnoreCase(str?: string, subStr?: string): boolean {
  if (!str) {
    return false
  }
  return includes(toLower(str), toLower(subStr || ""))
}
