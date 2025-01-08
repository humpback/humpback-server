package db

import (
	"humpback/types"
	"slices"
)

func UpdateNodeStatus(nodeId string, status string, lastUpdate int64) error {
	node, err := GetDataById[types.Node](BucketNodes, nodeId)
	if err != nil {
		return err
	}
	node.Status = status
	node.UpdatedAt = lastUpdate
	return SaveData[*types.Node](BucketNodes, nodeId, node)
}

func GetNodeById(nodeId string) (*types.Node, error) {
	return GetDataById[types.Node](BucketNodes, nodeId)
}

func GetGroupByNodeId(nodeId string) []string {
	groups := make([]string, 0)
	ng, err := GetDataByQuery[types.NodesGroups](BucketNodesGroups, func(key string, nodesGroups interface{}) bool {
		ngp := nodesGroups.(types.NodesGroups)
		return slices.Contains(ngp.Nodes, nodeId)
	})

	if err == nil {
		for _, v := range ng {
			groups = append(groups, v.GroupID)
		}
	}
	return groups
}
