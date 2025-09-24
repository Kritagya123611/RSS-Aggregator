package products

import (
	"encoding/json"
	//"fmt"
	"net/http"
)

type NewProduct struct {
	ID       string  `json:"id"`
	NAME     string  `json:"name"`
	PRICE    float64 `json:"price"`
	CATEGORY string  `json:"category"`
}

var ListProducts = []NewProduct{
	{ID: "1", NAME: "Laptop", PRICE: 75000, CATEGORY: "Electronics"},
	{ID: "2", NAME: "Phone", PRICE: 30000, CATEGORY: "Electronics"},
	{ID: "3", NAME: "Headphones", PRICE: 2000, CATEGORY: "Accessories"},
}

func AddProduct(w http.ResponseWriter, r *http.Request){
	var body NewProduct
	err:=json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
    http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
    return
}
	ListProducts=append(ListProducts, body)
	json.NewEncoder(w).Encode(body)
}