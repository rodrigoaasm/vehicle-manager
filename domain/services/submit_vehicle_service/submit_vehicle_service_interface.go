package submitvehicleservice

type ISubmitVehicleService interface {
	Submit(category, name, cor, serie string) error
}
