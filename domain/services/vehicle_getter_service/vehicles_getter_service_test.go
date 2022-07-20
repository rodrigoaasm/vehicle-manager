package vehiclegetterservice_test

import (
	"demo/domain/entities"
	vehiclegetterservice "demo/domain/services/vehicle_getter_service"
	"demo/external/datasource/mock/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var vehicleRepositoryMemo = repositories.VehicleRepositoryMemo{}
var vehicleGetterService = vehiclegetterservice.VehicleGetterService{VehicleRepository: vehicleRepositoryMemo}

func TestGetAllVehicle(t *testing.T) {
	result, err := vehicleGetterService.GetAllVehicle()

	require.Nil(t, err, "should not return an error  when TestGetAllVehicle was executed")
	assert.Equal(t, len(result), 3, "should return all vehicles")
}

func TestGetVehicleById(t *testing.T) {
	result, err := vehicleGetterService.GetVehicleById("a154")

	require.Nil(t, err, "should not return an error when GetVehicleById was executed")
	assert.Equal(t, result, entities.NewCar("a154", "VW GOL", "black", "14885511T125T"), "should return one car")
}
