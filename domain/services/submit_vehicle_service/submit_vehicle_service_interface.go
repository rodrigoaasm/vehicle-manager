package submitvehicleservice

import (
	"demo/domain/domainerror"
)

type ISubmitVehicleService interface {
	Submit(category, name, cor, serie, licensePlate string) *domainerror.DomainError
}
