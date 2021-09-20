package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/dto"
	"github.com/d-kuznetsov/calendar-backend/service"
)

func CreateEvent(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	userId, err := svc.ParseToken(token)
	if err != nil {
		throwUnauthorizedErr(wtr)
		return
	}

	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)
	eventData.UserId = userId

	eventId, err := svc.CreateEvent(eventData)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	wtr.Write([]byte(eventId))
	// wtr.WriteHeader(http.StatusNoContent)
}

func GetEvents(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	userId, err := svc.ParseToken(token)
	if err != nil {
		throwUnauthorizedErr(wtr)
		return
	}

	req.ParseForm()
	periodParams := dto.PeriodParams{
		PeriodStart: req.Form["periodStart"][0],
		PeriodEnd:   req.Form["periodEnd"][0],
		UserId:      userId,
	}
	events, err := svc.GetEvents(periodParams)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	json.NewEncoder(wtr).Encode(events)
}

func UpdateEvent(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	_, err := svc.ParseToken(token)
	if err != nil {
		throwUnauthorizedErr(wtr)
		return
	}

	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)

	err = svc.UpdateEvent(eventData)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	json.NewEncoder(wtr).Encode(true)
	// wtr.WriteHeader(http.StatusNoContent)
}

func DeleteEvent(wtr http.ResponseWriter, req *http.Request, svc service.IService) {
	token := extractToken(req)
	_, err := svc.ParseToken(token)
	if err != nil {
		throwUnauthorizedErr(wtr)
		return
	}

	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)

	err = svc.DeleteEvent(eventData.Id)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	json.NewEncoder(wtr).Encode(true)
	// wtr.WriteHeader(http.StatusNoContent)
}
