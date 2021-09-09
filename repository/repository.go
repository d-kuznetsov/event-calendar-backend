package repository

import (
	"errors"

	"github.com/d-kuznetsov/calendar-backend/models"
)

type IRepository interface {
	CreateUser(name, email, hashedPassword string) (string, error)
	GetUserByEmail(email string) (models.User, error)
}

var ErrNoUsersFound = errors.New("repository error: no users found")
