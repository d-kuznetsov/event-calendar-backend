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
	hdlr := handler.CreateHandler(svc)

	router := mux.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            true,
	})
	router.HandleFunc("/register", hdlr.Register).Methods("POST")
	router.HandleFunc("/login", hdlr.Login).Methods("POST")
	router.HandleFunc("/create-event", hdlr.CreateEvent).Methods("POST")
	router.HandleFunc("/update-event", hdlr.UpdateEvent).Methods("POST")
	router.HandleFunc("/delete-event", hdlr.DeleteEvent).Methods("POST")
	router.HandleFunc("/user-events", hdlr.GetEvents).Methods("GET")

	http.ListenAndServe(":8080", corsHandler.Handler(router))
}
