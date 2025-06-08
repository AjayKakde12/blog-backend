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

func Login(c *gin.Context) {
	var credentials models.Credentials
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		fmt.Println("Invalid data in request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please check user details"})
		return
	}
	token, err := services.LoginUser(credentials)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
