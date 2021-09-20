package repository

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/dto"
)

type IRepository interface {
	CreateUser(name, email, hashedPassword string) (string, error)
	GetUserByEmail(email string) (dto.User, error)
	CreateEvent(params EventOpts) (string, error)
	GetUserEvents(params struct {
		PeriodStart string
		PeriodEnd   string
		UserId      string
	}) ([]dto.Event, error)
	UpdateEvent(params EventOpts) error
	DeleteEventById(id string) error
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
