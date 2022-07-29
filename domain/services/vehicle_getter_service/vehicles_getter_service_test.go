package vehiclegetterservice_test

import (
	vehiclegetterservice "demo/domain/services/vehicle_getter_service"
	"demo/external/datasource/mock/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var vehicleRepositoryMemo = repositories.NewVehicleRepositoryMemo()
var vehicleGetterService = vehiclegetterservice.NewVehicleTurnService(vehicleRepositoryMemo)

func TestGetAllVehicle(t *testing.T) {
	result, err := vehicleGetterService.GetAllVehicle()

	require.Nil(t, err, "should not return an error  when TestGetAllVehicle was executed")
	assert.Equal(t, len(result), 3, "should return all vehicles")
}

func TestGetVehicleById(t *testing.T) {
	result, err := vehicleGetterService.GetVehicleById("a154")

	require.Nil(t, err, "should not return an error when GetVehicleById was executed")
	assert.Equal(t, result, vehiclegetterservice.VehicleGetterOutput{
		Id:           "a154",
		Category:     "car",
		Name:         "VW GOL",
		Color:        "black",
		Serie:        "14885511T125T",
		LicensePlate: "ABC1234",
		Travelled:    20000,
		Status:       false,
	}, "should return one car")
}
