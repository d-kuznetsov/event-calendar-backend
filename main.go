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
		AllowedHeaders:   []string{"*"},
		Debug:            true,
	})
	router.HandleFunc("/register", handler.CreateHandler(handler.RegisterHandler, svc)).Methods("POST")
	router.HandleFunc("/login", handler.CreateHandler(handler.LoginHandler, svc)).Methods("POST")
	router.HandleFunc("/create-event", handler.CreateHandler(handler.CreateEventHandler, svc)).Methods("POST")
	router.HandleFunc("/user-events", handler.CreateHandler(handler.GetEventsHandler, svc)).Methods("GET")
	router.HandleFunc("/update-event", handler.CreateHandler(handler.UpdateEventHandler, svc)).Methods("POST")
	router.HandleFunc("/delete-event", handler.CreateHandler(handler.DeleteEventHandler, svc)).Methods("POST")

	http.ListenAndServe(":8080", corsHandler.Handler(router))
}
