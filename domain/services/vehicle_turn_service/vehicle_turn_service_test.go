package vehicleturnservice_test

import (
	"demo/domain/entities"
	vehicleturnservice "demo/domain/services/vehicle_turn_service"
	"demo/external/datasource/mock/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var vehicleRepositoryMemo = repositories.NewVehicleRepositoryMemo()
var vehicleTurnService = vehicleturnservice.NewVehicleTurnService(vehicleRepositoryMemo)

func TestTurn(t *testing.T) {
	payload := vehicleturnservice.VehicleTurnInput{
		Id:                   "id",
		Status:               "on",
		StatusAutomaticPilot: "on",
	}
	err := vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Error(),
		"Vehicle Not found",
		"should return an error when vehicle does not exist",
	)

	payload = vehicleturnservice.VehicleTurnInput{
		Id:                   "a154",
		Status:               "on",
		StatusAutomaticPilot: "on",
	}
	err = vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Error(),
		"This type of vehicle does not have autopilot",
		"should return an error when the vehicle is not a truck",
	)

	payload = vehicleturnservice.VehicleTurnInput{
		Id:                   "a154",
		Status:               "invalid",
		StatusAutomaticPilot: "",
	}
	err = vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Error(),
		"The Status value is invalid! Valid values ​​(on/off).",
		"should return an error when the status value is invalid",
	)

	payload = vehicleturnservice.VehicleTurnInput{
		Id:                   "a156",
		Status:               "on",
		StatusAutomaticPilot: "invalid",
	}
	err = vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Error(),
		"The StatusAutomaticPilot value is invalid! Valid values ​​(on/off).",
		"should return an error when the StatusAutomaticPilot value is invalid",
	)

	payload = vehicleturnservice.VehicleTurnInput{
		Id:                   "a154",
		Status:               "on",
		StatusAutomaticPilot: "",
	}
	errTurnCar := vehicleTurnService.Turn(payload)
	require.Nil(t, errTurnCar)
	assert.Equal(t, true, vehicleRepositoryMemo.LastSavedVehicle.GetStatus(), "should turn on the car")

	payload = vehicleturnservice.VehicleTurnInput{
		Id:                   "a156",
		Status:               "on",
		StatusAutomaticPilot: "on",
	}
	errTurnTruck := vehicleTurnService.Turn(payload)
	savedVehicle := vehicleRepositoryMemo.LastSavedVehicle.(*entities.Truck)

	require.Nil(t, errTurnTruck)
	assert.Equal(t, true, savedVehicle.GetStatus(), "should turn on the truck")
	assert.Equal(t, true, savedVehicle.GetAutomaticPilotStatus(), "should turn on automatic pilot in the truck")
}
