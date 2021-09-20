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
	DeleteEventById(id string) error
	GetUserEvents(params struct {
		PeriodStart string
		PeriodEnd   string
		UserId      string
	}) ([]dto.Event, error)
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
