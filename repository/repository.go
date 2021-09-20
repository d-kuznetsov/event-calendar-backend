package repository

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/dto"
)

type IRepository interface {
	CreateUser(userData dto.User) (string, error)
	GetUserByEmail(email string) (dto.User, error)
	CreateEvent(eventData dto.Event) (string, error)
	UpdateEvent(eventData dto.Event) error
	DeleteEvent(id string) error
	GetEvents(params dto.PeriodParams) ([]dto.Event, error)
}

type EventOpts struct {
	Id        string
	Date      string
	StartTime string
	EndTime   string
	Content   string
	UserId    string
}

var ErrNoUsersFound = errors.New("repository error: no users found")
var ErrNoEventsFound = errors.New("repository error: no events found")
