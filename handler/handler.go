package handler

import (
	"net/http"

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
