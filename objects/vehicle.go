package objects

type IVehicle interface {
	TurnOn()
	GetStatus() bool
}

type Vehicle struct {
	Id     string
	Name   string
	Cor    string
	Serie  string
	status bool
}

func (vehicle *Vehicle) TurnOn() {
	vehicle.status = true
}

func (vehicle *Vehicle) GetStatus() bool {
	return vehicle.status
}
