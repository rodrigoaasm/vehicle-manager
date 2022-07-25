package submitvehicleservice

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/domain/interfaces"
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type SubmitVehicleService struct {
	VehicleRepository interfaces.IVehicleRepository
}

func NewSubmitVehicleService(vehicleRepository interfaces.IVehicleRepository) *SubmitVehicleService {
	return &SubmitVehicleService{VehicleRepository: vehicleRepository}
}

func (service *SubmitVehicleService) Submit(category, name, cor, serie, licensePlate string) error {
	// Validation
	if len(name) < 3 || len(name) > 25 {
		return errors.New("The name must be greater than 25 or less than 3.")
	}

	matchPlate, errPlate := regexp.MatchString("[A-Z]{3}[0-9][0-9A-Z][0-9]{2}", licensePlate)
	if errPlate != nil || !matchPlate {
		return errors.New("License Plate invalid")
	}

	// gen uuid
	id, errGenUUID := uuid.NewUUID()
	if errGenUUID != nil {
		return errors.New("Could not generate an id")
	}

	// create entity
	var vehicle abstract.IVehicle
	if category == "car" {
		vehicle = entities.NewCar(id.String(), name, cor, serie, licensePlate, false)
	} else if category == "truck" {
		vehicle = entities.NewTrunk(id.String(), name, cor, serie, licensePlate, false, false)
	} else {
		return errors.New("Category unknown")
	}

	err := service.VehicleRepository.SaveVehicle(vehicle)
	return err
}
