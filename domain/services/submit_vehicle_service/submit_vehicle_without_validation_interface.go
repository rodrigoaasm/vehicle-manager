package submitvehicleservice

import (
	"demo/domain/domainerror"
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/domain/interfaces"

	"github.com/google/uuid"
)

type SubmitVehicleWithOutValidationService struct {
	VehicleRepository interfaces.IVehicleRepository
}

func NewSubmitVehicleWithOutValidationService(vehicleRepository interfaces.IVehicleRepository) *SubmitVehicleWithOutValidationService {
	return &SubmitVehicleWithOutValidationService{VehicleRepository: vehicleRepository}
}

func (service *SubmitVehicleWithOutValidationService) Submit(category, name, cor, serie, licensePlate string, travelled float32) *domainerror.DomainError {
	// gen uuid
	id, errGenUUID := uuid.NewUUID()
	if errGenUUID != nil {
		return domainerror.New(domainerror.DEPENDENCY, "Could not generate an id")
	}

	// create entity
	var vehicle abstract.IVehicle
	if category == "car" {
		vehicle = entities.NewCar(id.String(), name, cor, serie, licensePlate, travelled, false)
	} else if category == "truck" {
		vehicle = entities.NewTrunk(id.String(), name, cor, serie, licensePlate, travelled, false, false)
	} else {
		return domainerror.New(domainerror.INVALID_DATA, "Category unknown")
	}

	err := service.VehicleRepository.SaveVehicle(vehicle)
	return err
}
