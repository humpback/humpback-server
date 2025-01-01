package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/api/handle/models"
	"humpback/api/middleware"
	"humpback/common/response"
)

func RouteUser(router *gin.RouterGroup) {
	router.POST("/login", login)
}

func login(c *gin.Context) {
	body := new(models.UserLoginReqInfo)
	if !middleware.BindAndCheckBody(c, body) {
		return
	}
	c.JSON(http.StatusOK, response.NewRespSucceed())
}
