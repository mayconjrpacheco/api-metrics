package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"./routes"
)

func main() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	routes.AllRoutes(subRouter)
	log.Fatal(http.ListenAndServe(":8080", router))
}