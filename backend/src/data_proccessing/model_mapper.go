package data_proccessing

import (
	"errors"
	"log/slog"
	"net/url"
	"online_store_api/src/util"
	"reflect"
)

func MapToModel[T any](params url.Values) (T, error) {
	var result T

	var numFound int = 0

	var reflectedModelValue = reflect.ValueOf(&result).Elem()
	var reflectedModelType = reflectedModelValue.Type()
	for i := range reflectedModelValue.NumField() {

		// check if parameter for field was provided
		var reflectedFieldType = reflectedModelType.Field(i)
		var reflectedFieldName = reflectedFieldType.Tag.Get(util.JSON_TAG)
		var value = params.Get(reflectedFieldName)
		if value == "" {
			continue
		}

		// check if we can set model field
		var reflectedFieldValue = reflectedModelValue.Field(i)
		if !reflectedFieldValue.CanSet() {
			slog.Error("Field is immutable", "field", reflectedFieldName)
			continue
		}

		// check if provided value is convertible to model value
		var convertedValue, err = util.ConvertFromString(value, reflectedFieldValue.Type())
		if err != nil {
			return result, err
		}

		// convert and set
		reflectedFieldValue.Set(convertedValue)
		slog.Info("converted to field successfully", "field", reflectedFieldName, "value", convertedValue)
		numFound++
	}

	// some of the provided parameters have been invalid
	if numFound != len(params) {
		return result, errors.New("invalid parameter(s)")
	}

	return result, nil
}
