package data_proccessing

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"online_store_api/src/authentication"
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

	var errorCode int = http.StatusOK

	switch request.Method {
	case "GET":
		errorCode = handler.handleGet(writer, request)
	case "POST":
		errorCode = handler.handlePost(request)
	default:
		http.Error(writer, "method not supported", http.StatusMethodNotAllowed)
		slog.Error("HTTP method not supported", "method", request.Method)
		return
	}

	if errorCode != http.StatusOK {
		http.Error(writer, http.StatusText(errorCode), errorCode)
		return
	}

	slog.Info("sucessfuly proccessed request", "Method", request.Method, "URL", request.URL)
}

func (handler *ProductsHandler) handleGet(writer http.ResponseWriter, request *http.Request) int {
	err := authentication.CheckAuthorization(request, authentication.User)
	if err != nil {
		slog.Info("authorization error", "error", err)
		return http.StatusUnauthorized
	}

	data := ParseURL(request)
	preparedQuery := db.BuildSelectQuery(data, PRODUCTS_TABLE_NAME)

	dataSet, err := handler.database.Read(preparedQuery)
	if err != nil {
		slog.Error("database transaction failed", "error", err.Error())
		return http.StatusInternalServerError
	}

	enc := json.NewEncoder(writer)
	enc.SetIndent("", "\t")
	enc.Encode(dataSet)

	return http.StatusOK
}

func (handler *ProductsHandler) handlePost(request *http.Request) int {
	err := authentication.CheckAuthorization(request, authentication.Admin)
	if err != nil {
		slog.Info("authorization error", "error", err)
		return http.StatusUnauthorized
	}

	data, err := ParseBody(request)
	if err != nil {
		slog.Error("json parse failed", "error", err.Error())
		return http.StatusBadRequest
	}

	preparedQuery := db.BuildInsertQuery(data, PRODUCTS_TABLE_NAME)
	err = handler.database.Write(preparedQuery)
	if err != nil {
		slog.Error("database transaction failed", "error", err.Error())
		return http.StatusInternalServerError
	}

	slog.Info("transaction with query successful", "query", preparedQuery)
	return http.StatusOK
}
