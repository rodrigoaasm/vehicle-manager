package objects

type Car struct {
	Vehicle
}

func NewCar(id string, name string, cor string, serie string) *Car {
	car := Car{
		Vehicle: Vehicle{
			Id:    id,
			Name:  name,
			Cor:   cor,
			Serie: serie,
		},
	}
	car.status = false

	return &car
}
