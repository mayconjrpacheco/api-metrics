package routes

import (
	"github.com/gorilla/mux"
	"../controllers"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
}