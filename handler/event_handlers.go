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
	eventOpts := service.EventOpts{}
	json.NewDecoder(req.Body).Decode(&eventOpts)
	err = svc.UpdateEvent(eventOpts)
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
	var opts struct{ Id string }
	json.NewDecoder(req.Body).Decode(&opts)
	err = svc.DeleteEventById(opts.Id)
	if err != nil {
		throw500Error(wtr)
		return
	}
	json.NewEncoder(wtr).Encode(true)
}
