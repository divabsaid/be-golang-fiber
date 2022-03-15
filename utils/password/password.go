package password

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytepwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytepwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password string, hashedpwd string) (bool, error) {
	bytehashedpwd := []byte(hashedpwd)
	bytepwd := []byte(password)
	err := bcrypt.CompareHashAndPassword(bytehashedpwd, bytepwd)
	if err != nil {
		return false, err
	}

	return true, nil
}
