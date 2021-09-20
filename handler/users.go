package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/dto"
	"github.com/d-kuznetsov/calendar-backend/service"
)

func RegisterHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	var userData dto.User
	json.NewDecoder(req.Body).Decode(&userData)
	userId, err := svc.Register(userData)
	if err == service.ErrUserAlreadyExists {
		throwBadReqErr(wtr, "User with this email already exists")
		return
	} else if err != nil {
		throwIntServerErr(wtr)
		return
	}
	token, err := svc.CreateToken(userId)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}
	resData := responseData{
		Token: token,
		Name:  userData.Name,
	}
	json.NewEncoder(wtr).Encode(resData)
}

func LoginHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	var userData dto.User
	json.NewDecoder(req.Body).Decode(&userData)
	user, err := svc.Login(userData)
	if err == service.ErrUserDoesNotExist {
		throwBadReqErr(wtr, "Incorrect email or password")
		return
	} else if err != nil {
		throwIntServerErr(wtr)
		return
	}
	token, err := svc.CreateToken(user.Id)
	if err != nil {
		throwIntServerErr(wtr)
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
