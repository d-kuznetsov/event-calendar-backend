package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/d-kuznetsov/calendar-backend/repository"
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

type IService interface {
	Register(name, email, password string) (string, error)
}

type Service struct {
	Repository repository.IRepository
}

func CreateService(repo repository.IRepository) IService {
	return &Service{repo}
}

var ErrUserExists = errors.New("service error: user exists")

func (service *Service) Register(name, email, password string) (string, error) {
	_, err := service.Repository.GetUserByEmail(email)
	if err == nil {
		return "", ErrUserExists
	}
	if err != nil && err != repository.ErrNoUsersFound {
		return "", err
	}
	id, err := service.Repository.CreateUser(name, email, password)
	if err != nil {
		return "", err
	}
	return generateToken(id, name, email)
}
