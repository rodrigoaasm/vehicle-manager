package submitvehicleservice_test

import (
	submitvehicleservice "demo/domain/services/submit_vehicle_service"
	"demo/external/datasource/mock/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var vehicleRepositoryMemo = repositories.NewVehicleRepositoryMemo()
var vehicleGetterService = submitvehicleservice.NewSubmitVehicleService(vehicleRepositoryMemo)

func TestSubmit(t *testing.T) {
	errNameLess := vehicleGetterService.Submit("car", "VW", "black", "14885511T125T", "ABC1234")
	assert.Equal(
		t,
		errNameLess.Message,
		"The name must be greater than 25 or less than 3.",
		"should return an error when name is less 3",
	)

	errNameGreat := vehicleGetterService.Submit("car", "123456789-123456789-123456", "black", "14885511T125T", "ABC1234")
	assert.Equal(
		t,
		errNameGreat.Message,
		"The name must be greater than 25 or less than 3.",
		"should return an error when name is greater 25",
	)

	errPlate := vehicleGetterService.Submit("car", "VW GOL", "black", "14885511T125T", "ABCD234")
	assert.Equal(
		t,
		errPlate.Message,
		"License Plate invalid",
		"should return an error when License Plate is invalid",
	)

	errCategory := vehicleGetterService.Submit("invalid", "VW GOL", "black", "14885511T124T", "ABC1224")
	assert.Equal(
		t,
		errCategory.Message,
		"Category unknown",
		"should return an error when category is unknown",
	)

	err := vehicleGetterService.Submit("car", "VW GOL", "black", "14885511T225T", "ABC1224")
	require.Nil(t, err, "should not return an error when submit a car")
}
