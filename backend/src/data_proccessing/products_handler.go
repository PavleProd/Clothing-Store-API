package data_proccessing

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"online_store_api/src/db"
	"online_store_api/src/model"
)

type ProductsHandler struct {
	database  *db.DatabaseManager
	tableName string
}

func NewProductsHandler(database *db.DatabaseManager, tableName string) *ProductsHandler {
	return &ProductsHandler{
		database:  database,
		tableName: tableName,
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
	product, err := MapToModel[model.Product](request.URL.Query())
	if err != nil {
		return http.StatusBadRequest, err
	}

	query := db.BuildReadQuery(product, handler.tableName)

	dataSet, err := handler.database.Read(query)
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
	var product model.Product

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		slog.Error("json parse failed", "error", err.Error())
		return http.StatusBadRequest, err
	}

	query, err := db.BuildInsertQuery(product, handler.tableName)
	if err != nil {
		slog.Error("building insert querty failed", "error", err.Error())
		return http.StatusInternalServerError, err
	}

	err = handler.database.Write(query)
	if err != nil {
		slog.Error("database transaction failder", "error", err.Error())
		return http.StatusInternalServerError, err
	}

	slog.Info("transaction with query successful", "query", query)
	return http.StatusOK, nil
}
