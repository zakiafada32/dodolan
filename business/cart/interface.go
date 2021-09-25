package cart

import "github.com/zakiafada32/retail/business/product"

type Service interface {
	AddCartItem(userId string, productId uint32, quantity uint32) error
}

type Repository interface {
	FindProductById(productId uint32) product.Product
	AddCartItem(userId string, productId uint32, quantity uint32) error
}
