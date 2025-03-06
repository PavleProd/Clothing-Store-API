package main

import (
	"log"
	"net/http"
	"online_store_api/src/data_proccessing"
	"online_store_api/src/db"
	"os"
)

var database *db.DatabaseManager

func initRoutes() {
	var productsHandler = data_proccessing.NewProductsHandler(database)
	http.Handle("/api/v1/products", productsHandler)

	var loginHandler = data_proccessing.NewLoginHandler(database)
	http.Handle("/api/v1/login", loginHandler)
}

func initDB() {
	database = &db.DatabaseManager{}

	err := database.Connect(os.Getenv("STORE_DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	defer database.Close()

	initRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
