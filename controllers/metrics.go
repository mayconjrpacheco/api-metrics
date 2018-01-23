package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Spa struct {
	Name string
	Url string
}

var spa Spa

func Metrics(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request metrics")
	_ = json.NewDecoder(req.Body).Decode(&spa)

	fmt.Printf("%s\n", spa.Url)
	resp, err := http.Get(spa.Url)
	if err != nil {
		json.NewEncoder(responseWriter).Encode(err.Error())
	}

	json.NewEncoder(responseWriter).Encode(resp.Header)
}