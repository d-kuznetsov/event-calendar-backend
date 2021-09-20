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
	Login(userData dto.User) (dto.User, error)
	CreateEvent(eventData dto.Event) (string, error)
	UpdateEvent(eventData dto.Event) error
	DeleteEvent(id string) error
	GetEvents(params dto.PeriodParams) ([]dto.Event, error)
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

func (service *Service) Login(userData dto.User) (dto.User, error) {
	applicant, err := service.repository.GetUserByEmail(userData.Email)
	if err == repository.ErrNoUsersFound {
		return dto.User{}, ErrUserDoesNotExist
	} else if err != nil {
		return dto.User{}, err
	}
	if applicant.Password != userData.Password {
		return dto.User{}, ErrUserDoesNotExist
	}
	return applicant, err
}

func (service *Service) CreateEvent(eventData dto.Event) (string, error) {
	return service.repository.CreateEvent(eventData)
}

func (service *Service) GetEvents(params dto.PeriodParams) ([]dto.Event, error) {
	return service.repository.GetEvents(params)
}

func (service *Service) UpdateEvent(eventData dto.Event) error {
	return service.repository.UpdateEvent(eventData)
}

func (service *Service) DeleteEvent(id string) error {
	return service.repository.DeleteEvent(id)
}
