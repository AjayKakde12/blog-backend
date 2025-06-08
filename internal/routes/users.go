package routes

import (
	"blog/root/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/user", handlers.CreateUser)
}
