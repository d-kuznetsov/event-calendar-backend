package service

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/entities"
	"github.com/d-kuznetsov/calendar-backend/repository"
)

type IService interface {
	CreateToken(id string) (string, error)
	ParseToken(token string) (string, error)
	Register(name, email, password string) (string, error)
	Login(email, password string) (entities.User, error)
	CreateEvent(params EventOpts) (string, error)
}

type EventOpts = repository.EventOpts

var ErrUserAlreadyExists = errors.New("service error: user already exists")
var ErrUserDoesNotExist = errors.New("service error: user does not exist")

type Service struct {
	repository repository.IRepository
}

func CreateService(repo repository.IRepository) IService {
	return &Service{repo}
}

func (service *Service) CreateToken(id string) (string, error) {
	return generateToken(id)
}

func (service *Service) ParseToken(token string) (string, error) {
	return parseToken(token)
}

func (service *Service) Register(name, email, password string) (string, error) {
	_, err := service.repository.GetUserByEmail(email)
	if err == nil {
		return "", ErrUserAlreadyExists
	}
	if err != nil && err != repository.ErrNoUsersFound {
		return "", err
	}
	userId, err := service.repository.CreateUser(name, email, password)
	if err != nil {
		return "", err
	}
	return userId, err
}

func (service *Service) Login(email, password string) (entities.User, error) {
	user, err := service.repository.GetUserByEmail(email)
	if err == repository.ErrNoUsersFound {
		return entities.User{}, ErrUserDoesNotExist
	} else if err != nil {
		return entities.User{}, err
	}
	if user.Password != password {
		return entities.User{}, ErrUserDoesNotExist
	}
	return user, err
}

func (service *Service) CreateEvent(params EventOpts) (string, error) {
	return service.repository.CreateEvent(params)
}
