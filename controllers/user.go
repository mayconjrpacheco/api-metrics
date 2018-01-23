package controllers

import (
	"net/http"
	"../models"
	"encoding/json"
)

func CreateUser(responseWriter http.ResponseWriter, req *http.Request) {
	userCreate := new(models.User)
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&userCreate)

	json.NewEncoder(responseWriter).Encode(userCreate)
}