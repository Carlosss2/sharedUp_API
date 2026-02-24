package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("encryptky0")

type Claims struct {
	IdUser int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(IdUser int) (string, error) {
	expirationTime := time.Now().Add(4 * time.Hour)
	claims := &Claims{
		IdUser: IdUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}