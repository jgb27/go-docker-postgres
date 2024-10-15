package main

import (
	routes "go-docker-postgres/api/routes"
	database "go-docker-postgres/db"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.Routes(app)
	database.ConnectToDatabase()

	app.Run(":8080")
}
