package repositories

import (
	"demo/domain/domainerror"
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/external/datasource/myleveldb"
	"demo/external/datasource/myleveldb/models"
	"demo/external/utils"
	"encoding/json"
	"reflect"
	"strings"
)

type VehicleRepository struct {
	DB *myleveldb.Database
}

func NewVehicleRepository(DB *myleveldb.Database) *VehicleRepository {
	return &VehicleRepository{DB: DB}
}

func (repo VehicleRepository) transform(marshalledVehicle []byte) (abstract.IVehicle, *domainerror.DomainError) {
	payload := strings.Split(string(marshalledVehicle), "->")

	var vehicleModel models.VehicleModel
	errUnMarshal := json.Unmarshal([]byte(payload[1]), &vehicleModel)

	if errUnMarshal != nil {
		return nil, domainerror.New(domainerror.INVALID_DATA, errUnMarshal.Error())
	}

	if payload[0] == "*entities.Car" {
		vehicle := entities.NewCar(
			vehicleModel.Id,
			vehicleModel.Name,
			vehicleModel.Color,
			vehicleModel.Serie,
			vehicleModel.LicensePlate,
			vehicleModel.Status,
		)

		return vehicle, nil

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
		return vehicle, nil

	}

	return nil, domainerror.New(domainerror.DATABASE, "Invalid type")
}

func (repo VehicleRepository) fromEntityToModel(vehicle abstract.IVehicle) (models.VehicleModel, *domainerror.DomainError) {
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
		return vehicleModel, nil
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
		return vehicleModel, nil
	} else {
		return models.VehicleModel{}, domainerror.New(domainerror.DATABASE, "Type Invalid")
	}
}

func (repo VehicleRepository) checkExistence(vehicle models.VehicleModel) *domainerror.DomainError {
	iVehicle, errPlate := repo.GetVehicleByLicensePlate(vehicle.LicensePlate)
	if errPlate == nil && iVehicle.GetId() != vehicle.Id {
		return domainerror.New(domainerror.CONFLICT, "There is already a vehicle with this license plate")
	}

	iVehicle, errSerie := repo.GetVehicleBySerie(vehicle.Serie)
	if errSerie == nil && iVehicle.GetId() != vehicle.Id {
		return domainerror.New(domainerror.CONFLICT, "There is already a vehicle with this serie")
	}

	return nil
}

func (repo VehicleRepository) SaveVehicle(vehicle abstract.IVehicle) *domainerror.DomainError {
	objectType := reflect.TypeOf(vehicle).String()

	vehicleModel, errToModel := repo.fromEntityToModel(vehicle)
	if errToModel != nil {
		return errToModel
	}

	if err := repo.checkExistence(vehicleModel); err != nil {
		return err
	}

	marshalledVehicle, errMarshal := json.Marshal(vehicleModel)
	if errMarshal != nil {
		return domainerror.New(domainerror.DATABASE, errMarshal.Error())
	}
	marshalledData := append([]byte(objectType+"->"), marshalledVehicle...)
	if err := repo.DB.Data.Put([]byte(vehicle.GetId()), marshalledData, nil); err != nil {
		return domainerror.New(domainerror.DATABASE, err.Error())
	}

	if err := repo.DB.Index.Put([]byte("serie:"+vehicleModel.Serie), []byte(vehicleModel.Id), nil); err != nil {
		repo.DB.Data.Delete([]byte(vehicleModel.Id), nil)
		return domainerror.New(domainerror.DATABASE, err.Error())
	}

	if err := repo.DB.Index.Put([]byte("licensePlate:"+vehicleModel.LicensePlate), []byte(vehicleModel.Id), nil); err != nil {
		repo.DB.Data.Delete([]byte(vehicleModel.Id), nil)
		repo.DB.Index.Delete([]byte("serie:"+vehicleModel.Serie), nil)
		return domainerror.New(domainerror.DATABASE, err.Error())
	}

	return nil
}

func (repo VehicleRepository) GetAllVehicle() ([]abstract.IVehicle, *domainerror.DomainError) {
	iter := repo.DB.Data.NewIterator(nil, nil)
	vehicles := []abstract.IVehicle{}

	for iter.Next() {
		marshalledVehicle := iter.Value()
		vehicle, err := repo.transform(marshalledVehicle)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	if iter.Error() != nil {
		return nil, domainerror.New(domainerror.DATABASE, "Failed to get data")
	}

	return vehicles, nil
}

func (repo VehicleRepository) GetVehicleById(id string) (abstract.IVehicle, *domainerror.DomainError) {
	marshalledVehicle, errGet := repo.DB.Data.Get([]byte(id), nil)
	if errGet != nil {
		return nil, domainerror.New(domainerror.NOT_FOUND, "Vehicle not found")
	}
	return repo.transform(marshalledVehicle)
}

func (repo VehicleRepository) GetVehicleBySerie(serie string) (abstract.IVehicle, *domainerror.DomainError) {
	id, errGet := repo.DB.Index.Get([]byte("serie:"+serie), nil)
	if errGet != nil {
		return nil, domainerror.New(domainerror.NOT_FOUND, "Vehicle not found")
	}

	return repo.GetVehicleById(string(id))
}

func (repo VehicleRepository) GetVehicleByLicensePlate(serie string) (abstract.IVehicle, *domainerror.DomainError) {
	id, errGet := repo.DB.Index.Get([]byte("licensePlate:"+serie), nil)
	if errGet != nil {
		return nil, domainerror.New(domainerror.NOT_FOUND, "Vehicle not found")
	}

	return repo.GetVehicleById(string(id))
}
