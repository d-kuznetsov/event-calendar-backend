package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/service"
)

type HandlerFunc = func(wtr http.ResponseWriter, req *http.Request, svc service.IService)

func CreateHttpHandler(handlFun HandlerFunc, svc service.IService) http.HandlerFunc {
	return func(wtr http.ResponseWriter, req *http.Request) {
		handlFun(wtr, req, svc)
	}
}

func RegisterHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	len := req.ContentLength
	body := make([]byte, len)
	req.Body.Read(body)
	var regData RegistrationData
	json.Unmarshal(body, &regData)
	token, err := svc.Register(regData.Name, regData.Email, regData.Password)
	if err == service.ErrUserExists {
		http.Error(wtr, "User with this email already exists", http.StatusBadRequest)
	}
	wtr.Write([]byte(token))
}

func LoginHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	len := req.ContentLength
	body := make([]byte, len)
	req.Body.Read(body)
	var regData RegistrationData
	json.Unmarshal(body, &regData)
	token, err := svc.Login(regData.Email, regData.Password)
	if err == service.ErrUserDoesNotExist {
		http.Error(wtr, "Incorrect email or password", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(wtr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	wtr.Write([]byte(token))
}

type RegistrationData struct {
	Name     string
	Email    string
	Password string
}
