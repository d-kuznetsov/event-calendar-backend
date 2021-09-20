package service

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/dto"
	"github.com/d-kuznetsov/calendar-backend/repository"
)

type IService interface {
	CreateToken(id string) (string, error)
	ParseToken(token string) (string, error)
	Register(userData dto.User) (string, error)
	Login(email, password string) (dto.User, error)
	CreateEvent(params EventOpts) (string, error)
	GetUserEvents(params struct {
		PeriodStart string
		PeriodEnd   string
		UserId      string
	}) ([]dto.Event, error)
	UpdateEvent(params EventOpts) error
	DeleteEventById(id string) error
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

func (service *Service) Register(userData dto.User) (string, error) {
	_, err := service.repository.GetUserByEmail(userData.Email)
	if err == nil {
		return "", ErrUserAlreadyExists
	}
	if err != nil && err != repository.ErrNoUsersFound {
		return "", err
	}
	userId, err := service.repository.CreateUser(userData)
	if err != nil {
		return "", err
	}
	return userId, err
}

func (service *Service) Login(email, password string) (dto.User, error) {
	user, err := service.repository.GetUserByEmail(email)
	if err == repository.ErrNoUsersFound {
		return dto.User{}, ErrUserDoesNotExist
	} else if err != nil {
		return dto.User{}, err
	}
	if user.Password != password {
		return dto.User{}, ErrUserDoesNotExist
	}
	return user, err
}

func (service *Service) CreateEvent(params EventOpts) (string, error) {
	return service.repository.CreateEvent(params)
}

func (service *Service) GetUserEvents(params struct {
	PeriodStart string
	PeriodEnd   string
	UserId      string
}) ([]dto.Event, error) {
	return service.repository.GetUserEvents(params)
}

func (service *Service) UpdateEvent(params EventOpts) error {
	return service.repository.UpdateEvent(params)
}

func (service *Service) DeleteEventById(id string) error {
	return service.repository.DeleteEventById(id)
}
