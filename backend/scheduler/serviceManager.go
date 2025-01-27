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
	ServiceInfo      *types.Service
	availableNodes   []string
	unavailableNodes []string
	CheckInterval    int64
	IsNeedCheckAll   atomic.Value
	isNeedQuit       atomic.Value
	isReconcile      atomic.Value
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

	if sm.ServiceInfo.Status == types.ServiceStatusRunning &&
		sm.ServiceInfo.Deployment.Replicas != len(sm.ServiceInfo.Containers) {

		slog.Info("[Service Manager] Service change status to NotReady......", "ServiceId", sm.ServiceInfo.ServiceId)
		sm.ServiceInfo.Status = types.ServiceStatusNotReady
	}

	// service被disabled后就删除全部容器
	if !sm.ServiceInfo.IsEnabled {
		for _, c := range sm.ServiceInfo.Containers {
			nodeId := c.NodeId
			containerName := c.ContainerName
			err := sm.DeleteContainer(nodeId, containerName)
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
		slog.Info("[Service Manager] Service is running ok......", "ServiceId", sm.ServiceInfo.ServiceId)
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
			containerName := c.ContainerName
			slog.Info("[Service Manager] Remove un-need container......", "ServiceId", sm.ServiceInfo.ServiceId, "ContainerName", c.ContainerName)
			err := sm.DeleteContainer(nodeId, containerName)
			if err != nil {
				c.ErrorMsg = err.Error()
			}
			slog.Info("[Service Manager] After remove container......", "ServiceId", sm.ServiceInfo.ServiceId, "Container Count", len(sm.ServiceInfo.Containers))
		}

		if sm.ServiceInfo.Deployment.Replicas > len(sm.ServiceInfo.Containers) {
			sm.StartNextContainer()
		} else {
			slog.Info("[Service Manager] Service change status to running......", "ServiceId", sm.ServiceInfo.ServiceId)
			sm.ServiceInfo.Status = types.ServiceStatusRunning
		}

		db.SaveService(sm.ServiceInfo)
	}

}

func (sm *ServiceManager) DeleteContainer(nodeId string, containerName string) error {
	if err := RemoveNodeContainer(nodeId, containerName); err != nil {
		return err
	}

	sm.ServiceInfo.Containers = lo.Filter(sm.ServiceInfo.Containers, func(cs *types.ContainerStatus, index int) bool {
		return cs.ContainerName != containerName
	})
	return nil
}

func (sm *ServiceManager) CheckNodeStatus() {
	groupId := sm.ServiceInfo.GroupId

	isNeedSave := false
	nodes, err := db.GetNodesByGroupId(groupId)

	if err != nil {
		slog.Error("[Service Manager] Get offline nodes error", "ServiceId", sm.ServiceInfo.ServiceId, "error", err.Error())
		return
	}

	sm.GetMatchedNodes(nodes)

	if sm.ServiceInfo.Deployment.Mode == types.DeployModeGlobal {
		sm.ServiceInfo.Deployment.Replicas = len(sm.availableNodes)
	}

	for _, c := range sm.ServiceInfo.Containers {
		if slices.Contains(sm.unavailableNodes, c.NodeId) {
			isNeedSave = true
			c.Status = types.ContainerStatusWarning
			c.ErrorMsg = "Node is offline"
		}
	}

	if isNeedSave {
		db.SaveService(sm.ServiceInfo)
	}
}

func (sm *ServiceManager) GetMatchedNodes(nodes []*types.Node) {
	sm.availableNodes = make([]string, 0)
	sm.unavailableNodes = make([]string, 0)

	for _, n := range nodes {
		if sm.ServiceInfo.Deployment.Placements != nil {
			for _, p := range sm.ServiceInfo.Deployment.Placements {
				if isPlacementMatched(n, p) && n.Status == types.NodeStatusOnline {
					sm.availableNodes = append(sm.availableNodes, n.NodeId)
				} else {
					sm.unavailableNodes = append(sm.unavailableNodes, n.NodeId)
				}
			}
		} else {
			if n.Status == types.NodeStatusOnline {
				sm.availableNodes = append(sm.availableNodes, n.NodeId)
			} else {
				sm.unavailableNodes = append(sm.unavailableNodes, n.NodeId)
			}
		}
	}
}

func (sm *ServiceManager) IsContainerAllReady() bool {
	result := true
	for _, c := range sm.ServiceInfo.Containers {
		version := parseVersionByContainerId(c.ContainerName)
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
		slog.Info("[Service Manager] Service change status to NotReady......", "ServiceId", sm.ServiceInfo.ServiceId)
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

	nodeDeployed := make(map[string]bool)

	for _, c := range sm.ServiceInfo.Containers {
		version := parseVersionByContainerId(c.ContainerName)
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

		if sm.ServiceInfo.Deployment.Mode == types.DeployModeGlobal {
			if _, ok := nodeDeployed[c.NodeId]; ok {
				return c, true
			} else {
				nodeDeployed[c.NodeId] = true
			}
		}
	}

	return nil, false

}

func (sm *ServiceManager) StartNextContainer() {

	nodes, err := db.GetNodesByIds(sm.availableNodes)

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
				slog.Info("[Service Manager] Choose Node for new container in global deployment......", "ServiceId", sm.ServiceInfo.ServiceId, "NodeId", nodeId)
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
		slog.Info("[Service Manager] Choose Node for new container in replicas deployment......", "ServiceId", sm.ServiceInfo.ServiceId, "NodeId", nodeId)
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
		ct.ContainerId = cs.ContainerId
		if ct.Status == types.ContainerStatusRunning {
			ct.ErrorMsg = ""
		}
		db.SaveService(sm.ServiceInfo)

		slog.Info("[Service Manager] Container status changed......", "ServiceId", sm.ServiceInfo.ServiceId, "ContainerName", ct.ContainerName)
	}

	if !ok {
		sm.ServiceInfo.Containers = append(sm.ServiceInfo.Containers, &cs)
		slog.Info("[Service Manager] New container found......", "ServiceId", sm.ServiceInfo.ServiceId, "ContainerName", cs.ContainerName)
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
