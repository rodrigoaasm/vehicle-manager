package mock

import (
	"net/http"
)

type ResponseWriterFake struct {
	Message string
	Status  int
}

func (w *ResponseWriterFake) Header() http.Header {
	return http.Header{}
}

func (w *ResponseWriterFake) Write(body []byte) (int, error) {
	w.Message = string(body)

	return 0, nil
}

func (w *ResponseWriterFake) WriteHeader(statusCode int) {
	w.Status = statusCode
}
