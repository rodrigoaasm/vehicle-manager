package objects

type Truck struct {
	Vehicle
	automaticPilotStatus bool
}

func NewTrunk(id string, name string, cor string, serie string) *Truck {
	trunck := Truck{
		Vehicle: Vehicle{
			Id:    id,
			Name:  name,
			Cor:   cor,
			Serie: serie,
		},
	}
	trunck.status = false
	trunck.automaticPilotStatus = false

	return &trunck
}

func (truck *Truck) TurnOnAutomaticPilot() {
	truck.automaticPilotStatus = true
}

func (truck *Truck) GetAutomaticPilotStatus() bool {
	return truck.automaticPilotStatus
}
