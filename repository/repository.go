package repository

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/models"
)

type IRepository interface {
	CreateUser(name, email, hashedPassword string) (string, error)
	GetUserByEmail(email string) (models.User, error)
	CreateEvent(opts EventOpts) (string, error)
}

type EventOpts struct {
	Date      string
	StartTime string
	EndTime   string
	Content   string
	UserId    string
}

var ErrNoUsersFound = errors.New("repository error: no users found")
