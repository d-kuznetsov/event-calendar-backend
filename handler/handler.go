package handler

import (
	"net/http"
	"net/mail"
	"regexp"

	"github.com/d-kuznetsov/calendar-backend/service"
)

type handler struct {
	service service.IService
}

func CreateHandler(svc service.IService) handler {
	return handler{
		service: svc,
	}
}

func throwBadReqErr(wtr http.ResponseWriter, errMsg string) {
	http.Error(wtr, errMsg, http.StatusBadRequest)
}

func throwUnauthorizedErr(wtr http.ResponseWriter) {
	http.Error(wtr, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func throwIntServerErr(wtr http.ResponseWriter) {
	http.Error(wtr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func extractToken(req *http.Request) string {
	prefix := "Bearer "
	authHeader := req.Header.Get("Authorization")
	if len(authHeader) < len(prefix) {
		return authHeader
	}
	return authHeader[len(prefix):]
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

var dateRegexp = regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`)

func isDateValid(date string) bool {
	res := dateRegexp.MatchString(date)
	return res
}

var timeRegexp = regexp.MustCompile(`^([01]\d|2[0-3]):([0-5]\d)$`)

func isTimeValid(time string) bool {
	res := timeRegexp.MatchString(time)
	return res
}
