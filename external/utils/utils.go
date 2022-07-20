package utils

import "reflect"

func IsThisType[T any](object any) bool {
	var tempObject T

	objectType := reflect.TypeOf(object).String()
	genericType := reflect.TypeOf(tempObject).String()

	return objectType == genericType || objectType == "*"+genericType
}
