package controllers

import (
	httpadapter "demo/app/adapters/http_adapter"
	"demo/domain/domainerror"
	vehicleupdateservice "demo/domain/services/vehicle_update_service"
	"encoding/json"
	"net/http"
)

type VehicleUpdateBodyRequest struct {
	Id                   string  `json:"id"`
	Status               string  `json:"status"`
	StatusAutomaticPilot string  `json:"automaticPilot"`
	Travelled            float32 `json:"travelled"`
}

type VehicleUpdateController struct {
	VehicleTurnService vehicleupdateservice.IVehicleUpdateService
}

func (controller VehicleUpdateController) Handle(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Header().Set("Content-Type", "application/json")

	var data VehicleUpdateBodyRequest
	errDecode := json.NewDecoder(req.Body).Decode(&data)
	if errDecode != nil {
		httpadapter.BackError(resWriter, domainerror.New(domainerror.INVALID_DATA, errDecode.Error()))
		return
	}

	errSubmit := controller.VehicleTurnService.Turn(vehicleupdateservice.VehicleUpdateInput{
		Id:                   data.Id,
		Status:               data.Status,
		StatusAutomaticPilot: data.StatusAutomaticPilot,
		Travelled:            data.Travelled,
	})
	if errSubmit != nil {
		httpadapter.BackError(resWriter, errSubmit)
	}

	resWriter.WriteHeader(204)
}
