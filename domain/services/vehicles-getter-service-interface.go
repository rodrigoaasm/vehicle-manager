package services

import "demo/domain/entities/abstract"

type IVehicleGetterService interface {
	GetAllVehicle() ([]abstract.IVehicle, error)
	GetVehicleById(id string) (abstract.IVehicle, error)
}
