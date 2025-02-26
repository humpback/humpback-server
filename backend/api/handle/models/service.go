package models

import (
	"fmt"
	"slices"
	"strings"

	"humpback/common/enum"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/common/verify"
	"humpback/pkg/utils"
	"humpback/types"
)

type ServiceCreateReqInfo struct {
	GroupId     string `json:"groupId"`
	ServiceName string `json:"serviceName"`
	Description string `json:"description"`
}

func (s *ServiceCreateReqInfo) Check() error {
	if err := verify.CheckIsEmpty(s.GroupId, locales.CodeGroupIdNotEmpty); err != nil {
		return err
	}
	if err := verify.CheckRequiredAndLengthLimit(s.ServiceName, enum.LimitServiceName.Min, enum.LimitServiceName.Max, locales.CodeServiceNameNotEmpty, locales.CodeServiceNameLimitLength); err != nil {
		return err
	}
	if err := verify.CheckLengthLimit(s.Description, enum.LimitDescription.Min, enum.LimitDescription.Max, locales.CodeDescriptionLimitLength); err != nil {
		return err
	}
	return nil
}

func (s *ServiceCreateReqInfo) NewServiceInfo() *types.Service {
	nowT := utils.NewActionTimestamp()
	return &types.Service{
		ServiceId:   fmt.Sprintf("%s%s", s.GroupId, utils.GenerateRandomStringWithLength(8)),
		GroupId:     s.GroupId,
		ServiceName: s.ServiceName,
		Description: s.Description,
		Version:     utils.GenerateRandomStringWithLength(5),
		Action:      "",
		IsEnabled:   false,
		IsDelete:    false,
		Status:      types.ServiceStatusNotReady,
		Meta:        nil,
		Deployment:  nil,
		Containers:  make([]*types.ContainerStatus, 0),
		CreatedAt:   nowT,
		UpdatedAt:   nowT,
	}
}

type ServiceQueryFilterInfo struct {
	Status   string `json:"status"`
	Schedule string `json:"schedule"`
}

type ServiceQueryReqInfo struct {
	types.QueryInfo
	UserInfo   *types.User             `json:"-"`
	FilterInfo *ServiceQueryFilterInfo `json:"-"`
}

func (s *ServiceQueryReqInfo) Check() error {
	s.CheckBase()
	s.FilterInfo = new(ServiceQueryFilterInfo)
	if err := ParseMapToStructConvert(s.Filter, s.FilterInfo); err != nil {
		return err
	}

	if s.FilterInfo != nil {
		if s.FilterInfo.Status != "" && slices.Index([]string{
			types.SwitchEnabled,
			types.SwitchDisabled,
			types.ServiceStatusRunning,
			types.ServiceStatusNotReady,
			types.ServiceStatusFailed,
		}, s.FilterInfo.Status) == -1 {
			return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
		}
		if s.FilterInfo.Schedule != "" && slices.Index([]string{"Yes", "No"}, s.FilterInfo.Schedule) == -1 {
			return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
		}
	}

	if s.Keywords != "" && s.Mode != "keywords" {
		return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
	}
	return nil
}

func (s *ServiceQueryReqInfo) QueryFilter(services []*types.Service) []*types.Service {
	result := make([]*types.Service, 0)
	for _, service := range services {
		if s.filter(service) {
			result = append(result, service)
		}
	}
	s.sort(result)
	return result
}

func (s *ServiceQueryReqInfo) filter(service *types.Service) bool {
	if s.FilterInfo != nil {
		switch s.FilterInfo.Status {
		case types.SwitchEnabled:
			if !service.IsEnabled {
				return false
			}
		case types.SwitchDisabled:
			if service.IsEnabled {
				return false
			}
		case types.ServiceStatusRunning, types.ServiceStatusNotReady, types.ServiceStatusFailed:
			if !service.IsEnabled || service.Status != s.FilterInfo.Status {
				return false
			}
		}

		if s.FilterInfo.Schedule == "Yes" && (service.Deployment == nil || service.Deployment.Type != types.DeployTypeSchedule) {
			return false
		}

		if s.FilterInfo.Schedule == "No" && (service.Deployment == nil || service.Deployment.Type != types.DeployTypeBackground) {
			return false
		}
	}
	if strings.Contains(strings.ToLower(service.ServiceName), strings.ToLower(s.Keywords)) {
		return true
	}
	if service.Meta != nil && strings.Contains(strings.ToLower(service.Meta.Image), strings.ToLower(s.Keywords)) {
		return true
	}
	return false
}

func (s *ServiceQueryReqInfo) sort(list []*types.Service) []*types.Service {
	var sortField = []string{"serviceName", "updatedAt", "createdAt"}
	if s.SortInfo == nil || s.SortInfo.Field == "" || slices.Index(sortField, s.SortInfo.Field) == -1 {
		return list
	}
	slices.SortFunc(list, func(a, b *types.Service) int {
		switch s.SortInfo.Field {
		case "serviceName":
			return types.QuerySortOrder(s.SortInfo.Order, strings.ToLower(a.ServiceName), strings.ToLower(b.ServiceName))
		case "updatedAt":
			return types.QuerySortOrder(s.SortInfo.Order, a.UpdatedAt, b.UpdatedAt)
		case "createdAt":
			return types.QuerySortOrder(s.SortInfo.Order, a.CreatedAt, b.CreatedAt)
		}
		return 1
	})
	return list
}
