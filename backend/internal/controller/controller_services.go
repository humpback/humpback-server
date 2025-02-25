package controller

import (
	"fmt"
	"slices"
	"strings"

	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/internal/db"
	"humpback/types"
)

func ServiceQuery(groupId string, queryInfo *models.ServiceQueryReqInfo) (*response.QueryResult[types.Service], error) {
	services, err := db.ServicesGetByPrefix(fmt.Sprintf("%s-", groupId))
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	result := queryInfo.QueryFilter(services)
	return response.NewQueryResult[types.Service](
		len(result),
		types.QueryPagination[types.Service](queryInfo.PageInfo, result),
	), nil
}

func ServiceTotal(groupId string) (int, error) {
	total, err := db.ServiceGetTotalByPrefix(fmt.Sprintf("%s-", groupId))
	if err != nil {
		if err == db.ErrBucketNotExist {
			return 0, response.NewBadRequestErr(locales.CodeGroupNotExist)
		}
		return 0, response.NewRespServerErr(err.Error())
	}
	return total, nil
}

func ServiceCreate(info *models.ServiceCreateReqInfo) (string, error) {
	services, err := db.ServicesGetByPrefix(fmt.Sprintf("%s-", info.GroupId))
	if err != nil {
		if err == db.ErrBucketNotExist {
			return "", response.NewBadRequestErr(locales.CodeGroupNotExist)
		}
		return "", response.NewRespServerErr(err.Error())
	}
	if slices.IndexFunc(services, func(service *types.Service) bool {
		return strings.ToLower(service.ServiceName) == strings.ToLower(info.ServiceName)
	}) != -1 {
		return "", response.NewBadRequestErr(locales.CodeServiceNameAlreadyExist)
	}
	newService := info.NewServiceInfo()
	if err = db.ServiceUpdate(newService); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	return newService.ServiceId, nil
}
