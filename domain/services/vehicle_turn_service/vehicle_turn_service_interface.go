package vehicleturnservice

import "demo/domain/domainerror"

type VehicleTurnInput struct {
	Id                   string
	Status               string
	StatusAutomaticPilot string
}

type IVehicleTurnService interface {
	Turn(vehicle VehicleTurnInput) *domainerror.DomainError
}
