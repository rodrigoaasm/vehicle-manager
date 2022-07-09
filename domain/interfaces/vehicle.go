package interfaces

type IVehicle interface {
	GetId() string
	TurnOn()
	GetStatus() bool
}
