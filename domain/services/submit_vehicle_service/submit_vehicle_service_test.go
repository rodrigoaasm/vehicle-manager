package submitvehicleservice_test

import (
	"demo/domain/entities"
	submitvehicleservice "demo/domain/services/submit_vehicle_service"
	"demo/external/datasource/mock/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var vehicleRepositoryMemo = repositories.VehicleRepositoryMemo{}
var vehicleGetterService = submitvehicleservice.SubmitVehicleService{VehicleRepository: vehicleRepositoryMemo}

func TestSubmit(t *testing.T) {
	car := entities.NewCar("", "VW GOL", "black", "14885511T125T")
	err := vehicleGetterService.Submit("car", "VW GOL", "black", "14885511T125T")

	require.Nil(t, err, "should not return an error when submit a car")

	registeredCar, err := vehicleRepositoryMemo.GetVehicleById("")
	require.Nil(t, err, "should not return an error when get registered car")
	assert.Equal(t, car, registeredCar)
}
