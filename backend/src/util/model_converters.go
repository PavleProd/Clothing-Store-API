package util

import (
	"log/slog"
	"reflect"
)

func MapToModel[T any](data DataRecord) (T, error) {
	var result T

	var numFound int = 0

	var reflectedModelValue = reflect.ValueOf(&result).Elem()
	var reflectedModelType = reflectedModelValue.Type()
	for i := range reflectedModelValue.NumField() {

		// check if parameter for field was provided
		var reflectedFieldType = reflectedModelType.Field(i)
		var reflectedFieldName = reflectedFieldType.Tag.Get(JSON_TAG)
		var value, ok = data[reflectedFieldName]
		if !ok {
			continue
		}

		// check if we can set model field
		var reflectedFieldValue = reflectedModelValue.Field(i)
		if !reflectedFieldValue.CanSet() {
			slog.Error("Field is immutable", "field", reflectedFieldName)
			continue
		}

		// check if provided value is convertible to model value
		var convertedValue, err = ConvertFromString(value, reflectedFieldValue.Type())
		if err != nil {
			return result, err
		}

		// convert and set
		reflectedFieldValue.Set(convertedValue)
		slog.Info("converted to field successfully", "field", reflectedFieldName, "value", convertedValue)
		numFound++
	}

	return result, nil
}

type SlicedField struct {
	Name  string
	Value any
	Tag   string
}

func GetModelSlicedFields[T any](model T) []SlicedField {
	var slicedFields []SlicedField

	var reflectedModelValue = reflect.ValueOf(&model).Elem()
	var reflectedModelType = reflectedModelValue.Type()
	for i := range reflectedModelValue.NumField() {
		var reflectedFieldType = reflectedModelType.Field(i)
		var reflectedFieldValue = reflectedModelValue.Field(i)

		var slicedField = SlicedField{
			Name:  reflectedFieldType.Name,
			Value: reflectedFieldValue.Interface(),
			Tag:   reflectedFieldType.Tag.Get(JSON_TAG),
		}

		slicedFields = append(slicedFields, slicedField)
	}

	return slicedFields
}
