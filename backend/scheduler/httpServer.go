package scheduler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HumpbackScheduler struct {
	// nodeCtrl    *NodeControllerInter
	// serviceCtrl *ServiceControllerInter
}

func (scheduler *HumpbackScheduler) Start() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
