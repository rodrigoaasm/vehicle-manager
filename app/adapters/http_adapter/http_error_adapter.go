package httpadapter

import (
	"demo/domain/domainerror"
	"encoding/json"
	"net/http"
)

func BackError(resWriter http.ResponseWriter, domainError *domainerror.DomainError) {
	code := dict[string(domainError.Type)]
	if code == 0 {
		resWriter.WriteHeader(500)
		domainError.Type = domainerror.UNKNOWN
	} else {
		resWriter.WriteHeader(code)
	}
	json.NewEncoder(resWriter).Encode(domainError)
}

var dict = map[string]int{
	string(domainerror.DATABASE):     500,
	string(domainerror.DEPENDENCY):   500,
	string(domainerror.UNKNOWN):      500,
	string(domainerror.INVALID_DATA): 400,
	string(domainerror.NOT_FOUND):    400,
	string(domainerror.CONFLICT):     400,
}
