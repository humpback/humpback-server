package scheduler

import (
	"strings"

	"humpback/internal/db"
	"humpback/internal/node"
	"humpback/types"
)

// Service管理入口，每个service一个Manager
type ServiceController struct {
	ServiceCtrls        map[string]*ServiceManager
	NodeChangeChan      chan types.NodeSimpleInfo
	ContainerChangeChan chan types.ContainerStatus
	ContainerRemoveChan chan types.ContainerStatus
	ServiceChangeChan   chan types.ServiceChangeInfo
}

func NewServiceController(nodeChan chan types.NodeSimpleInfo, containerChan chan types.ContainerStatus, serviceChan chan types.ServiceChangeInfo) *ServiceController {
	sc := &ServiceController{
		ServiceCtrls:        make(map[string]*ServiceManager),
		NodeChangeChan:      nodeChan,
		ContainerChangeChan: containerChan,
		ServiceChangeChan:   serviceChan,
		ContainerRemoveChan: make(chan types.ContainerStatus, 100),
	}

	go sc.HandleNodeChanged()
	go sc.HandleContainerChanged()
	go sc.HandleServiceChange()
	go sc.HandleContainerRemove()

	return sc
}

// RestoreServiceManager 重启时恢复服务
func (sc *ServiceController) RestoreServiceManager() {
	svcs, err := db.ServicesGetAll()
	if err != nil {
		panic(err)
	}

	for _, svc := range svcs {
		if svc.IsEnabled && !svc.IsDelete {
			sm := NewServiceManager(svc)
			sc.ServiceCtrls[svc.ServiceId] = sm
		}
	}
}

func (sc *ServiceController) HandleServiceChange() {
	for serviceInfo := range sc.ServiceChangeChan {
		if serviceManager, ok := sc.ServiceCtrls[serviceInfo.ServiceId]; ok {

			if serviceInfo.Version != serviceManager.ServiceInfo.Version {
				serviceManager.IsNeedCheckAll.Store(true)
			} else if serviceInfo.Action == types.ServiceActionDisable ||
				serviceInfo.Action == types.ServiceActionDelete {
				serviceManager.IsNeedCheckAll.Store(true)
				delete(sc.ServiceCtrls, serviceInfo.ServiceId)
			} else {
				go serviceManager.DoServiceAction(serviceInfo.Action)
			}

		} else {
			svc, err := db.ServiceGetById(serviceInfo.ServiceId)
			if err == nil && svc.IsEnabled && !svc.IsDelete {
				sm := NewServiceManager(svc)
				sc.ServiceCtrls[svc.ServiceId] = sm
			}
		}
	}
}

func (sc *ServiceController) HandleNodeChanged() {
	for nodeInfo := range sc.NodeChangeChan {
		sc.HandleNodeStatusChanged(nodeInfo)
	}
}

// 机器上下线时需要通知该机器所属的Group，去检查Group中所有service的状态
func (sc *ServiceController) HandleNodeStatusChanged(nodeInfo types.NodeSimpleInfo) {
	groups, _ := db.GroupsGetByNodeId(nodeInfo.NodeId)
	for _, g := range groups {
		for _, serviceManager := range sc.ServiceCtrls {
			if serviceManager.ServiceInfo.GroupId == g.GroupId {
				serviceManager.IsNeedCheckAll.Store(true)
			}
		}
	}
}

func (sc *ServiceController) HandleContainerChanged() {
	for containerStatus := range sc.ContainerChangeChan {
		serviceId, version := getServiceIdByContainerId(containerStatus.ContainerName)
		if serviceId != "" {
			serviceManager, ok := sc.ServiceCtrls[serviceId]
			serviceManager.RLock()
			currentVersion := serviceManager.ServiceInfo.Version
			serviceManager.RUnlock()
			if ok && currentVersion == version {
				go serviceManager.UpdateContainerWhenChanged(containerStatus)
			} else {
				sc.ContainerRemoveChan <- containerStatus
			}

		}
	}
}

func (sc *ServiceController) HandleContainerRemove() {
	for containerStatus := range sc.ContainerRemoveChan {
		if containerStatus.State != types.ContainerStatusRemoved {
			node.RemoveNodeContainer(containerStatus.NodeId, containerStatus.ContainerId)
		}
	}
}

func getServiceIdByContainerId(containerName string) (string, string) {
	serviceId := ""
	version := ""
	splits := strings.Split(containerName, "-")
	if len(splits) == 4 && strings.EqualFold(splits[0], "humpback") {
		serviceId = splits[1]
		version = splits[2]
	}
	return serviceId, version
}
