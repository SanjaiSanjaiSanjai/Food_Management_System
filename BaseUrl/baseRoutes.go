package baseurl

import (
	"github.com/gin-gonic/gin"
)

func ApiGroup(routes *gin.Engine) *gin.RouterGroup {
	return routes.Group("/v1/api")
}
