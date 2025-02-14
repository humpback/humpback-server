package controller

import (
	"time"

	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/internal/db"
	"humpback/types"
)

func NodeCreate(nodes models.NodesCreateReqInfo) error {
	nodeList, err := db.NodesGetAll()
	if err != nil {
		return response.NewRespServerErr(err.Error())
	}
	addNodes := nodes.NewNodesInfo()
	for _, node := range addNodes {
		for _, existNode := range nodeList {
			if node.IpAddress == existNode.IpAddress {
				return response.NewBadRequestErr(locales.CodeNodesIpAddressAlreadyExist)
			}
		}
	}
	if err = db.NodesUpdate(addNodes); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	return nil
}

func NodeUpdateLabel(info *models.NodeUpdateLabelReqInfo) (string, error) {
	node, err := Node(info.NodeId)
	if err != nil {
		return "", err
	}
	node.Labels = info.Labels
	node.UpdatedAt = time.Now().UnixMilli()
	if err = db.NodeUpdate(node); err != nil {
		return "", response.NewRespServerErr(err.Error())
	}
	//todo 发送消息到scheduler，注：是否检查已经在group中
	return info.NodeId, nil
}

func NodeUpdateSwitch(nodeId string, enabled bool) (string, error) {
	node, err := Node(nodeId)
	if err != nil {
		return "", err
	}
	if node.IsEnable != enabled {
		node.IsEnable = enabled
		if err = db.NodeUpdate(node); err != nil {
			return "", response.NewRespServerErr(err.Error())
		}
		//todo 发送消息到scheduler,注：是否检查已经在group中
	}
	return nodeId, nil
}

func NodeDelete(id string) error {
	node, err := db.NodeGetById(id)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil
		}
		return response.NewRespServerErr(err.Error())
	}
	if err = db.NodeDelete(node.NodeId); err != nil {
		return response.NewRespServerErr(err.Error())
	}
	if node.IsEnable {
		//todo 发送消息到scheduler，注：是否检查已经在group中
	}
	return nil
}

func Node(id string) (*types.Node, error) {
	node, err := db.NodeGetById(id)
	if err != nil {
		if err == db.ErrKeyNotExist {
			return nil, response.NewBadRequestErr(locales.CodeNodesNotExist)
		}
		return nil, response.NewRespServerErr(err.Error())
	}
	return node, nil
}

func NodesQuery(queryInfo *models.NodeQueryReqInfo) (*response.QueryResult[types.Node], error) {
	nodes, err := db.NodesGetAll()
	if err != nil {
		return nil, response.NewRespServerErr(err.Error())
	}
	result := queryInfo.QueryFilter(nodes)
	return response.NewQueryResult[types.Node](
		len(result),
		types.QueryPagination[types.Node](queryInfo.PageInfo, result),
	), nil
}
