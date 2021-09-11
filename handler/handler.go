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
	var regData Credentials
	json.Unmarshal(body, &regData)
	userId, err := svc.Register(regData.Name, regData.Email, regData.Password)
	if err == service.ErrUserExists {
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

func CreateEventHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	userId, err := svc.ParseToken(token)
	if err != nil {
		throw401Error(wtr)
		return
	}
	len := req.ContentLength
	body := make([]byte, len)
	req.Body.Read(body)
	eventOpts := service.EventOpts{}
	json.Unmarshal(body, &eventOpts)
	eventOpts.UserId = userId
	eventId, err := svc.CreateEvent(eventOpts)
	if err != nil {
		throw500Error(wtr)
		return
	}
	wtr.Write([]byte(eventId))
}

func throw400Error(wtr http.ResponseWriter, errMsg string) {
	http.Error(wtr, errMsg, http.StatusBadRequest)
}

func throw401Error(wtr http.ResponseWriter) {
	http.Error(wtr, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func throw500Error(wtr http.ResponseWriter) {
	http.Error(wtr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func extractToken(req *http.Request) string {
	prefix := "Bearer "
	authHeader := req.Header.Get("Authorization")
	return authHeader[len(prefix):]
}
