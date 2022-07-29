package repositories

import (
	"demo/domain/domainerror"
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/external/utils"
)

var vehicles = []abstract.IVehicle{
	entities.NewCar("a154", "VW GOL", "black", "14885511T125T", "ABC1234", false),
	entities.NewCar("a155", "VW GOL", "pink", "12885511T125T", "ABC1234", false),
	entities.NewTrunk("a156", "VW CONSTELLATION", "white", "2885511T125E", "ABC1234", false, false),
}

type VehicleRepositoryMemo struct {
	LastSavedVehicle abstract.IVehicle
}

func NewVehicleRepositoryMemo() *VehicleRepositoryMemo {
	return &VehicleRepositoryMemo{}
}

func (repo *VehicleRepositoryMemo) SaveVehicle(vehicle abstract.IVehicle) *domainerror.DomainError {
	repo.LastSavedVehicle = vehicle
	vehicles = append(vehicles, vehicle)
	return nil
}

func (repo *VehicleRepositoryMemo) GetAllVehicle() ([]abstract.IVehicle, *domainerror.DomainError) {
	return vehicles, nil
}

func (repo *VehicleRepositoryMemo) GetVehicleById(id string) (abstract.IVehicle, *domainerror.DomainError) {
	for _, vehicle := range vehicles {

		if vehicle.GetId() == id {
			return vehicle, nil
		}
	}

	return nil, domainerror.New(domainerror.NOT_FOUND, "NOT_FOUND")
}

func (repo *VehicleRepositoryMemo) transform(iVehicle abstract.IVehicle) (entities.Vehicle, *domainerror.DomainError) {
	if utils.IsThisType[entities.Car](iVehicle) {
		return iVehicle.(*entities.Car).Vehicle, nil
	} else if utils.IsThisType[entities.Truck](iVehicle) {
		return iVehicle.(*entities.Truck).Vehicle, nil
	} else {
		return entities.Vehicle{}, domainerror.New(domainerror.INVALID_DATA, "Type invalid")
	}

}

func (repo *VehicleRepositoryMemo) GetVehicleBySerie(serie string) (abstract.IVehicle, *domainerror.DomainError) {
	for _, iVehicle := range vehicles {
		vehicle, err := repo.transform(iVehicle)
		if err != nil {
			return nil, domainerror.New(domainerror.NOT_FOUND, "NOT_FOUND")
		}

		if vehicle.Serie == serie {
			return iVehicle, nil
		}
	}

	return nil, domainerror.New(domainerror.NOT_FOUND, "NOT_FOUND")
}

func (repo *VehicleRepositoryMemo) GetVehicleByLicensePlate(licensePlate string) (abstract.IVehicle, *domainerror.DomainError) {
	for _, iVehicle := range vehicles {
		vehicle, err := repo.transform(iVehicle)
		if err != nil {
			return nil, domainerror.New(domainerror.NOT_FOUND, "NOT_FOUND")
		}

		if vehicle.LicensePlate == licensePlate {
			return iVehicle, nil
		}
	}

	return nil, domainerror.New(domainerror.NOT_FOUND, "NOT_FOUND")
}
