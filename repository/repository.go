package repository

import (
	"github.com/d-kuznetsov/calendar-backend/models"
)

type Repository interface {
	CreateUser(name, email, hashedPassword string) (string, error)
	GetUserByEmail(emaul string) (models.User, error)
}
