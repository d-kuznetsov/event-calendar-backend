package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/service"
)

func RegisterHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	len := req.ContentLength
	body := make([]byte, len)
	req.Body.Read(body)
	var regData Credentials
	json.Unmarshal(body, &regData)
	userId, err := svc.Register(regData.Name, regData.Email, regData.Password)
	if err == service.ErrUserAlreadyExists {
		throw400Error(wtr, "User with this email already exists")
		return
	} else if err != nil {
		throw500Error(wtr)
		return
	}
	token, err := svc.CreateToken(userId)
	if err != nil {
		throw500Error(wtr)
		return
	}
	wtr.Write([]byte(token))
}

func LoginHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	len := req.ContentLength
	body := make([]byte, len)
	req.Body.Read(body)
	var loginData Credentials
	json.Unmarshal(body, &loginData)
	userId, err := svc.Login(loginData.Email, loginData.Password)
	if err == service.ErrUserDoesNotExist {
		throw400Error(wtr, "Incorrect email or password")
		return
	} else if err != nil {
		throw500Error(wtr)
		return
	}
	token, err := svc.CreateToken(userId)
	if err != nil {
		throw500Error(wtr)
		return
	}
	wtr.Write([]byte(token))
}

type Credentials struct {
	Name     string
	Email    string
	Password string
}
