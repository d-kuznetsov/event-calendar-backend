package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/dto"
	"github.com/d-kuznetsov/calendar-backend/service"
)

func CreateEventHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	userId, err := svc.ParseToken(token)
	if err != nil {
		throw401Error(wtr)
		return
	}
	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)
	eventData.UserId = userId
	eventId, err := svc.CreateEvent(eventData)
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
	req.ParseForm()
	var body = struct {
		PeriodStart string
		PeriodEnd   string
		UserId      string
	}{
		PeriodStart: req.Form["periodStart"][0],
		PeriodEnd:   req.Form["periodEnd"][0],
		UserId:      userId,
	}
	events, err := svc.GetUserEvents(body)
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
	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)
	err = svc.UpdateEvent(eventData)
	if err != nil {
		throw500Error(wtr)
		return
	}
	json.NewEncoder(wtr).Encode(true)
}

func DeleteEventHandler(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	_, err := svc.ParseToken(token)
	if err != nil {
		throw401Error(wtr)
		return
	}
	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)
	err = svc.DeleteEventById(eventData.Id)
	if err != nil {
		throw500Error(wtr)
		return
	}
	json.NewEncoder(wtr).Encode(true)
}
