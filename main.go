package main

import (
	routes "go-docker-postgres/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.Routes(app)

	app.Run(":8080")
}
