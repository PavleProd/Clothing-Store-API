package data_proccessing

import (
	"encoding/json"
	"log"
	"net/http"
	"online_store_api/src/db"
	"online_store_api/src/model"
	"online_store_api/src/util"
)

type ProductsHandler struct {
	database *db.DatabaseManager
}

func NewProductsHandler(database *db.DatabaseManager) *ProductsHandler {
	return &ProductsHandler{
		database: database,
	}
}

func (handler *ProductsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if handler.database == nil {
		log.Println("database not available")
		return
	}

	switch request.Method {
	case "GET":
		handler.handleGet(writer, request)
	default:
		http.Error(writer, "method not supported", http.StatusMethodNotAllowed)
		log.Printf("HTTP method \"%v\" not supported\n", request.Method)
	}
}

func (handler *ProductsHandler) handleGet(writer http.ResponseWriter, request *http.Request) {
	product, err := MapToModel[model.Product](request.URL.Query())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query := db.BuildReadQuery(product, util.PRODUCTS_TABLE_NAME)
	log.Printf("PROCCESSED QUERY: %v", query)

	dataSet, err := handler.database.Read(query)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(writer)
	enc.SetIndent("", "\t")
	enc.Encode(dataSet)
}
