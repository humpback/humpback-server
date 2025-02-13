package controller

import (
	"humpback/api/handle/models"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/internal/db"
)

func NodeCreate(nodes models.NodesCreateReqInfo) error {
	existNodes, err := db.NodesGetAll()
	if err != nil {
		return response.NewRespServerErr(err.Error())
	}
	addNodes := nodes.NewNodesInfo()
	for _, node := range addNodes {
		for _, existNode := range existNodes {
			if node.IpAddress == existNode.IpAddress {
				return response.NewBadRequestErr(locales.CodeNodesIpAddressAlreadyExist)
			}
		}
	}
	return db.NodesAdd(addNodes)
}

func NodeUpdateLabel(info *models.NodeUpdateLabelReqInfo) error {
	return nil
}

func NodeDelete(id string) error {
	return nil
}

func NodeQuery(queryInfo *models.NodeQueryReqInfo) error {
	return nil
}
