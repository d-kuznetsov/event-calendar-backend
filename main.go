package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/d-kuznetsov/calendar-backend/handler"
	"github.com/d-kuznetsov/calendar-backend/repository/mongodb"
	"github.com/d-kuznetsov/calendar-backend/service"
)

func main() {
	client := mongodb.CreateClient("mongodb://localhost:27017")
	repo := mongodb.CreateRepository(client, "calendar")
	svc := service.CreateService(repo)

	router := mux.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            false,
	})
	router.HandleFunc("/register", handler.CreateHttpHandler(handler.RegisterHandler, svc)).Methods("POST")
	router.HandleFunc("/login", handler.CreateHttpHandler(handler.LoginHandler, svc)).Methods("POST")
	router.HandleFunc("/create-event", handler.CreateHttpHandler(handler.CreateEventHandler, svc)).Methods("POST")
	router.HandleFunc("/user-events", handler.CreateHttpHandler(handler.GetUserEventsHandler, svc)).Methods("GET")

	http.ListenAndServe(":8080", corsHandler.Handler(router))
}
