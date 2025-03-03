package main

import (
	"encoding/json"
	"log"
	"net/http"
	"online_store_api/src/data_proccessing"
	"online_store_api/src/db"
	"online_store_api/src/model"
	"online_store_api/src/util"
	"os"
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
		return
	}

	query := db.BuildReadQuery(product, util.PRODUCTS_TABLE_NAME)
	log.Printf("PROCCESSED QUERY: %v", query)

	dataSet, err := StoreDb.Read(query)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(dataSet)
}

func initRoutes() {
	http.HandleFunc("/api/v1/products", getProductsHandler)
}

func main() {
	initRoutes()

	StoreDb = &db.ConnectionManager{}

	err := StoreDb.Connect(os.Getenv("STORE_DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer StoreDb.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
