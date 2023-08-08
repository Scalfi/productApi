package repository

import "prodocutApi/model"

type ProductDB struct {
	products map[int]model.Product
}

func NewProductRepository() *ProductDB {
	return &ProductDB{}
}

func (repo *ProductDB) GetProducts() map[int]model.Product {
	return repo.products
}
func (repo *ProductDB) GetProduct() {

}
func (repo *ProductDB) SaveProduct() {

}
func (repo *ProductDB) UpdateProduct() {

}
func (repo *ProductDB) DeleteProduct() {

}
