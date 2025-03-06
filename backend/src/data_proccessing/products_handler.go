package data_proccessing

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"online_store_api/src/db"
)

const PRODUCTS_TABLE_NAME = "products"

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

	var err error = nil
	var errorCode int = http.StatusOK

	switch request.Method {
	case "GET":
		errorCode, err = handler.handleGet(writer, request)
	case "POST":
		errorCode, err = handler.handlePost(writer, request)
	default:
		http.Error(writer, "method not supported", http.StatusMethodNotAllowed)
		slog.Error("HTTP method not supported", "method", request.Method)
		return
	}

	if err != nil {
		http.Error(writer, err.Error(), errorCode)
		return
	}

	slog.Info("sucessfuly proccessed request", "Method", request.Method, "URL", request.URL)
}

func (handler *ProductsHandler) handleGet(writer http.ResponseWriter, request *http.Request) (int, error) {
	data := ConvertURL(request)
	preparedQuery := db.BuildSelectQuery(data, PRODUCTS_TABLE_NAME)

	dataSet, err := handler.database.Read(preparedQuery)
	if err != nil {
		slog.Error(err.Error())
		return http.StatusInternalServerError, err
	}

	enc := json.NewEncoder(writer)
	enc.SetIndent("", "\t")
	enc.Encode(dataSet)

	return http.StatusOK, nil
}

func (handler *ProductsHandler) handlePost(writer http.ResponseWriter, request *http.Request) (int, error) {
	data, err := ConvertBody(request)
	if err != nil {
		slog.Error("json parse failed", "error", err.Error())
		return http.StatusBadRequest, err
	}

	preparedQuery := db.BuildInsertQuery(data, PRODUCTS_TABLE_NAME)
	err = handler.database.Write(preparedQuery)
	if err != nil {
		slog.Error("database transaction failed", "error", err.Error())
		return http.StatusInternalServerError, err
	}

	slog.Info("transaction with query successful", "query", preparedQuery)
	return http.StatusOK, nil
}
