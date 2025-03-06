package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"online_store_api/src/model"
	"online_store_api/src/util"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("very-secret-key:)")

func CreateToken(userData util.DataRecord) (string, error) {
	user, err := util.MapToModel[model.User](userData)
	if err != nil {
		return "", err
	}

	var payload = NewPayload(user)

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

func CheckAuthorization(request *http.Request, minRole int8) error {
	tokenString, err := getTokenString(request)
	if err != nil {
		return err
	}

	payload, err := getPayload(tokenString)
	if err != nil {
		return err
	}

	err = validatePayload(payload, minRole)
	return err
}

func getTokenString(request *http.Request) (string, error) {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return authHeader, fmt.Errorf("authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return authHeader, fmt.Errorf("authorization header invalid")
	}

	return parts[1], nil
}

func getPayload(tokenString string) (Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return Payload{}, err
	}

	payload, ok := token.Claims.(*Payload)
	if !ok || !token.Valid {
		return Payload{}, errors.New("invalid token")
	}

	return *payload, nil
}

func validatePayload(payload Payload, minRole int8) error {
	if payload.Expiration < time.Now().Unix() {
		return errors.New("token expired")
	}

	if minRole > payload.Role {
		return errors.New("access not authorized")
	}

	return nil
}
