package service

import (
	"prodocutApi/gateway"
	"prodocutApi/model"
)

type ProductServiceInterface interface {
	GetProductsService() map[int]model.Product
	GetProductService()
	SaveProductService()
	UpdateProductService()
	DeleteProductService()
}

type Productservice struct {
	repository gateway.ProductGateway
}

func (service *Productservice) GetProductsService() {

}

func (service *Productservice) GetProductService() {

}
func (service *Productservice) SaveProductService() {

}
func (service *Productservice) UpdateProductService() {

}
func (service *Productservice) DeleteProductService() {

}
