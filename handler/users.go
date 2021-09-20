package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/dto"
	"github.com/d-kuznetsov/calendar-backend/service"
)

func (hdlr *handler) Register(wtr http.ResponseWriter, req *http.Request) {
	var userData dto.User
	json.NewDecoder(req.Body).Decode(&userData)
	userId, err := hdlr.service.Register(userData)
	if err == service.ErrIncorrectData {
		throwBadReqErr(wtr, err.Error())
		return
	} else if err == service.ErrUserAlreadyExists {
		throwBadReqErr(wtr, "User with this email already exists")
		return
	} else if err != nil {
		throwIntServerErr(wtr)
		return
	}

	token, err := hdlr.service.CreateToken(userId)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	json.NewEncoder(wtr).Encode(struct {
		Token string `json:"token"`
		Name  string `json:"name"`
	}{
		Token: token, Name: userData.Name,
	})
}

func (hdlr *handler) Login(wtr http.ResponseWriter, req *http.Request) {
	var userData dto.User
	json.NewDecoder(req.Body).Decode(&userData)
	user, err := hdlr.service.Login(userData)
	if err == service.ErrIncorrectData {
		throwBadReqErr(wtr, err.Error())
		return
	} else if err == service.ErrUserDoesNotExist {
		throwBadReqErr(wtr, "Incorrect email or password")
		return
	} else if err != nil {
		throwIntServerErr(wtr)
		return
	}

	token, err := hdlr.service.CreateToken(user.Id)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	json.NewEncoder(wtr).Encode(struct {
		Token string `json:"token"`
		Name  string `json:"name"`
	}{
		Token: token, Name: user.Name,
	})
}
