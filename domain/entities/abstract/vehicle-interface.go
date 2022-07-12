package abstract

type IVehicle interface {
	GetId() string
	TurnOn()
	GetStatus() bool
}
