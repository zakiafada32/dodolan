package cart

type Service interface {
	UpdateCartItem(userId string, cartItem CartItem) error
	// Checkout(userId string) error
	// FindCartItem(userId string) (Cart, error)
	// DeleteCart(userId string, productsId []uint32) error
}

type Repository interface {
	UpdateCartItem(userId string, productId uint32, quantity uint32) error
	// FindCartItem(userId string, cartItem CartItem) ([]CartItem, error)
	// DeleteCartItem(userId string, cartItem CartItem) error
	// Checkout(userId string) error
}
