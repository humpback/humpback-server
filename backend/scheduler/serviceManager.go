package scheduler

import (
	"log/slog"
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
	isNeedQuit     atomic.Value
	isReconcile    atomic.Value
	sync.RWMutex
}

type NodesScore struct {
	NodeId      string
	DeployCount int
	DeployUsage float32
	CPUUsage    float32
	MemoryUsage float32
	Score       float32
}

func NewServiceManager(svc *types.Service) *ServiceManager {
	sm := &ServiceManager{
		ServiceInfo:   svc,
		CheckInterval: int64(config.BackendArgs().ServiceCheckInterval),
	}

	sm.isNeedQuit.Store(false)
	sm.isReconcile.Store(false)
	sm.IsNeedCheckAll.Store(true)

	go sm.CheckService()

	return sm
}

// Reconcile 服务状态变化时，调用该方法，让服务逐步趋于预期状态
func (sm *ServiceManager) Reconcile() {
	sm.Lock()
	defer sm.Unlock()

	sm.isReconcile.Store(true)
	defer func() {
		sm.isReconcile.Store(false)
	}()

	if sm.IsNeedCheckAll.Load().(bool) {
		sm.IsNeedCheckAll.Store(false)
		svc, err := db.GetServiceById(sm.ServiceInfo.ServiceId)
		if err != nil {
			slog.Error("[Service Manager] Get service error", "ServiceId", sm.ServiceInfo.ServiceId, "error", err.Error())
			return
		}

		sm.ServiceInfo = svc
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
		sm.isNeedQuit.Store(true)
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
			slog.Info("[Service Manager] Wait pending container......", "ServiceId", sm.ServiceInfo.ServiceId)
			return
		}

		// 先选一个容器做删除
		if c, ok := sm.TryToDeleteOne(); ok {
			nodeId := c.NodeId
			containerId := c.ContainerId
			err := RemoveNodeContainer(nodeId, containerId)
			if err != nil {
				c.ErrorMsg = err.Error()
			} else {
				// 再选一个节点做调度
				sm.StartNextContainer()
			}
		} else if sm.ServiceInfo.Deployment.Replicas > len(sm.ServiceInfo.Containers) {
			sm.StartNextContainer()
		} else {
			sm.ServiceInfo.Status = types.ServiceStatusRunning
		}
		db.SaveService(sm.ServiceInfo)
	}

}

func (sm *ServiceManager) CheckNodeStatus() {
	groupId := sm.ServiceInfo.GroupId

	isNeedSave := false
	nodes, totalNodes, err := db.GetOfflineNodesByGroupId(groupId)

	if err != nil {
		slog.Error("[Service Manager] Get offline nodes error", "ServiceId", sm.ServiceInfo.ServiceId, "error", err.Error())
		return
	}

	if sm.ServiceInfo.Deployment.Mode == types.DeployModeGlobal {
		sm.ServiceInfo.Deployment.Replicas = totalNodes
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
		if isContainerFailed(c.Status) || isContainerRemoved(c.Status) {
			return c, true
		}
	}

	return nil, false

}

func (sm *ServiceManager) StartNextContainer() {

	groupId := sm.ServiceInfo.GroupId
	nodes, err := db.GetOnlineNodesByGroupId(groupId)

	if err != nil {
		slog.Error("[Service Manager] Start Service error", "ServiceId", sm.ServiceInfo.ServiceId, "error", err.Error())
		return
	}

	if len(nodes) == 0 {
		slog.Error("[Service Manager] Start Service error: No available nodes", "ServiceId", sm.ServiceInfo.ServiceId)
		return
	}

	nodeId := sm.ChooseNextNodes(nodes)

	if nodeId == "" {
		slog.Error("[Service Manager] Start Service error: No available nodes", "ServiceId", sm.ServiceInfo.ServiceId)
		return
	}

	cerr := StartNewContainer(nodeId, sm.ServiceInfo)
	if cerr != nil {
		slog.Error("[Service Manager] Start New Container error", "ServiceId", sm.ServiceInfo.ServiceId, "error", cerr.Error())
		return
	}

	db.SaveService(sm.ServiceInfo)

}

func (sm *ServiceManager) ChooseNextNodes(nodes []*types.Node) (nodeId string) {

	if sm.ServiceInfo.Deployment.Mode == types.DeployModeGlobal {

		deployedNodes := lo.Map(sm.ServiceInfo.Containers, func(c *types.ContainerStatus, index int) string {
			return c.NodeId
		})

		for _, n := range nodes {
			if !slices.Contains(deployedNodes, n.NodeId) {
				nodeId = n.NodeId
				break
			}
		}

	} else {
		totalReplicas := sm.ServiceInfo.Deployment.Replicas

		nodeUsage := make(map[string]*NodesScore)
		for _, n := range nodes {
			nu := &NodesScore{
				NodeId:      n.NodeId,
				CPUUsage:    n.CPUUsage,
				MemoryUsage: n.MemoryUsage,
				DeployCount: 0,
			}
			nodeUsage[n.NodeId] = nu
		}

		for _, c := range sm.ServiceInfo.Containers {
			if n, ok := nodeUsage[c.NodeId]; ok {
				n.DeployCount++
			}
		}

		var maxScore float32

		for nId, nu := range nodeUsage {
			nu.DeployUsage = float32(nu.DeployCount) / float32(totalReplicas)
			nu.Score = (1-nu.CPUUsage)*100*0.3 + (1-nu.MemoryUsage)*100*0.2 + (1-nu.DeployUsage)*100*0.5
			if nu.Score > maxScore {
				maxScore = nu.Score
				nodeId = nId
			}
		}
	}
	return
}

// UpdateContainerWhenChanged 如果容器状态有变化，就保存DB
// 然后等定时检查起来后，重新Reconcile Service
func (sm *ServiceManager) UpdateContainerWhenChanged(cs types.ContainerStatus) {

	sm.Lock()
	defer sm.Unlock()

	ct, ok := lo.Find(sm.ServiceInfo.Containers, func(c *types.ContainerStatus) bool {
		return c.ContainerName == cs.ContainerName
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
		slog.Info("[Service Manager] Check service......", "ServiceId", sm.ServiceInfo.ServiceId)
		if sm.isNeedQuit.Load().(bool) {
			ticker.Stop()
			slog.Info("[Service Manager] Service is disabled", "ServiceId", sm.ServiceInfo.ServiceId)
			return
		}
		if !sm.isReconcile.Load().(bool) {
			sm.Reconcile()
		} else {
			slog.Info("[Service Manager] Service is busy", "ServiceId", sm.ServiceInfo.ServiceId)
		}
	}
}
