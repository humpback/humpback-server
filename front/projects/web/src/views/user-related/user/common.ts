import { NewPageInfo, NewSortInfo, QueryInfo } from "@/types"
import { find, omitBy } from "lodash-es"

export const sortOptions = ["username", "updatedAt", "createdAt"]

export const defaultSort = NewSortInfo("username", "asc")
export const defaultPage = NewPageInfo(1, 20)
export const defaultFilter = { role: 0 }

export class QueryUserInfo extends QueryInfo {
  constructor(queryInfo: any) {
    super(queryInfo, ["username", "configNameemail", "phone"], defaultPage, defaultSort, sortOptions, defaultFilter)
    const role = queryInfo["role"] ? Number(queryInfo["role"]).valueOf() : defaultFilter.role
    this.filter.role = !isNaN(role) && find([0, UserRole.User, UserRole.Admin, UserRole.SupperAdmin], x => x === role) ? role : defaultFilter.role
  }

  getQuery() {
    return {
      query: Object.assign({}, this.filter.role !== defaultFilter.role ? { role: this.filter.role } : {}, this.getBaseQuery())
    }
  }

  getSearch() {
    return omitBy(this, (value, key) => key.startsWith("_"))
  }
}
