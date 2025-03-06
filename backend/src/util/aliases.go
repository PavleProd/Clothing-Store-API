package util

import "os"

type DataInterface = map[string]any
type DataRecord = map[string]string
type DataSet = []DataRecord

const JSON_TAG = "json"

var STORE_DB_URL = os.Getenv("STORE_DB_URL")
var LOGIN_DB_URL = os.Getenv("LOGIN_DB_URL")
