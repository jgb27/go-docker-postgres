package routes

import (
	controllers "go-docker-postgres/api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()
	routes := router.Group("/tweet")
	{
		routes.GET("/", tweetController.FindAll)
		routes.POST("/", tweetController.Create)
		routes.DELETE("/:id", tweetController.Delete)
	}

	return routes
}
