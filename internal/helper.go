package internal

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(input string) (string, error) {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(byteHash), err
}
