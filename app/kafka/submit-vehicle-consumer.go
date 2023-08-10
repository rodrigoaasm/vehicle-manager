package kafka

import (
	"demo/domain/domainerror"
	submitvehicleservice "demo/domain/services/submit_vehicle_service"
	"encoding/json"
	"io"
)

type PayloadKafka struct {
	Data io.Reader
}

func (p PayloadKafka) ack(ErrorKafka interface{}) {
	return
}

func registerTopic(topic string, handler func(PayloadKafka)) {

}

// ------------------------------------------------------

type SubmitVehiclePayload struct {
	Category     string  `json:"category"`
	Name         string  `json:"name"`
	Cor          string  `json:"cor"`
	Serie        string  `json:"serie"`
	LicensePlate string  `json:"licensePlate`
	Travelled    float32 `json:"travelled"`
}

type SubmitVehicleConsumer struct {
	SubmitVehicleService submitvehicleservice.ISubmitVehicleService
}

func (consumer SubmitVehicleConsumer) Consume() {
	registerTopic("device.data", consumer.Handle)
}

func (consumer SubmitVehicleConsumer) Handle(payload PayloadKafka) {
	data := SubmitVehiclePayload{}
	errDecode := json.NewDecoder(payload.Data).Decode(&data)
	if errDecode != nil {
		payload.ack(domainerror.New(domainerror.INVALID_DATA, errDecode.Error()))
		return
	}

	errSubmit := consumer.SubmitVehicleService.Submit(
		data.Category, data.Name, data.Cor, data.Serie, data.LicensePlate, data.Travelled,
	)

	if errSubmit != nil {
		payload.ack(errSubmit)
	}

	payload.ack(nil)
}
