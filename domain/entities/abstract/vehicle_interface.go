package abstract

type IVehicle interface {
	GetId() string
	TurnOn()
	TurnOff()
	GetStatus() bool
}
