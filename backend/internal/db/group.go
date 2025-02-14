package db

import (
	"slices"

	"humpback/types"
)

func GroupGetByNodeId(nodeId string) []string {
	groups := make([]string, 0)
	ng, err := GetDataByQuery[types.NodesGroups](BucketNodesGroups, func(key string, nodesGroups interface{}) bool {
		ngp := nodesGroups.(*types.NodesGroups)
		return slices.Contains(ngp.Nodes, nodeId)
	})

	if err == nil {
		for _, v := range ng {
			groups = append(groups, v.GroupId)
		}
	}
	return groups
}
