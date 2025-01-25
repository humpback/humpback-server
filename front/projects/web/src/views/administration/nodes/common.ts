import { NewPageInfo, NewSortInfo, QueryInfo } from "@/types"
import { find, omitBy } from "lodash-es"

export const sortOptions = ["ip", "hostname", "updatedAt", "createdAt"]

export const defaultSort = NewSortInfo("ip", "asc")
export const defaultPage = NewPageInfo(1, 20)
export const defaultFilter = { group: "", status: "" }

export const statusOptions = [
  { label: "label.all", value: "" },
  { label: "label.enabled", value: "enabled" },
  { label: "label.disabled", value: "disabled" },
  { label: "label.online", value: "online" },
  { label: "label.offline", value: "offline" }
]

export class QueryNodesInfo extends QueryInfo {
  constructor(queryInfo: any, groupOptions: any[]) {
    super(queryInfo, ["keywords", "label"], defaultPage, defaultSort, sortOptions, defaultFilter)
    this.filter.group = queryInfo["group"] ? (queryInfo["group"] as string) : ""
    const statusInfo = find(statusOptions, x => x.value === (queryInfo["status"] as string))
    this.filter.status = statusInfo?.value || ""
  }

  getQuery() {
    return {
      query: Object.assign({}, { group: this.filter.group || undefined, status: this.filter.status || undefined }, this.getBaseQuery())
    }
  }

  getSearch() {
    return omitBy(this, (value, key) => key.startsWith("_"))
  }
}
