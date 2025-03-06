package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int8   `json:"role"`
}
