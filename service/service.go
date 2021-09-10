package service

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/repository"
)

type IService interface {
	Register(name, email, password string) (string, error)
}

type Service struct {
	repository repository.IRepository
}

func CreateService(repo repository.IRepository) IService {
	return &Service{repo}
}

var ErrUserExists = errors.New("service error: user exists")

func (service *Service) Register(name, email, password string) (string, error) {
	_, err := service.repository.GetUserByEmail(email)
	if err == nil {
		return "", ErrUserExists
	}
	if err != nil && err != repository.ErrNoUsersFound {
		return "", err
	}
	id, err := service.repository.CreateUser(name, email, password)
	if err != nil {
		return "", err
	}
	return generateToken(id, name, email)
}
