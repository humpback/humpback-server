package middleware

import (
	"github.com/gin-gonic/gin"
	"humpback/common/locales"
	"humpback/common/response"
)

type CheckInterface interface {
	Check() error
}

func BindAndCheckBody(c *gin.Context, obj CheckInterface) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		AbortErr(c, response.NewBadRequestErr(locales.CodeRequestParamsInvalid, err.Error()))
		return false
	}
	if err := obj.Check(); err != nil {
		AbortErr(c, err)
		return false
	}
	return true
}

func BindBody(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		AbortErr(c, response.NewBadRequestErr(locales.CodeRequestParamsInvalid, err.Error()))
		return false
	}
	return true
}

func CheckBody(c *gin.Context, body CheckInterface) bool {
	if err := body.Check(); err != nil {
		AbortErr(c, err)
		return false
	}
	return true
}

func AbortErr(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}
