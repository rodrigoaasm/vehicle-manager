package app

import (
	"demo/app/kafka"
	"demo/app/web/controllers"
	submitvehicleservice "demo/domain/services/submit_vehicle_service"
	vehiclegetterservice "demo/domain/services/vehicle_getter_service"
	vehicleupdateservice "demo/domain/services/vehicle_update_service"
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
	submitVehicleServiceWithOut := submitvehicleservice.NewSubmitVehicleWithOutValidationService(vehicleRepo)
	vehicleGetterService := vehiclegetterservice.NewVehicleTurnService(vehicleRepo)
	vehicleUpdateService := vehicleupdateservice.NewVehicleUpdateService(vehicleRepo)

	// controllers
	submitVehicleController := controllers.SubmitVehicleController{
		SubmitVehicleService: submitVehicleService,
	}
	vehicleGetterController := controllers.VehicleGetterController{
		VehicleGetterService: vehicleGetterService,
	}
	vehicleTurnersController := controllers.VehicleUpdateController{
		VehicleTurnService: vehicleUpdateService,
	}

	// consumers
	submitVehicleConsumer := kafka.SubmitVehicleConsumer{
		SubmitVehicleService: submitVehicleServiceWithOut,
	}
	submitVehicleConsumer.Consume()

	apiRouter.HandleFunc("/vehicle", vehicleGetterController.Handle).Methods("GET")
	apiRouter.HandleFunc("/vehicle", submitVehicleController.Handle).Methods("POST")
	apiRouter.HandleFunc("/vehicle", vehicleTurnersController.Handle).Methods("PATCH")
}
