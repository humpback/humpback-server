package handle

import (
	"github.com/gin-gonic/gin"
)

func RouteTeam(router *gin.RouterGroup) {
	router.POST("", teamCreate)
	router.PUT("", teamUpdate)
	router.GET("/:id", team)
	router.POST("/query", teams)
	router.GET("/query-by-user/:userId", teamsByUserId)
	router.DELETE("/:id", teamDelete)
}

func teamCreate(c *gin.Context) {

}

func teamUpdate(c *gin.Context) {

}

func team(c *gin.Context) {

}

func teams(c *gin.Context) {

}

func teamsByUserId(c *gin.Context) {

}

func teamDelete(c *gin.Context) {

}
