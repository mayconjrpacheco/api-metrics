package routes

import (
	"github.com/gorilla/mux"
	"../handlers"
)

func AllRoutes(router *mux.Router) {
	router.HandleFunc("/metrics/", handlers.Metrics).Methods("POST")
}