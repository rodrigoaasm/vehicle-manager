package vehicleupdateservice

import "demo/domain/domainerror"

type VehicleUpdateInput struct {
	Id                   string
	Status               string
	StatusAutomaticPilot string
	Travelled            float32
}

type IVehicleUpdateService interface {
	Turn(vehicle VehicleUpdateInput) *domainerror.DomainError
}
