package repositories

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"errors"
)

var vehicles = []abstract.IVehicle{
	entities.NewCar("a154", "VW GOL", "black", "14885511T125T"),
	entities.NewCar("a155", "VW GOL", "pink", "12885511T125T"),
	entities.NewTrunk("a156", "VW CONSTELLATION", "white", "2885511T125E"),
}

type VehicleRepositoryMemo struct {
}

func (repo VehicleRepositoryMemo) saveVehicle(vehicle abstract.IVehicle) error {
	return nil
}

func (repo VehicleRepositoryMemo) GetAllVehicle() ([]abstract.IVehicle, error) {
	return vehicles, nil
}

func (repo VehicleRepositoryMemo) GetVehicleById(id string) (abstract.IVehicle, error) {
	for _, vehicle := range vehicles {

		if vehicle.GetId() == id {
			return vehicle, nil
		}
	}

	return nil, errors.New("NOT_FOUND")
}