package app

import (
	"demo/app/controllers"
	submitvehicleservice "demo/domain/services/submit_vehicle_service"
	vehiclegetterservice "demo/domain/services/vehicle_getter_service"
	"demo/external/datasource/myleveldb"
	"demo/external/datasource/myleveldb/repositories"
	"log"

	"github.com/gorilla/mux"
)

func CreateApp(apiRouter *mux.Router) {

	db, err := myleveldb.CreateLevelDB("data/db")

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	vehicleRepo := repositories.VehicleRepository{
		DB: db,
	}

	// services
	submitVehicleService := submitvehicleservice.SubmitVehicleService{
		VehicleRepository: vehicleRepo,
	}
	vehicleGetterService := vehiclegetterservice.VehicleGetterService{
		VehicleRepository: vehicleRepo,
	}

	// controllers
	submitVehicleController := controllers.SubmitVehicleController{
		SubmitVehicleService: submitVehicleService,
	}
	vehicleGetterController := controllers.VehicleGetterController{
		VehicleGetterService: vehicleGetterService,
	}

	apiRouter.HandleFunc("/vehicle", vehicleGetterController.Handle).Methods("GET")
	apiRouter.HandleFunc("/vehicle", submitVehicleController.Handle).Methods("POST")
}
