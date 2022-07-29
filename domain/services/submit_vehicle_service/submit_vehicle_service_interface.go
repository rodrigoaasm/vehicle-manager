package submitvehicleservice

import (
	"demo/domain/domainerror"
)

type ISubmitVehicleService interface {
	Submit(category, name, cor, serie, licensePlate string, travelled float32) *domainerror.DomainError
}
