package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/d-kuznetsov/event-calendar-backend/config"

	"github.com/dgrijalva/jwt-go"
)

func (service *Service) CreateToken(id string) (string, error) {
	return generateToken(id)
}

func (service *Service) ParseToken(token string) (string, error) {
	return parseToken(token)
}

var signingKey = config.GetConfig().SigningKey

const tokenTTL = 12 * time.Hour

func generateToken(payload string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		Subject:   payload,
	})
	return token.SignedString([]byte(signingKey))
}

func parseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, err error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("token error: unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	//token.Claims.Valid()

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token error: receiving claims from token")
	}

	return claims["sub"].(string), nil
}
