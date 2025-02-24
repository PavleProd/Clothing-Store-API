package main

import (
	"io"
	"log"
	"net/http"
)

func getProductsHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "majaba")
}

func main() {
	http.HandleFunc("/products", getProductsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
