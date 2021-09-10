package main

import (
	_ "fmt"
	_ "log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/d-kuznetsov/calendar-backend/handler"
	"github.com/d-kuznetsov/calendar-backend/repository"
	"github.com/d-kuznetsov/calendar-backend/service"
)

func main() {
	router := mux.NewRouter()
	client := repository.CreateMongoClient("mongodb://localhost:27017")
	repo := repository.CreateMongoRepo(client, "calendar")
	svc := service.CreateService(repo)
	router.HandleFunc("/register", handler.CreateHttpHandler(handler.RegisterHandler, svc)).Methods("POST")
	http.ListenAndServe(":8080", router)
}
