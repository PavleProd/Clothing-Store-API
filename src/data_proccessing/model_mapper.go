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

	var modelReflectedValue = reflect.ValueOf(&result).Elem()
	var modelReflectedType = modelReflectedValue.Type()
	for i := range modelReflectedValue.NumField() {

		// check if parameter for field was provided
		var modelTypeField = modelReflectedType.Field(i)
		var modelTagName = modelTypeField.Tag.Get("json")
		var value = params.Get(modelTagName)
		if value == "" {
			continue
		}

		// check if we can set model field
		var modelValueField = modelReflectedValue.Field(i)
		if !modelValueField.CanSet() {
			log.Printf("Tag value can't be set %v", modelTagName)
			continue
		}

		// check if provided value is convertible to model value
		var convertedValue, err = util.Convert(value, modelValueField.Type())
		if err != nil {
			return result, err
		}

		// convert and set
		modelValueField.Set(convertedValue)
		log.Printf("converted successfully: %v", modelTagName)
		numFound++
	}

	// some of the provided parameters have been invalid
	if numFound != len(params) {
		return result, errors.New("invalid parameter(s)")
	}

	return result, nil
}
