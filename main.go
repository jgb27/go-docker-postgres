package main

import (
	routes "go-docker-postgres/api/routes"
	database "go-docker-postgres/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDatabase()
	app := gin.Default()
	routes.Routes(app)
	app.Run(":8080")
}
