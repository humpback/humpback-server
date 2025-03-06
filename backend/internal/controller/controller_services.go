package controller

import (
	"slices"
	"strings"

	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/internal/db"
	"humpback/pkg/utils"
	"humpback/types"
)

func ServiceQuery(groupId string, queryInfo *models.ServiceQueryReqInfo) (*response.QueryResult[types.Service], error) {
	services, err := db.ServicesGetByPrefix(groupId)
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
	total, err := db.ServiceGetTotalByPrefix(groupId)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return 0, response.NewBadRequestErr(locales.CodeGroupNotExist)
		}
		return 0, response.NewRespServerErr(err.Error())
	}
	return total, nil
}

func ServiceCreate(info *models.ServiceCreateReqInfo) (string, error) {
	services, err := db.ServicesGetByPrefix(info.GroupId)
	if err != nil {
		if err == db.ErrKeyNotExist {
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

func ServiceUpdate(serviceChangeChan chan types.ServiceChangeInfo, info *models.ServiceUpdateReqInfo) (string, error) {
	service, err := Service(info.GroupId, info.ServiceId)
	if err != nil {
		return "", err
	}
	switch info.Type {
	case models.ServiceUpdateBasicInfo:
		service.Description = info.Description
	case models.ServiceUpdateApplication:
		service.Meta = info.MetaInfo
	case models.ServiceUpdateDeployment:
		service.Deployment = info.DeploymentInfo
	}
	service.UpdatedAt = utils.NewActionTimestamp()
	if service.IsEnabled {
		service.Version = utils.GenerateRandomStringWithLength(5)
	}
	if err = db.ServiceUpdate(service); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	//if service.IsEnabled {
	//	sendServiceEvent(serviceChangeChan, service.ServiceId, service.Version, service.Action)
	//}
	//todo 检查状态后，往schedule发送消息
	return service.ServiceId, nil
}

func Service(groupId, serviceId string) (*types.Service, error) {
	service, err := db.ServiceGetById(serviceId)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeServiceNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	if service.GroupId != groupId {
		return nil, response.NewBadRequestErr(locales.CodeServiceNotExist)
	}
	return service, nil
}

func ServiceOperate(info *models.ServiceOperateReqInfo) (string, error) {
	service, err := Service(info.GroupId, info.ServiceId)
	if err != nil {
		return "", err
	}
	switch info.Aciton {
	case types.ServiceActionEnable:
		if service.IsEnabled {
			return "", response.NewBadRequestErr(locales.CodeServiceIsEnabled)
		}
		service.Version = utils.GenerateRandomStringWithLength(5)
		//todo 判断enable和disable状态是否不一致
		service.IsEnabled = true
	case types.ServiceActionDisable:
		if service.IsEnabled {
			return "", response.NewBadRequestErr(locales.CodeServiceIsDisabled)
		}
		service.IsEnabled = false
	case types.ServiceActionStart, types.ServiceActionRestart, types.ServiceActionStop:
		if !service.IsEnabled {
			return "", response.NewBadRequestErr(locales.CodeServiceIsNotEnable)
		}
		service.Action = info.Aciton
	}
	service.UpdatedAt = utils.NewActionTimestamp()
	if err = db.ServiceUpdate(service); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	//todo 检查状态后，往schedule发送消息
	return service.ServiceId, nil
}

func ServiceDelete(groupId, serviceId string) error {
	service, err := db.ServiceGetById(serviceId)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil
		}
		return response.NewRespServerErr(err.Error())
	}
	if service.GroupId != groupId {
		return nil
	}
	if err = db.ServiceDelete(serviceId); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	if service.IsEnabled {
		//todo 检查状态后，往schedule发送消息
	}
	return nil
}
