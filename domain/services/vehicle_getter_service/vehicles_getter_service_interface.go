package vehiclegetterservice

type IVehicleGetterService interface {
	GetAllVehicle() ([]VehicleGetterOutput, error)
	GetVehicleById(id string) (VehicleGetterOutput, error)
}
