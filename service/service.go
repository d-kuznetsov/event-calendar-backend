package service

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/repository"
)

type IService interface {
	CreateToken(id string) (string, error)
	ParseToken(token string) (string, error)
	Register(name, email, password string) (string, error)
	Login(email, password string) (string, error)
	CreateEvent(params EventOpts) (string, error)
}

type EventOpts = repository.EventOpts

type Service struct {
	repository repository.IRepository
}

func CreateService(repo repository.IRepository) IService {
	return &Service{repo}
}

var ErrUserExists = errors.New("service error: user exists")
var ErrUserDoesNotExist = errors.New("service error: user does not exist")

func (service *Service) CreateToken(id string) (string, error) {
	return generateToken(id)
}

func (service *Service) ParseToken(token string) (string, error) {
	return parseToken(token)
}

func (service *Service) Register(name, email, password string) (string, error) {
	_, err := service.repository.GetUserByEmail(email)
	if err == nil {
		return "", ErrUserExists
	}
	if err != nil && err != repository.ErrNoUsersFound {
		return "", err
	}
	userId, err := service.repository.CreateUser(name, email, password)
	if err != nil {
		return "", err
	}
	return generateToken(userId)
}

func (service *Service) Login(email, password string) (string, error) {
	user, err := service.repository.GetUserByEmail(email)
	if err == repository.ErrNoUsersFound {
		return "", ErrUserDoesNotExist
	} else if err != nil {
		return "", err
	}
	if user.Email != email {
		return "", ErrUserDoesNotExist
	}
	return generateToken(user.Id)
}

func (service *Service) CreateEvent(params EventOpts) (string, error) {
	return service.repository.CreateEvent(params)
}
