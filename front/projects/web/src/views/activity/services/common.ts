import { NewPageInfo, NewSortInfo, QueryInfo } from "@/types"
import { cloneDeep, omitBy } from "lodash-es"

export const sortOptions = ["createdAt"]

export const defaultSort = NewSortInfo("createdAt", "desc")
export const defaultPage = NewPageInfo(1, 20)
export const defaultFilter = { group: "", user: "", action: "", startAt: 0, endAT: 0 }

export class QueryActivityServiceInfo extends QueryInfo {
  constructor(queryInfo: any) {
    super(queryInfo, ["keywords"], defaultPage, defaultSort, sortOptions, cloneDeep(defaultFilter))
    this.filter.group = queryInfo["group"] ? queryInfo["group"] : defaultFilter.group
    this.filter.action = queryInfo["action"] ? queryInfo["action"] : defaultFilter.action
    this.filter.startAt = queryInfo["startAt"] ? Number(queryInfo["startAt"]).valueOf() : defaultFilter.startAt
    this.filter.endAT = queryInfo["endAT"] ? Number(queryInfo["endAT"]).valueOf() : defaultFilter.endAT
    this.filter.user = queryInfo["user"] ? queryInfo["user"] : defaultFilter.user
  }

  urlQuery() {
    return {
      query: Object.assign(
        {},
        {
          group: this.filter.group || undefined,
          user: this.filter.user || undefined,
          startAt: this.filter.startAt || undefined,
          endAt: this.filter.endAt || undefined,
          action: this.filter.action || undefined
        },
        this.getBaseQuery()
      )
    }
  }

  searchParams() {
    return omitBy(this, (value, key) => key.startsWith("_"))
  }
}
