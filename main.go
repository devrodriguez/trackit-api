package main

import (
	"net/http"

	"github.com/devrodriguez/first-class-api-go/routes"
	"github.com/gorilla/mux"
)

func main() {
	port := ":3001"
	router := mux.NewRouter()

	router.HandleFunc("/checks", routes.GetChecks).Methods(http.MethodGet)
	router.HandleFunc("/checks", routes.CreateCheck).Methods(http.MethodPost)

	http.ListenAndServe(port, router)
}
