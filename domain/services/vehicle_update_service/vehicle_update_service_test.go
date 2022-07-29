package vehicleupdateservice_test

import (
	"demo/domain/entities"
	vehicleupdateservice "demo/domain/services/vehicle_update_service"
	"demo/external/datasource/mock/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var vehicleRepositoryMemo = repositories.NewVehicleRepositoryMemo()
var vehicleTurnService = vehicleupdateservice.NewVehicleUpdateService(vehicleRepositoryMemo)

func TestTurn(t *testing.T) {
	payload := vehicleupdateservice.VehicleUpdateInput{
		Id:                   "id",
		Status:               "on",
		StatusAutomaticPilot: "on",
	}
	err := vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Message,
		"Vehicle Not found",
		"should return an error when vehicle does not exist",
	)

	payload = vehicleupdateservice.VehicleUpdateInput{
		Id:                   "a154",
		Status:               "on",
		StatusAutomaticPilot: "on",
	}
	err = vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Message,
		"This type of vehicle does not have autopilot",
		"should return an error when the vehicle is not a truck",
	)

	payload = vehicleupdateservice.VehicleUpdateInput{
		Id:                   "a154",
		Status:               "invalid",
		StatusAutomaticPilot: "",
	}
	err = vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Message,
		"The Status value is invalid! Valid values ​​(on/off).",
		"should return an error when the status value is invalid",
	)

	payload = vehicleupdateservice.VehicleUpdateInput{
		Id:                   "a156",
		Status:               "on",
		StatusAutomaticPilot: "invalid",
	}
	err = vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Message,
		"The StatusAutomaticPilot value is invalid! Valid values ​​(on/off).",
		"should return an error when the StatusAutomaticPilot value is invalid",
	)

	payload = vehicleupdateservice.VehicleUpdateInput{
		Id:                   "a154",
		Status:               "on",
		StatusAutomaticPilot: "",
	}
	errTurnCar := vehicleTurnService.Turn(payload)
	require.Nil(t, errTurnCar)
	assert.Equal(t, true, vehicleRepositoryMemo.LastSavedVehicle.GetStatus(), "should turn on the car")

	payload = vehicleupdateservice.VehicleUpdateInput{
		Id:                   "a156",
		Status:               "on",
		StatusAutomaticPilot: "on",
	}
	errTurnTruck := vehicleTurnService.Turn(payload)
	savedVehicle := vehicleRepositoryMemo.LastSavedVehicle.(*entities.Truck)
	require.Nil(t, errTurnTruck)
	assert.Equal(t, true, savedVehicle.GetStatus(), "should turn on the truck")
	assert.Equal(t, true, savedVehicle.GetAutomaticPilotStatus(), "should turn on automatic pilot in the truck")

	payload = vehicleupdateservice.VehicleUpdateInput{
		Id:                   "a156",
		Status:               "off",
		StatusAutomaticPilot: "on",
		Travelled:            0,
	}
	err = vehicleTurnService.Turn(payload)
	assert.Equal(
		t,
		err.Message,
		"The new value of travelled is less than current value of travelled.",
		"should return an error when the travelled value is invalid",
	)

	payload = vehicleupdateservice.VehicleUpdateInput{
		Id:                   "a156",
		Status:               "off",
		StatusAutomaticPilot: "on",
		Travelled:            200,
	}
	errTurnTravelled := vehicleTurnService.Turn(payload)
	require.Nil(t, errTurnTravelled)
	assert.Equal(t, false, savedVehicle.GetStatus(), "should turn off the truck")
	assert.Equal(t, float32(200), savedVehicle.Travelled, "should update value of travelled")
}
