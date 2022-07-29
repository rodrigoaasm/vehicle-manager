package entities

type Car struct {
	Vehicle
}

func NewCar(id, name, color, serie, licensePlate string, travelled float32, status bool) *Car {
	car := Car{
		Vehicle: Vehicle{
			Id:           id,
			Name:         name,
			Color:        color,
			Serie:        serie,
			status:       status,
			LicensePlate: licensePlate,
			Travelled:    travelled,
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

func (car *Car) TurnOff() {
	car.status = false
}

func (car *Car) GetStatus() bool {
	return car.status
}
