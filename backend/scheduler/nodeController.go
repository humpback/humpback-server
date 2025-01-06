package scheduler

import (
	"humpback/config"
	"humpback/internal/db"
	"humpback/types"
	"log"
	"sync"
	"time"

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

func (nc *NodeController) CheckNodesCore() {
	nc.Lock()
	defer nc.Unlock()

	currentTime := time.Now().Unix()

	//Node 状态有变化时需要重新保存DB
	for nodeId, nodeInfo := range nc.NodesInfo {

		if nodeInfo.Status == types.NodeStatusOnline {
			if currentTime-nodeInfo.LastHeartbeat > nc.CheckInterval {
				log.Printf("Node %s is not responding. Last heartbeat: %d", nodeId, nodeInfo.LastHeartbeat)

				nodeInfo.Status = types.NodeStatusOffline
				nc.NodeHeartbeatChan <- *nodeInfo
			}
		}

		if nodeInfo.Status == types.NodeStatusOffline {
			if currentTime-nodeInfo.LastHeartbeat < nc.CheckInterval &&
				nodeInfo.OnlineThreshold >= nc.CheckThreshold {
				log.Printf("need report online node [%s]", nodeId)

				nodeInfo.Status = types.NodeStatusOnline
				nc.NodeHeartbeatChan <- *nodeInfo
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
		n, err := db.GetDataById[types.Node](db.BucketNodes, nodeId)
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

func (nc *NodeController) CheckContainers(healthInfo types.HealthInfo) {
	for _, container := range healthInfo.ContainerList {
		container.NodeId = healthInfo.NodeId
		nc.ContainerChangeChan <- container
	}
}
