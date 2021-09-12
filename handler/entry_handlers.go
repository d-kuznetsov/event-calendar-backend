package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/service"
)

func RegisterHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	var regData Credentials
	json.NewDecoder(req.Body).Decode(&regData)
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
	resData := responseData{
		Token: token,
		Name:  regData.Name,
	}
	json.NewEncoder(wtr).Encode(resData)
}

func LoginHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	var loginData Credentials
	json.NewDecoder(req.Body).Decode(&loginData)
	user, err := svc.Login(loginData.Email, loginData.Password)
	if err == service.ErrUserDoesNotExist {
		throw400Error(wtr, "Incorrect email or password")
		return
	} else if err != nil {
		throw500Error(wtr)
		return
	}
	token, err := svc.CreateToken(user.Id)
	if err != nil {
		throw500Error(wtr)
		return
	}
	resData := responseData{
		Token: token,
		Name:  user.Name,
	}
	json.NewEncoder(wtr).Encode(resData)
}

type Credentials struct {
	Name     string
	Email    string
	Password string
}

type responseData struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}
