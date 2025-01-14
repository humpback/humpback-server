package scheduler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"humpback/config"
	"humpback/types"

	"github.com/gin-gonic/gin"
)

type HumpbackScheduler struct {
	httpSrv             *http.Server
	nodeCtrl            *NodeController
	serviceCtrl         *ServiceController
	NodeHeartbeatChan   chan NodeSimpleInfo
	ContainerChangeChan chan types.ContainerStatus
	ServiceChangeChan   chan string
}

func NewHumpbackScheduler() *HumpbackScheduler {
	hs := &HumpbackScheduler{}
	hs.NodeHeartbeatChan = make(chan NodeSimpleInfo, 100)
	hs.ContainerChangeChan = make(chan types.ContainerStatus, 100)
	hs.ServiceChangeChan = make(chan string, 100)
	hs.serviceCtrl = NewServiceController(hs.NodeHeartbeatChan, hs.ContainerChangeChan, hs.ServiceChangeChan)
	hs.nodeCtrl = NewNodeController(hs.NodeHeartbeatChan, hs.ContainerChangeChan)

	NewCacheManager()

	return hs
}

func doHealth(c *gin.Context) {
	payload := types.HealthInfo{}

	if c.BindJSON(&payload) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	nodeId := MatchNodeWithIpAddress(payload.HostInfo.IpAddress)
	if nodeId == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "node not found"})
		return
	}
	payload.NodeId = nodeId

	sc := c.MustGet("scheduler").(*HumpbackScheduler)
	sc.nodeCtrl.HeartBeat(payload)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (scheduler *HumpbackScheduler) Start() {
	scheduler.serviceCtrl.RestoreServiceManager()
	go func() {
		e := gin.Default()

		e.Use(func(c *gin.Context) {
			c.Set("scheduler", scheduler)
			c.Next()
		})

		e.POST("/health", doHealth)

		e.GET("/mock/nodes", mockNodes)

		e.GET("/nodes", getAllNodes)

		e.GET("/services", getAllServices)

		e.GET("/mock/service/gateway", mockGatewayServices)

		listeningAddress := fmt.Sprintf("%s:%s", config.NodeArgs().HostIp, config.BackendArgs().BackendPort)
		slog.Info("[Scheduler] listening...", "Address", listeningAddress)
		scheduler.httpSrv = &http.Server{
			Addr:    listeningAddress,
			Handler: e,
		}
		if err := scheduler.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("[Scheduler] start failed", "Address", listeningAddress, "error", err)
		}
	}()
}

func (scheduler *HumpbackScheduler) Close(c context.Context) error {
	return scheduler.httpSrv.Shutdown(c)
}
