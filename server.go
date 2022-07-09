package main

import (
	"demo/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/vehicle", services.GetVehicle).Methods("GET")
	log.Fatal(http.ListenAndServe(":7000", router))
}
