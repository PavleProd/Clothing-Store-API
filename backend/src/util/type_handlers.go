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

func IsDefaultValue(value any) bool {
	return value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}
