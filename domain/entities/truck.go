package entities

type Truck struct {
	Vehicle
	automaticPilotStatus bool
}

func NewTrunk(id string, name string, cor string, serie string) *Truck {
	truck := Truck{
		Vehicle: Vehicle{
			Id:     id,
			Name:   name,
			Cor:    cor,
			Serie:  serie,
			status: false,
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
