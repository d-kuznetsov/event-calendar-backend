package repository

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/entities"
)

type IRepository interface {
	CreateUser(name, email, hashedPassword string) (string, error)
	GetUserByEmail(email string) (entities.User, error)
	CreateEvent(opts EventOpts) (string, error)
	GetUserEvents(userId string) ([]entities.Event, error)
}

type EventOpts struct {
	Date      string
	StartTime string
	EndTime   string
	Content   string
	UserId    string
}

var ErrNoUsersFound = errors.New("repository error: no users found")
