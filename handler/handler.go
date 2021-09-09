package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/models"
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
	var user models.User
	json.Unmarshal(body, &user)
	token, err := svc.Register(user.Name, user.Email, user.Password)
	if err == service.ErrUserExists {
		http.Error(wtr, "User with this email already exists", http.StatusBadRequest)
	}
	wtr.Write([]byte(token))
}
