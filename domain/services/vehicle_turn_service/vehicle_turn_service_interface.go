package vehicleturnservice

type VehicleTurnInput struct {
	Id                   string
	Status               string
	StatusAutomaticPilot string
}

type IVehicleTurnService interface {
	Turn(vehicle VehicleTurnInput) error
}
