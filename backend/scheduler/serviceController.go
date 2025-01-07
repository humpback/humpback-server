package scheduler

import (
	"slices"
	"strings"

	"humpback/internal/db"
	"humpback/types"
)

// Service管理入口，每个service一个Manager
type ServiceController struct {
	ServiceCtrls        map[string]*ServiceManager
	NodeChangeChan      chan NodeSimpleInfo
	ContainerChangeChan chan types.ContainerStatus
}

func NewServiceController(nodeChan chan NodeSimpleInfo, containerChan chan types.ContainerStatus) *ServiceController {
	sc := &ServiceController{
		ServiceCtrls:        make(map[string]*ServiceManager),
		NodeChangeChan:      nodeChan,
		ContainerChangeChan: containerChan,
	}

	go sc.HandleNodeChanged()
	go sc.HandleContainerChanged()

	return sc
}

// RestoreServiceManager 重启时恢复服务
func (sc *ServiceController) RestoreServiceManager() {
	svcs, err := db.GetDataAll[types.Service](db.BucketServices)
	if err != nil {
		panic(err)
	}

	for _, svc := range svcs {
		if svc.IsEnabled {
			sm := NewServiceManager(svc)
			sc.ServiceCtrls[svc.ServiceId] = sm
		}
	}
}

func (sc *ServiceController) HandleNodeChanged() {
	for nodeInfo := range sc.NodeChangeChan {
		sc.HandleNodeStatusChanged(nodeInfo)
	}
}

// 机器上下线时需要通知该机器所属的Group，去检查Group中所有service的状态
func (sc *ServiceController) HandleNodeStatusChanged(nodeInfo NodeSimpleInfo) {
	groupIds := GetGroupByNodeId(nodeInfo.NodeId)
	for _, gId := range groupIds {
		for _, serviceManager := range sc.ServiceCtrls {
			if serviceManager.ServiceInfo.GroupId == gId {
				go serviceManager.Reconcile()
			}
		}
	}
}

func GetGroupByNodeId(nodeId string) []string {
	groups := make([]string, 0)
	ng, err := db.GetDataByQuery[types.NodesGroups](db.BucketNodesGroups, func(key string, nodesGroups interface{}) bool {
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

func (sc *ServiceController) HandleContainerChanged() {
	for containerStatus := range sc.ContainerChangeChan {
		serviceId := getServiceIdByContainerId(containerStatus.ContainerId)
		if serviceId != "" {
			serviceManager, ok := sc.ServiceCtrls[serviceId]
			if ok {
				go serviceManager.UpdateContainerWhenChanged(containerStatus)
			}
		}
	}
}

func getServiceIdByContainerId(containerId string) string {
	serviceId := ""
	splits := strings.Split(containerId, "-")
	if len(splits) == 4 {
		serviceId = splits[1]
	}
	return serviceId
}
