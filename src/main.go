package main

import (
	"encoding/json"
	"log"
	"net/http"
	"online_store_api/src/data_proccessing"
	"online_store_api/src/db"
	"online_store_api/src/model"
	"online_store_api/src/util"
)

var StoreDb *db.ConnectionManager

func getProductsHandler(w http.ResponseWriter, req *http.Request) {
	if StoreDb == nil {
		log.Panicln("db not available")
		return
	}

	product, err := data_proccessing.MapToModel[model.Product](req.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	query := db.BuildReadQuery(product, util.PRODUCTS_TABLE_NAME)
	log.Printf("PROCCESSED QUERY: %v", query)

	dataSet, err := StoreDb.Read(query)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dataSet)
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
