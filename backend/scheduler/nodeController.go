package scheduler

import (
	"log"
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
}

type NodeController struct {
	NodesInfo           map[string]*NodeSimpleInfo
	NodeHeartbeatChan   chan NodeSimpleInfo
	ContainerChangeChan chan types.ContainerStatus
	CheckInterval       int64
	CheckThreshold      int
	sync.RWMutex
}

func NewNodeController(nodeChan chan NodeSimpleInfo, containerChan chan types.ContainerStatus) *NodeController {
	nc := &NodeController{
		NodesInfo:           make(map[string]*NodeSimpleInfo),
		NodeHeartbeatChan:   nodeChan,
		ContainerChangeChan: containerChan,
		CheckInterval:       int64(config.BackendArgs().CheckInterval),
		CheckThreshold:      config.BackendArgs().CheckThreshold,
	}

	go nc.CheckNodes()

	return nc
}

func (nc *NodeController) CheckNodes() {
	interval := nc.CheckInterval
	time.Sleep(time.Duration(rand.Int31n(int32(interval))) * time.Second)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	for range ticker.C {
		log.Println("check nodes......")
		nc.CheckNodesCore()
	}
}

// 机器上下线时需要通知该机器所属的Group，去检查Group中所有service的状态
func (nc *NodeController) CheckNodesCore() {
	nc.Lock()
	defer nc.Unlock()

	currentTime := time.Now().Unix()
	isNeedSave := false
	for nodeId, nodeInfo := range nc.NodesInfo {

		if nodeInfo.Status == types.NodeStatusOnline {
			if currentTime-nodeInfo.LastHeartbeat > nc.CheckInterval {
				log.Printf("Node %s is not responding. Last heartbeat: %d", nodeId, nodeInfo.LastHeartbeat)

				nodeInfo.Status = types.NodeStatusOffline
				isNeedSave = true
				nc.NodeHeartbeatChan <- *nodeInfo
			}
		}

		if nodeInfo.Status == types.NodeStatusOffline {
			if currentTime-nodeInfo.LastHeartbeat < nc.CheckInterval &&
				nodeInfo.OnlineThreshold >= nc.CheckThreshold {
				log.Printf("need report online node [%s]", nodeId)

				nodeInfo.Status = types.NodeStatusOnline
				isNeedSave = true
				nc.NodeHeartbeatChan <- *nodeInfo
			}
		}

		if isNeedSave {
			err := db.UpdateNodeStatus(nodeId, nodeInfo.Status, nodeInfo.LastHeartbeat)
			if err != nil {
				log.Printf("update node status failed: %s", err)
			}
		}
	}
}

func (nc *NodeController) HeartBeat(healthInfo types.HealthInfo) {
	nc.Lock()
	defer nc.Unlock()

	nodeId := healthInfo.NodeId
	ts := time.Now().Unix()
	if n, ok := nc.NodesInfo[nodeId]; ok {
		if n.Status == types.NodeStatusOffline && ts-n.LastHeartbeat < nc.CheckInterval {
			n.OnlineThreshold++
		} else {
			n.OnlineThreshold = 1
			nc.CheckContainers(healthInfo)
		}
		n.LastHeartbeat = ts
	} else {
		n, err := db.GetNodeById(nodeId)
		if err == nil {
			nc.NodesInfo[nodeId] = &NodeSimpleInfo{
				NodeId:          n.NodeID,
				IpAddress:       n.IpAddress,
				Status:          types.NodeStatusOffline,
				LastHeartbeat:   ts,
				OnlineThreshold: 1,
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
