package vehicleupdateservice

import (
	"demo/domain/domainerror"
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/domain/interfaces"
	"demo/external/utils"
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

type VehicleUpdateService struct {
	VehicleRepository interfaces.IVehicleRepository
}

func NewVehicleUpdateService(vehicleRepository interfaces.IVehicleRepository) *VehicleUpdateService {
	return &VehicleUpdateService{VehicleRepository: vehicleRepository}
}

func (service VehicleUpdateService) callTurn(attrName string, attrValue string, vehicle abstract.IVehicle) *domainerror.DomainError {
	if attrValue != "" {
		attrMethod := methods[attrName]

		if method := attrMethod[attrValue]; method != "" {
			v := reflect.ValueOf(vehicle)
			m := v.MethodByName(method)
			m.Call([]reflect.Value{})
			return nil
		} else {
			return domainerror.New(domainerror.INVALID_DATA, "The "+attrName+" value is invalid! Valid values ​​(on/off).")
		}
	}

	return nil
}

func (service VehicleUpdateService) checkTravelled(newTravelled float32, oldTravelled float32) *domainerror.DomainError {
	if newTravelled <= oldTravelled {
		return domainerror.New(domainerror.INVALID_DATA, "The new value of travelled is less than current value of travelled.")
	}

	return nil
}

func (service VehicleUpdateService) Turn(payload VehicleUpdateInput) *domainerror.DomainError {
	iVehicle, err := service.VehicleRepository.GetVehicleById(payload.Id)
	if err != nil {
		return domainerror.New(domainerror.NOT_FOUND, "Vehicle Not found")
	}

	if utils.IsThisType[entities.Truck](iVehicle) {
		truck := iVehicle.(*entities.Truck)
		// travelled update
		if err = service.checkTravelled(payload.Travelled, truck.Travelled); err != nil && payload.Status == "off" {
			return err
		}
		truck.Travelled = payload.Travelled

		// StatusAutomaticPilot on/off value for method
		if err = service.callTurn("StatusAutomaticPilot", payload.StatusAutomaticPilot, iVehicle); err != nil {
			return err
		}
	} else if payload.StatusAutomaticPilot != "" {
		return domainerror.New(domainerror.INVALID_DATA, "This type of vehicle does not have autopilot")
	} else {
		car := iVehicle.(*entities.Car)
		// travelled update
		if err = service.checkTravelled(payload.Travelled, car.Travelled); err != nil && payload.Status == "off" {
			return err
		}
		car.Travelled = payload.Travelled
	}

	// Status on/off value for method
	if err = service.callTurn("Status", payload.Status, iVehicle); err != nil {
		return err
	}

	err = service.VehicleRepository.SaveVehicle(iVehicle)
	return err
}
