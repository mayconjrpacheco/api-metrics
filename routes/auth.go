package routes

import (
	"github.com/gorilla/mux"
	"../controllers"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth", controllers.CreateUser).Methods("POST")
}