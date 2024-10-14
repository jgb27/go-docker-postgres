package routes

import (
	controllers "go-docker-postgres/api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) *gin.RouterGroup {

	routes := router.Group("/tweet")
	{
		routes.GET("/", controllers.NewTweetController().FindAll)
		routes.POST("/", controllers.NewTweetController().Create)
		routes.DELETE("/:id", controllers.NewTweetController().Delete)
	}

	return routes
}
