package main

import (
	"log"
	"net/http"
	"online_store_api/data_proccessing"
	"online_store_api/model"
)

func getProductsHandler(w http.ResponseWriter, req *http.Request) {
	var product, err = data_proccessing.MapToModel[model.Product](req.URL.Query())
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
	log.Fatal(http.ListenAndServe(":8080", nil))

}
