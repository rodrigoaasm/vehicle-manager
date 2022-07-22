package repositories_test

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/external/datasource/myleveldb"
	"demo/external/datasource/myleveldb/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var db, errDatabase = myleveldb.CreateLevelDB("test/data")
var vehicleRepository = repositories.VehicleRepository{DB: db}

var car = entities.NewCar("a154", "VW GOL", "black", "14885511T125T", "ABC1452")
var truck = entities.NewTrunk("a156", "VW CONSTELLATION", "white", "2885511T125E", "ABC1452")

func TestCreateLevelDB(t *testing.T) {
	require.Nil(t, errDatabase, "should create a local database")
}

func TestSaveVehicle(t *testing.T) {
	err := vehicleRepository.SaveVehicle(car)
	require.Nil(t, err, "should save the Car")

	err = vehicleRepository.SaveVehicle(truck)
	require.Nil(t, err, "should save the Truck")
}

func TestGetVehicleById(t *testing.T) {
	result, _ := vehicleRepository.GetVehicleById("a154")
	assert.Equal(t, car, result, "should return the Car with id 'a154'")

	result, _ = vehicleRepository.GetVehicleById("a156")
	assert.Equal(t, truck, result, "should return the Truck with id 'a156'")

	_, err := vehicleRepository.GetVehicleById("a1561")
	require.Error(t, err, "should return an error when the id does not exist")
}

func TestGetAllVehicle(t *testing.T) {
	result, _ := vehicleRepository.GetAllVehicle()
	vehicles := []abstract.IVehicle{car, truck}

	assert.Equal(t, vehicles, result, "should return all vehicles")
}
