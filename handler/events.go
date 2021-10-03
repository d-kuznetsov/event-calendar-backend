package handler

import (
	"encoding/json"
	"net/http"

	"github.com/d-kuznetsov/event-calendar-backend/dto"
	"github.com/d-kuznetsov/event-calendar-backend/service"
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
	eventData.UserId = userId

	eventId, err := hdlr.service.CreateEvent(eventData)
	if err == service.ErrIncorrectData {
		throwBadReqErr(wtr, err.Error())
		return
	} else if err != nil {
		throwIntServerErr(wtr, err)
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

	events, err := hdlr.service.GetEvents(periodParams)
	if err == service.ErrIncorrectData {
		throwBadReqErr(wtr, err.Error())
		return
	} else if err != nil {
		throwIntServerErr(wtr, err)
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

	err = hdlr.service.UpdateEvent(eventData)
	if err == service.ErrIncorrectData {
		throwBadReqErr(wtr, err.Error())
		return
	} else if err != nil {
		throwIntServerErr(wtr, err)
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

	err = hdlr.service.DeleteEvent(eventData.Id)
	if err == service.ErrIncorrectData {
		throwBadReqErr(wtr, err.Error())
		return
	} else if err != nil {
		throwIntServerErr(wtr, err)
		return
	}

	json.NewEncoder(wtr).Encode(true)
	// wtr.WriteHeader(http.StatusNoContent)
}
