package scheduler

import (
	"context"
	"fmt"
	"humpback/config"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HumpbackScheduler struct {
	engine  *gin.Engine
	httpSrv *http.Server
}

func NewHumpbackScheduler() *HumpbackScheduler {
	return &HumpbackScheduler{}
}

func (scheduler *HumpbackScheduler) Start() {
	go func() {
		scheduler.engine = gin.Default()

		scheduler.engine.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		listeningAddress := fmt.Sprintf("%s:%s", config.NodeArgs().HostIp, config.NodeArgs().BackendPort)
		slog.Info("[Api] listening...", "Address", listeningAddress)
		scheduler.httpSrv = &http.Server{
			Addr:    listeningAddress,
			Handler: scheduler.engine,
		}
		if err := scheduler.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(fmt.Sprintf("listening %s failed: %s", listeningAddress, err))
		}
	}()
}

func (scheduler *HumpbackScheduler) Close(c context.Context) error {
	return scheduler.httpSrv.Shutdown(c)
}
