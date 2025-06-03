package main

import (
	"log"
	"net/http"

	"sensors/db"
	"sensors/handlers"
	"sensors/middleware"
	"sensors/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.LoadEnv()
	db.Init(utils.DBPath)

	r := mux.NewRouter()

	r.Handle("/temperature", middleware.Auth(http.HandlerFunc(handlers.CreateTemperature))).Methods("POST")
	r.HandleFunc("/temperature", handlers.GetTemperature).Methods("GET")

	log.Printf("Server running at :%s\n", utils.Port)
	log.Fatal(http.ListenAndServe(":"+utils.Port, r))
}
