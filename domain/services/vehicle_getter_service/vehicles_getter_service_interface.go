package vehiclegetterservice

import "demo/domain/domainerror"

type VehicleGetterOutput struct {
	Id             string  `json:"id"`
	Category       string  `json:"category"`
	Name           string  `json:"name"`
	Color          string  `json:"color"`
	Serie          string  `json:"serie"`
	LicensePlate   string  `json:"licensePlate"`
	Travelled      float32 `json:"travelled"`
	Status         bool    `json:"status"`
	AutomaticPilot bool    `json:"automaticPilot"`
}

type IVehicleGetterService interface {
	GetAllVehicle() ([]VehicleGetterOutput, *domainerror.DomainError)
	GetVehicleById(id string) (VehicleGetterOutput, *domainerror.DomainError)
}
