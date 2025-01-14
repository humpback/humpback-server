package types

type QueryInfo struct {
	Keywords string         `json:"keywords"`
	Mode     string         `json:"mode"`
	Filter   map[string]any `json:"filter"`
	PageInfo *PageInfo      `json:"pageInfo"`
	SortInfo *SortInfo      `json:"sortInfo"`
}

func (q *QueryInfo) CheckBase() {
	if q.Filter == nil {
		q.Filter = make(map[string]any)
	}
	if q.PageInfo == nil {
		if q.PageInfo.Index < 1 {
			q.PageInfo.Index = 1
		}
		if q.PageInfo.Size < 1 {
			q.PageInfo.Size = 20
		}
	}
	if q.SortInfo == nil {
		if q.SortInfo.Field == "" {
			q.SortInfo = nil
		} else if q.SortInfo.Order != "asc" && q.SortInfo.Order != "desc" {
			q.SortInfo.Order = "asc"
		}
	}
}

type SortInfo struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

type PageInfo struct {
	Index int `json:"index"`
	Size  int `json:"size"`
}

type QueryResult[T any] struct {
	Total int  `json:"total"`
	Data  []*T `json:"data"`
}
