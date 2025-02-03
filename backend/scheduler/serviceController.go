package scheduler

import (
	"strings"

	"humpback/internal/db"
	"humpback/types"
)

// Service管理入口，每个service一个Manager
type ServiceController struct {
	ServiceCtrls        map[string]*ServiceManager
	NodeChangeChan      chan NodeSimpleInfo
	ContainerChangeChan chan types.ContainerStatus
	ContainerRemoveChan chan types.ContainerStatus
	ServiceChangeChan   chan string
}

func NewServiceController(nodeChan chan NodeSimpleInfo, containerChan chan types.ContainerStatus, serviceChan chan string) *ServiceController {
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
	svcs, err := db.GetAllService()
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

func (sc *ServiceController) HandleServiceChange() {
	for serviceId := range sc.ServiceChangeChan {
		if serviceManager, ok := sc.ServiceCtrls[serviceId]; ok {
			serviceManager.IsNeedCheckAll.Store(true)
		} else {
			svc, err := db.GetServiceById(serviceId)
			if err == nil && svc.IsEnabled {
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
func (sc *ServiceController) HandleNodeStatusChanged(nodeInfo NodeSimpleInfo) {
	groupIds := db.GetGroupByNodeId(nodeInfo.NodeId)
	for _, gId := range groupIds {
		for _, serviceManager := range sc.ServiceCtrls {
			if serviceManager.ServiceInfo.GroupId == gId {
				serviceManager.IsNeedCheckAll.Store(true)
			}
		}
	}
}

func (sc *ServiceController) HandleContainerChanged() {
	for containerStatus := range sc.ContainerChangeChan {
		serviceId := getServiceIdByContainerId(containerStatus.ContainerName)
		if serviceId != "" {
			serviceManager, ok := sc.ServiceCtrls[serviceId]
			if ok {
				go serviceManager.UpdateContainerWhenChanged(containerStatus)
			} else {
				sc.ContainerRemoveChan <- containerStatus
			}
		}
	}
}

func (sc *ServiceController) HandleContainerRemove() {
	for containerStatus := range sc.ContainerRemoveChan {
		RemoveNodeContainer(containerStatus.NodeId, containerStatus.ContainerName)
	}
}

func getServiceIdByContainerId(containerName string) string {
	serviceId := ""
	splits := strings.Split(containerName, "-")
	if len(splits) == 4 {
		serviceId = splits[1]
	}
	return serviceId
}
