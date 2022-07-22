package repositories

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

type VehicleRepository struct {
	DB *leveldb.DB
}

func (repo VehicleRepository) transform(marshalledVehicle []byte) (abstract.IVehicle, error) {
	payload := strings.Split(string(marshalledVehicle), "->")
	if payload[0] == "*entities.Car" {
		var vehicle entities.Car
		errUnMarshal := json.Unmarshal([]byte(payload[1]), &vehicle)
		return &vehicle, errUnMarshal

	} else if payload[0] == "*entities.Truck" {
		var vehicle entities.Truck
		errUnMarshal := json.Unmarshal([]byte(payload[1]), &vehicle)
		return &vehicle, errUnMarshal

	}

	return nil, errors.New("Invalid entity")
}

func (repo VehicleRepository) SaveVehicle(vehicle abstract.IVehicle) error {
	marshalledVehicle, errMarshal := json.Marshal(vehicle)
	if errMarshal != nil {
		return errMarshal
	}

	// Add entity type to unmarshalled entity
	objectType := reflect.TypeOf(vehicle).String()
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
