package models

import (
	"slices"
	"strings"
	"time"

	"humpback/common/enum"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/common/verify"
	"humpback/pkg/utils"
	"humpback/types"
)

type RegistryCreateReqInfo struct {
	RegistryName string `json:"registryName"`
	URL          string `json:"url"`
	IsDefault    bool   `json:"isDefault"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func (r *RegistryCreateReqInfo) Check() error {
	if r.Username != "" {
		r.Username = utils.RSADecrypt(r.Username)
	}
	if r.Password != "" {
		r.Password = utils.RSADecrypt(r.Password)
	}
	if err := verify.CheckRequiredAndLengthLimit(r.RegistryName, enum.LimitRegistryName.Min, enum.LimitRegistryName.Max, locales.CodeRegistryNameNotEmpty, locales.CodeRegistryNameLimitLength); err != nil {
		return err
	}
	if err := verify.CheckRequiredAndLengthLimit(r.URL, 0, enum.LimitRegistryUrl.Max, locales.CodeRegistryUrlNotEmpty, locales.CodeRegistryUrlLimitLength); err != nil {
		return err
	}
	if err := verify.CheckLengthLimit(r.Username, 0, enum.LimitUsername.Max, locales.CodeRegistryUsernameLimitLength); err != nil {
		return err
	}
	if err := verify.CheckLengthLimit(r.Password, 0, enum.LimitPassword.Max, locales.CodeRegistryPasswordLimitLength); err != nil {
		return err
	}
	return nil
}

func (r *RegistryCreateReqInfo) NewRegistryInfo() *types.Registry {
	nowT := time.Now().UnixMilli()
	return &types.Registry{
		RegistryId:   utils.NewGuidStr(),
		RegistryName: r.RegistryName,
		URL:          r.URL,
		IsDefault:    r.IsDefault,
		Username:     r.Username,
		Password:     r.Password,
		CreatedAt:    nowT,
		UpdatedAt:    nowT,
	}
}

type RegistryUpdateReqInfo struct {
	RegistryId string `json:"registryId"`
	RegistryCreateReqInfo
}

func (r *RegistryUpdateReqInfo) Check() error {
	if err := verify.CheckIsEmpty(r.RegistryId, locales.CodeRegistryIdNotEmpty); err != nil {
		return err
	}
	return r.RegistryCreateReqInfo.Check()
}

func (r *RegistryUpdateReqInfo) NewRegistryInfo(oldInfo *types.Registry) *types.Registry {
	return &types.Registry{
		RegistryId:   oldInfo.RegistryId,
		RegistryName: r.RegistryName,
		URL:          r.URL,
		IsDefault:    r.IsDefault,
		Username:     r.Username,
		Password:     r.Password,
		CreatedAt:    oldInfo.CreatedAt,
		UpdatedAt:    time.Now().UnixMilli(),
	}
}

type RegistryQueryReqInfo struct {
	types.QueryInfo
}

func (c *RegistryQueryReqInfo) Check() error {
	c.QueryInfo.CheckBase()
	if c.Keywords != "" && slices.Index(c.keywordsModes(), c.Mode) == -1 {
		return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
	}
	return nil
}

func (c *RegistryQueryReqInfo) keywordsModes() []string {
	return []string{"registryName"}
}

func (c *RegistryQueryReqInfo) QueryFilter(registrys []*types.Registry) []*types.Registry {
	result := make([]*types.Registry, 0)
	for _, registry := range registrys {
		if c.filter(registry) {
			result = append(result, registry)
		}
	}
	c.sort(result)
	return result
}

func (c *RegistryQueryReqInfo) filter(info *types.Registry) bool {
	if c.Keywords != "" && c.Mode == "registryName" {
		return strings.Contains(strings.ToLower(info.RegistryName), strings.ToLower(c.Keywords))
	}
	return true
}

func (c *RegistryQueryReqInfo) sort(list []*types.Registry) []*types.Registry {
	var sortField = []string{"registryName", "updatedAt", "createdAt"}
	if c.SortInfo == nil || c.SortInfo.Field == "" || slices.Index(sortField, c.SortInfo.Field) == -1 {
		return list
	}
	slices.SortFunc(list, func(a, b *types.Registry) int {
		switch c.SortInfo.Field {
		case "registryName":
			return types.QuerySortOrder(c.SortInfo.Order, strings.ToLower(a.RegistryName), strings.ToLower(b.RegistryName))
		case "updatedAt":
			return types.QuerySortOrder(c.SortInfo.Order, a.UpdatedAt, b.UpdatedAt)
		case "createdAt":
			return types.QuerySortOrder(c.SortInfo.Order, a.CreatedAt, b.CreatedAt)
		}
		return 1
	})
	return list
}
