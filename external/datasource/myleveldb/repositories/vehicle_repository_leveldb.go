package repositories

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/external/datasource/myleveldb/models"
	"demo/external/utils"
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

type VehicleRepository struct {
	DB *leveldb.DB
}

func NewVehicleRepository(DB *leveldb.DB) *VehicleRepository {
	return &VehicleRepository{DB: DB}
}

func (repo VehicleRepository) transform(marshalledVehicle []byte) (abstract.IVehicle, error) {
	payload := strings.Split(string(marshalledVehicle), "->")

	var vehicleModel models.VehicleModel
	errUnMarshal := json.Unmarshal([]byte(payload[1]), &vehicleModel)

	if payload[0] == "*entities.Car" {
		vehicle := entities.NewCar(
			vehicleModel.Id,
			vehicleModel.Name,
			vehicleModel.Color,
			vehicleModel.Serie,
			vehicleModel.LicensePlate,
			vehicleModel.Status,
		)

		return vehicle, errUnMarshal

	} else if payload[0] == "*entities.Truck" {
		vehicle := entities.NewTrunk(
			vehicleModel.Id,
			vehicleModel.Name,
			vehicleModel.Color,
			vehicleModel.Serie,
			vehicleModel.LicensePlate,
			vehicleModel.Status,
			vehicleModel.StatusAutomaticPilot,
		)
		return vehicle, errUnMarshal

	}

	return nil, errors.New("Invalid entity")
}

func (repo VehicleRepository) SaveVehicle(vehicle abstract.IVehicle) error {

	objectType := reflect.TypeOf(vehicle).String()

	// build model
	var vehicleModel models.VehicleModel

	if utils.IsThisType[entities.Truck](vehicle) {
		truck := vehicle.(*entities.Truck)
		vehicleModel = models.VehicleModel{
			Id:                   truck.Id,
			Name:                 truck.Name,
			Color:                truck.Color,
			Serie:                truck.Serie,
			LicensePlate:         truck.LicensePlate,
			Status:               truck.GetStatus(),
			StatusAutomaticPilot: truck.GetAutomaticPilotStatus(),
		}
	} else if utils.IsThisType[entities.Car](vehicle) {
		car := vehicle.(*entities.Car)
		vehicleModel = models.VehicleModel{
			Id:           car.Id,
			Name:         car.Name,
			Color:        car.Color,
			Serie:        car.Serie,
			LicensePlate: car.LicensePlate,
			Status:       car.GetStatus(),
		}
	} else {
		return errors.New("Type Invalid")
	}

	marshalledVehicle, errMarshal := json.Marshal(vehicleModel)
	if errMarshal != nil {
		return errMarshal
	}

	// Add entity type to unmarshalled entity
	marshalledData := append([]byte(objectType+"->"), marshalledVehicle...)

	// Save
	errStore := repo.DB.Put([]byte(vehicle.GetId()), marshalledData, nil)
	if errStore != nil {
		return errStore
	}

	return nil
}

func (repo VehicleRepository) GetAllVehicle() ([]abstract.IVehicle, error) {
	iter := repo.DB.NewIterator(nil, nil)
	vehicles := []abstract.IVehicle{}

	for iter.Next() {
		marshalledVehicle := iter.Value()
		vehicle, err := repo.transform(marshalledVehicle)
		if err != nil {
			return nil, errors.New("Failed to transform one element")
		}

		vehicles = append(vehicles, vehicle)
	}

	if iter.Error() != nil {
		return nil, errors.New("Failed to get data")
	}

	return vehicles, nil
}

func (repo VehicleRepository) GetVehicleById(id string) (abstract.IVehicle, error) {
	marshalledVehicle, errGet := repo.DB.Get([]byte(id), nil)
	if errGet != nil {
		return nil, errGet
	}
	return repo.transform(marshalledVehicle)
}
