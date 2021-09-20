package service

import (
	"net/mail"

	"github.com/d-kuznetsov/calendar-backend/dto"
	"github.com/d-kuznetsov/calendar-backend/repository"
)

func (service *Service) Register(userData dto.User) (string, error) {
	if userData.Name == "" || userData.Password == "" || !isEmailValid(userData.Email) {
		return "", ErrIncorrectData
	}

	_, err := service.repository.GetUserByEmail(userData.Email)
	if err == nil {
		return "", ErrUserAlreadyExists
	}
	if err != nil && err != repository.ErrNoUsersFound {
		return "", err
	}
	userId, err := service.repository.CreateUser(userData)

	return userId, err
}

func (service *Service) Login(userData dto.User) (dto.User, error) {
	if userData.Email == "" || userData.Password == "" {
		return dto.User{}, ErrIncorrectData
	}

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

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
