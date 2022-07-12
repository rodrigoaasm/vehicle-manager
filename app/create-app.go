package app

import (
	"demo/app/controllers"
	"demo/domain/services"
	"demo/external/datasource/mock/repositories"

	"github.com/gorilla/mux"
)

func CreateApp(apiRouter *mux.Router) {

	vehicleRepo := repositories.VehicleRepositoryMemo{}

	vehicleGetterService := services.VehicleGetterService{
		VehicleRepository: vehicleRepo,
	}
	vehicleGetterController := controllers.VehicleGetterController{
		VehicleGetterService: vehicleGetterService,
	}

	apiRouter.HandleFunc("/vehicle", vehicleGetterController.Handle).Methods("GET")
}
