package util

import (
	"errors"
	"reflect"
	"strconv"
)

func Convert(value string, targetType reflect.Type) (reflect.Value, error) {

	var reflectedValue reflect.Value
	var err error

	switch targetType.Kind() {
	case reflect.String:
		reflectedValue = reflect.ValueOf(value)
	case reflect.Float32, reflect.Float64:
		var value, e = strconv.ParseFloat(value, targetType.Bits())
		reflectedValue, err = reflect.ValueOf(value), e
	case reflect.Bool:
		var value, e = strconv.ParseBool(value)
		reflectedValue, err = reflect.ValueOf(value), e
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		var value, e = strconv.ParseInt(value, 10, targetType.Bits())
		reflectedValue, err = reflect.ValueOf(value), e
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		var value, e = strconv.ParseUint(value, 10, targetType.Bits())
		reflectedValue, err = reflect.ValueOf(value), e
	default:
		err = errors.New("unexpected reflected type")
	}

	return reflectedValue.Convert(targetType), err
}

func IsDefaultOrZeroValueExcludingBool(value any) bool {
	if reflect.TypeOf(value) == reflect.TypeOf(true) {
		return false
	}

	return value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
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
