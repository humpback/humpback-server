package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/internal/controller"
)

func RouteActivity(router *gin.RouterGroup) {
	router.POST("/query", activityQuery)
}

func activityQuery(c *gin.Context) {
	body := new(models.ActivityQueryReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	body.UserInfo = middleware.GetUserInfo(c)
	result, err := controller.ActivityQuery(body)
	if err != nil {
		middleware.AbortErr(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
