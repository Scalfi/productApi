package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type serve struct {
}

var products = make(map[int]Product)
var ID int

func main() {

	err := http.ListenAndServe(":3001", &serve{})

	if err != nil {
		log.Fatal(err)
	}
}

func (s *serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.GETHandler(w, r)
	case http.MethodPost:
		s.POSTHandler(w, r)
	case http.MethodPut:
		s.PUTHandler(w, r)
	case http.MethodDelete:
		s.DELETEHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}

func (s *serve) GETHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func (s *serve) POSTHandler(w http.ResponseWriter, r *http.Request) {
	newProduct := Product{}
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The payload is invalid")
		return
	}

	ID += 1
	products[ID] = newProduct

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(products[ID])
}

func (s *serve) PUTHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

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

func (s *serve) DELETEHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

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
