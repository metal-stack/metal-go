package metalgo

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/metal-stack/metal-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// AttemptGenericError can be used for trying to extract our default HTTPErrorResponse from
// any error response returned by the generated swagger client (which always has a specific type).
//
// This is useful for printing out errors nicely without having to implement big error type switches.
//
// The function returns the original error when the error does not contain a HTTPErrorResponse.
func AttemptGenericError(swaggerErr error) error {
	obj, err := getField(swaggerErr, "Payload")
	if err != nil {
		return swaggerErr
	}
	httperr, ok := obj.(*models.HttperrorsHTTPErrorResponse)
	if ok {
		return httperrors.FromDefaultResponse(httperr.Statuscode, httperr.Message, swaggerErr)
	}
	return swaggerErr
}

// copied over from https://github.com/oleiade/reflections
func getField(obj interface{}, name string) (interface{}, error) {
	if !hasValidType(obj, []reflect.Kind{reflect.Struct, reflect.Ptr}) {
		return nil, errors.New("Cannot use GetField on a non-struct interface")
	}

	objValue := reflectValue(obj)
	field := objValue.FieldByName(name)
	if !field.IsValid() {
		return nil, fmt.Errorf("No such field: %s in obj", name)
	}

	return field.Interface(), nil
}

func hasValidType(obj interface{}, types []reflect.Kind) bool {
	for _, t := range types {
		if reflect.TypeOf(obj).Kind() == t {
			return true
		}
	}

	return false
}

func reflectValue(obj interface{}) reflect.Value {
	var val reflect.Value

	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		val = reflect.ValueOf(obj).Elem()
	} else {
		val = reflect.ValueOf(obj)
	}

	return val
}
