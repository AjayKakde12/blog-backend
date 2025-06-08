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

func validateUser(credentials models.Credentials) (models.Credentials, error) {
	userDetails, err := credentials.GetUserByUserName()
	if err != nil {
		return models.Credentials{}, err
	}
	return userDetails, err
}

func LoginUser(credentials models.Credentials) (string, error) {
	var token string
	userDetails, err := validateUser(credentials)
	if err != nil {
		return token, err
	}
	err = internal.VerifyHash(credentials.Password, userDetails.Password)
	if err != nil {
		return token, err
	}
	token, err = internal.GenerateJWTToken(userDetails.ID, credentials.UserName, credentials.Role)
	if err != nil {
		return token, err
	}

	return token, nil
}
