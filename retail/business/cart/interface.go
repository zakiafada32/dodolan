package cart

type Service interface {
	Update(userId string, cartItem CartItem) error
	FindAll(userId string) (Cart, error)
	DeleteCartItem(userId string, productsId []uint32) error
	Checkout(userId string, paymentId uint32, courier uint32) error
}

type Repository interface {
	Update(userId string, productId uint32, quantity uint32) error
	FindAll(userId string) ([]CartItemAtt, error)
	DeleteCartItem(userId string, productsId []uint32) error
	Checkout(userId string, paymentId uint32, courier uint32, cart Cart) error
}
