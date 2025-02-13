package db

import (
	"encoding/json"
	"fmt"
	"slices"

	bolt "go.etcd.io/bbolt"
	"humpback/common/response"
	"humpback/types"
)

func NodesGetAll() ([]*types.Node, error) {
	return GetDataAll[types.Node](BucketNodes)
}

func NodesAdd(nodes []*types.Node) error {
	if err := TransactionUpdates(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketNodes))
		if bucket == nil {
			return ErrBucketNotExist
		}
		for _, node := range nodes {
			data, err := json.Marshal(node)
			if err != nil {
				return fmt.Errorf("failed to encode node data: %s", err)
			}
			if err = bucket.Put([]byte(node.NodeId), data); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}

func NodeDelete(id string) error {
	if err := DeleteData(BucketNodes, id); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}

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

func GetAllEnabledNodes() ([]*types.Node, error) {
	nodes, err := GetDataByQuery[types.Node](BucketNodes, func(key string, node interface{}) bool {
		return node.(*types.Node).IsEnable
	})
	return nodes, err
}

func GetGroupByNodeId(nodeId string) []string {
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

func GetNodesByGroupId(groupId string) ([]*types.Node, error) {
	ng, err := GetDataById[types.NodesGroups](BucketNodesGroups, groupId)
	if err != nil {
		return nil, err
	} else {
		nodes := make([]*types.Node, 0)
		for _, v := range ng.Nodes {
			node, err := GetDataById[types.Node](BucketNodes, v)
			if err == nil {
				nodes = append(nodes, node)
			}
		}
		return nodes, nil
	}
}

func GetNodesByIds(nodeIds []string) ([]*types.Node, error) {
	nodes := make([]*types.Node, 0)
	for _, v := range nodeIds {
		node, err := GetDataById[types.Node](BucketNodes, v)
		if err == nil {
			nodes = append(nodes, node)
		}
	}
	return nodes, nil
}
