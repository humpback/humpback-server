package scheduler

import (
	"strings"

	"humpback/internal/db"
	"humpback/types"
)

type ServiceChangeInfo struct {
	ServiceId string
	Action    string
	Version   string
}

// Service管理入口，每个service一个Manager
type ServiceController struct {
	ServiceCtrls        map[string]*ServiceManager
	NodeChangeChan      chan NodeSimpleInfo
	ContainerChangeChan chan types.ContainerStatus
	ContainerRemoveChan chan types.ContainerStatus
	ServiceChangeChan   chan ServiceChangeInfo
}

func NewServiceController(nodeChan chan NodeSimpleInfo, containerChan chan types.ContainerStatus, serviceChan chan ServiceChangeInfo) *ServiceController {
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
func (sc *ServiceController) HandleNodeStatusChanged(nodeInfo NodeSimpleInfo) {
	groupIds := db.GroupGetByNodeId(nodeInfo.NodeId)
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
		if containerStatus.Status != types.ContainerStatusRemoved {
			RemoveNodeContainer(containerStatus.NodeId, containerStatus.ContainerId)
		}
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
