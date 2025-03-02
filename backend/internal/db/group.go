package db

import (
	"slices"
	"strings"

	"humpback/types"
)

func GroupsGetAll() ([]*types.NodesGroups, error) {
	return GetDataAll[types.NodesGroups](BucketNodesGroups)
}

func GroupGetById(id string) (*types.NodesGroups, error) {
	return GetDataById[types.NodesGroups](BucketNodesGroups, id)
}

func GroupGetByNodeId(nodeId string) []string {
	groups := make([]string, 0)
	ng, err := GetDataByQuery[types.NodesGroups](BucketNodesGroups, func(key string, value interface{}) bool {
		ngp := value.(*types.NodesGroups)
		return slices.Contains(ngp.Nodes, nodeId)
	})

	if err == nil {
		for _, v := range ng {
			groups = append(groups, v.GroupId)
		}
	}
	return groups
}

func GroupsGetByName(name string, isLower bool) ([]*types.NodesGroups, error) {
	groups, err := GroupsGetAll()
	if err != nil {
		return nil, err
	}
	var result []*types.NodesGroups
	for _, group := range groups {
		if isLower && strings.EqualFold(group.GroupName, name) {
			result = append(result, group)
		}
		if !isLower && group.GroupName == name {
			result = append(result, group)
		}
	}
	return result, nil
}

func GroupUpdate(info *types.NodesGroups) error {
	return SaveData(BucketNodesGroups, info.GroupId, info)
}

func GroupDelete(id string) error {
	return DeleteData(BucketNodesGroups, id)
}
