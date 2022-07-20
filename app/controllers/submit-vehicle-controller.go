package controllers

import (
	"encoding/json"
	"net/http"

	submitvehicleservice "demo/domain/services/submit_vehicle_service"
)

type SubmitVehicleBodyRequest struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Cor      string `json:"cor"`
	Serie    string `json:"serie"`
}

type SubmitVehicleController struct {
	SubmitVehicleService submitvehicleservice.ISubmitVehicleService
}

func (controller SubmitVehicleController) Handle(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Header().Set("Content-Type", "application/json")

	var data SubmitVehicleBodyRequest
	errDecode := json.NewDecoder(req.Body).Decode(&data)
	if errDecode != nil {
		http.Error(resWriter, errDecode.Error(), 400)
		return
	}

	errSubmit := controller.SubmitVehicleService.Submit(
		data.Category, data.Name, data.Cor, data.Serie,
	)
	if errSubmit != nil {
		http.Error(resWriter, errSubmit.Error(), 500)
	}

	resWriter.WriteHeader(201)
}
