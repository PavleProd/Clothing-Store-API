package main

import (
	"log"
	"net/http"
	"online_store_api/src/data_proccessing"
	"online_store_api/src/db"
	"os"
)

var StoreDb *db.DatabaseManager

func initRoutes() {
	var productsHandler = data_proccessing.NewProductsHandler(StoreDb)
	http.Handle("/api/v1/products", productsHandler)
}

func initDB() {
	StoreDb = &db.DatabaseManager{}

	err := StoreDb.Connect(os.Getenv("STORE_DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	defer StoreDb.Close()

	initRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
