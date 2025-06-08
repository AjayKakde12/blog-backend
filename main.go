package main

import (
	"blog/root/internal/database"
	"blog/root/internal/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Not able to load environment variables")
	}

	server := gin.Default()

	database.InitDatabase()

	routes.RegisterRoutes(server)
	address := fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT"))

	server.Run(address)
}
