package gateway

type ProductGateway interface {
	GetProducts() map[int]ProductGateway
	GetProduct()
	SaveProduct()
	UpdateProduct()
	DeleteProduct()
}
