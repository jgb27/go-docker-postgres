package routes

import (
	controllers "go-docker-postgres/api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) *gin.RouterGroup {
	routes := router.Group("/tweet")
	{
		routes.GET("/", controllers.FindAll)
		routes.POST("/", controllers.Create)
		routes.DELETE("/:id", controllers.Delete)
	}

	return routes
}
