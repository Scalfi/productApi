package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = make(map[int]Product)
var ID int

func GETHandlerProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GETHandlerOneProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(GetField(r, 0))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products[id])
}

func POSTHandlerProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := Product{}
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The payload is invalid")
		return
	}

	ID += 1
	newProduct.ID = ID
	products[ID] = newProduct

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(products[ID])
}

func PUTHandlerProduct(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	w.Header().Set("Content-Type", "application/json")

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("URL param 'key' is missing")
		return
	}

	updateProduct := Product{}

	err := json.NewDecoder(r.Body).Decode(&updateProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The payload is invalid")
		return
	}

	key, _ := strconv.Atoi(keys[0])
	product := products[key]

	product.Name = updateProduct.Name
	product.Price = updateProduct.Price

	products[key] = product
	json.NewEncoder(w).Encode(products[key])
}

func DELETEHandlerProduct(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	w.Header().Set("Content-Type", "application/json")

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("URL param 'key' is missing")
		return
	}
	key, _ := strconv.Atoi(keys[0])

	if _, ok := products[key]; ok {
		delete(products, key)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}
