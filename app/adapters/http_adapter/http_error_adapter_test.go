package httpadapter_test

import (
	httpadapter "demo/app/adapters/http_adapter"
	"demo/app/mock"
	"demo/domain/domainerror"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackError(t *testing.T) {
	// Known
	responseWriterFake := mock.ResponseWriterFake{}
	domainError := domainerror.New(domainerror.INVALID_DATA, "Error fake")

	httpadapter.BackError(&responseWriterFake, domainError)
	assert.Equal(t, 400, responseWriterFake.Status)
	assert.Equal(t, "{\"type\":\"INVALID_DATA\",\"message\":\"Error fake\"}\n", responseWriterFake.Message)

	// Unknown
	responseWriterFake = mock.ResponseWriterFake{}
	domainError = domainerror.New(domainerror.TEST, "Error fake")

	httpadapter.BackError(&responseWriterFake, domainError)
	assert.Equal(t, 500, responseWriterFake.Status)
	assert.Equal(t, "{\"type\":\"UNKNOWN\",\"message\":\"Error fake\"}\n", responseWriterFake.Message)
}
