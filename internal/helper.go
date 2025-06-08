package internal

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "TOP_SECRET_KEY"

func Hash(input string) (string, error) {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(byteHash), err
}

func VerifyHash(inputString, hashedString string) error {
	fmt.Println(hashedString, inputString)
	return bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(inputString))
}

func GenerateJWTToken(id int64, user_name string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        id,
		"user_name": user_name,
		"role":      role,
	})
	return token.SignedString([]byte(secretKey))
}
