package handler

import (
	"net/http"

	"github.com/d-kuznetsov/calendar-backend/service"
)

type HandlerFunc = func(wtr http.ResponseWriter, req *http.Request, svc service.IService)

func CreateHttpHandler(handlFun HandlerFunc, svc service.IService) http.HandlerFunc {
	return func(wtr http.ResponseWriter, req *http.Request) {
		handlFun(wtr, req, svc)
	}
}

func throw400Error(wtr http.ResponseWriter, errMsg string) {
	http.Error(wtr, errMsg, http.StatusBadRequest)
}

func throw401Error(wtr http.ResponseWriter) {
	http.Error(wtr, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func throw500Error(wtr http.ResponseWriter) {
	http.Error(wtr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func extractToken(req *http.Request) string {
	prefix := "Bearer "
	authHeader := req.Header.Get("Authorization")
	return authHeader[len(prefix):]
}
