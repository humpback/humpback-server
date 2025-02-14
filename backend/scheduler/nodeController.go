package scheduler

import (
	"log/slog"
	"sync"
	"time"
	
	"humpback/config"
	"humpback/internal/db"
	"humpback/types"
	
	"golang.org/x/exp/rand"
)

type NodeSimpleInfo struct {
	NodeId          string
	IpAddress       string
	Status          string
	LastHeartbeat   int64
	OnlineThreshold int
	CPUUsage        float32
	MemoryUsage     float32
}

type NodeController struct {
	NodesInfo           map[string]*NodeSimpleInfo
	NodeHeartbeatChan   chan NodeSimpleInfo
	ContainerChangeChan chan types.ContainerStatus
	CheckInterval       int64
	CheckThreshold      int
	ThresholdInvterval  int64
	sync.RWMutex
}

func NewNodeController(nodeChan chan NodeSimpleInfo, containerChan chan types.ContainerStatus) *NodeController {
	nc := &NodeController{
		NodesInfo:           make(map[string]*NodeSimpleInfo),
		NodeHeartbeatChan:   nodeChan,
		ContainerChangeChan: containerChan,
		CheckInterval:       int64(config.BackendArgs().CheckInterval),
		CheckThreshold:      config.BackendArgs().CheckThreshold,
		ThresholdInvterval:  int64(config.BackendArgs().CheckInterval) * int64(config.BackendArgs().CheckThreshold),
	}

	go nc.CheckNodes()

	return nc
}

func (nc *NodeController) RestoreNodes() {
	nc.Lock()
	defer nc.Unlock()

	nodes, err := db.NodesGetAllEnabled()
	if err != nil {
		slog.Info("[Node Controller] restore nodes failed", "error", err)
		return
	}

	for _, node := range nodes {
		nc.NodesInfo[node.NodeId] = &NodeSimpleInfo{
			NodeId:          node.NodeId,
			IpAddress:       node.IpAddress,
			Status:          node.Status,
			LastHeartbeat:   node.UpdatedAt,
			OnlineThreshold: 1,
			CPUUsage:        node.CPUUsage,
			MemoryUsage:     node.MemoryUsage,
		}
	}
}

func (nc *NodeController) CheckNodes() {
	interval := nc.CheckInterval
	time.Sleep(time.Duration(rand.Int31n(int32(interval))) * time.Second)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	for range ticker.C {
		slog.Info("[Node Controller] Check nodes......")
		nc.CheckNodesCore()
	}
}

// 机器上下线时需要通知该机器所属的Group，去检查Group中所有service的状态
func (nc *NodeController) CheckNodesCore() {
	nc.Lock()
	defer nc.Unlock()

	currentTime := time.Now().Unix()
	for nodeId, nodeInfo := range nc.NodesInfo {

		if nodeInfo.Status == types.NodeStatusOnline {
			if currentTime-nodeInfo.LastHeartbeat > nc.ThresholdInvterval {
				slog.Info("[Node Controller] Node is not responding.", "nodeId", nodeId, "Last heartbeat", nodeInfo.LastHeartbeat)
				nodeInfo.Status = types.NodeStatusOffline
				nc.NodeHeartbeatChan <- *nodeInfo
			}
		}

		if nodeInfo.Status == types.NodeStatusOffline {
			if currentTime-nodeInfo.LastHeartbeat < nc.ThresholdInvterval &&
				nodeInfo.OnlineThreshold >= nc.CheckThreshold {
				slog.Info("[Node Controller] need report online node", "nodeId", nodeId)
				nodeInfo.Status = types.NodeStatusOnline
				nc.NodeHeartbeatChan <- *nodeInfo
			}
		}

		err := db.NodeUpdateStatus(nodeId, nodeInfo.Status, nodeInfo.LastHeartbeat, nodeInfo.CPUUsage, nodeInfo.MemoryUsage)
		if err != nil {
			slog.Info("[Node Controller] update node status to DB failed", "error", err)
		}
	}
}

func (nc *NodeController) HeartBeat(healthInfo types.HealthInfo) {
	nc.Lock()
	defer nc.Unlock()

	nodeId := healthInfo.NodeId
	ts := time.Now().Unix()
	if n, ok := nc.NodesInfo[nodeId]; ok {
		n.CPUUsage = healthInfo.HostInfo.CPUUsage
		n.MemoryUsage = healthInfo.HostInfo.MemoryUsage
		if n.Status == types.NodeStatusOffline && ts-n.LastHeartbeat < nc.ThresholdInvterval {
			n.OnlineThreshold++
		} else {
			n.OnlineThreshold = 1
			nc.CheckContainers(healthInfo)
		}
		n.LastHeartbeat = ts
	} else {
		n, err := db.NodeGetById(nodeId)
		if err == nil {
			nc.NodesInfo[nodeId] = &NodeSimpleInfo{
				NodeId:          n.NodeId,
				IpAddress:       n.IpAddress,
				Status:          types.NodeStatusOffline,
				LastHeartbeat:   ts,
				OnlineThreshold: 1,
				CPUUsage:        healthInfo.HostInfo.CPUUsage,
				MemoryUsage:     healthInfo.HostInfo.MemoryUsage,
			}
		}
	}
}

// 对于在线的机器，检查容器状态
func (nc *NodeController) CheckContainers(healthInfo types.HealthInfo) {
	for _, container := range healthInfo.ContainerList {
		container.NodeId = healthInfo.NodeId
		nc.ContainerChangeChan <- container
	}
}
