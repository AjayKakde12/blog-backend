package handlers

import (
	"blog/root/internal/models"
	"blog/root/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.Users
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("Invalid data in request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please check user details"})
		return
	}
	err = services.CreateUser(&user)
	if err != nil {
		fmt.Println("failed to insert user's data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create new user. Try again."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
