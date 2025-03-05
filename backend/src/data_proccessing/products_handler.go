package data_proccessing

import (
	"encoding/json"
	"log/slog"
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
		slog.Error("database not available")
		return
	}

	switch request.Method {
	case "GET":
		handler.handleGet(writer, request)
		slog.Error("HTTP method not supported", "method", request.Method)
	case "POST":
		handler.handlePost(writer, request)
	default:
		http.Error(writer, "method not supported", http.StatusMethodNotAllowed)
		slog.Error("HTTP method not supported", "method", request.Method)
		return
	}

	slog.Info("processed %v query: %v", request.Method, request.URL)
}

func (handler *ProductsHandler) handleGet(writer http.ResponseWriter, request *http.Request) {
	product, err := MapToModel[model.Product](request.URL.Query())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query := db.BuildReadQuery(product, util.PRODUCTS_TABLE_NAME)

	dataSet, err := handler.database.Read(query)
	if err != nil {
		slog.Error(err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(writer)
	enc.SetIndent("", "\t")
	enc.Encode(dataSet)
}

func (handler *ProductsHandler) handlePost(writer http.ResponseWriter, request *http.Request) {

}
