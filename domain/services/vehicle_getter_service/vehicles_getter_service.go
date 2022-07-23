package vehiclegetterservice

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/domain/interfaces"
	"demo/external/utils"
	"errors"
)

type VehicleGetterService struct {
	VehicleRepository interfaces.IVehicleRepository
}

func NewVehicleTurnService(vehicleRepository interfaces.IVehicleRepository) *VehicleGetterService {
	return &VehicleGetterService{VehicleRepository: vehicleRepository}
}

func (service VehicleGetterService) transform(vehicle abstract.IVehicle) (VehicleGetterOutput, error) {
	if utils.IsThisType[entities.Car](vehicle) {
		car := vehicle.(*entities.Car)

		return VehicleGetterOutput{
			Id:           car.Id,
			Category:     "car",
			Name:         car.Name,
			Color:        car.Color,
			Serie:        car.Serie,
			LicensePlate: car.LicensePlate,
			Status:       car.GetStatus(),
		}, nil

	} else if utils.IsThisType[entities.Truck](vehicle) {
		truck := vehicle.(*entities.Truck)

		return VehicleGetterOutput{
			Id:             truck.Id,
			Category:       "truck",
			Name:           truck.Name,
			Color:          truck.Color,
			Serie:          truck.Serie,
			LicensePlate:   truck.LicensePlate,
			Status:         truck.GetStatus(),
			AutomaticPilot: truck.GetAutomaticPilotStatus(),
		}, nil
	}

	return VehicleGetterOutput{}, errors.New("Type unknown")
}

func (service VehicleGetterService) GetAllVehicle() ([]VehicleGetterOutput, error) {
	vehicles, err := service.VehicleRepository.GetAllVehicle()
	outputs := []VehicleGetterOutput{}

	for _, vehicle := range vehicles {
		output, err := service.transform(vehicle)

		if err != nil {
			return nil, errors.New("Object format fails")
		}

		outputs = append(outputs, output)
	}

	return outputs, err
}

func (service VehicleGetterService) GetVehicleById(id string) (VehicleGetterOutput, error) {
	vehicle, err := service.VehicleRepository.GetVehicleById(id)

	if err != nil {
		return VehicleGetterOutput{}, err
	}

	return service.transform(vehicle)
}
