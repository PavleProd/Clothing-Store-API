package db

type Credentials struct {
	Username string
	Password string
}

var Admin = Credentials{"postgres", "admin"}
