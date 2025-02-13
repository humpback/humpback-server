package models

import (
	"regexp"
	"time"

	"humpback/common/enum"
	"humpback/common/locales"
	"humpback/common/response"
	"humpback/common/verify"
	"humpback/pkg/utils"
	"humpback/types"
)

type NodesCreateReqInfo []string

func (n *NodesCreateReqInfo) Check() error {
	if len(*n) == 0 {
		return response.NewBadRequestErr(locales.CodeNodesNotEmpty)
	}
	ipRule := regexp.MustCompile(enum.RegularIpAddress)
	for _, ipAddress := range *n {
		if ipRule.MatchString(ipAddress) {
			return response.NewBadRequestErr(locales.CodeNodesIpAddressInvalid)
		}
	}
	return nil
}

func (n *NodesCreateReqInfo) NewNodeInfo() []*types.Node {
	result := make([]*types.Node, 0)
	nowT := time.Now().UnixMilli()
	for _, ip := range *n {
		result = append(result, &types.Node{
			NodeId:      utils.NewGuidStr(),
			Name:        "",
			IpAddress:   ip,
			Port:        8677,
			Status:      types.NodeStatusOffline,
			IsEnable:    true,
			CreatedAt:   nowT,
			UpdatedAt:   nowT,
			CPUUsage:    0,
			CPU:         0,
			MemoryUsage: 0,
			MemoryTotal: 0,
			MemoryUsed:  0,
			Labels:      make(map[string]string),
		})
	}
	return result
}

type NodeUpdateLabelReqInfo struct {
	NodeId string            `json:"nodeId"`
	Labels map[string]string `json:"labels"`
}

func (n *NodeUpdateLabelReqInfo) Check() error {
	if err := verify.CheckIsEmpty(n.NodeId, locales.CodeNodesIdNotEmpty); err != nil {
		return err
	}
	//todo 检查labels
	return nil
}
