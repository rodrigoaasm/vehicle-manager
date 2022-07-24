package repositories_test

import (
	"demo/domain/entities"
	"demo/domain/entities/abstract"
	"demo/external/datasource/myleveldb"
	"demo/external/datasource/myleveldb/repositories"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var db, errDatabase = myleveldb.NewDatabase("test/" + strconv.Itoa(rand.Intn(100)))
var vehicleRepository = repositories.NewVehicleRepository(db)

var car = entities.NewCar("a157", "VW GOL", "black", "14885511T125T", "ABC1452", true)
var truck = entities.NewTrunk("a156", "VW CONSTELLATION", "white", "2885511T125E", "ABC1451", false, false)

func TestSaveVehicle(t *testing.T) {
	err := vehicleRepository.SaveVehicle(car)
	result, _ := vehicleRepository.GetVehicleById("a157")
	require.Nil(t, err, "should save the Car")
	assert.Equal(t, car, result, "should return a turned on car")

	err = vehicleRepository.SaveVehicle(truck)
	require.Nil(t, err, "should save the Truck")
}

func TestGetVehicleById(t *testing.T) {
	result, _ := vehicleRepository.GetVehicleById("a157")
	assert.Equal(t, car, result, "should return the Car with id 'a154'")

	result, _ = vehicleRepository.GetVehicleById("a156")
	assert.Equal(t, truck, result, "should return the Truck with id 'a156'")

	_, err := vehicleRepository.GetVehicleById("a1561")
	require.Error(t, err, "should return an error when the id does not exist")
}

func TestGetVehicleBySerie(t *testing.T) {
	result, _ := vehicleRepository.GetVehicleBySerie("14885511T125T")
	assert.Equal(t, car, result, "should return the vehicle with serie '14885511T125T'")

	_, err := vehicleRepository.GetVehicleBySerie("invalid")
	require.Error(t, err, "should return an error when the serie does not exist")
}

func TestGetVehicleByLicensePlate(t *testing.T) {
	result, _ := vehicleRepository.GetVehicleByLicensePlate("ABC1452")
	assert.Equal(t, car, result, "should return the vehicle with serie '14885511T125T'")

	_, err := vehicleRepository.GetVehicleByLicensePlate("invalid")
	require.Error(t, err, "should return an error when the serie does not exist")
}

func TestGetAllVehicle(t *testing.T) {
	result, _ := vehicleRepository.GetAllVehicle()
	vehicles := []abstract.IVehicle{truck, car}

	assert.Equal(t, vehicles, result, "should return all vehicles")
}
