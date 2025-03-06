package data_proccessing

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"online_store_api/src/db"
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
	data := ConvertURL(request)
	query := db.BuildSelectQuery(data, handler.tableName)

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
	data, err := ConvertBody(request)
	if err != nil {
		slog.Error("json parse failed", "error", err.Error())
		return http.StatusBadRequest, err
	}

	query := db.BuildInsertQuery(data, handler.tableName)
	log.Println(query)
	err = handler.database.Write(query)
	if err != nil {
		slog.Error("database transaction failed", "error", err.Error())
		return http.StatusInternalServerError, err
	}

	slog.Info("transaction with query successful", "query", query)
	return http.StatusOK, nil
}
