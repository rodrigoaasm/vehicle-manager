package entities

type Car struct {
	Vehicle
}

func NewCar(id string, name string, cor string, serie string) *Car {
	car := Car{
		Vehicle: Vehicle{
			Id:     id,
			Name:   name,
			Cor:    cor,
			Serie:  serie,
			status: false,
		},
	}

	return &car
}

func (car *Car) GetId() string {
	return car.Id
}

func (car *Car) TurnOn() {
	car.status = true
}

func (car *Car) GetStatus() bool {
	return car.status
}
