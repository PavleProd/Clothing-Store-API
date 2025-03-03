package util

import "os"

type DataRecord = map[string]string
type DataSet = []DataRecord

const JSON_TAG = "json"

const PRODUCTS_TABLE_NAME = "products"

var STORE_DB_URL = os.Getenv("STORE_DB_URL")
