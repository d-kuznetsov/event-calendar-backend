package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/service"
)

func CreateEventHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	userId, err := svc.ParseToken(token)
	if err != nil {
		throw401Error(wtr)
		return
	}
	eventOpts := service.EventOpts{}
	json.NewDecoder(req.Body).Decode(&eventOpts)
	eventOpts.UserId = userId
	eventId, err := svc.CreateEvent(eventOpts)
	if err != nil {
		throw500Error(wtr)
		return
	}
	wtr.Write([]byte(eventId))
}

func GetUserEventsHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	userId, err := svc.ParseToken(token)
	if err != nil {
		throw401Error(wtr)
		return
	}
	events, err := svc.GetUserEvents(userId)
	if err != nil {
		throw500Error(wtr)
		return
	}
	json.NewEncoder(wtr).Encode(events)
}

func UpdateEventHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	_, err := svc.ParseToken(token)
	if err != nil {
		throw401Error(wtr)
		return
	}
	eventOpts := service.EventOpts{}
	json.NewDecoder(req.Body).Decode(&eventOpts)
	err = svc.UpdateEvent(eventOpts)
	if err != nil {
		throw500Error(wtr)
		return
	}
	json.NewEncoder(wtr).Encode(true)
}
