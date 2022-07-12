package main

import (
	"demo/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	app.CreateApp(router)

	log.Fatal(http.ListenAndServe(":7000", router))
}
