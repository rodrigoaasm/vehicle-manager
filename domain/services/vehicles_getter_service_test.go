package services_test

import (
	"demo/domain/entities"
	"demo/domain/services"
	"demo/external/datasource/mock/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var vehicleRepositoryMemo = repositories.VehicleRepositoryMemo{}
var vehicleGetterService = services.VehicleGetterService{VehicleRepository: vehicleRepositoryMemo}

func TestGetAllVehicle_1(t *testing.T) {
	result, err := vehicleGetterService.GetAllVehicle()

	require.Nil(t, err, "should not return an error")
	assert.Equal(t, len(result), 3, "should return all vehicles")
}

func TestGetAllVehicle_2(t *testing.T) {
	result, err := vehicleGetterService.GetVehicleById("a154")

	require.Nil(t, err, "should not return an error")
	assert.Equal(t, result, entities.NewCar("a154", "VW GOL", "black", "14885511T125T"), "should return one car")
}
