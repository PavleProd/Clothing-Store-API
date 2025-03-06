package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ConvertFromString(value string, targetType reflect.Type) (reflect.Value, error) {

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

	if !reflectedValue.Type().ConvertibleTo(targetType) {
		return reflectedValue, fmt.Errorf("cannot convert %v to %v", value, targetType)
	}

	return reflectedValue.Convert(targetType), err
}

func ConvertToString(value any) (string, error) {
	var result string

	var reflectedValue = reflect.ValueOf(value)
	switch reflectedValue.Kind() {
	case reflect.String:
		result = reflectedValue.String()
	case reflect.Bool:
		result = fmt.Sprint(reflectedValue.Bool())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		result = fmt.Sprint(reflectedValue.Int())
	case reflect.Float32, reflect.Float64:
		result = fmt.Sprint(reflectedValue.Float())
	default:
		return "", fmt.Errorf("cannot convert %v to string", value)
	}

	return result, nil
}
