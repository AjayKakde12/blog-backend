package main

import (
	"blog/root/internal/database"
	"blog/root/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	database.InitDatabase()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
