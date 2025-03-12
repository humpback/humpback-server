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

	"github.com/google/go-cmp/cmp"
)

func ServiceQuery(groupId string, queryInfo *models.ServiceQueryReqInfo) (*response.QueryResult[types.Service], error) {
	services, err := db.ServicesGetValidByPrefix(groupId)
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
	services, err := db.ServicesGetValidByPrefix(groupId)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return 0, response.NewBadRequestErr(locales.CodeGroupNotExist)
		}
		return 0, response.NewRespServerErr(err.Error())
	}
	return len(services), nil
}

func ServiceCreate(info *models.ServiceCreateReqInfo) (string, error) {
	if err := serviceCheckNameExist(info.GroupId, info.ServiceName); err != nil {
		return "", err
	}
	newService := info.NewServiceInfo()
	if err := db.ServiceUpdate(newService); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	return newService.ServiceId, nil
}

func ServiceClone(info *models.ServiceCloneReqInfo) (string, error) {
	serviceInfo, err := Service(info.GroupId, info.ServiceId)
	if err != nil {
		return "", err
	}
	if err = serviceCheckNameExist(info.NewGroupId, info.ServiceName); err != nil {
		return "", err
	}
	nowT := utils.NewActionTimestamp()
	if info.GroupId != info.NewGroupId {
		serviceInfo.Deployment.Placements = make([]*types.PlacementInfo, 0)
	}
	newService := &types.Service{
		ServiceId:   utils.NewServiceId(info.NewGroupId),
		GroupId:     info.NewGroupId,
		ServiceName: info.ServiceName,
		Description: info.Description,
		Version:     "",
		Action:      "",
		IsEnabled:   false,
		IsDelete:    false,
		Status:      "",
		Meta:        serviceInfo.Meta,
		Deployment:  serviceInfo.Deployment,
		Containers:  make([]*types.ContainerStatus, 0),
		CreatedAt:   nowT,
		UpdatedAt:   nowT,
	}
	if err = db.ServiceUpdate(newService); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	return newService.ServiceId, nil
}

func serviceCheckNameExist(groupId, serviceName string) error {
	services, err := db.ServicesGetValidByPrefix(groupId)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return response.NewBadRequestErr(locales.CodeGroupNotExist)
		}
		return response.NewRespServerErr(err.Error())
	}
	if slices.IndexFunc(services, func(service *types.Service) bool {
		return strings.ToLower(service.ServiceName) == strings.ToLower(serviceName)
	}) != -1 {
		return response.NewBadRequestErr(locales.CodeServiceNameAlreadyExist)
	}
	return nil
}

func ServiceUpdate(svcChan chan types.ServiceChangeInfo, info *models.ServiceUpdateReqInfo) (string, error) {
	service, err := Service(info.GroupId, info.ServiceId)
	if err != nil {
		return "", err
	}

	switch info.Type {
	case models.ServiceUpdateBasicInfo:
		service.Description = info.Description
	case models.ServiceUpdateApplication:
		{
			if service.Meta != nil {
				info.MetaInfo.Envs = service.Meta.Envs
			}
			if service.IsEnabled && !cmp.Equal(service.Meta, info.MetaInfo) {
				service.Version = utils.NewVersionId()
			}
			service.Meta = info.MetaInfo
		}
	case models.ServiceUpdateDeployment:
		{
			if service.IsEnabled && (service.Deployment == nil || !cmp.Equal(service.Deployment.Schedule, info.DeploymentInfo.Schedule)) {
				service.Version = utils.NewVersionId()
			}
			service.Deployment = info.DeploymentInfo
		}
	}
	service.Action = types.ServiceActionDispatch
	service.UpdatedAt = utils.NewActionTimestamp()
	if err = db.ServiceUpdate(service); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	if service.IsEnabled {
		sendServiceEvent(svcChan, service.ServiceId, service.Version, service.Action)
	}
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
	if service.GroupId != groupId || service.IsDelete {
		return nil, response.NewBadRequestErr(locales.CodeServiceNotExist)
	}
	return service, nil
}

func ServiceOperate(svcChan chan types.ServiceChangeInfo, info *models.ServiceOperateReqInfo) (string, error) {
	service, err := Service(info.GroupId, info.ServiceId)
	if err != nil {
		return "", err
	}

	switch info.Action {
	case types.ServiceActionEnable:
		if service.IsEnabled {
			return "", response.NewBadRequestErr(locales.CodeServiceIsEnabled)
		}
		service.Version = utils.NewVersionId()
		service.IsEnabled = true
	case types.ServiceActionDisable:
		if !service.IsEnabled {
			return "", response.NewBadRequestErr(locales.CodeServiceIsDisabled)
		}
		service.IsEnabled = false
	case types.ServiceActionStart, types.ServiceActionRestart, types.ServiceActionStop:
		if !service.IsEnabled {
			return "", response.NewBadRequestErr(locales.CodeServiceIsNotEnable)
		}
		service.Action = info.Action
	}
	service.UpdatedAt = utils.NewActionTimestamp()
	if err = db.ServiceUpdate(service); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	sendServiceEvent(svcChan, service.ServiceId, service.Version, info.Action)
	return service.ServiceId, nil
}

func ServiceSoftDelete(svcChan chan types.ServiceChangeInfo, groupId, serviceId string) error {
	service, err := db.ServiceGetById(serviceId)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil
		}
		return response.NewRespServerErr(err.Error())
	}
	if service.GroupId != groupId || service.IsDelete {
		return nil
	}
	service.IsDelete = true
	if err = db.ServiceUpdate(service); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	if service.IsEnabled {
		sendServiceEvent(svcChan, service.ServiceId, service.Version, types.ServiceActionDelete)
	}
	return nil
}
