package controllers

import (
	"encoding/json"
	"net/http"

	vehiclegetterservice "demo/domain/services/vehicle_getter_service"
)

type VehicleGetterController struct {
	VehicleGetterService vehiclegetterservice.IVehicleGetterService
}

func (controller VehicleGetterController) Handle(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Header().Set("Content-Type", "application/json")
	id := req.URL.Query().Get("id")

	if len(id) > 0 {
		result, err := controller.VehicleGetterService.GetVehicleById(id)

		if err != nil {
			http.Error(resWriter, err.Error(), 500)
			return
		}

		json.NewEncoder(resWriter).Encode(result)
	} else {
		result, _ := controller.VehicleGetterService.GetAllVehicle()
		json.NewEncoder(resWriter).Encode(result)
	}
}
