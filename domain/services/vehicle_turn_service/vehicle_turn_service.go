package vehicleturnservice

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/domain/interfaces"
	"demo/external/utils"
	"errors"
	"reflect"
)

var methods = map[string]map[string]string{
	"Status": {
		"on":  "TurnOn",
		"off": "TurnOff",
	},
	"StatusAutomaticPilot": {
		"on":  "TurnOnAutomaticPilot",
		"off": "TurnOffAutomaticPilot",
	},
}

type VehicleTurnService struct {
	VehicleRepository interfaces.IVehicleRepository
}

func NewVehicleTurnService(vehicleRepository interfaces.IVehicleRepository) *VehicleTurnService {
	return &VehicleTurnService{VehicleRepository: vehicleRepository}
}

func (service VehicleTurnService) callTurn(attrName string, attrValue string, vehicle abstract.IVehicle) error {
	if attrValue != "" {
		attrMethod := methods[attrName]

		if method := attrMethod[attrValue]; method != "" {
			v := reflect.ValueOf(vehicle)
			m := v.MethodByName(method)
			m.Call([]reflect.Value{})
			return nil
		} else {
			return errors.New("The " + attrName + " value is invalid! Valid values ​​(on/off).")
		}
	}

	return nil
}

func (service VehicleTurnService) Turn(payload VehicleTurnInput) error {
	vehicle, err := service.VehicleRepository.GetVehicleById(payload.Id)
	if err != nil {
		return errors.New("Vehicle Not found")
	}

	if utils.IsThisType[entities.Truck](vehicle) {
		// StatusAutomaticPilot on/off value for method
		if err = service.callTurn("StatusAutomaticPilot", payload.StatusAutomaticPilot, vehicle); err != nil {
			return err
		}
	} else if payload.StatusAutomaticPilot != "" {
		return errors.New("This type of vehicle does not have autopilot")
	}

	// Status on/off value for method
	if err = service.callTurn("Status", payload.Status, vehicle); err != nil {
		return err
	}

	err = service.VehicleRepository.SaveVehicle(vehicle)
	return err
}
