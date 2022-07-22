package submitvehicleservice

type ISubmitVehicleService interface {
	Submit(category, name, cor, serie, licensePlate string) error
}
