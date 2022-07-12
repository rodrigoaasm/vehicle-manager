package interfaces

import "demo/domain/entities/abstract"

type VehicleRepository interface {
	GetAllVehicle() ([]abstract.IVehicle, error)
	GetVehicleById(id string) (abstract.IVehicle, error)
}
