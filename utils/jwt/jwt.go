package jwt

import (
	"be-golang-fiber/entity/user"
	"be-golang-fiber/utils"
	"be-golang-fiber/utils/config_variable"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWTToken(u *user.UserLoginModel) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["admin"] = false
	if u.RoleID == 1 {
		claims["admin"] = true
	}
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config_variable.Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetClaims(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(config_variable.Secret), nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil
}

func VerifyAdminToken(token string) (bool, error) {
	token, err := GetTokenfromHeader(token)
	if err != nil {
		return false, err
	}
	claims, err := GetClaims(token)
	if err != nil {
		return false, err
	}
	// valid, err := verifyTokenValidity(claims, token)
	// if !valid || err != nil {
	// 	return valid, err
	// }
	admin, _ := claims["admin"].(bool)
	if !admin {
		return admin, errors.New(utils.UNAUTHORIZED)
	}
	return admin, nil
}

func GetIDfromToken(token string) (id int, err error) {
	token, err = GetTokenfromHeader(token)
	if err != nil {
		return id, err
	}
	claims, err := GetClaims(token)
	if err != nil {
		return id, err
	}
	// valid, err := verifyTokenValidity(claims, token)
	// if !valid || err != nil {
	// 	return id, err
	// }
	idClaim, _ := claims["id"].(float64)
	id = int(idClaim)
	return id, nil
}

func GetTokenfromHeader(token string) (string, error) {
	if token == "" {
		return "", errors.New(utils.AUTH_REQUIRED)
	}
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]
	return token, nil
}
