package service

import (
	"errors"

	"github.com/d-kuznetsov/event-calendar-backend/dto"
	"github.com/d-kuznetsov/event-calendar-backend/repository"
)

type IService interface {
	Register(userData dto.User) (string, error)
	Login(userData dto.User) (dto.User, error)
	CreateToken(id string) (string, error)
	ParseToken(token string) (string, error)
	CreateEvent(eventData dto.Event) (string, error)
	UpdateEvent(eventData dto.Event) error
	DeleteEvent(id string) error
	GetEvents(params dto.PeriodParams) ([]dto.Event, error)
}

var ErrUserAlreadyExists = errors.New("service error: user already exists")
var ErrUserDoesNotExist = errors.New("service error: user does not exist")
var ErrIncorrectData = errors.New("service error: incorrect data")

type Service struct {
	repository repository.IRepository
}

func CreateService(repo repository.IRepository) IService {
	return &Service{repo}
}
