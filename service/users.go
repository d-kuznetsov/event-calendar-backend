package service

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"

	"github.com/d-kuznetsov/calendar-backend/dto"
	"github.com/d-kuznetsov/calendar-backend/repository"
)

func (service *Service) Register(userData dto.User) (string, error) {
	if userData.Name == "" || userData.Password == "" || !checkEmail(userData.Email) {
		return "", ErrIncorrectData
	}

	_, err := service.repository.GetUserByEmail(userData.Email)
	if err == nil {
		return "", ErrUserAlreadyExists
	}
	if err != nil && err != repository.ErrNoUsersFound {
		return "", err
	}

	pwdHash, err := hashPassword(userData.Password)
	if err != nil {
		return "", err
	}
	userData.Password = pwdHash

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
	if !checkPasswordHash(userData.Password, applicant.Password) {
		return dto.User{}, ErrUserDoesNotExist
	}
	return applicant, err
}

func checkEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
