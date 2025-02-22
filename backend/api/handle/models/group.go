package models

import (
	"slices"
	"strings"

	"github.com/jinzhu/copier"
	"golang.org/x/exp/maps"
	"humpback/common/enum"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/common/verify"
	"humpback/pkg/utils"
	"humpback/types"
)

type GroupCreateReqInfo struct {
	GroupName   string   `json:"groupName"`
	Description string   `json:"description"`
	Users       []string `json:"users"`
	Teams       []string `json:"teams"`
}

func (g *GroupCreateReqInfo) Check() error {
	if err := verify.CheckRequiredAndLengthLimit(g.GroupName, enum.LimitGroupName.Min, enum.LimitGroupName.Max, locales.CodeGroupNameNotEmpty, locales.CodeGroupNameLimitLength); err != nil {
		return err
	}
	if err := verify.CheckLengthLimit(g.Description, enum.LimitDescription.Min, enum.LimitDescription.Max, locales.CodeDescriptionLimitLength); err != nil {
		return err
	}
	return nil
}

func (g *GroupCreateReqInfo) NewGroupInfo() *types.NodesGroups {
	nowT := utils.NewActionTimestamp()
	return &types.NodesGroups{
		GroupId:     utils.NewGuidStr(),
		GroupName:   g.GroupName,
		Description: g.Description,
		Users:       g.Users,
		Teams:       g.Teams,
		Nodes:       make([]string, 0),
		CreatedAt:   nowT,
		UpdatedAt:   nowT,
	}
}

type GroupUpdateReqInfo struct {
	GroupId string `json:"groupId"`
	GroupCreateReqInfo
}

func (g *GroupUpdateReqInfo) Check() error {
	if err := verify.CheckIsEmpty(g.GroupId, locales.CodeGroupIdNotEmpty); err != nil {
		return err
	}
	return g.GroupCreateReqInfo.Check()
}

func (g *GroupUpdateReqInfo) NewGroupInfo(oldInfo *types.NodesGroups) *types.NodesGroups {
	return &types.NodesGroups{
		GroupId:     oldInfo.GroupId,
		GroupName:   g.GroupName,
		Description: g.Description,
		Users:       g.Users,
		Teams:       g.Teams,
		Nodes:       oldInfo.Nodes,
		CreatedAt:   oldInfo.CreatedAt,
		UpdatedAt:   utils.NewActionTimestamp(),
	}
}

type GroupUpdateNodesReqInfo struct {
	GroupId  string   `json:"groupId"`
	IsDelete bool     `json:"isDelete"`
	Nodes    []string `json:"nodes"`
}

func (g *GroupUpdateNodesReqInfo) Check() error {
	if err := verify.CheckIsEmpty(g.GroupId, locales.CodeGroupIdNotEmpty); err != nil {
		return err
	}
	nodeList := make([]string, 0)
	for _, node := range g.Nodes {
		if strings.TrimSpace(node) != "" {
			nodeList = append(nodeList, node)
		}
	}
	g.Nodes = nodeList
	if len(g.Nodes) == 0 {
		return response.NewBadRequestErr(locales.CodeNodesNotEmpty)
	}
	return nil
}

func (g *GroupUpdateNodesReqInfo) NewGroupInfo(oldInfo *types.NodesGroups) *types.NodesGroups {
	tempNodes := make(map[string]bool)
	for _, id := range oldInfo.Nodes {
		tempNodes[id] = true
	}
	for _, id := range g.Nodes {
		if g.IsDelete {
			delete(tempNodes, id)
		} else {
			tempNodes[id] = true
		}
	}
	newInfo := &types.NodesGroups{}
	copier.Copy(newInfo, oldInfo)
	newInfo.Nodes = maps.Keys(tempNodes)
	newInfo.UpdatedAt = utils.NewActionTimestamp()
	return newInfo
}

type GroupQueryReqInfo struct {
	UserInfo *types.User `json:"-"`
	types.QueryInfo
}

func (g *GroupQueryReqInfo) Check() error {
	g.QueryInfo.CheckBase()
	if g.Keywords != "" && g.Mode != "groupName" {
		return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
	}
	return nil
}

func (g *GroupQueryReqInfo) QueryFilter(groups []*types.NodesGroups) []*types.NodesGroups {
	result := make([]*types.NodesGroups, 0)
	for _, group := range groups {
		if g.UserInfo.InGroup(group) && strings.Contains(strings.ToLower(group.GroupName), strings.ToLower(g.Keywords)) {
			result = append(result, group)
		}
	}
	g.sort(result)
	return result
}

func (g *GroupQueryReqInfo) sort(list []*types.NodesGroups) []*types.NodesGroups {
	var sortField = []string{"groupName", "updatedAt", "createdAt"}
	if g.SortInfo == nil || g.SortInfo.Field == "" || slices.Index(sortField, g.SortInfo.Field) == -1 {
		return list
	}
	slices.SortFunc(list, func(a, b *types.NodesGroups) int {
		switch g.SortInfo.Field {
		case "groupName":
			return types.QuerySortOrder(g.SortInfo.Order, strings.ToLower(a.GroupName), strings.ToLower(b.GroupName))
		case "updatedAt":
			return types.QuerySortOrder(g.SortInfo.Order, a.UpdatedAt, b.UpdatedAt)
		case "createdAt":
			return types.QuerySortOrder(g.SortInfo.Order, a.CreatedAt, b.CreatedAt)
		}
		return 1
	})
	return list
}

type GroupQueryNodesReqInfo struct {
	types.QueryInfo
}

func (g *GroupQueryNodesReqInfo) Check() error {
	g.CheckBase()
	if g.Keywords != "" && g.Mode != "keywords" {
		return response.NewBadRequestErr(locales.CodeRequestParamsInvalid)
	}
	return nil
}

func (g *GroupQueryNodesReqInfo) QueryFilter(nodes []*types.Node) []*types.Node {
	result := make([]*types.Node, 0)
	for _, node := range nodes {
		if strings.Contains(node.IpAddress, g.Keywords) || strings.Contains(strings.ToLower(node.Name), strings.ToLower(g.Keywords)) {
			result = append(result, node)
		}
	}
	g.sort(result)
	return result
}

func (g *GroupQueryNodesReqInfo) sort(list []*types.Node) []*types.Node {
	var sortField = []string{"ipAddress", "name", "updatedAt", "createdAt"}
	if g.SortInfo == nil || g.SortInfo.Field == "" || slices.Index(sortField, g.SortInfo.Field) == -1 {
		return list
	}
	slices.SortFunc(list, func(a, b *types.Node) int {
		switch g.SortInfo.Field {
		case "ipAddress":
			return types.QuerySortOrder(g.SortInfo.Order, strings.ToLower(a.IpAddress), strings.ToLower(b.IpAddress))
		case "name":
			return types.QuerySortOrder(g.SortInfo.Order, strings.ToLower(a.Name), strings.ToLower(b.Name))
		case "updatedAt":
			return types.QuerySortOrder(g.SortInfo.Order, a.UpdatedAt, b.UpdatedAt)
		case "createdAt":
			return types.QuerySortOrder(g.SortInfo.Order, a.CreatedAt, b.CreatedAt)
		}
		return 1
	})
	return list
}
