package data_proccessing

import (
	"fmt"
	"log/slog"
	"net/http"
	"online_store_api/src/authentication"
	"online_store_api/src/db"
)

const USERS_DB_NAME = "users"

type LoginHandler struct {
	database *db.DatabaseManager
}

func NewLoginHandler(database *db.DatabaseManager) *LoginHandler {
	return &LoginHandler{
		database: database,
	}
}

func (handler *LoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if handler.database == nil {
		slog.Error("database not available")
		return
	}

	var errorCode int = http.StatusOK

	switch request.Method {
	case "POST":
		errorCode = handler.handlePost(writer, request)
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

func (handler *LoginHandler) handlePost(writer http.ResponseWriter, request *http.Request) int {
	data, err := ParseBody(request)
	if err != nil {
		slog.Info("json parse failed", "error", err.Error())
		return http.StatusBadRequest
	}

	var preparedQuery db.PreparedQuery = db.BuildSelectQuery(data, USERS_DB_NAME)
	users, err := handler.database.Read(preparedQuery)
	if err != nil {
		slog.Info("database transaction failed", "error", err.Error())
		return http.StatusInternalServerError
	}

	if len(users) == 0 {
		slog.Info("user entered wrong credentials")
		return http.StatusBadRequest
	} else if len(users) > 1 {

		slog.Error("database invalid, more than 1 unique users", "users", users)
		return http.StatusInternalServerError
	}

	token, err := authentication.CreateToken(users[0])
	if err != nil {
		slog.Error("jwt authentication failed", "error", err.Error())
		return http.StatusInternalServerError
	}

	fmt.Fprint(writer, token)
	return http.StatusOK
}
