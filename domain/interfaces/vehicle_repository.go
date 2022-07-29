package interfaces

import (
	"demo/domain/domainerror"
	"demo/domain/entities/abstract"
)

type IVehicleRepository interface {
	GetAllVehicle() ([]abstract.IVehicle, *domainerror.DomainError)
	GetVehicleById(id string) (abstract.IVehicle, *domainerror.DomainError)
	SaveVehicle(vehicle abstract.IVehicle) *domainerror.DomainError
	GetVehicleByLicensePlate(licensePlate string) (abstract.IVehicle, *domainerror.DomainError)
	GetVehicleBySerie(serie string) (abstract.IVehicle, *domainerror.DomainError)
}
