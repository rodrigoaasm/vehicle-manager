package main

import (
	"log"
	"net/http"

	"demo/app/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/vehicle", controllers.GetVehicle).Methods("GET")
	log.Fatal(http.ListenAndServe(":7000", router))
}
