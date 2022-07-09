package controllers

import (
	"encoding/json"
	"net/http"

	"demo/domain/services"
)

func GetVehicle(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Header().Set("Content-Type", "application/json")
	id := req.URL.Query().Get("id")

	if len(id) > 0 {
		result, err := services.GetVehicleById(id)

		if err != nil {
			resWriter.WriteHeader(http.StatusNotFound)
			json.NewEncoder(resWriter)
			return
		}

		json.NewEncoder(resWriter).Encode(result)
	} else {
		result, _ := services.GetAllVehicle()
		json.NewEncoder(resWriter).Encode(result)
	}
}
