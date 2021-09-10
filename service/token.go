package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("test_key")

const tokenTTL = 12 * time.Hour

type tokenClaims struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func generateToken(id, name, email string) (string, error) {
	claims := tokenClaims{
		id,
		name,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}
