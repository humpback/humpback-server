package scheduler

import (
	"log"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	"humpback/config"
	"humpback/internal/db"
	"humpback/types"

	"github.com/samber/lo"
	"golang.org/x/exp/rand"
)

type ServiceManager struct {
	ServiceInfo    *types.Service
	CheckInterval  int64
	IsNeedCheckAll atomic.Value
	isNeedQuit     bool
	isReconcile    bool
	sync.RWMutex
}

func NewServiceManager(svc *types.Service) *ServiceManager {
	sm := &ServiceManager{
		ServiceInfo:   svc,
		CheckInterval: int64(config.BackendArgs().ServiceCheckInterval),
	}

	sm.IsNeedCheckAll.Store(true)

	go sm.CheckService()

	return sm
}

// Reconcile 服务状态变化时，调用该方法，让服务逐步趋于预期状态
func (sm *ServiceManager) Reconcile() {
	sm.Lock()
	defer sm.Unlock()

	sm.isReconcile = true
	defer func() {
		sm.isReconcile = false
	}()

	if sm.IsNeedCheckAll.Load().(bool) {
		sm.IsNeedCheckAll.Store(false)
		if svc, err := db.GetServiceById(sm.ServiceInfo.ServiceId); err == nil {
			sm.ServiceInfo = svc
		}

		sm.CheckNodeStatus()
	}

	// service被disabled后就删除全部容器
	if !sm.ServiceInfo.IsEnabled {
		for _, c := range sm.ServiceInfo.Containers {
			nodeId := c.NodeId
			containerId := c.ContainerId
			err := RemoveNodeContainer(nodeId, containerId)
			if err != nil {
				c.ErrorMsg = err.Error()
			}
			db.SaveService(sm.ServiceInfo)
		}
		sm.isNeedQuit = true
		return
	}

	// 所有容器都是正常的，就不需要再做任何操作
	if sm.ServiceInfo.Status == types.ServiceStatusRunning && sm.IsContainerAllReady() {
		return
	}

	// 服务状态不正常，就开始尝试调度
	if sm.ServiceInfo.Status == types.ServiceStatusNotReady {

		// 如果有容器正在启动，就不再继续
		if sm.HasPendingContainer() {
			return
		}

		// 先选一个容器做删除
		if c, ok := sm.TryToDeleteOne(); ok {
			nodeId := c.NodeId
			containerId := c.ContainerId
			err := RemoveNodeContainer(nodeId, containerId)
			if err != nil {
				c.ErrorMsg = err.Error()
				db.SaveService(sm.ServiceInfo)
				return
			}

			// 再选一个节点做调度
			sm.StartNextContainer(c)
		}
	}

}

func (sm *ServiceManager) CheckNodeStatus() {
	groupId := sm.ServiceInfo.GroupId

	isNeedSave := false
	nodes, err := db.GetOfflineNodesByGroupId(groupId)

	if err != nil {
		log.Printf("[Service Manager] Check Service [%s] offline nodes error: %s", sm.ServiceInfo.ServiceId, err.Error())
		return
	}

	for _, c := range sm.ServiceInfo.Containers {
		if slices.Contains(nodes, c.NodeId) {
			isNeedSave = true
			c.Status = types.ContainerStatusWarning
			c.ErrorMsg = "Node is offline"
		}
	}

	if isNeedSave {
		db.SaveService(sm.ServiceInfo)
	}
}

func (sm *ServiceManager) IsContainerAllReady() bool {
	result := true
	for _, c := range sm.ServiceInfo.Containers {
		version := parseVersionByContainerId(c.ContainerId)
		if version == sm.ServiceInfo.Version {
			if isContainerExited(c.Status) && sm.ServiceInfo.Deployment.Type == types.DeployTypeSchedule {
				continue
			}
			if isContainerRunning(c.Status) {
				continue
			}
		}

		result = false
		break
	}
	if !result {
		sm.ServiceInfo.Status = types.ServiceStatusNotReady
	}
	return result
}

func (sm *ServiceManager) HasPendingContainer() bool {
	for _, c := range sm.ServiceInfo.Containers {
		if isContainerStarting(c.Status) {
			return true
		}
	}
	return false
}

func (sm *ServiceManager) TryToDeleteOne() (*types.ContainerStatus, bool) {

	for _, c := range sm.ServiceInfo.Containers {
		version := parseVersionByContainerId(c.ContainerId)
		if version != sm.ServiceInfo.Version {
			return c, true
		}
		if isContainerExited(c.Status) {
			if sm.ServiceInfo.Deployment.Type == types.DeployTypeBackground {
				return c, true
			}
		}
		if isContainerFailed(c.Status) {
			return c, true
		}
	}

	return nil, false

}

func (sm *ServiceManager) StartNextContainer(rmc *types.ContainerStatus) {

}

// UpdateContainerWhenChanged 如果容器状态有变化，就保存DB
// 然后等定时检查起来后，重新Reconcile Service
func (sm *ServiceManager) UpdateContainerWhenChanged(cs types.ContainerStatus) {

	sm.Lock()
	defer sm.Unlock()

	ct, ok := lo.Find(sm.ServiceInfo.Containers, func(c *types.ContainerStatus) bool {
		return c.ContainerId == cs.ContainerId
	})

	if ok && (ct.Status != cs.Status || ct.StartAt != cs.StartAt) {
		ct.Status = cs.Status
		ct.StartAt = cs.StartAt
		db.SaveService(sm.ServiceInfo)
	}

	if !ok {
		sm.ServiceInfo.Containers = append(sm.ServiceInfo.Containers, &cs)
		db.SaveService(sm.ServiceInfo)
	}

}

// 定时检查服务状态，看是否满足预期
func (sm *ServiceManager) CheckService() {
	interval := sm.CheckInterval
	time.Sleep(time.Duration(rand.Int31n(int32(interval))) * time.Second)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	for range ticker.C {
		log.Printf("check service [%s]......", sm.ServiceInfo.ServiceId)
		if sm.isNeedQuit {
			ticker.Stop()
			log.Printf("service [%s] is disabled", sm.ServiceInfo.ServiceId)
			return
		}
		if !sm.isReconcile {
			sm.Reconcile()
		} else {
			log.Printf("service [%s] is busy", sm.ServiceInfo.ServiceId)
		}
	}
}
