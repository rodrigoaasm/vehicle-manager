package controllers

import (
	httpadapter "demo/app/adapters/http_adapter"
	"demo/domain/domainerror"
	vehicleturnservice "demo/domain/services/vehicle_turn_service"
	"encoding/json"
	"net/http"
)

type VehicleTurnersBodyRequest struct {
	Id                   string `json:"id"`
	Status               string `json:"status"`
	StatusAutomaticPilot string `json:"automaticPilot"`
}

type VehicleTurnersController struct {
	VehicleTurnService vehicleturnservice.IVehicleTurnService
}

func (controller VehicleTurnersController) Handle(resWriter http.ResponseWriter, req *http.Request) {
	var data VehicleTurnersBodyRequest
	errDecode := json.NewDecoder(req.Body).Decode(&data)
	if errDecode != nil {
		httpadapter.BackError(resWriter, domainerror.New(domainerror.INVALID_DATA, errDecode.Error()))
		return
	}

	errSubmit := controller.VehicleTurnService.Turn(vehicleturnservice.VehicleTurnInput{
		Id:                   data.Id,
		Status:               data.Status,
		StatusAutomaticPilot: data.StatusAutomaticPilot,
	})
	if errSubmit != nil {
		httpadapter.BackError(resWriter, errSubmit)
	}

	resWriter.WriteHeader(204)
}
