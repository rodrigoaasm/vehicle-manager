package services

import (
	"errors"

	"demo/domain/entities"
	"demo/domain/interfaces"
)

var vehicles = []interfaces.IVehicle{
	entities.NewCar("a154", "VW GOL", "black", "14885511T125T"),
	entities.NewCar("a155", "VW GOL", "pink", "12885511T125T"),
	entities.NewTrunk("a156", "VW CONSTELLATION", "white", "2885511T125E"),
}

func GetAllVehicle() ([]interfaces.IVehicle, error) {
	return vehicles, nil
}

func GetVehicleById(id string) (interfaces.IVehicle, error) {
	for _, vehicle := range vehicles {

		if vehicle.GetId() == id {
			return vehicle, nil
		}
	}

	return nil, errors.New("NOT_FOUND")
}
