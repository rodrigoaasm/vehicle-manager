package controllers

import (
	"encoding/json"
	"net/http"

	httpadapter "demo/app/adapters/http_adapter"
	"demo/domain/domainerror"
	submitvehicleservice "demo/domain/services/submit_vehicle_service"
)

type SubmitVehicleBodyRequest struct {
	Category     string `json:"category"`
	Name         string `json:"name"`
	Cor          string `json:"cor"`
	Serie        string `json:"serie"`
	LicensePlate string `json:"licensePlate`
}

type SubmitVehicleController struct {
	SubmitVehicleService submitvehicleservice.ISubmitVehicleService
}

func (controller SubmitVehicleController) Handle(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Header().Set("Content-Type", "application/json")

	var data SubmitVehicleBodyRequest
	errDecode := json.NewDecoder(req.Body).Decode(&data)
	if errDecode != nil {
		httpadapter.BackError(resWriter, domainerror.New(domainerror.INVALID_DATA, errDecode.Error()))
		return
	}

	errSubmit := controller.SubmitVehicleService.Submit(
		data.Category, data.Name, data.Cor, data.Serie, data.LicensePlate,
	)
	if errSubmit != nil {
		httpadapter.BackError(resWriter, errSubmit)
	}

	resWriter.WriteHeader(201)
}
