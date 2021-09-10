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
		throw400Error(wtr, "User with this email already exists")
		return
	} else if err != nil {
		throw500Error(wtr)
		return
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
		throw400Error(wtr, "Incorrect email or password")
		return
	} else if err != nil {
		throw500Error(wtr)
		return
	}
	wtr.Write([]byte(token))
}

type RegistrationData struct {
	Name     string
	Email    string
	Password string
}

func throw400Error(wtr http.ResponseWriter, errMsg string) {
	http.Error(wtr, errMsg, http.StatusBadRequest)
}

func throw500Error(wtr http.ResponseWriter) {
	http.Error(wtr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
