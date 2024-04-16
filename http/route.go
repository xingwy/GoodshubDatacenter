package http

import "github.com/gin-gonic/gin"

// Init 路由
func Reginster(instance *gin.Engine) *gin.Engine {
	api := instance.Group("/sh-goodshub-datacenter")

	debug := api.Group("/debug")
	{
		debug.POST("/test", Test)
	}

	return instance
}
