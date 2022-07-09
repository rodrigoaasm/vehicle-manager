package services

import (
	"demo/objects"
	"encoding/json"
	"log"
	"net/http"
)

var vehicles = []objects.IVehicle{
	objects.NewCar("a154", "VW GOL", "black", "14885511T125T"),
	objects.NewCar("a155", "VW GOL", "pink", "12885511T125T"),
	objects.NewTrunk("a156", "VW CONSTELLATION", "white", "2885511T125E"),
}

func GetVehicle(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Header().Set("Content-Type", "application/json")

	id := req.URL.Query().Get("id")

	if len(id) > 0 {
		for _, vehicle := range vehicles {
			var objectVehicle, ok = vehicle.(*objects.Car)

			if !ok {
				log.Printf("Vehicle conversion failed")
				resWriter.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(resWriter)
				return
			}

			if objectVehicle.Id == id {
				json.NewEncoder(resWriter).Encode(objectVehicle)
				return
			}
		}

		resWriter.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resWriter)
		return
	}

	json.NewEncoder(resWriter).Encode(vehicles)
}
