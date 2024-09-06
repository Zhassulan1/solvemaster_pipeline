package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// setting up endpoints
	router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	router.HandleFunc("/update", Update).Methods("GET")
	http.Handle("/", router)
	// Update()
	// starting listener
	http.ListenAndServe(":8080", router)
}
