package data_proccessing

import (
	"encoding/json"
	"net/http"
	"online_store_api/src/util"
)

func ParseURL(request *http.Request) util.DataRecord {
	var urlParams = request.URL.Query()
	var result = util.DataRecord{}
	for key, valueSlice := range urlParams {
		if len(valueSlice) == 0 {
			continue
		}

		result[key] = valueSlice[0]
	}

	return result
}

func ParseBody(request *http.Request) (util.DataRecord, error) {
	var result = util.DataRecord{}

	var dataInterface = util.DataInterface{}
	err := json.NewDecoder(request.Body).Decode(&dataInterface)
	if err != nil {
		return result, err
	}

	for key, value := range dataInterface {
		convertedValue, err := util.ConvertToString(value)
		if err != nil {
			return result, err
		}

		result[key] = convertedValue
	}

	return result, nil
}
