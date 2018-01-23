package routes

import (
	"github.com/gorilla/mux"
)

func AllRoutes(router *mux.Router) {
	UserRoutes(router)
}