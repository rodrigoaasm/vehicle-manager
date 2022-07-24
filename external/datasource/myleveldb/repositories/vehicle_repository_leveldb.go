package repositories

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/external/datasource/myleveldb"
	"demo/external/datasource/myleveldb/models"
	"demo/external/utils"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

type VehicleRepository struct {
	DB *myleveldb.Database
}

func NewVehicleRepository(DB *myleveldb.Database) *VehicleRepository {
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
	marshalledData := append([]byte(objectType+"->"), marshalledVehicle...)
	if errStore := repo.DB.Data.Put([]byte(vehicle.GetId()), marshalledData, nil); errStore != nil {
		return errStore
	}

	if errIndexSerie := repo.DB.Index.Put([]byte("serie:"+vehicleModel.Serie), []byte(vehicleModel.Id), nil); errIndexSerie != nil {
		repo.DB.Data.Delete([]byte(vehicleModel.Id), nil)
		return errIndexSerie
	}

	if errIndexSerie := repo.DB.Index.Put([]byte("licensePlate:"+vehicleModel.LicensePlate), []byte(vehicleModel.Id), nil); errIndexSerie != nil {
		repo.DB.Data.Delete([]byte(vehicleModel.Id), nil)
		repo.DB.Index.Delete([]byte("serie:"+vehicleModel.Serie), nil)
		return errIndexSerie
	}

	return nil
}

func (repo VehicleRepository) GetAllVehicle() ([]abstract.IVehicle, error) {
	iter := repo.DB.Data.NewIterator(nil, nil)
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
	marshalledVehicle, errGet := repo.DB.Data.Get([]byte(id), nil)
	if errGet != nil {
		return nil, errGet
	}
	return repo.transform(marshalledVehicle)
}

func (repo VehicleRepository) GetVehicleBySerie(serie string) (abstract.IVehicle, error) {
	id, errGet := repo.DB.Index.Get([]byte("serie:"+serie), nil)
	if errGet != nil {
		return nil, errGet
	}

	return repo.GetVehicleById(string(id))
}

func (repo VehicleRepository) GetVehicleByLicensePlate(serie string) (abstract.IVehicle, error) {
	id, errGet := repo.DB.Index.Get([]byte("licensePlate:"+serie), nil)
	if errGet != nil {
		return nil, errGet
	}

	return repo.GetVehicleById(string(id))
}
