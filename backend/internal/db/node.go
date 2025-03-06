package db

import (
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"

	"humpback/types"
)

func NodesGetAll() ([]*types.Node, error) {
	return GetDataAll[types.Node](BucketNodes)
}

func NodeUpdateStatus(nodeInfo *types.NodeSimpleInfo) error {
	node, err := GetDataById[types.Node](BucketNodes, nodeInfo.NodeId)
	if err != nil {
		return err
	}
	node.Name = nodeInfo.Name
	node.Port = nodeInfo.Port
	node.Status = nodeInfo.Status
	node.UpdatedAt = nodeInfo.LastHeartbeat
	node.CPUUsage = nodeInfo.CPUUsage
	node.MemoryUsage = nodeInfo.MemoryUsage
	return SaveData(BucketNodes, nodeInfo.NodeId, node)
}

func NodeGetById(nodeId string) (*types.Node, error) {
	return GetDataById[types.Node](BucketNodes, nodeId)
}

func NodesGetAllEnabled() ([]*types.Node, error) {
	nodes, err := GetDataByQuery[types.Node](BucketNodes, func(key string, node interface{}) bool {
		return node.(*types.Node).IsEnable
	})
	return nodes, err
}

func NodesGetEnabledByGroupId(groupId string) ([]*types.Node, error) {
	ng, err := GroupGetById(groupId)
	if err != nil {
		return nil, err
	}
	nodes := make([]*types.Node, 0)
	for _, v := range ng.Nodes {
		node, err := GetDataById[types.Node](BucketNodes, v)
		if err == nil && node.IsEnable {
			nodes = append(nodes, node)
		}
	}
	return nodes, nil
}

func NodesGetByIds(ids []string, ignoreNotExist bool) ([]*types.Node, error) {
	return GetDataByIds[types.Node](BucketNodes, ids, ignoreNotExist)
}

func NodeUpdate(node *types.Node) error {
	return SaveData[*types.Node](BucketNodes, node.NodeId, node)
}

func NodesUpdate(nodes []*types.Node) error {
	return TransactionUpdates(func(tx *bolt.Tx) error {
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
	})
}

func NodeDelete(id string) error {
	return DeleteData(BucketNodes, id)
}
