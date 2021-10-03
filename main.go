package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/d-kuznetsov/event-calendar-backend/config"
	"github.com/d-kuznetsov/event-calendar-backend/handler"
	"github.com/d-kuznetsov/event-calendar-backend/repository/mongodb"
	"github.com/d-kuznetsov/event-calendar-backend/service"
)

func main() {
	cfg := config.LoadConfig()
	client := mongodb.CreateClient(cfg.DbUri)
	repo := mongodb.CreateRepository(client, cfg.DbName)
	svc := service.CreateService(repo)
	hdlr := handler.CreateHandler(svc)

	router := mux.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{cfg.ClientUri},
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

	http.ListenAndServe(cfg.ServerUri, corsHandler.Handler(router))
}
