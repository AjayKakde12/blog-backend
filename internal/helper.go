package internal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("SECRET_KEY")

func Hash(input string) (string, error) {
	var byteHash []byte
	cost, err := strconv.ParseInt(os.Getenv("HASH_COST"), 10, 64)
	if err != nil {
		return "", err
	}
	byteHash, err = bcrypt.GenerateFromPassword([]byte(input), int(cost))
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
