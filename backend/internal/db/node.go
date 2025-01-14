package db

import (
	"humpback/types"
	"slices"
)

func UpdateNodeStatus(nodeId string, status string, lastUpdate int64, cpuUsage float32, memoryUsage float32) error {
	node, err := GetDataById[types.Node](BucketNodes, nodeId)
	if err != nil {
		return err
	}
	node.Status = status
	node.UpdatedAt = lastUpdate
	node.CPUUsage = cpuUsage
	node.MemoryUsage = memoryUsage
	return SaveData(BucketNodes, nodeId, node)
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
			groups = append(groups, v.GroupId)
		}
	}
	return groups
}

func GetOfflineNodesByGroupId(groupId string) ([]string, int, error) {
	ng, err := GetDataById[types.NodesGroups](BucketNodesGroups, groupId)
	if err != nil {
		return nil, 0, err
	} else {
		nodes := make([]string, 0)
		for _, v := range ng.Nodes {
			node, err := GetDataById[types.Node](BucketNodes, v)
			if err == nil && node.Status == types.NodeStatusOffline {
				nodes = append(nodes, node.NodeId)
			}
		}
		return nodes, len(ng.Nodes), nil
	}
}

func GetOnlineNodesByGroupId(groupId string) ([]*types.Node, error) {
	ng, err := GetDataById[types.NodesGroups](BucketNodesGroups, groupId)
	if err != nil {
		return nil, err
	} else {
		nodes := make([]*types.Node, 0)
		for _, v := range ng.Nodes {
			node, err := GetDataById[types.Node](BucketNodes, v)
			if err == nil && node.Status == types.NodeStatusOnline {
				nodes = append(nodes, node)
			}
		}
		return nodes, nil
	}
}
