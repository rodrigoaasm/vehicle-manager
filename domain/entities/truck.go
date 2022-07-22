package entities

type Truck struct {
	Vehicle
	automaticPilotStatus bool
}

func NewTrunk(id, name, color, serie, licensePlate string) *Truck {
	truck := Truck{
		Vehicle: Vehicle{
			Id:           id,
			Name:         name,
			Color:        color,
			Serie:        serie,
			status:       false,
			LicensePlate: licensePlate,
		},
		automaticPilotStatus: false,
	}

	return &truck
}

func (truck *Truck) GetId() string {
	return truck.Id
}

func (truck *Truck) TurnOn() {
	truck.status = true
}

func (truck *Truck) GetStatus() bool {
	return truck.status
}

func (truck *Truck) TurnOnAutomaticPilot() {
	truck.automaticPilotStatus = true
}

func (truck *Truck) GetAutomaticPilotStatus() bool {
	return truck.automaticPilotStatus
}
