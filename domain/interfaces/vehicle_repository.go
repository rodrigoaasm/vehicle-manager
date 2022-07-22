package interfaces

import "demo/domain/entities/abstract"

type IVehicleRepository interface {
	GetAllVehicle() ([]abstract.IVehicle, error)
	GetVehicleById(id string) (abstract.IVehicle, error)
	SaveVehicle(vehicle abstract.IVehicle) error
}
