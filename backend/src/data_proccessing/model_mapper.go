package data_proccessing

import (
	"errors"
	"log"
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
			log.Printf("Tag value can't be set %v", reflectedFieldName)
			continue
		}

		// check if provided value is convertible to model value
		var convertedValue, err = util.Convert(value, reflectedFieldValue.Type())
		if err != nil {
			return result, err
		}

		// convert and set
		reflectedFieldValue.Set(convertedValue)
		log.Printf("converted successfully: %v", reflectedFieldName)
		numFound++
	}

	// some of the provided parameters have been invalid
	if numFound != len(params) {
		return result, errors.New("invalid parameter(s)")
	}

	return result, nil
}
