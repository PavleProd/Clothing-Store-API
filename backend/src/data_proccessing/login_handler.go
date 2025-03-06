package data_proccessing

import (
	"log/slog"
	"net/http"
	"online_store_api/src/db"
)

type LoginHandler struct {
	database *db.DatabaseManager
}

func (handler *LoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if handler.database == nil {
		slog.Error("database not available")
		return
	}

	var errorCode int = http.StatusOK
	var err error = nil

	switch request.Method {
	case "GET":
		errorCode, err = handler.handleGet(writer, request)
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

func (handler *LoginHandler) handleGet(writer http.ResponseWriter, request *http.Request) (int, error) {
	return 0, nil
}
