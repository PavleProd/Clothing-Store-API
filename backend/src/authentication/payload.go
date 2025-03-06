package authentication

import (
	"online_store_api/src/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	jwt.RegisteredClaims

	Username   string `json:"Username"`
	Role       int8   `json:"Role"`
	Expiration int64  `json:"Expiration"`
}

func NewPayload(user model.User) *Payload {
	return &Payload{
		Username:   user.Username,
		Role:       user.Role,
		Expiration: time.Now().Add(time.Hour * 24).Unix(),
	}
}
