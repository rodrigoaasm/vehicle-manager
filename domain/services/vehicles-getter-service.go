package services

import (
	"demo/domain/entities/abstract"
	"demo/domain/interfaces"
)

type VehicleGetterService struct {
	VehicleRepository interfaces.VehicleRepository
}

func (service VehicleGetterService) GetAllVehicle() ([]abstract.IVehicle, error) {
	result, err := service.VehicleRepository.GetAllVehicle()

	return result, err
}

func (service VehicleGetterService) GetVehicleById(id string) (abstract.IVehicle, error) {
	result, err := service.VehicleRepository.GetVehicleById(id)

	return result, err
}
