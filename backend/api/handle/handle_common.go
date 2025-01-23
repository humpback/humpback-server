package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"humpback/common/enum"
)

func RouteCommon(router *gin.RouterGroup) {
	router.GET("/config", ruleConfig)
}

func ruleConfig(c *gin.Context) {
	data := map[string]any{
		"RuleLengthLimit": enum.RuleLength,
		"RuleFormat":      enum.RuleFormat,
		"RsaPublicKey":    enum.PublicKey,
	}
	c.JSON(http.StatusOK, data)
}
