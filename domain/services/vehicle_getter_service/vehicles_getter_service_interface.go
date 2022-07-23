package vehiclegetterservice

type VehicleGetterOutput struct {
	Id             string `json:"id"`
	Category       string `json:"category"`
	Name           string `json:"name"`
	Color          string `json:"color"`
	Serie          string `json:"serie"`
	LicensePlate   string `json:"licensePlate"`
	Status         bool   `json:"status"`
	AutomaticPilot bool   `json:"automaticPilot"`
}

type IVehicleGetterService interface {
	GetAllVehicle() ([]VehicleGetterOutput, error)
	GetVehicleById(id string) (VehicleGetterOutput, error)
}
