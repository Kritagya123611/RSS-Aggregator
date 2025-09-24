package products

import (
	"encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
)

type Product struct{
	ID string `json:"id"`
	NAME string `json:"name"`
	PRICE float64 `json:"price"`
}

var Products = []Product{
	{ID: "1", NAME: "Laptop", PRICE: 75000},
	{ID: "2", NAME: "Phone", PRICE: 30000},
	{ID: "3", NAME: "Headphones", PRICE: 2000},
}

func AllProducts(w http.ResponseWriter,r *http.Request ){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

func GetProductByID(w http.ResponseWriter,r *http.Request){
	id:=chi.URLParam(r,"id")
	for _,p:=range Products{
		if p.ID==id{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}