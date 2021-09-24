package product

type Service interface {
	CreateNewProduct(product Product) error
}

type Repository interface {
	CreateNewProduct(product Product) error
}
