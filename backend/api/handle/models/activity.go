package models

import (
	"strconv"
	"strings"

	"humpback/common/locales"
	"humpback/common/response"
	"humpback/common/verify"
	"humpback/types"
)

type ActivityQueryFilterInfo struct {
	UserId   string               `json:"userId"`
	StartAt  int64                `json:"startAt"`
	EndAt    int64                `json:"endAt"`
	GroupId  string               `json:"groupId"`
	Action   types.ActivityAction `json:"action"`
	Operator string               `json:"operator"`
	Type     string               `json:"type"`
}

type ActivityQueryReqInfo struct {
	UserInfo *types.User `json:"-"`
	types.QueryInfo
	FilterInfo *ActivityQueryFilterInfo `json:"-"`
}

func (a *ActivityQueryReqInfo) Check() error {
	a.CheckBase()
	a.FilterInfo = new(ActivityQueryFilterInfo)
	if len(a.Filter) == 0 {
		return response.NewBadRequestErr(locales.CodeActivityTypeNotEmpty)
	}
	if err := ParseMapToStructConvert(a.Filter, a.FilterInfo); err != nil {
		return err
	}
	if err := verify.CheckIsEmpty(a.FilterInfo.Type, locales.CodeActivityTypeNotEmpty); err != nil {
		return err
	}
	if a.FilterInfo.StartAt > 0 && a.FilterInfo.EndAt < a.FilterInfo.StartAt {
		return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
	}
	return nil
}

func (a *ActivityQueryReqInfo) isValid(key string) (bool, error) {
	keys := strings.Split(key, "-")
	if a.FilterInfo.StartAt > 0 || a.FilterInfo.EndAt > 0 {
		timestamp, err := strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			return false, err
		}
		if a.FilterInfo.StartAt < 0 && timestamp < a.FilterInfo.StartAt {
			return false, nil
		}
		if a.FilterInfo.EndAt > 0 && timestamp > a.FilterInfo.EndAt {
			return false, nil
		}
	}

	//todo  过滤不同类型的查询情况
	return true, nil
}
