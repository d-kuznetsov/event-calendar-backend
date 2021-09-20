package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/dto"
)

func (hdlr *handler) CreateEvent(wtr http.ResponseWriter, req *http.Request) {
	token := extractToken(req)
	userId, err := hdlr.service.ParseToken(token)
	if err != nil {
		throwUnauthorizedErr(wtr)
		return
	}

	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)

	if !isDateValid(eventData.Date) ||
		!isTimeValid(eventData.StartTime) ||
		!isTimeValid(eventData.EndTime) ||
		eventData.StartTime > eventData.EndTime ||
		eventData.Content == "" {
		throwBadReqErr(wtr, "Incorrect data")
		return
	}

	eventData.UserId = userId
	eventId, err := hdlr.service.CreateEvent(eventData)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	wtr.Write([]byte(eventId))
	// wtr.WriteHeader(http.StatusNoContent)
}

func (hdlr *handler) GetEvents(wtr http.ResponseWriter, req *http.Request) {
	token := extractToken(req)
	userId, err := hdlr.service.ParseToken(token)
	if err != nil {
		throwUnauthorizedErr(wtr)
		return
	}

	req.ParseForm()
	periodParams := dto.PeriodParams{
		PeriodStart: req.Form.Get("periodStart"),
		PeriodEnd:   req.Form.Get("periodEnd"),
		UserId:      userId,
	}

	if !isDateValid(periodParams.PeriodStart) ||
		!isDateValid(periodParams.PeriodEnd) ||
		periodParams.PeriodStart > periodParams.PeriodEnd ||
		periodParams.UserId == "" {
		throwBadReqErr(wtr, "Incorrect data")
		return
	}

	events, err := hdlr.service.GetEvents(periodParams)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	json.NewEncoder(wtr).Encode(events)
}

func (hdlr *handler) UpdateEvent(wtr http.ResponseWriter, req *http.Request) {
	token := extractToken(req)
	_, err := hdlr.service.ParseToken(token)
	if err != nil {
		throwUnauthorizedErr(wtr)
		return
	}

	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)

	if eventData.Id == "" ||
		!isDateValid(eventData.Date) ||
		!isTimeValid(eventData.StartTime) ||
		!isTimeValid(eventData.EndTime) ||
		eventData.StartTime > eventData.EndTime ||
		eventData.Content == "" {
		throwBadReqErr(wtr, "Incorrect data")
		return
	}

	err = hdlr.service.UpdateEvent(eventData)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	json.NewEncoder(wtr).Encode(true)
	// wtr.WriteHeader(http.StatusNoContent)
}

func (hdlr *handler) DeleteEvent(wtr http.ResponseWriter, req *http.Request) {
	token := extractToken(req)
	_, err := hdlr.service.ParseToken(token)
	if err != nil {
		throwUnauthorizedErr(wtr)
		return
	}

	var eventData dto.Event
	json.NewDecoder(req.Body).Decode(&eventData)

	if eventData.Id == "" {
		throwBadReqErr(wtr, "Incorrect data")
		return
	}

	err = hdlr.service.DeleteEvent(eventData.Id)
	if err != nil {
		throwIntServerErr(wtr)
		return
	}

	json.NewEncoder(wtr).Encode(true)
	// wtr.WriteHeader(http.StatusNoContent)
}
