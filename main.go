package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"online_store_api/model"
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

func MapToProduct(params url.Values) (model.Product, error) {
	var result = model.Product{}

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
		var convertedValue, err = Convert(value, modelValueField.Type())
		if err != nil {
			return result, err
		}

		// convert and set
		modelValueField.Set(convertedValue)
		log.Printf("Converted Successfully %v", modelTagName)
		numFound++
	}

	// some of the provided parameters have been invalid
	if numFound != len(params) {
		return result, errors.New("invalid parameter(s)")
	}

	return result, nil
}

func GetRequestProduct(req *http.Request) (model.Product, error) {
	var product, err = MapToProduct(req.URL.Query())
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func getProductsHandler(w http.ResponseWriter, req *http.Request) {
	var product, err = GetRequestProduct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Println(product)
}

func main() {
	http.HandleFunc("/products", getProductsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
