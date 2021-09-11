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
