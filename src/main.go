package main

import (
	"log"
	"net/http"
	"online_store_api/src/data_proccessing"
	"online_store_api/src/db"
	"online_store_api/src/model"
)

var StoreDb *db.ConnectionManager

func getProductsHandler(w http.ResponseWriter, req *http.Request) {
	if StoreDb == nil {
		log.Panicln("db not available")
		return
	}

	dataSet, err := StoreDb.Query("SELECT * FROM products") // TODO: QUERY BUILDER
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(dataSet) // TODO: OBRADI REZULTAT

	product, err := data_proccessing.MapToModel[model.Product](req.URL.Query(), "json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Println(product)
}

func initRoutes() {
	http.HandleFunc("/api/v1/products", getProductsHandler)
}

func main() {
	initRoutes()

	StoreDb = &db.ConnectionManager{DatabaseName: "store"}

	err := StoreDb.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer StoreDb.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
