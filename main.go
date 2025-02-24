package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"online_store_api/model"
	"reflect"
)

func MapToProduct(params url.Values) (model.Product, error) {
	var result = model.Product{}

	var numFound int = 0

	var modelReflectedValue = reflect.ValueOf(&result).Elem()
	var modelReflectedType = reflect.TypeOf(result)
	for i := range modelReflectedType.NumField() {

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
		var reflectedValue = reflect.ValueOf(value)
		if !reflectedValue.CanConvert(modelValueField.Type()) {
			log.Printf("Can't convert value to tag %v", modelTagName)
			continue
		}

		// convert and set
		modelValueField.Set(reflectedValue.Convert(modelValueField.Type()))
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
