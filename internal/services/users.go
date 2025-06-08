package services

import (
	"blog/root/internal"
	"blog/root/internal/models"
	"fmt"
	"time"
)

func CreateUser(user *models.Users) error {
	var err error
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password, err = internal.Hash(user.Password)
	if err != nil {
		fmt.Println("Failed to hash a password:", err)
		return err
	}
	err = user.Create()
	if err != nil {
		return err
	}
	return err
}
