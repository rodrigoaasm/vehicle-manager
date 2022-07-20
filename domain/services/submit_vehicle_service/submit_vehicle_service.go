package submitvehicleservice

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/domain/interfaces"
	"errors"

	"github.com/google/uuid"
)

type SubmitVehicleService struct {
	VehicleRepository interfaces.IVehicleRepository
}

func (service SubmitVehicleService) Submit(category, name, cor, serie string) error {

	var vehicle abstract.IVehicle
	id, errGenUUID := uuid.NewUUID()

	if errGenUUID != nil {
		return errors.New("Could not generate an id")
	}

	if category == "car" {
		vehicle = entities.NewCar(id.String(), name, cor, serie)
	} else if category == "truck" {
		vehicle = entities.NewTrunk(id.String(), name, cor, serie)
	} else {
		return errors.New("Category unknown")
	}

	err := service.VehicleRepository.SaveVehicle(vehicle)

	return err
}
