package app

import (
	"demo/app/controllers"
	submitvehicleservice "demo/domain/services/submit_vehicle_service"
	vehiclegetterservice "demo/domain/services/vehicle_getter_service"
	vehicleturnservice "demo/domain/services/vehicle_turn_service"
	"demo/external/datasource/myleveldb"
	"demo/external/datasource/myleveldb/repositories"
	"log"

	"github.com/gorilla/mux"
)

func CreateApp(apiRouter *mux.Router) {

	db, err := myleveldb.NewDatabase("db/")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	vehicleRepo := repositories.NewVehicleRepository(db)

	// services
	submitVehicleService := submitvehicleservice.NewSubmitVehicleService(vehicleRepo)
	vehicleGetterService := vehiclegetterservice.NewVehicleTurnService(vehicleRepo)
	vehicleTurnService := vehicleturnservice.NewVehicleTurnService(vehicleRepo)

	// controllers
	submitVehicleController := controllers.SubmitVehicleController{
		SubmitVehicleService: submitVehicleService,
	}
	vehicleGetterController := controllers.VehicleGetterController{
		VehicleGetterService: vehicleGetterService,
	}
	vehicleTurnersController := controllers.VehicleTurnersController{
		VehicleTurnService: vehicleTurnService,
	}

	apiRouter.HandleFunc("/vehicle", vehicleGetterController.Handle).Methods("GET")
	apiRouter.HandleFunc("/vehicle", submitVehicleController.Handle).Methods("POST")
	apiRouter.HandleFunc("/vehicle", vehicleTurnersController.Handle).Methods("PATCH")
}
