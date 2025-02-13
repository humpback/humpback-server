package handle

import (
	"github.com/gin-gonic/gin"
	"humpback/api/middleware"
)

func RouteNodes(router *gin.RouterGroup) {
	router.POST("", middleware.CheckAdminPermissions(), nodesCreate)
	router.PUT("/labels", middleware.CheckAdminPermissions(), nodeUpdateLabels)
	router.GET("/info/:id", registryGet)
	router.POST("/query", registryQuery)
	router.DELETE("/:id", middleware.CheckAdminPermissions(), nodeDelete)
}

func nodesCreate(c *gin.Context) {

}
