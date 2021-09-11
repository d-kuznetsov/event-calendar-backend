package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/d-kuznetsov/calendar-backend/handler"
	"github.com/d-kuznetsov/calendar-backend/repository/mongodb"
	"github.com/d-kuznetsov/calendar-backend/service"
)

func main() {
	client := mongodb.CreateClient("mongodb://localhost:27017")
	repo := mongodb.CreateRepository(client, "calendar")
	svc := service.CreateService(repo)

	router := mux.NewRouter()
	router.HandleFunc("/register", handler.CreateHttpHandler(handler.RegisterHandler, svc)).Methods("POST")
	router.HandleFunc("/login", handler.CreateHttpHandler(handler.LoginHandler, svc)).Methods("POST")
	router.HandleFunc("/create-event", handler.CreateHttpHandler(handler.CreateEventHandler, svc)).Methods("POST")

	http.ListenAndServe(":8080", router)
}
